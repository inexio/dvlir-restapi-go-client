package dvlirclient

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

/*
DvLIRClient is an implementation of the client specified for dvlir
*/
type DvLIRClient struct {
	client
}

/*
NewDvLIRClient generates a new dvlir api-client object which can be used to communicate with the DvLIR-API
*/
func NewDvLIRClient(ipAddress string, password string) (*DvLIRClient, error) {
	if ipAddress == "" || password == "" {
		return nil, errors.New("invalid IP address or invalid password")
	}

	clientData := clientData{ipAddress: ipAddress, password: password, resty: resty.New()}
	newClient := client{&clientData}
	return &DvLIRClient{newClient}, nil
}

/*
Login performs a login via the dvlir api-client
*/
func (d *DvLIRClient) Login() error {
	if !d.isValid() {
		return &NotValidError{}
	}

	res, err := d.get("/getSID.txt?pwd=", "", d.password)
	if err != nil {
		return errors.Wrap(err, "Error during login request")
	}
	if res.String() == "" {
		err = errors.New("Error during login request")
		return err
	}
	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return err
	}

	d.sessionId = res.String()

	return err
}

/*
Logout performs a logout via the dvlir api-client
*/
func (d *DvLIRClient) Logout() error {
	if !d.isValid() {
		return &NotValidError{}
	}
	_, err := d.get("/getSID.txt", "", "")
	if err != nil {
		return errors.Wrap(err, "Error during logout")
	}

	return err
}

/*
GetDataFile a performs GetDataFile operation via the dvlir api-client
*/
func (d *DvLIRClient) GetDataFile(lines int) (DataLines, error) {
	var empty DataLines
	var splitted []string
	var fileLines DataLines
	var fileLine DataLine
	if !d.isValid() {
		return empty, &NotValidError{}
	}

	var err error
	if lines < 1 || lines > 14400 {
		err = errors.New("Line number has to be either <0 or >14400")
		return empty, errors.Wrap(err, "Invalid line number: "+strconv.Itoa(lines))
	}

	linesE := url.QueryEscape(strconv.Itoa(lines))

	path := "/daten.csv?sid=" + d.sessionId + "&lines=" + linesE
	f, err := d.get(path, "", "")
	if err != nil {
		return empty, errors.Wrap(err, "Error during GetDataFile request")
	}

	check := f.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return empty, err
	}

	splitted = d.LineSplitter(f.String())
	for i := 0; i < len(splitted); i++ {
		fileLine = d.DataLineConversion(splitted[i])
		fileLines = append(fileLines, fileLine)
	}

	return fileLines, err
}

/*
GetMomentaryValues a performs GetMomentaryValue operation via the dvlir api-client
*/
func (d *DvLIRClient) GetMomentaryValues() (MomentaryValues, error) {
	var empty MomentaryValues
	var values MomentaryValues
	if !d.isValid() {
		return empty, &NotValidError{}
	}

	path := "/data.txt?sid=" + d.sessionId
	v, err := d.get(path, "", "")
	if err != nil {
		return empty, errors.Wrap(err, "Error during GetMomentaryValues request")
	}

	check := v.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return empty, err
	}

	val := d.HashTagSplitter(v.String())

	values.MeterNumber = val[0]
	values.OBISNum = val[1]
	values.MomentaryPower = val[2]
	values.MeterReadingAP = val[3]
	values.MeterReadingAM = val[4]
	for i := 5; i <= 13; i++ {
		values.MeterReadingsAP[i-5] = val[i]
	}
	for j := 14; j <= 22; j++ {
		values.MeterReadingsAM[j-14] = val[j]
	}
	values.Status = val[23]
	values.SavingInterval = val[24]

	return values, err
}

