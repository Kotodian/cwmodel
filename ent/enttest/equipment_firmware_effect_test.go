package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel/ent/equipmentfirmwareeffect"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFirmwareEffect(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	firmware, err := cli.Firmware.Get(ctx, 343755557568517)
	assert.Nil(t, err)
	assert.NotNil(t, firmware)

	err = cli.EquipmentFirmwareEffect.Update().
		SetEquipment(equip).
		Where(equipmentfirmwareeffect.RequestIDEQ(1666080165)).
		SetState(6).
		Exec(ctx)
	assert.Nil(t, err)
}
