package enttest

import (
	"context"
	"strconv"
	"testing"
	"time"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/gokit/datasource"
	"github.com/Kotodian/gokit/id"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()
	cli.OrderInfo.Use(func(next ent.Mutator) ent.Mutator {
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

	conn, err := equip.QueryConnector().Unique(false).
		Where(connector.SerialEQ("1")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, conn)

	transactionID := strconv.FormatInt(time.Now().Unix(), 10)
	err = cli.OrderInfo.Create().
		SetEquipment(equip).
		SetConnector(conn).
		SetOrderStartTime(time.Now().Unix()).
		SetChargeStartTime(time.Now().Unix()).
		SetTransactionID(transactionID).
		SetCallerOrderID(id.Next().String()).
		SetState(0).
		SetAuthorizationID("1").
		SetAuthorizationMode(1).
		SetPriceSchemeReleaseID(-1).
		SetChargeStartElectricity(1000).
		SetOffline(false).
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryOrderInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()

	equip, err := cli.Equipment.Get(ctx, 336379858853894)
	assert.Nil(t, err)
	assert.NotNil(t, equip)

	conn, err := equip.QueryConnector().Unique(false).
		Where(connector.SerialEQ("1")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, conn)

	orders, err := conn.QueryOrderInfo().Unique(false).All(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, orders)
	t.Log(len(orders))
}

func TestUpdateOrderInfo(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()
	cli.OrderInfo.Use(func(next ent.Mutator) ent.Mutator {
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

	conn, err := equip.QueryConnector().Unique(false).
		Where(connector.SerialEQ("1")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, conn)

	orders, err := conn.QueryOrderInfo().Unique(false).All(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, orders)

	lo.ForEach(orders, func(order *cwmodel.OrderInfo, _ int) {
		stopTime := time.Now().Unix()
		order, err = order.Update().
			SetChargeFinalTime(stopTime).
			SetOrderFinalTime(stopTime).
			SetChargeFinalElectricity(1000).Save(ctx)
		assert.Nil(t, err)
		assert.Equal(t, stopTime, *order.ChargeFinalTime)
		assert.Equal(t, stopTime, *order.OrderFinalTime)
		assert.Equal(t, float64(1000), *order.ChargeFinalElectricity)
	})
}
