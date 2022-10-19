package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
}
