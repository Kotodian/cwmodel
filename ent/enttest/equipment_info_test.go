package enttest

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestQueryEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", "root:jqcsms@uat123@tcp(192.168.0.4:3306)/jx-csms?parseTime=True")
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
	cli := Open(t, "mysql", "root:jqcsms@uat123@tcp(192.168.0.4:3306)/jx-csms?parseTime=True")
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	err = cli.EquipmentInfo.Create().SetEquipment(equip).
		SetAccessPod("jx-ac-ohf-cluster-0").
		SetEquipmentSn(equip.Sn).
		SetFirmwareID(307040015638533).
		SetModelID(237955140000011).
		SetState(false).
		SetManufacturerID(237955190000010).Exec(ctx)
	assert.Nil(t, err)
}
