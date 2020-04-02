package dvlirclient

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestDvLIRClient_LoginLogout covers:
	- Login
	- Logout
*/
func TestDvLIRClient_LoginLogout(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")
	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_GetDataFile covers:
	- Login
	- GetDataFile
	- Logout
*/
func TestDvLIRClient_GetDataFile(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	file, err := dvlirClient.GetDataFile(10)
	if !assert.NoError(t, err, "Error during GetDataFile request") {
		return
	}
	if !assert.IsType(t, DataLines{}, file, "Return value isn't of type DataLines") {
		return
	}
	if !assert.NotEmpty(t, file, "Data file is empty") {
		return
	}
	fmt.Println(file)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_GetMomentaryValue covers:
	- Login
	- GetMomentaryValue
	- Logout
*/
func TestDvLIRClient_GetMomentaryValues(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	value, err := dvlirClient.GetMomentaryValues()
	if !assert.NoError(t, err, "Error during GetMomentaryValues request") {
		return
	}
	if !assert.IsType(t, MomentaryValues{}, value, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, value, "Momentary data is empty") {
		return
	}

	fmt.Println(value)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_GetGeneralInformation covers:
	- Login
	- GetGeneralInformation
	- Logout
*/
func TestDvLIRClient_GetGeneralInformation(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	info, err := dvlirClient.GetGeneralInformation()
	if !assert.NoError(t, err, "Error during GetGeneralInformation request") {
		return
	}
	if !assert.IsType(t, GeneralInfo{}, info, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, info, "General information is empty") {
		return
	}

	fmt.Println(info)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_GetNetworkInformation covers:
	- Login
	- GetNetworkInformation
	- Logout
*/
func TestDvLIRClient_GetNetworkInformation(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	info, err := dvlirClient.GetNetworkInformation()
	if !assert.NoError(t, err, "Error during GetNetworkInformation request") {
		return
	}
	if !assert.IsType(t, NetworkInfo{}, info, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, info, "Network information is empty") {
		return
	}

	fmt.Println(info)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_GetSystemInformation covers:
	- Login
	- GetSystemInformation
	- Logout
*/
func TestDvLIRClient_GetSystemInformation(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	sysinfo, err := dvlirClient.GetSystemInformation()
	if !assert.NoError(t, err, "Error during GetSystemInformation request") {
		return
	}
	if !assert.IsType(t, SystemInfo{}, sysinfo, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, sysinfo, "System information is empty") {
		return
	}

	fmt.Println(sysinfo)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_Blink covers:
	- Login
	- Blink
	- Logout
*/
func TestDvLIRClient_Blink(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	num, err := dvlirClient.Blink(500, 10)
	if !assert.NoError(t, err, "Error during Blink request") {
		return
	}
	if !assert.Equal(t, 123, num, "Device didn't return correct return value") {
		return
	}

	fmt.Println(num)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_NTPServerTest covers:
	- Login
	- NTPServerTest
	- Logout
*/
func TestDvLIRClient_NTPServerTest(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	code, err := dvlirClient.NTPServerTest("de.pool.ntp.org")
	if !assert.NoError(t, err, "Error while testing NTP server") {
		return
	}
	if code == 2 {
		fmt.Println("Please wait at least 30 seconds before repeating this request")
	}
	if !assert.Equal(t, 1, code, "Device didn't return correct return value") {
		return
	}

	fmt.Println(code)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_ChangeNetworkSettings covers:
	- Login
	- ChangeNetworkSettings
	- Logout
*/
func TestDvLIRClient_ChangeNetworkSettings(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	res, err := dvlirClient.ChangeNetworkSettings("no", "", "", "", "", "", "", "")
	if !assert.NoError(t, err, "Error while while changing network settings") {
		return
	}

	fmt.Println(res)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_ChangeSavingInterval covers:
	- Login
	- ChangeSavingInterval
	- Logout
*/
func TestDvLIRClient_ChangeSavingInterval(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	res, err := dvlirClient.ChangeSavingInterval("min")
	if !assert.NoError(t, err, "Error while while changing saving interval") {
		return
	}

	fmt.Println(res)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_AllowResetWithPwd covers:
	- Login
	- AllowResetWithPwd
	- Logout
*/
func TestDvLIRClient_AllowResetWithPwd(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	res, err := dvlirClient.AllowResetWithPwd("Yes")
	if !assert.NoError(t, err, "Error while while changing saving interval") {
		return
	}

	fmt.Println(res)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_ResetAll covers:
	- Login
	- GetSystemInformation
	- ResetAll
	- Logout
*/
func TestDvLIRClient_ResetAll(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}
	code, err := dvlirClient.GetSystemInformation()
	if !assert.NoError(t, err, "Error during GetSystemInformation request") {
		return
	}
	if !assert.IsType(t, SystemInfo{}, code, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, code, "System information is empty") {
		return
	}

	res, err := dvlirClient.ResetAll(code.ResetCode)
	if !assert.NoError(t, err, "Error during ResetAll request") {
		return
	}

	fmt.Println(res)
	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_DeleteData covers:
	- Login
	- GetSystemInformation
	- DeleteData
	- Logout
*/
func TestDvLIRClient_DeleteData(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	code, err := dvlirClient.GetSystemInformation()
	if !assert.NoError(t, err, "Error during GetSystemInformation request") {
		return
	}
	if !assert.IsType(t, SystemInfo{}, code, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, code, "System information is empty") {
		return
	}

	res, err := dvlirClient.DeleteData(code.DeleteCode)
	if !assert.NoError(t, err, "Error during DeleteData request") {
		return
	}

	fmt.Println(res)
	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_ChangePassword covers:
	- Login
	- ChangePassword
	- Logout
*/
func TestDvLIRClient_ChangePassword(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	res, err := dvlirClient.ChangePassword("<Old password>", "<New password>", "<New password>")
	if !assert.NoError(t, err, "Error during ChangePassword request") {
		return
	}

	fmt.Println(res)
	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_UploadFirmware covers:
	- Login
	- UploadFirmware
	- Logout
*/
func TestDvLIRClient_UploadFirmware(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")
	file := viper.GetString("Firmware")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	res, err := dvlirClient.UploadFirmware(file)
	if !assert.NoError(t, err, "Error during UploadFirmware request") {
		return
	}

	fmt.Println(res)
	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_Restart covers:
	- Login
	- AllowResetWithPwd
	- Restart
	- Logout
*/
func TestDvLIRClient_Restart(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()
	if !assert.NoError(t, err, "Error during Login") {
		return
	}

	_, err = dvlirClient.AllowResetWithPwd("Yes")
	if !assert.NoError(t, err, "Error during AllowResetWithPwd request") {
		return
	}

	res, err := dvlirClient.Restart()
	if !assert.NoError(t, err, "Error during Restart request") {
		return
	}

	fmt.Println(res)
	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}

/*
TestDvLIRClient_Restart covers:
	- Login
	- GetDataFile
	- GetMomentaryValue
	- Logout
*/
func TestDvLIRClient_Setup(t *testing.T) {
	ip := viper.GetString("IPAddress")
	pw := viper.GetString("Password")

	dvlirClient, err := NewDvLIRClient(ip, pw)
	if !assert.NoError(t, err, "Error while creating Api client") {
		return
	}

	err = dvlirClient.Login()

	file, err := dvlirClient.GetDataFile(10)
	if !assert.NoError(t, err, "Error during GetDataFile request") {
		return
	}
	if !assert.IsType(t, DataLines{}, file, "Return value isn't of type DataLines") {
		return
	}
	if !assert.NotEmpty(t, file, "Data file is empty") {
		return
	}

	value, err := dvlirClient.GetMomentaryValues()
	if !assert.NoError(t, err, "Error during GetMomentaryValues request") {
		return
	}
	if !assert.IsType(t, MomentaryValues{}, value, "Return value isn't of type MomentaryValues") {
		return
	}
	if !assert.NotEmpty(t, value, "Momentary data is empty") {
		return
	}

	fmt.Println("Serial number electric meter: " + file[1].MeterNumber)
	fmt.Println("Serial number network adapter: " + file[1].DvLIRSn)
	fmt.Println("Meter reading: " + file[1].OneEightZero)
	fmt.Println("Adapter status: " + value.Status)

	defer func() {
		err = dvlirClient.Logout()
		if !assert.NoError(t, err, "Error during Logout") {
			return
		}
	}()
}
