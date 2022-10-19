package enttest

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentQuery(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	defer cli.Close()
	equipments, err := cli.Equipment.Query().All(context.TODO())
	assert.Nil(t, err)
	t.Log(equipments)
}
