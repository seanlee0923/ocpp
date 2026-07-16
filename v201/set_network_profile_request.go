// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v201

import (
	"ocpp-go/internal/validation"
	"ocpp-go/protocol"
)

var _ protocol.Payload = SetNetworkProfileRequest{}

var schemaSetNetworkProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "configurationSlot": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectionData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "apn": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "apn": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "apnUserName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "apnPassword": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "simPin": &validation.Schema{Type: "integer", AllowAdditional: true}, "preferredNetwork": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 6, HasMaxLength: true}, "useOnlyPreferredNetwork": &validation.Schema{Type: "boolean", AllowAdditional: true}, "apnAuthentication": &validation.Schema{Type: "string", Enum: []string{"CHAP", "NONE", "PAP", "AUTO"}}}, Required: []string{"apn", "apnAuthentication"}}, "ocppVersion": &validation.Schema{Type: "string", Enum: []string{"OCPP12", "OCPP15", "OCPP16", "OCPP20"}}, "ocppTransport": &validation.Schema{Type: "string", Enum: []string{"JSON", "SOAP"}}, "ocppCsmsUrl": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "messageTimeout": &validation.Schema{Type: "integer", AllowAdditional: true}, "securityProfile": &validation.Schema{Type: "integer", AllowAdditional: true}, "ocppInterface": &validation.Schema{Type: "string", Enum: []string{"Wired0", "Wired1", "Wired2", "Wired3", "Wireless0", "Wireless1", "Wireless2", "Wireless3"}}, "vpn": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}, "server": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 512, HasMaxLength: true}, "user": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "group": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "password": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 20, HasMaxLength: true}, "key": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"IKEv2", "IPSec", "L2TP", "PPTP"}}}, Required: []string{"server", "user", "password", "key", "type"}}}, Required: []string{"ocppVersion", "ocppTransport", "ocppCsmsUrl", "messageTimeout", "securityProfile", "ocppInterface"}}}, Required: []string{"configurationSlot", "connectionData"}}

type SetNetworkProfileRequest struct {
	CustomData        *SetNetworkProfileRequestCustomData              `json:"customData,omitempty"`
	ConfigurationSlot int                                              `json:"configurationSlot"`
	ConnectionData    SetNetworkProfileRequestNetworkConnectionProfile `json:"connectionData"`
}

type SetNetworkProfileRequestNetworkConnectionProfile struct {
	CustomData      *SetNetworkProfileRequestCustomData       `json:"customData,omitempty"`
	Apn             *SetNetworkProfileRequestAPN              `json:"apn,omitempty"`
	OCPPVersion     SetNetworkProfileRequestOCPPVersionEnum   `json:"ocppVersion"`
	OCPPTransport   SetNetworkProfileRequestOCPPTransportEnum `json:"ocppTransport"`
	OCPPCSMSUrl     string                                    `json:"ocppCsmsUrl"`
	MessageTimeout  int                                       `json:"messageTimeout"`
	SecurityProfile int                                       `json:"securityProfile"`
	OCPPInterface   SetNetworkProfileRequestOCPPInterfaceEnum `json:"ocppInterface"`
	Vpn             *SetNetworkProfileRequestVPN              `json:"vpn,omitempty"`
}

type SetNetworkProfileRequestVPN struct {
	CustomData *SetNetworkProfileRequestCustomData `json:"customData,omitempty"`
	Server     string                              `json:"server"`
	User       string                              `json:"user"`
	Group      *string                             `json:"group,omitempty"`
	Password   string                              `json:"password"`
	Key        string                              `json:"key"`
	Type       SetNetworkProfileRequestVPNEnum     `json:"type"`
}

type SetNetworkProfileRequestVPNEnum string

const (
	SetNetworkProfileRequestVPNEnumIKEv2 SetNetworkProfileRequestVPNEnum = "IKEv2"
	SetNetworkProfileRequestVPNEnumIPSec SetNetworkProfileRequestVPNEnum = "IPSec"
	SetNetworkProfileRequestVPNEnumL2TP  SetNetworkProfileRequestVPNEnum = "L2TP"
	SetNetworkProfileRequestVPNEnumPPTP  SetNetworkProfileRequestVPNEnum = "PPTP"
)

type SetNetworkProfileRequestOCPPInterfaceEnum string

