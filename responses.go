package dvlirclient

/*
DataLines is an array of DataLines.
*/
type DataLines []DataLine

/*
DataLine contains the contents of a single line of the daten.csv file
*/
type DataLine struct {
	Index        string `json:"index"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	DvLIRSn      string `json:"dvlir_sn"`
	MeterNumber  string `json:"meter_number"`
	OneEightZero string `json:"one_eight_zero"`
	OneEightOne  string `json:"one_eight_one"`
	OneEightTwo  string `json:"one_eight_two"`
	TwoEightZero string `json:"two_eight_zero"`
	TwoEightOne  string `json:"two_eight_one"`
	TwoEightTwo  string `json:"two_eight_two"`
	Power        string `json:"power"`
	Status       string `json:"status"`
}

/*
MomentaryValues contains the response of the api in case of a GetMomentaryValues request
*/
type MomentaryValues struct {
	MeterNumber     string    `json:"meter_number"`
	OBISNum         string    `json:"obis_num"`
	MomentaryPower  string    `json:"momentary_power"`
	MeterReadingAP  string    `json:"meter_reading_ap"`
	MeterReadingAM  string    `json:"meter_reading_am"`
	MeterReadingsAP [9]string `json:"meter_readings_ap"`
	MeterReadingsAM [9]string `json:"meter_readings_am"`
	Status          string    `json:"status"`
	SavingInterval  string    `json:"saving_interval"`
}

/*
GeneralInfo contains the response of the api in case of a GetGeneralInformation request
*/
type GeneralInfo struct {
	ServerIDMeter    string `json:"server_id_meter"`
	MeterNumber      string `json:"meter_number"`
	ManufacturerCode string `json:"manufacturer_code"`
	IPAddress        string `json:"ip_address"`
	Gateway          string `json:"gateway"`
	DNSServer        string `json:"dns_server"`
	NetworkName      string `json:"network_name"`
	MACAddress       string `json:"mac_address"`
	SavingInterval   string `json:"saving_interval"`
	Date             string `json:"date"`
	Time             string `json:"time"`
	DeviceSn         string `json:"device_sn"`
	FirmwareVersion  string `json:"firmware_version"`
}

/*
NetworkInfo contains the response of the api in case of a GetNetworkInformation request
*/
type NetworkInfo struct {
	DHCPServer string `json:"dhcp_server"`
	IPAddress  string `json:"ip_address"`
	SubnetMask string `json:"subnet_mask"`
	Gateway    string `json:"gateway"`
	DNSServer  string `json:"dns_server"`
	NTPServer  string `json:"ntp_server"`
	NTPName    string `json:"ntp_name"`
}

/*
SystemInfo contains the response of the api in case of a GetSystemInformation request
*/
type SystemInfo struct {
	SavingInterval      string `json:"saving_interval"`
	ResetCode           string `json:"reset_code"`
	DeleteCode          string `json:"delete_code"`
	ResetWithDefaultPwd string `json:"reset_with_default_pwd"`
}
