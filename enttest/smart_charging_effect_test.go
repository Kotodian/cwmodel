package enttest

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/types"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"
)

func TestCreateSmartChargingEffect(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	ctx := context.TODO()
	defer cli.Close()

	cli.SmartChargingEffect.Use(func(next ent.Mutator) ent.Mutator {
		type UpdateBy interface {
			SetUpdatedBy(value datasource.UUID)
		}
		type CreateBy interface {
			SetCreatedBy(value datasource.UUID)
		}
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op() == ent.OpUpdate || m.Op() == ent.OpCreate || m.Op() == ent.OpUpdateOne {
				if m.Op() == ent.OpCreate {
					if cb, ok := m.(CreateBy); ok {
						cb.SetCreatedBy(datasource.UUID(999999))
					}
				}
				if ub, ok := m.(UpdateBy); ok {
					ub.SetUpdatedBy(99999)
				}
			}
			return next.Mutate(ctx, m)
		})
	})

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	connector, err := equip.QueryConnector().Unique(false).Where(connector.SerialEQ("1")).Where(connector.EvseSerial("1")).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, connector)

	err = cli.SmartChargingEffect.Create().
		SetEquipment(equip).
		SetConnectorID(connector.ID).
		SetSmartID(time.Now().Unix()).
		SetPid(id.Next()).
		SetUnit("W").
		SetEquipmentSn(equip.Sn).
		SetSpec([]types.ChargingSchedulePeriod{{From: 0, To: 1600, Limit: 1}}).
		Exec(ctx)
	assert.Nil(t, err)

}
func TestQuerySmartChargingEffect(t *testing.T) {

}
