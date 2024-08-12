package readme

import (
	"io"

	"github.com/tektoncd/catlin/pkg/app"
	"go.uber.org/zap"
)

type testApp struct {
	log   *zap.Logger
	steam *app.Stream
}

var _ app.CLI = (*testApp)(nil)

func NewTestApp(in io.Reader, out, err io.Writer) *testApp {
	log, _ := zap.NewDevelopment()
	return &testApp{
		log:   log,
		steam: &app.Stream{In: in, Out: out, Err: err},
	}

}

func (t *testApp) Logger() *zap.Logger {
	return t.log
}

func (t *testApp) Stream() *app.Stream {
	//Needs to set to verify output
	return nil
}
