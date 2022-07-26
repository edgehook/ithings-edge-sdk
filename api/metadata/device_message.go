package metadata

const (
	DEVICE_STATUS_ONLINE  = "online"
	DEVICE_STATUS_OFFLINE = "offline"
)

/*
* DeviceSpecMeta
* this is for edge device or edge gateway
 */
type DeviceSpecMeta struct {
	DeviceID                 string `json:"id"`
	DeviceOS                 string `json:"os,omitempty"`
	DeviceCatagory           string `json:"catagory,omitempty"`
	DeviceIdentificationCode string `json:"id_code,omitempty"`

	//device should be started or stopped.
	State string `json:"state,omitempty"`
	// Additional metadata like tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`

	// Required: The protocol configuration used to connect to the device.
	//this should be a json string
	Protocol string `json:"proto_conf,omitempty"`

	// Required: List of device services.
	Services []*DeviceServiceSpec `json:"svcs,omitempty"`
}

/*
* Single device report message.
 */
type ReportDeviceMessage struct {
	DeviceID string `json:"d_id"`
	// Required: List of device services.
	Services []*TwinProperty `json:"svcs"`
}

/*
* Devices report message.
 */
type ReportDevicesMessage struct {
	Devices []*ReportDeviceMessage `json:"devs,omitempty"`
}

// single device status message
type DeviceStatusMessage struct {
	DeviceID string `json:"d_id"`
	//Required:
	//device status
	// offline: device offline online: device online
	Status string `json:"stat"`
	// some error message info of this device.
	// +optional
	ErrorMessage string `json:"err_msg,omitempty"`
	//Timestamp
	Timestamp int64 `json:"ts,omitempty"`
}

// devices status message
type DevicesStatusMessage struct {
	DevicesStatus []*DeviceStatusMessage `json:"devs_stat"`
}

//report event
type ReportEventMsg struct {
	DeviceID     string `json:"d_id"`
	ServiceName  string `json:"svc_n"`
	EventName    string `json:"event_n"`
	Details      string `json:"details,omitempty"`
	Timestamp    int64  `json:"ts,omitempty"`
	ErrorMessage string `json:"err_msg,omitempty"`
}

//device desired twin update message.
type DeviceDesiredTwinsUpdateMessage struct {
	DeviceID string `json:"d_id"`
	//desired twins we want to set. and it's just used for updateDesiredTwins API.
	//Optional:
	DesiredTwins []*TwinProperty `json:"desired_twins,omitempty"`
}

type ProtocolSpec struct {
	ID string `json:"-"`
	//Name: specifies the mapper name.
	Type string `json:"type"`

	//ID: specifies the mapper ID.
	Spec string `json:"spec"`

	DeviceStatus map[string]*DeviceStatusMessage `json:"-"`
}
