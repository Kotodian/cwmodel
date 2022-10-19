package enttest

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvse(t *testing.T) {
	cli := Open(t, "mysql", dsn)
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
