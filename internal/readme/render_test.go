package readme

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"gotest.tools/v3/assert"
)

func TestCommand(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		expectedError  string
	}{
		{
			name:          "missing args",
			expectedError: "",
		},
		{
			name: "test render",
			args: []string{"basic.yaml"},
			expectedError: `
yaml
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app := NewTestApp(nil, nil, nil)
			errBuffer := bytes.NewBuffer(nil)
			outBuffer := bytes.NewBuffer(nil)

			command := Command(app)
			command.SetOut(outBuffer)
			command.SetErr(errBuffer)
			//command.SetIn()
			command.SetArgs(test.args)
			err := command.Execute()
			fmt.Println(err, errBuffer.String(), outBuffer.String())
		})
	}
}

func TestOrgPackageName(t *testing.T) {
	tests := []struct {
		name                string
		remoteURLs          []string
		expectedPackageName string
		expectedError       string
	}{
		{
			name:                "test working one",
			remoteURLs:          []string{"git@github.com:stakater-tekton-catalog/code-linting-mvn.git"},
			expectedPackageName: "stakater-tekton-catalog/code-linting-mvn",
			expectedError:       "",
		},
		{
			name:                "test non-git remote URL",
			remoteURLs:          []string{"https://github.com/stakater-tekton-catalog/code-linting-mvn.git"},
			expectedPackageName: "",
			expectedError:       "no support for url based remote",
		},
		{
			name:                "test invalid remote URL format",
			remoteURLs:          []string{"git@github.com:stakater-tekton-catalog/code-linting-mvn"},
			expectedPackageName: "stakater-tekton-catalog/code-linting-mvn",
			expectedError:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			repo, err := git.Init(memory.NewStorage(), memfs.New())
			assert.NilError(t, err)
			remotes := map[string]*config.RemoteConfig{
				"origin": {
					Name: "origin",
					URLs: test.remoteURLs,
				}}
			err = repo.SetConfig(&config.Config{Remotes: remotes})
			assert.NilError(t, err)

			packageName, err := orgPackageName(repo)
			if test.expectedError != "" {
				assert.Error(t, err, test.expectedError)
			} else {
				assert.NilError(t, err)
				assert.Equal(t, test.expectedPackageName, packageName)
			}
		})
	}
}
