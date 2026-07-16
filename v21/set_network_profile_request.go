// Code generated from the official OCPP JSON Schema. DO NOT EDIT.

package v21

import (
	"github.com/seanlee0923/ocpp/internal/validation"
	"github.com/seanlee0923/ocpp/protocol"
)

var _ protocol.Payload = SetNetworkProfileRequest{}

var schemaSetNetworkProfileRequest = &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"configurationSlot": &validation.Schema{Type: "integer", AllowAdditional: true}, "connectionData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"apn": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"apn": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "apnUserName": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "apnPassword": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 64, HasMaxLength: true}, "simPin": &validation.Schema{Type: "integer", AllowAdditional: true}, "preferredNetwork": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 6, HasMaxLength: true}, "useOnlyPreferredNetwork": &validation.Schema{Type: "boolean", AllowAdditional: true}, "apnAuthentication": &validation.Schema{Type: "string", Enum: []string{"PAP", "CHAP", "NONE", "AUTO"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"apn", "apnAuthentication"}}, "ocppVersion": &validation.Schema{Type: "string", Enum: []string{"OCPP12", "OCPP15", "OCPP16", "OCPP20", "OCPP201", "OCPP21"}}, "ocppInterface": &validation.Schema{Type: "string", Enum: []string{"Wired0", "Wired1", "Wired2", "Wired3", "Wireless0", "Wireless1", "Wireless2", "Wireless3", "Any"}}, "ocppTransport": &validation.Schema{Type: "string", Enum: []string{"SOAP", "JSON"}}, "messageTimeout": &validation.Schema{Type: "integer", AllowAdditional: true}, "ocppCsmsUrl": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "securityProfile": &validation.Schema{Type: "integer", AllowAdditional: true, Minimum: 0, HasMinimum: true}, "identity": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 48, HasMaxLength: true}, "basicAuthPassword": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 64, HasMaxLength: true}, "vpn": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"server": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 2000, HasMaxLength: true}, "user": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "group": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 50, HasMaxLength: true}, "password": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 64, HasMaxLength: true}, "key": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}, "type": &validation.Schema{Type: "string", Enum: []string{"IKEv2", "IPSec", "L2TP", "PPTP"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"server", "user", "password", "key", "type"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"ocppInterface", "ocppTransport", "messageTimeout", "ocppCsmsUrl", "securityProfile"}}, "customData": &validation.Schema{Type: "object", Properties: map[string]*validation.Schema{"vendorId": &validation.Schema{Type: "string", AllowAdditional: true, MaxLength: 255, HasMaxLength: true}}, Required: []string{"vendorId"}, AllowAdditional: true}}, Required: []string{"configurationSlot", "connectionData"}}

type SetNetworkProfileRequest struct {
	ConfigurationSlot int                                              `json:"configurationSlot"`
	ConnectionData    SetNetworkProfileRequestNetworkConnectionProfile `json:"connectionData"`
	CustomData        *SetNetworkProfileRequestCustomData              `json:"customData,omitempty"`
}

type SetNetworkProfileRequestNetworkConnectionProfile struct {
	Apn               *SetNetworkProfileRequestAPN              `json:"apn,omitempty"`
	OCPPVersion       *SetNetworkProfileRequestOCPPVersionEnum  `json:"ocppVersion,omitempty"`
	OCPPInterface     SetNetworkProfileRequestOCPPInterfaceEnum `json:"ocppInterface"`
	OCPPTransport     SetNetworkProfileRequestOCPPTransportEnum `json:"ocppTransport"`
	MessageTimeout    int                                       `json:"messageTimeout"`
	OCPPCSMSUrl       string                                    `json:"ocppCsmsUrl"`
	SecurityProfile   int                                       `json:"securityProfile"`
	Identity          *string                                   `json:"identity,omitempty"`
	BasicAuthPassword *string                                   `json:"basicAuthPassword,omitempty"`
	Vpn               *SetNetworkProfileRequestVPN              `json:"vpn,omitempty"`
	CustomData        *SetNetworkProfileRequestCustomData       `json:"customData,omitempty"`
}

type SetNetworkProfileRequestVPN struct {
	Server     string                              `json:"server"`
	User       string                              `json:"user"`
	Group      *string                             `json:"group,omitempty"`
	Password   string                              `json:"password"`
	Key        string                              `json:"key"`
	Type       SetNetworkProfileRequestVPNEnum     `json:"type"`
	CustomData *SetNetworkProfileRequestCustomData `json:"customData,omitempty"`
}

