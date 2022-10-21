package enttest

import (
	"context"
	"testing"
	"time"

	"github.com/Kotodian/cwmodel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestQueryEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	info, err := equip.QueryEquipmentInfo().Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestCreateEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	err = cli.EquipmentInfo.Create().SetEquipment(equip).
		SetAccessPod("jx-ac-ohf-cluster-0").
		SetEquipmentSn(equip.Sn).
		SetFirmwareID(307040015638533).
		SetModelID(237955140000011).
		SetState(false).
		SetEvseNumber(0).
		SetAlarmNumber(0).
		SetRegisterDatetime(time.Now().Unix()).
		SetRemoteAddress("127.0.0.1").
		SetManufacturerID(237955190000010).Exec(ctx)
	assert.Nil(t, err)
}

func TestUpdateEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	info, err := equip.QueryEquipmentInfo().
		Unique(false).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = info.Update().
		SetAccessPod("jx-ac-ocpp-cluster-1").
		SetRemoteAddress("127.0.0.1").
		Save(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "jx-ac-ocpp-cluster-1", info.AccessPod)
	assert.Equal(t, "127.0.0.1", info.RemoteAddress)
}
