package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel"
	"github.com/stretchr/testify/assert"
)

func TestQueryManufacturer(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	ctx := context.TODO()
	defer cli.Close()

	manufacturer, err := cli.Manufacturer.Query().Unique(false).All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, manufacturer)
}
