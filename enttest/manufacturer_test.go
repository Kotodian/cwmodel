package enttest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryManufacturer(t *testing.T) {
	cli := Open(t, "mysql", dsn)
	cli = cli.Debug()
	ctx := context.TODO()
	defer cli.Close()

	manufacturer, err := cli.Manufacturer.Query().All(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, manufacturer)
}
