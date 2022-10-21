package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentQuery(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()

	equipments, err := cli.Equipment.Query().Unique(false).All(context.TODO())
	assert.Nil(t, err)
	t.Log(equipments)
}
