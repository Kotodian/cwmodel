package enttest

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvse(t *testing.T) {
	cli := Open(t, "mysql", "root:jqcsms@uat123@tcp(192.168.0.4:3306)/jx-csms?parseTime=True")
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.Evse.Create().
		SetEquipment(equip).
		SetSerial("1").
		SetConnectorNumber(1).
		Exec(ctx)
	assert.Nil(t, err)
}
