package enttest

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"
)

func TestCreateEquipmentLog(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	cli.EquipmentLog.Use(func(next ent.Mutator) ent.Mutator {
		type UpdateBy interface {
			SetUpdatedAt(value int64)
			SetUpdatedBy(value datasource.UUID)
		}
		type CreateBy interface {
			SetCreatedAt(value int64)
			SetCreatedBy(value datasource.UUID)
		}
		type ID interface {
			SetID(value datasource.UUID)
		}
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op() == ent.OpUpdate || m.Op() == ent.OpCreate || m.Op() == ent.OpUpdateOne {
				if m.Op() == ent.OpCreate {
					if cb, ok := m.(CreateBy); ok {
						m.(ID).SetID(id.Next())
						cb.SetCreatedAt(time.Now().Unix())
						cb.SetCreatedBy(datasource.UUID(999999))
					}
				}
				if ub, ok := m.(UpdateBy); ok {
					ub.SetUpdatedBy(99999)
					ub.SetUpdatedAt(time.Now().Unix())
				}
			}
			return next.Mutate(ctx, m)
		})
	})

	err = cli.EquipmentLog.Create().
		SetEquipment(equip).
		SetRequestID(time.Now().Unix()).
		SetDataLink(id.Next()).
		SetState(0).
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryEquipmentLog(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	logs, err := equip.QueryEquipmentLog().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
}