const (
	SetNetworkProfileRequestOCPPInterfaceEnumWired0    SetNetworkProfileRequestOCPPInterfaceEnum = "Wired0"
	SetNetworkProfileRequestOCPPInterfaceEnumWired1    SetNetworkProfileRequestOCPPInterfaceEnum = "Wired1"
	SetNetworkProfileRequestOCPPInterfaceEnumWired2    SetNetworkProfileRequestOCPPInterfaceEnum = "Wired2"
	SetNetworkProfileRequestOCPPInterfaceEnumWired3    SetNetworkProfileRequestOCPPInterfaceEnum = "Wired3"
	SetNetworkProfileRequestOCPPInterfaceEnumWireless0 SetNetworkProfileRequestOCPPInterfaceEnum = "Wireless0"
	SetNetworkProfileRequestOCPPInterfaceEnumWireless1 SetNetworkProfileRequestOCPPInterfaceEnum = "Wireless1"
	SetNetworkProfileRequestOCPPInterfaceEnumWireless2 SetNetworkProfileRequestOCPPInterfaceEnum = "Wireless2"
	SetNetworkProfileRequestOCPPInterfaceEnumWireless3 SetNetworkProfileRequestOCPPInterfaceEnum = "Wireless3"
)

type SetNetworkProfileRequestOCPPTransportEnum string

const (
	SetNetworkProfileRequestOCPPTransportEnumJSON SetNetworkProfileRequestOCPPTransportEnum = "JSON"
	SetNetworkProfileRequestOCPPTransportEnumSOAP SetNetworkProfileRequestOCPPTransportEnum = "SOAP"
)

type SetNetworkProfileRequestOCPPVersionEnum string

const (
	SetNetworkProfileRequestOCPPVersionEnumOCPP12 SetNetworkProfileRequestOCPPVersionEnum = "OCPP12"
	SetNetworkProfileRequestOCPPVersionEnumOCPP15 SetNetworkProfileRequestOCPPVersionEnum = "OCPP15"
	SetNetworkProfileRequestOCPPVersionEnumOCPP16 SetNetworkProfileRequestOCPPVersionEnum = "OCPP16"
	SetNetworkProfileRequestOCPPVersionEnumOCPP20 SetNetworkProfileRequestOCPPVersionEnum = "OCPP20"
)

type SetNetworkProfileRequestAPN struct {
	CustomData              *SetNetworkProfileRequestCustomData           `json:"customData,omitempty"`
	Apn                     string                                        `json:"apn"`
	ApnUserName             *string                                       `json:"apnUserName,omitempty"`
	ApnPassword             *string                                       `json:"apnPassword,omitempty"`
	SimPin                  *int                                          `json:"simPin,omitempty"`
	PreferredNetwork        *string                                       `json:"preferredNetwork,omitempty"`
	UseOnlyPreferredNetwork *bool                                         `json:"useOnlyPreferredNetwork,omitempty"`
	ApnAuthentication       SetNetworkProfileRequestAPNAuthenticationEnum `json:"apnAuthentication"`
}

type SetNetworkProfileRequestAPNAuthenticationEnum string

const (
	SetNetworkProfileRequestAPNAuthenticationEnumCHAP SetNetworkProfileRequestAPNAuthenticationEnum = "CHAP"
	SetNetworkProfileRequestAPNAuthenticationEnumNONE SetNetworkProfileRequestAPNAuthenticationEnum = "NONE"
	SetNetworkProfileRequestAPNAuthenticationEnumPAP  SetNetworkProfileRequestAPNAuthenticationEnum = "PAP"
	SetNetworkProfileRequestAPNAuthenticationEnumAUTO SetNetworkProfileRequestAPNAuthenticationEnum = "AUTO"
)

type SetNetworkProfileRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

func (SetNetworkProfileRequest) ActionName() string { return "SetNetworkProfile" }

func (SetNetworkProfileRequest) Version() protocol.Version { return protocol.OCPP201 }

func (SetNetworkProfileRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetNetworkProfileRequest) SchemaName() string { return "SetNetworkProfileRequest.json" }

func (message SetNetworkProfileRequest) Validate() error {
	return validation.Validate("SetNetworkProfileRequest.json", schemaSetNetworkProfileRequest, message)
}

func (SetNetworkProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetNetworkProfileRequest.json", schemaSetNetworkProfileRequest, data)
}
