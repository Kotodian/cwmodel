package enttest

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/gokit/datasource"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestQueryEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	info, err := equip.QueryEquipmentInfo().Unique(false).Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}

func TestCreateEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	cli.EquipmentInfo.Use(func(next ent.Mutator) ent.Mutator {
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
	err = cli.EquipmentInfo.Create().SetEquipment(equip).
		SetAccessPod("jx-ac-ohf-cluster-0").
		SetEquipmentSn(equip.Sn).
		SetFirmwareID(307040015638533).
		SetModelID(237955140000011).
		SetState(false).
		SetEvseNumber(0).
		SetAlarmNumber(0).
		SetRegisterDatetime(time.Now().Unix()).
		SetRemoteAddress("127.0.0.1").
		SetManufacturerID(237955190000010).Exec(ctx)
	assert.Nil(t, err)
}

func TestUpdateEquipmentInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()
	cli.EquipmentInfo.Use(func(next ent.Mutator) ent.Mutator {
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
	equip, err := cli.Equipment.Get(ctx, 244667116421190)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	info, err := equip.QueryEquipmentInfo().
		Unique(false).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, info)

	info, err = info.Update().
		SetAccessPod("jx-ac-ocpp-cluster-1").
		SetRemoteAddress("127.0.0.1").
		Save(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "jx-ac-ocpp-cluster-1", info.AccessPod)
	assert.Equal(t, "127.0.0.1", info.RemoteAddress)
}
