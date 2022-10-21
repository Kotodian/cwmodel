package enttest

import (
	"context"
	"testing"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/cwmodel/equipmentfirmwareeffect"
	"github.com/Kotodian/gokit/datasource"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFirmwareEffect(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	cli.EquipmentFirmwareEffect.Use(func(next ent.Mutator) ent.Mutator {
		type UpdateBy interface {
			SetUpdatedBy(value datasource.UUID)
		}
		type CreateBy interface {
			SetCreatedBy(value datasource.UUID)
		}
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op() == ent.OpUpdate || m.Op() == ent.OpCreate {
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

	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	err = cli.EquipmentFirmwareEffect.Update().
		SetEquipment(equip).
		Where(equipmentfirmwareeffect.RequestIDEQ(1666080165)).
		SetState(6).
		Exec(ctx)
	assert.Nil(t, err)
}
