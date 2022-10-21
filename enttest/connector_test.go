package enttest

import (
	"context"
	"testing"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/gokit/datasource"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateConnector(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	cli.Connector.Use(func(next ent.Mutator) ent.Mutator {
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

	e, err := equip.QueryEvse().Unique(false).
		Where(evse.SerialEQ("1")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, e)

	err = cli.Connector.Create().
		SetEquipment(equip).
		SetEvse(e).
		SetEvseSerial(e.Serial).
		SetSerial("1").
		SetEquipmentSn(equip.Sn).
		SetCurrentState(0).
		SetBeforeState(0).
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryConnector(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	connectors, err := equip.QueryConnector().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, connectors)

}

func TestUpdateConnector(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	cli.Connector.Use(func(next ent.Mutator) ent.Mutator {
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

	connectors, err := equip.QueryConnector().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, connectors)

	for _, connector := range connectors {
		connector, err = connector.Update().
			SetBeforeState(1).
			SetCurrentState(1).
			Save(ctx)
		assert.Nil(t, err)
		assert.Equal(t, 1, connector.BeforeState)
		assert.Equal(t, 1, connector.CurrentState)
	}
}
