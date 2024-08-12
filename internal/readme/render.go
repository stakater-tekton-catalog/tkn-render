package readme

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/tektoncd/catlin/pkg/app"
	"github.com/tektoncd/catlin/pkg/parser"
)

//go:embed readme.tmpl
var readmeTmpl embed.FS

func Command(cli app.CLI) *cobra.Command {
	var url string
	cmd := &cobra.Command{
		Use:  "render",
		Args: validResourcePath(),

		RunE: func(cmd *cobra.Command, args []string) error {
			return render(cli, args, url)
		},
	}

	cmd.PersistentFlags().StringVar(&url, "url", "https://raw.githubusercontent.com", "Base url")

	return cmd
}

func validResourcePath() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least 1 path to tekton resource yaml but received none")
		}

		return nil
	}
}

func render(cli app.CLI, args []string, url string) error {
	filename := args[0]

	printStdErr := func(err error) {
		fmt.Fprint(cli.Stream().Err, err.Error()+"\n")
	}

	r, err := os.Open(filename)
	if err != nil {
		return err
	}

	res, err := parser.ForReader(r).Parse()
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get wd")
	}

	repo, err := git.PlainOpen(wd)
	if err != nil {
		return fmt.Errorf("failed to open repo: %w", err)
	}

	tmpl, err := template.New("readme.tmpl").
		Funcs(template.FuncMap{
			"GetUrl": func() string {
				path, err := orgPackageName(repo)
				if err != nil {
					printStdErr(err)
				}
				return fmt.Sprintf("%s/%s/%s/%s", url, path, res.Unstructured.GetLabels()["app.kubernetes.io/version"], filename)
			},
			"Versions": func() (t []string) {
				//return t
				tags, err := repo.Tags()
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(tags.ForEach(func(ref *plumbing.Reference) error {
					t = append(t, strings.ReplaceAll(ref.Name().String(), "refs/tags/", ""))
					return nil
				}))
				printStdErr(fmt.Errorf("%s", t))
				return t
			},
		}).
		ParseFS(readmeTmpl, "readme.tmpl")
	if err != nil {
		return fmt.Errorf("failed to template parse fs: %w", err)
	}

	err = tmpl.Execute(cli.Stream().Out, res.Unstructured.Object)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func orgPackageName(repo *git.Repository) (string, error) {
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", fmt.Errorf("remote 'origin' not found: %w", err)
	}

	if len(remote.Config().URLs) == 0 {
		return "", fmt.Errorf("no remote urls configured")
	}

	url := remote.Config().URLs[0]
	if !strings.HasPrefix(url, "git@") {
		return "", fmt.Errorf("no support for url based remote")
	}

	paths := strings.Split(url, ":")

	return strings.TrimSuffix(paths[1], ".git"), nil
}

//Todo should we use first arg or flag to say were the file is located in repo?
