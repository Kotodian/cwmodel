package enttest

import (
	"context"
	"testing"
	"time"

	"github.com/Kotodian/cwmodel"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderEvent(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()

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
