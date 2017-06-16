package protocol

const ADDRESS_SEPARATOR = "~"

// NewSignalAddress returns a new signal address.
func NewSignalAddress(name, deviceID string) *SignalAddress {
	addr := SignalAddress{
		name:     name,
		deviceID: deviceID,
	}

	return &addr
}

// SignalAddress is a combination of a name and a device ID.
type SignalAddress struct {
	name     string
	deviceID string
}

// Name returns the signal address's name.
func (s *SignalAddress) Name() string {
	return s.name
}

// DeviceID returns the signal address's device ID.
func (s *SignalAddress) DeviceID() string {
	return s.deviceID
}

// String returns a string of both the address name and device id.
func (s *SignalAddress) String() string {
	return s.name + ADDRESS_SEPARATOR + s.deviceID
}
