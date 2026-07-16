// Package customdata keeps OCPP 2.0.1/2.1 customData vendor extension
// fields intact across Go marshal/unmarshal instead of silently dropping
// them, since generated CustomData structs only declare VendorID even
// though the official schema allows arbitrary additional properties.
package customdata

import "encoding/json"

// Marshal encodes vendorId together with any additional vendor-defined
// properties collected by Unmarshal. If extra also contains a "vendorId"
// key, vendorID always wins.
func Marshal(vendorID string, extra map[string]json.RawMessage) ([]byte, error) {
	out := make(map[string]json.RawMessage, len(extra)+1)
	for key, value := range extra {
		out[key] = value
	}
	encodedVendorID, err := json.Marshal(vendorID)
	if err != nil {
		return nil, err
	}
	out["vendorId"] = encodedVendorID
	return json.Marshal(out)
}

// Unmarshal splits vendorId from any additional vendor-defined properties so
// they are preserved instead of discarded.
func Unmarshal(data []byte, vendorID *string, extra *map[string]json.RawMessage) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if value, ok := raw["vendorId"]; ok {
		if err := json.Unmarshal(value, vendorID); err != nil {
			return err
		}
		delete(raw, "vendorId")
	}
	if len(raw) > 0 {
		*extra = raw
	} else {
		*extra = nil
	}
	return nil
}