/*
GetGeneralInformation a performs GetGeneralInformation operation via the dvlir api-client
*/
func (d *DvLIRClient) GetGeneralInformation() (GeneralInfo, error) {
	var empty GeneralInfo
	var info GeneralInfo
	if !d.isValid() {
		return empty, &NotValidError{}
	}

	path := "/info.txt?sid=" + d.sessionId
	i, err := d.get(path, "", "")
	if err != nil {
		return empty, errors.Wrap(err, "Error during GetGeneralInformation request")
	}

	check := i.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return empty, err
	}

	information := d.HashTagSplitter(i.String())
	info.ServerIDMeter = information[0]
	info.MeterNumber = information[1]
	info.ManufacturerCode = information[2]
	info.IPAddress = information[3]
	info.Gateway = information[4]
	info.DNSServer = information[5]
	info.NetworkName = information[6]
	info.MACAddress = information[7]
	info.SavingInterval = information[8]
	info.Date = information[9]
	info.Time = information[10]
	info.DeviceSn = information[11]
	info.FirmwareVersion = information[12]

	return info, err
}

/*
GetNetworkInformation a performs GetNetworkInformation operation via the dvlir api-client
*/
func (d *DvLIRClient) GetNetworkInformation() (NetworkInfo, error) {
	var empty NetworkInfo
	var net NetworkInfo
	if !d.isValid() {
		return empty, &NotValidError{}
	}

	path := "/network.txt?sid=" + d.sessionId
	n, err := d.get(path, "", "")
	if err != nil {
		return empty, errors.Wrap(err, "Error during GetGeneralInformation request")
	}

	check := n.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return empty, err
	}

	info := d.HashTagSplitter(n.String())
	net.DHCPServer = info[0]
	net.IPAddress = info[1]
	net.SubnetMask = info[2]
	net.Gateway = info[3]
	net.DNSServer = info[4]
	net.NTPServer = info[5]
	net.NTPName = info[6]

	return net, err
}

/*
GetSystemInformation a performs GetSystemInformation operation via the dvlir api-client
*/
func (d *DvLIRClient) GetSystemInformation() (SystemInfo, error) {
	var empty SystemInfo
	var system SystemInfo
	if !d.isValid() {
		return empty, &NotValidError{}
	}

	path := "/system.txt?sid=" + d.sessionId
	s, err := d.get(path, "", "")
	if err != nil {
		return empty, errors.Wrap(err, "Error during GetSystemInformation request")
	}

	check := s.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return empty, err
	}

	sys := d.HashTagSplitter(s.String())
	system.SavingInterval = sys[0]
	system.ResetCode = sys[1]
	system.DeleteCode = sys[2]
	system.ResetWithDefaultPwd = sys[3]

	return system, err
}

/*
Blink a performs Blink operation via the dvlir api-client
*/
func (d *DvLIRClient) Blink(blink int, pause int) (response int, err error) {
	if !d.isValid() {
		return 0, &NotValidError{}
	}
	if blink < 1 || blink > 10000 {
		err = errors.New("Value of blink is outside of 1..10000")
		return 0, err
	}
	if pause < 1 || pause > 1000 {
		err = errors.New("Value of pause is outside of 1..1000")
		return 0, err
	}

	pauseE := url.QueryEscape(strconv.Itoa(pause))
	blinkE := url.QueryEscape(strconv.Itoa(blink))

	path := "/blink.cmd?sid=" + d.sessionId + "&ledPause=" + pauseE + "&ledBlink=" + blinkE
	resp, err := d.get(path, "", "")
	if err != nil {
		return 0, errors.Wrap(err, "Error during Blink request")
	}
	res, err := strconv.Atoi(resp.String())
	if err != nil {
		return 0, errors.Wrap(err, "Error during conversion of response code from string to integer")
	}

	if len(resp.String()) > 3 {
		check := resp.String()[0:9]
		if check == "<!DOCTYPE" {
			err = errors.New("Login page was returned")
			return 0, err
		}
	}

	return res, err
}

/*
NTPServerTest a performs NTPServerTest operation via the dvlir api-client
*/
func (d *DvLIRClient) NTPServerTest(ntpName string) (int, error) {
	if !d.isValid() {
		return 0, &NotValidError{}
	}

	ntpNameE := url.QueryEscape(ntpName)

	path := "/ntpTest.cmd?sid=" + d.sessionId + "&ntpName=" + ntpNameE
	c, err := d.get(path, "", "")
	if err != nil {
		return 0, errors.Wrap(err, "Error during Blink request")
	}
	code, err := strconv.Atoi(c.String())
	if err != nil {
		return 0, errors.Wrap(err, "Error during conversion of response code")
	}

	switch code {
	case 0:
		err = errors.New("NTP-server can't be reached")
		return 0, err
	case 1:
		break
	case 2:
		err = errors.New("You have to wait at least 30 seconds until doing this again")
		return 2, err
	case 3:
		err = errors.New("NTP-servername is invalid")
		return 3, err
	default:
		err = errors.New("Error during NTPServerTest request")
		return 0, err
	}

	if len(c.String()) > 3 {
		check := c.String()[0:9]
		if check == "<!DOCTYPE" {
			err = errors.New("Login page was returned")
			return 0, err
		}
	}

	return code, err
}

