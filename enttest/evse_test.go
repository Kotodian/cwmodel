package enttest

import (
	"context"
	"testing"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/gokit/datasource"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvse(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	cli.Evse.Use(func(next ent.Mutator) ent.Mutator {
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

	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.Evse.Create().
		SetEquipment(equip).
		SetSerial("1").
		SetConnectorNumber(1).
		Exec(ctx)
	assert.Nil(t, err)
}
