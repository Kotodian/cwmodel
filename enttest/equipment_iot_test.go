package enttest

import (
	"context"
	"testing"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/gokit/datasource"
	"github.com/stretchr/testify/assert"
)

func TestCreateEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	cli.EquipmentIot.Use(func(next ent.Mutator) ent.Mutator {
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

	err = cli.EquipmentIot.Create().
		SetEquipment(equip).
		SetIccid("123456667").
		SetImei("123127841724").
		SetRemoteAddress("127.0.0.1").
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	equip, err := cli.Equipment.Get(ctx, 336379858853894)

	assert.Nil(t, err)
	assert.NotNil(t, equip)

	iot, err := equip.QueryEquipmentIot().Unique(false).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, iot)
}

func TestUpdateEquipmentIot(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	cli.EquipmentIot.Use(func(next ent.Mutator) ent.Mutator {
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
	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	iot, err := equip.QueryEquipmentIot().Unique(false).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, iot)

	iot, err = iot.Update().
		SetIccid("12345678").
		SetImei("1231278417245").Save(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "12345678", *iot.Iccid)
	assert.Equal(t, "1231278417245", *iot.Imei)
}
