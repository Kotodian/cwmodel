package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAlarm(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
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
