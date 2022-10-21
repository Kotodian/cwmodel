package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/cwmodel/appmoduleinfo"
	"github.com/stretchr/testify/assert"
)

func TestQueryAppModuleInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	appModuleInfo, err := cli.AppModuleInfo.Query().
		Unique(false).
		Where(appmoduleinfo.NameEQ("jx-coregw")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, appModuleInfo)
}
