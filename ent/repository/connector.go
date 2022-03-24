package repository

import (
	"context"

	"github.com/Kotodian/ent-practice/ent"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/gokit/datasource"
)

func FindConnectors(ctx context.Context, client *ent.Client, equipmentID datasource.UUID, fields ...string) ([]*ent.Connector, error) {
	query := client.Connector.Query()
	var sel *ent.ConnectorSelect
	if len(fields) > 0 {
		sel = query.Select(fields...)
	} else {
		sel = query.Select(connector.Columns...)
	}
	return sel.QueryEquipment().Where(equipment.IDEQ(equipmentID)).QueryConnectors().All(ctx)
}

func GetConnector(ctx context.Context, client *ent.Client, id datasource.UUID, fields ...string) (*ent.Connector, error) {
	query := client.Connector.Query()
	var sel *ent.ConnectorSelect
	if len(fields) > 0 {
		sel = query.Select(fields...)
	} else {
		sel = query.Select(connector.Columns...)
	}
	return sel.Where(connector.IDEQ(id)).Only(ctx)
}

func GetConnectorBySerial(ctx context.Context, client *ent.Client, equipmentID datasource.UUID, evseSerial string, connectorSerial string, fields ...string) (*ent.Connector, error) {
	query := client.Connector.Query()
	var sel *ent.ConnectorSelect
	if len(fields) > 0 {
		sel = query.Select(fields...)
	} else {
		sel = query.Select(connector.Columns...)
	}
	return sel.QueryEquipment().Where(equipment.IDEQ(equipmentID)).
				QueryConnectors().
				Where(connector.EvseSerialEQ(evseSerial), connector.SerialEQ(connectorSerial)).
				Only(ctx)
}

