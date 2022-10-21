package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel/equipmentfirmwareeffect"
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

	err = cli.EquipmentFirmwareEffect.Update().
		SetEquipment(equip).
		Where(equipmentfirmwareeffect.RequestIDEQ(1666080165)).
		SetState(6).
		Exec(ctx)
	assert.Nil(t, err)
}
