package model

type Device struct {
	ID          int
	Name        string
	Size        DeviceSize
	WarehouseID int
}

func NewDevice(id int, name string, size DeviceSize) (*Device, error) {
	if name == "" {
		return nil, ErrNameRequired
	}

	return &Device{
		ID:   id,
		Name: name,
		Size: size,
	}, nil
}

type DeviceSize int

const (
	DeviceSizeXS DeviceSize = iota
	DeviceSizeS
	DeviceSizeM
	DeviceSizeL
	DeviceSizeXL
)
