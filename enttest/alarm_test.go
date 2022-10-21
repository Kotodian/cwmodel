package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCreateAlarm(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.EquipmentAlarm.Create().
		SetDtcCode(16715009).
		SetCount(0).
		SetEquipment(equip).
		SetRemoteAddress("127.0.0.1").
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryAlarms(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	alarms, err := equip.QueryEquipmentAlarm().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, alarms)
	t.Log(len(alarms))
}

func TestUpdateAlarms(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	alarms, err := equip.QueryEquipmentAlarm().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, alarms)

	lo.ForEach(alarms, func(alarm *cwmodel.EquipmentAlarm, _ int) {
		newAlarm, err := alarm.Update().AddCount(1).Save(ctx)
		assert.Nil(t, err)
		assert.Equal(t, 1, newAlarm.Count-alarm.Count)
	})
}
