package test

import (
	_ "bs.mobgi.cc/bootstrap"
	"bs.mobgi.cc/cronJobs/jobs/app/logic"
	"testing"
)

func TestApp(t *testing.T) {
	err := logic.NewAppQueryLogic().AppQuery()
	t.Log(err)
}
