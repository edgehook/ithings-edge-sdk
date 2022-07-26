package metadata

import (
	"time"
)

// TwinProperty represents the device property for which an Expected/Actual state can be defined.
type TwinProperty struct {
	Service string `json:"svc"`
	// Required: The property name for which the desired/reported values are specified.
	// This property should be present in the device model.
	PropertyName string `json:"pn"`
	// Required: The value for this property.
	Value interface{} `json:"val,omitempty"`
	//the timestamp of the property collecting value
	// +optional
	Timestamp int64 `json:"ts,omitempty"`
	// some error message info when collecting this property.
	// +optional, just for reported.
	ErrorMessage string `json:"err_msg,omitempty"`

	WriteAble   bool    `json:"-"`
	MaxValue    float64 `json:"-"`
	MinValue    float64 `json:"-"`
	Unit        string  `json:"-"`
	DataType    string  `json:"-"`
	Description string  `json:"-"`
}

//New device twin property.
func NewTwinProperty(svc, propName, errMsg string, val interface{}) *TwinProperty {
	now := time.Now().Unix()

	return &TwinProperty{
		Service:      svc,
		PropertyName: propName,
		Value:        val,
		Timestamp:    now,
		ErrorMessage: errMsg,
	}
}

// DeviceServiceSpec is the  an instantation of a DeviceServiceModel.
type DeviceServiceSpec struct {
	Name       string                `json:"name"`
	Properties []*DevicePropertySpec `json:"props,omitempty"`
	Events     []*DeviceEventSpec    `json:"events,omitempty"`
	Commands   []*DeviceCommandSpec  `json:"cmds,omitempty"`
}

func (dss *DeviceServiceSpec) FindDevicePropertySpec(name string) *DevicePropertySpec {
	for _, p := range dss.Properties {
		if p.Name == name {
			return p
		}
	}

	return nil
}

func (dss *DeviceServiceSpec) FindDeviceEventSpec(name string) *DeviceEventSpec {
	for _, e := range dss.Events {
		if e.Name == name {
			return e
		}
	}

	return nil
}

func (dss *DeviceServiceSpec) FindDeviceCommandSpec(name string) *DeviceCommandSpec {
	for _, c := range dss.Commands {
		if c.Name == name {
			return c
		}
	}

	return nil
}

// DevicePropertySpec is an instantation of a DevicePropertyModel.
type DevicePropertySpec struct {
	*DevicePropertyModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"ac,omitempty"`
}

// DeviceEventSpec is an instantation of a DeviceEventModel.
type DeviceEventSpec struct {
	*DeviceEventModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"ac,omitempty"`
}

// DeviceCommandSpec is an instantation of a DeviceCommandModel.
type DeviceCommandSpec struct {
	*DeviceCommandModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"ac,omitempty"`
}
