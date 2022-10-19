package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.EquipmentIot.Create().
		SetEquipment(equip).
		SetIccid("123456667").
		SetImei("123127841724").
		SetRemoteAddress("127.0.0.1").
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	iot, err := equip.QueryEquipmentIot().Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, iot)
}
