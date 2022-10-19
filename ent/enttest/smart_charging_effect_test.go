package enttest

import (
	"context"
	"testing"
	"time"

	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/types"
	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"
)

func TestCreateSmartChargingEffect(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	ctx := context.TODO()
	defer cli.Close()

	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	connector, err := equip.QueryConnector().Where(connector.SerialEQ("1")).Where(connector.EvseSerial("1")).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, connector)

	err = cli.SmartChargingEffect.Create().
		SetEquipment(equip).
		SetConnector(connector).
		SetSmartID(time.Now().Unix()).
		SetPid(id.Next()).
		SetUnit("W").
		SetEquipmentSn(equip.Sn).
		SetSpec([]types.ChargingSchedulePeriod{{From: 0, To: 1600, Limit: 1}}).
		Exec(ctx)
	assert.Nil(t, err)

}
func TestQuerySmartChargingEffect(t *testing.T) {

}
