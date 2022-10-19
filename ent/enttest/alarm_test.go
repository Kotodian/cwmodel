package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAlarm(t *testing.T) {
	cli := Open(t, "mysql", "root:jqcsms@uat123@tcp(192.168.0.4:3306)/jx-csms?parseTime=True")
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	err = cli.EquipmentAlarm.Create().Exec(ctx)
	assert.Nil(t, err)
}
