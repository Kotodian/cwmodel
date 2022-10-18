package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAlarm(t *testing.T) {
	cli := Open(t, "mysql", "root:jqcsms@uat123@tcp(192.168.0.4:3306)/jx-csms?parseTime=True")
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Debug().Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)
	cli.Debug().EquipmentAlarm.Create()
}
