package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/stretchr/testify/assert"
)

func TestQueryEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Query().Unique(false).Where(equipment.IDEQ(336379858853894)).First(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
}
