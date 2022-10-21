package enttest

import (
	"context"
	"testing"

	"github.com/Kotodian/cwmodel"
	"github.com/Kotodian/cwmodel/firmware"
	"github.com/stretchr/testify/assert"
)

func TestQueryFirmwareTemplate(t *testing.T) {
	cli := Open(t, "mysql", dsn, WithOptions(cwmodel.Debug()))
	defer cli.Close()
	ctx := context.TODO()

	model, err := cli.Model.Get(ctx, 237955140000000)
	assert.Nil(t, err)
	assert.NotNil(t, model)

	manufacturer, err := cli.Manufacturer.Get(ctx, 237955190000000)
	assert.Nil(t, err)
	assert.NotNil(t, manufacturer)

	firm, err := cli.Firmware.Query().Unique(false).
		Where(firmware.ModelIDEQ(model.ID)).
		Where(firmware.ManufacturerIDEQ(manufacturer.ID)).
		Where(firmware.EquipVersionEQ("0.0.1")).
		Only(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, firm)
}
