package cwmodel

func (client *Client) Config() config {
	return client.config
}

func (e *Equipment) SetConfig(conf config) *Equipment {
	e.config = conf
	return e
}

func (e *EquipmentInfo) SetConfig(conf config) *EquipmentInfo {
	e.config = conf
	return e
}

func (e *Evse) SetConfig(conf config) *Evse {
	e.config = conf
	return e
}

func (c *Connector) SetConfig(conf config) *Connector {
	c.config = conf
	return c
}

func (e *EquipmentAlarm) SetConfig(conf config) *EquipmentAlarm {
	e.config = conf
	return e
}

func (e *EquipmentFirmwareEffect) SetConfig(conf config) *EquipmentFirmwareEffect {
	e.config = conf
	return e
}

func (e *EquipmentIot) SetConfig(conf config) *EquipmentIot {
	e.config = conf
	return e
}

func (e *EquipmentLog) SetConfig(conf config) *EquipmentLog {
	e.config = conf
	return e
}

func (e *Firmware) SetConfig(conf config) *Firmware {
	e.config = conf
	return e
}

func (m *Manufacturer) SetConfig(conf config) *Manufacturer {
	m.config = conf
	return m
}

func (m *Model) SetConfig(conf config) *Model {
	m.config = conf
	return m
}

func (o *OrderEvent) SetConfig(conf config) *OrderEvent {
	o.config = conf
	return o
}

func (o *OrderInfo) SetConfig(conf config) *OrderInfo {
	o.config = conf
	return o
}

func (r *Reservation) SetConfig(conf config) *Reservation {
	r.config = conf
	return r
}

func (s *SmartChargingEffect) SetConfig(conf config) *SmartChargingEffect {
	s.config = conf
	return s
}
