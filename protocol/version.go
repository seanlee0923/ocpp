package protocol

import "fmt"

// Version is an OCPP-J WebSocket subprotocol.
type Version string

const (
	OCPP16  Version = "ocpp1.6"
	OCPP201 Version = "ocpp2.0.1"
	OCPP21  Version = "ocpp2.1"
)

func (v Version) Valid() bool {
	switch v {
	case OCPP16, OCPP201, OCPP21:
		return true
	default:
		return false
	}
}

func ParseVersion(value string) (Version, error) {
	v := Version(value)
	if !v.Valid() {
		return "", fmt.Errorf("unsupported OCPP subprotocol %q", value)
	}
	return v, nil
}

func SupportedVersions() []Version {
	return []Version{OCPP21, OCPP201, OCPP16}
}
