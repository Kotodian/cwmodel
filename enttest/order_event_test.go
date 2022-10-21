package enttest

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent"
	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/gokit/datasource"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderEvent(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()
	cli.OrderEvent.Use(func(next ent.Mutator) ent.Mutator {
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
	orderInfo, err := cli.OrderInfo.Get(ctx, 302002886349637)
	assert.Nil(t, err)
	assert.NotNil(t, orderInfo)

	err = cli.OrderEvent.Create().
		SetOrderInfo(orderInfo).
		SetContent("测试").
		SetOccurrence(time.Now().Unix()).
		Exec(ctx)
	assert.Nil(t, err)
}

func TestQueryOrderEvent(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))

	ctx := context.TODO()
	defer cli.Close()

	orderInfo, err := cli.OrderInfo.Get(ctx, 302002886349637)
	assert.Nil(t, err)
	assert.NotNil(t, orderInfo)

	events, err := orderInfo.QueryOrderEvent().All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, events)
}