/*
ChangeNetworkSettings a performs ChangeNetworkSettings operation via the dvlir api-client
*/
func (d *DvLIRClient) ChangeNetworkSettings(dhcp, ip, sub, gw, dns, ntpName, ntpServer, setDt string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}
	path := "/network.cmd?sid=" + d.sessionId

	if dhcp != "" {
		if switchCase(dhcp) {
			dhcpE := url.QueryEscape(dhcp)
			path += "&dhcpServer=" + dhcpE
		} else {
			err := errors.New("Invalid input for dhcp (can only be 'Yes', 'yes', 'No', 'no')")
			return "", err
		}
	}
	if ip != "" {
		ipE := url.QueryEscape(ip)
		path += "&ip=" + ipE
	}
	if sub != "" {
		subE := url.QueryEscape(sub)
		path += "&sub=" + subE
	}
	if gw != "" {
		gwE := url.QueryEscape(gw)
		path += "&gw=" + gwE
	}
	if dns != "" {
		dnsE := url.QueryEscape(dns)
		path += "&dns=" + dnsE
	}
	if ntpName != "" {
		ntpNameE := url.QueryEscape(ntpName)
		path += "&ntpName=" + ntpNameE
	}
	if ntpServer != "" {
		if switchCase(ntpServer) {
			ntpServerE := url.QueryEscape(ntpServer)
			path += "&ntpServer=" + ntpServerE
		} else {
			err := errors.New("Invalid input for ntpServer (can only be 'Yes', 'yes', 'No', 'no')")
			return "", err
		}
	}
	if setDt != "" {
		setDtE := url.QueryEscape(setDt)
		path += "&setDt=" + setDtE
	}

	res, err := d.get(path, "", "")
	if err != nil {
		return "", errors.Wrap(err, "Error during ChangeNetworkSettings request")
	}

	if res.String() == "cmd=" {
		err = errors.New("Either no parameters were given or an unexpected error occurred")
		return "", errors.Wrap(err, "An error was returned")
	}

	if len(res.String()) > 4 {
		check := res.String()[0:9]
		if check == "<!DOCTYPE" {
			err = errors.New("Login page was returned")
			return "", err
		}
	}

	return res.String(), err
}

/*
ChangeSavingInterval a performs ChangeSavingInterval operation via the dvlir api-client
*/
func (d *DvLIRClient) ChangeSavingInterval(interval string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	if interval != "15min" && interval != "min" && interval != "sec" {
		err := errors.New("Invalid saving interval (can only be '15min', 'min', 'sec')")
		return "", err
	}

	intervalE := url.QueryEscape(interval)

	path := "/system.cmd?sid=" + d.sessionId + "&interval=" + intervalE
	res, err := d.get(path, "", "")
	if err != nil {
		return "", errors.Wrap(err, "Error during ChangeSavingInterval request")
	}
	if res.String() == "cmd=" {
		err = errors.New("Invalid interval")
		return "", err
	}

	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return "", err
	}

	return res.String(), err
}

/*
AllowResetWithPwd a performs AllowResetWithPwd operation via the dvlir api-client
*/
func (d *DvLIRClient) AllowResetWithPwd(allow string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	if allow != "Yes" && allow != "yes" && allow != "No" && allow != "no" {
		err := errors.New("Invalid input for allow (can only be 'Yes', 'yes', 'No', 'no')")
		return "", err
	}

	allowE := url.QueryEscape(allow)

	path := "/system.cmd?sid=" + d.sessionId + "&allowResetWithPwd=" + allowE
	res, err := d.get(path, "", "")
	if err != nil {
		return "", errors.Wrap(err, "Error during AllowResetWithPwd request")
	}
	if res.String() == "cmd=" {
		err = errors.New("Invalid input")
		return "", err
	}

	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return "", err
	}

	return res.String(), err
}

