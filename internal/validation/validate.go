package validation

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/tektoncd/catlin/pkg/parser"
	"github.com/tektoncd/catlin/pkg/validator"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/urfave/cli/v2"
)

func Validate(c *cli.Context) error {
	file, err := os.Open(c.Args().First())
	if err != nil {
		return err
	}
	defer file.Close()

	res, err := parser.ForReader(file).Parse()
	if err != nil {
		return err
	}
	result := validator.NewTaskValidator(res).Validate()
	fmt.Println(result)
	return nil
	//obj, g, err := file.Decode(c.Args().First())
	//if err != nil {
	//	return fmt.Errorf("failed to decode file: %w", err)
	//}
	//
	//if task, ok := obj.(*v1beta1.Task); ok {
	//	err := validateV1Beta1Task(c.Context, task)
	//	if err != nil {
	//		return fmt.Errorf("%s validation failed: %w", g.String(), err)
	//	}
	//}
	//
	////fmt.Println(obj.(*v1.Task).Spec.Params)
	//
	////fmt.Println(obj.(*v1beta1.Task).Spec.Params)
	////
	////checkErr(obj.(*v1.Task).ObjectMeta)
	//
	//return nil
}

func validateV1Beta1Task(ctx context.Context, task *v1beta1.Task) error {
	fieldError := task.Validate(ctx)
	if fieldError != nil {
		return fmt.Errorf(fieldError.Error())
	}
	err := validateLabels(task.Labels)
	if err != nil {
		return err
	}
	err = validateAnnotations(task.Annotations)
	if err != nil {
		return err
	}

	for _, param := range task.Spec.Params {
		err := validateParamName(param.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateLabels(labels map[string]string) error {
	return nil
}

func validateAnnotations(labels map[string]string) error {
	return nil
}

func validateParamName(name string) error {
	if screamingSnakeCase.FindString(name) == "" {
		return fmt.Errorf("faied to validate screaming snake case for param: %s", name)
	}
	return nil
}

var screamingSnakeCase = regexp.MustCompile(`[A-Z0-9]+(?:_[A-Z0-9]+)*`)
