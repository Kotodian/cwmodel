package enttest

import (
	"context"
	"testing"
	"time"

	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
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

	err = cli.Reservation.Create().
		SetEquipment(equip).
		SetConnector(connector).
		SetAuthorizationID(id.Next().String()).
		SetAuthorizationMode(1).
		SetAdditional(`[{"authorizationMode":1,"authorizationId":"299953297383429"}]`).
		SetReservationID(time.Now().Unix()).
		SetExpired(time.Now().Unix()).
		SetState(3).
		Exec(ctx)
	assert.Nil(t, err)
}
