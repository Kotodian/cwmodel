package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel/ent/evse"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnector(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	e, err := equip.QueryEvse().Where(evse.SerialEQ("1")).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, e)
	err = cli.Connector.Create().
		SetEquipment(equip).
		SetEvse(e).
		SetEvseSerial(e.Serial).
		SetSerial("1").
		SetEquipmentSn(equip.Sn).
		SetCurrentState(0).
		SetBeforeState(0).
		Exec(ctx)
	assert.Nil(t, err)
}