/*
ResetAll a performs ResetAll operation via the dvlir api-client
*/
func (d *DvLIRClient) ResetAll(rCode string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	rCodeE := url.QueryEscape(rCode)

	path := "/system.cmd?sid=" + d.sessionId + "&resetAll=" + rCodeE
	res, err := d.get(path, "", "")
	if err != nil {
		return "", errors.Wrap(err, "Error during ResetAll request")
	}
	if res.String() == "cmd=" {
		err = errors.New("Invalid input")
		return "", err
	}

	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return "", err
	}

	return res.String(), err
}

/*
DeleteData a performs DeleteData operation via the dvlir api-client
*/
func (d *DvLIRClient) DeleteData(code string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	dCodeE := url.QueryEscape(code)

	path := "/system.cmd?sid=" + d.sessionId + "&resetData=" + dCodeE
	res, err := d.get(path, "", "")
	if err != nil {
		return "", errors.Wrap(err, "Error during DeleteData request")
	}
	if res.String() == "cmd=" {
		err = errors.New("Invalid input")
		return "", err
	}

	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return "", err
	}

	return res.String(), err
}

/*
ChangePassword a performs ChangePassword operation via the dvlir api-client
*/
func (d *DvLIRClient) ChangePassword(pw1, pw2, pw3 string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}
	path := "http://" + d.ipAddress + "/password.cmd?sid=" + d.sessionId

	Pw1 := url.QueryEscape(pw1)
	Pw2 := url.QueryEscape(pw2)
	Pw3 := url.QueryEscape(pw3)

	res, err := http.PostForm(path, url.Values{"pw1": {Pw1}, "pw2": {Pw2}, "pw3": {Pw3}})
	if err != nil {
		return "", errors.Wrap(err, "Error during ChangePassword request")
	}

	var bodyString string
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			err = errors.New("Error while converting http response to string")
			return "", err
		}
		bodyString = string(bodyBytes)
	}

	if bodyString == "2" {
		err = errors.New("Current password is wrong")
		return "", err
	} else if bodyString == "3" {
		err = errors.New("New password does not match confirm new password")
		return "", err
	} else if bodyString == "4" {
		err = errors.New("Illegal character in new password")
	}

	if len(bodyString) > 3 {
		check := bodyString[0:9]
		if check == "<!DOCTYPE" {
			err := errors.New("Login page was returned")
			return "", err
		}
	}

	return bodyString, err

}

/*
UploadFirmware a performs UploadFirmware operation via the dvlir api-client
*/
func (d *DvLIRClient) UploadFirmware(filePath string) (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	path := "http://" + d.ipAddress + "/upload.cmd?sid=" + d.sessionId

	header := make(map[string]string)
	header["Content-Type"] = "multipart/form-data"

	filePathE := url.QueryEscape(filePath)

	resp, err := d.post(path, filePathE, header, nil, true)

	if err != nil {
		return "", errors.Wrap(err, "Error during UploadFirmware")
	}
	if resp.String() == "2" || resp.String() == "3" || resp.String() == "4" || resp.String() == "5" {
		err = errors.New("The following error occurred" + resp.String())
		return "", err
	}

	if len(resp.String()) > 3 {
		check := resp.String()[0:9]
		if check == "<!DOCTYPE" {
			err = errors.New("Login page was returned")
			return "", err
		}
	}

	return resp.String(), err
}

/*
Restart a performs Restart operation via the dvlir api-client
*/
func (d *DvLIRClient) Restart() (string, error) {
	if !d.isValid() {
		return "", &NotValidError{}
	}

	path := "/doReset.cmd?pwd="
	res, err := d.get(path, "", d.password)
	if err != nil {
		return "", err
	}
	if res.String() == "cmd=" {
		err = errors.New("Error during Restart")
		return "", err
	}

	check := res.String()[0:9]
	if check == "<!DOCTYPE" {
		err = errors.New("Login page was returned")
		return "", err
	}

	return res.String(), err
}