type SetNetworkProfileRequestVPNEnum string

const (
	SetNetworkProfileRequestVPNEnumIKEv2 SetNetworkProfileRequestVPNEnum = "IKEv2"
	SetNetworkProfileRequestVPNEnumIPSec SetNetworkProfileRequestVPNEnum = "IPSec"
	SetNetworkProfileRequestVPNEnumL2TP  SetNetworkProfileRequestVPNEnum = "L2TP"
	SetNetworkProfileRequestVPNEnumPPTP  SetNetworkProfileRequestVPNEnum = "PPTP"
)

type SetNetworkProfileRequestOCPPTransportEnum string

const (
	SetNetworkProfileRequestOCPPTransportEnumSOAP SetNetworkProfileRequestOCPPTransportEnum = "SOAP"
	SetNetworkProfileRequestOCPPTransportEnumJSON SetNetworkProfileRequestOCPPTransportEnum = "JSON"
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
	SetNetworkProfileRequestOCPPInterfaceEnumAny       SetNetworkProfileRequestOCPPInterfaceEnum = "Any"
)

type SetNetworkProfileRequestOCPPVersionEnum string

const (
	SetNetworkProfileRequestOCPPVersionEnumOCPP12  SetNetworkProfileRequestOCPPVersionEnum = "OCPP12"
	SetNetworkProfileRequestOCPPVersionEnumOCPP15  SetNetworkProfileRequestOCPPVersionEnum = "OCPP15"
	SetNetworkProfileRequestOCPPVersionEnumOCPP16  SetNetworkProfileRequestOCPPVersionEnum = "OCPP16"
	SetNetworkProfileRequestOCPPVersionEnumOCPP20  SetNetworkProfileRequestOCPPVersionEnum = "OCPP20"
	SetNetworkProfileRequestOCPPVersionEnumOCPP201 SetNetworkProfileRequestOCPPVersionEnum = "OCPP201"
	SetNetworkProfileRequestOCPPVersionEnumOCPP21  SetNetworkProfileRequestOCPPVersionEnum = "OCPP21"
)

type SetNetworkProfileRequestAPN struct {
	Apn                     string                                        `json:"apn"`
	ApnUserName             *string                                       `json:"apnUserName,omitempty"`
	ApnPassword             *string                                       `json:"apnPassword,omitempty"`
	SimPin                  *int                                          `json:"simPin,omitempty"`
	PreferredNetwork        *string                                       `json:"preferredNetwork,omitempty"`
	UseOnlyPreferredNetwork *bool                                         `json:"useOnlyPreferredNetwork,omitempty"`
	ApnAuthentication       SetNetworkProfileRequestAPNAuthenticationEnum `json:"apnAuthentication"`
	CustomData              *SetNetworkProfileRequestCustomData           `json:"customData,omitempty"`
}

type SetNetworkProfileRequestCustomData struct {
	VendorID string `json:"vendorId"`
}

type SetNetworkProfileRequestAPNAuthenticationEnum string

const (
	SetNetworkProfileRequestAPNAuthenticationEnumPAP  SetNetworkProfileRequestAPNAuthenticationEnum = "PAP"
	SetNetworkProfileRequestAPNAuthenticationEnumCHAP SetNetworkProfileRequestAPNAuthenticationEnum = "CHAP"
	SetNetworkProfileRequestAPNAuthenticationEnumNONE SetNetworkProfileRequestAPNAuthenticationEnum = "NONE"
	SetNetworkProfileRequestAPNAuthenticationEnumAUTO SetNetworkProfileRequestAPNAuthenticationEnum = "AUTO"
)

func (SetNetworkProfileRequest) ActionName() string { return "SetNetworkProfile" }

func (SetNetworkProfileRequest) Version() protocol.Version { return protocol.OCPP21 }

func (SetNetworkProfileRequest) Direction() protocol.PayloadDirection { return protocol.RequestPayload }

func (SetNetworkProfileRequest) SchemaName() string { return "SetNetworkProfileRequest.json" }

func (message SetNetworkProfileRequest) Validate() error {
	return validation.Validate("SetNetworkProfileRequest.json", schemaSetNetworkProfileRequest, message)
}

func (SetNetworkProfileRequest) ValidateJSON(data []byte) error {
	return validation.ValidateJSON("SetNetworkProfileRequest.json", schemaSetNetworkProfileRequest, data)
}
