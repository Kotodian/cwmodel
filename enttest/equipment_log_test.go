package enttest

import (
	"context"
	"testing"
	"time"

	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"
)

func TestCreateEquipmentLog(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.EquipmentLog.Create().
		SetEquipment(equip).
		SetRequestId(time.Now().Unix()).
		SetDataLink(id.Next()).
		SetState(0).
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryEquipmentLog(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	logs, err := equip.QueryEquipmentLog().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
}
