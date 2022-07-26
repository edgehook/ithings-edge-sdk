package metadata

type EventDefType int32

const (
	EventDef_InfoEvent    EventDefType = 0
	EventDef_WarningEvent EventDefType = 1
	EventDef_AlertEvent   EventDefType = 2
)

// Represents the type and data validation of a property.
// Only one of its members may be specified.
type PropertyType struct {
	// +optional
	Int *PropertyTypeInt64 `json:"int,omitempty"`
	// +optional
	String *PropertyTypeString `json:"string,omitempty"`
	// +optional
	Double *PropertyTypeDouble `json:"double,omitempty"`
	// +optional
	Float *PropertyTypeFloat `json:"float,omitempty"`
	// +optional
	Boolean *PropertyTypeBoolean `json:"boolean,omitempty"`
	// +optional
	Bytes *PropertyTypeBytes `json:"bytes,omitempty"`
}

type PropertyTypeInt64 struct {
	// +optional
	DefaultValue int64 `json:"default_value,omitempty"`
	// +optional
	Minimum int64 `json:"min,omitempty"`
	// +optional
	Maximum int64 `json:"max,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeString struct {
	// +optional
	DefaultValue string `json:"default_value,omitempty"`
}

type PropertyTypeDouble struct {
	// +optional
	DefaultValue float64 `json:"default_value,omitempty"`
	// +optional
	Minimum float64 `json:"min,omitempty"`
	// +optional
	Maximum float64 `json:"max,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeFloat struct {
	// +optional
	DefaultValue float32 `json:"default_value,omitempty"`
	// +optional
	Minimum float32 `json:"min,omitempty"`
	// +optional
	Maximum float32 `json:"max,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeBoolean struct {
	// +optional
	DefaultValue bool `json:"default_value,omitempty"`
}

type PropertyTypeBytes struct {
}

// device command param defination
type RequestParamDef struct {
	DataType PropertyType `json:"dt,omitempty"`
}

// The access mode for  a device property.
// +kubebuilder:validation:Enum=ReadWrite;ReadOnly
type PropertyAccessMode string

// Access mode constants for a device property.
const (
	ReadWrite PropertyAccessMode = "ReadWrite"
	ReadOnly  PropertyAccessMode = "ReadOnly"
)

// device property defination
type DevicePropertyModel struct {
	// Required: The device property name.
	Name        string  `json:"pn"`
	WriteAble   bool    `json:"rw"`
	MaxValue    float64 `json:"max"`
	MinValue    float64 `json:"min"`
	Unit        string  `json:"un,omitempty"`
	DataType    string  `json:"dt,omitempty"`
	Description string  `json:"desc,omitempty"`
}

type DeviceCommandModel struct {
	Name        string `json:"cn"`
	Description string `json:"desc,omitempty"`
	// params
	RequestParam map[string]string `json:"req_param,omitempty"`
}

// device event defination
type DeviceEventModel struct {
	Name        string  `json:"en"`
	EventType   string  `json:"et,omitempty"`
	MaxValue    float64 `json:"max,omitempty"`
	MinValue    float64 `json:"min,omitempty"`
	Unit        string  `json:"un,omitempty"`
	DataType    string  `json:"dt,omitempty"`
	Description string  `json:"desc,omitempty"`
}

//service describe a based module contains some properties, events
// and commands.
type DeviceServiceModel struct {
	Name           string                 `json:"name"`
	Description    string                 `json:"desc,omitempty"`
	PropertyModels []*DevicePropertyModel `json:"prop_models,omitempty"`
	EventModels    []*DeviceEventModel    `json:"event_models,omitempty"`
	CommandModels  []*DeviceCommandModel  `json:"cmd_models,omitempty"`
}

// DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device
// capabilities.
type DeviceModelSpec struct {
	// Required: List of device services.
	ServiceModels []*DeviceServiceModel `json:"svc_models,omitempty"`
}

// DeviceModel is the Schema for the device model API
type DeviceModel struct {
	//update accessconfig and update protocol
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Industry     string `json:"industry,omitempty"`
	DataFormat   string `json:"data_format,omitempty"`
	DeviceNumber int64  `json:"device_number,omitempty"`
	TagNumber    int64  `json:"tag_number,omitempty"`
	GroupID      string `json:"group_id,omitempty"`
	//who create the device by ID.
	Creator         string `json:"creator,omitempty"`
	CreateTimeStamp int64  `json:"create_timestamp,omitempty"`
	UpdateTimeStamp int64  `json:"update_timestamp,omitempty"`
	//Spec.
	//Spec DeviceModelSpec `json:"spec,omitempty"`
	*DeviceModelSpec `json:",inline,omitempty"`
}

// DeviceModelList contains a list of DeviceModel
type DeviceModelList struct {
	Items []DeviceModel `json:"items"`
}
