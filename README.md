# dvlir-restapi-go-client

## Description

Golang package - client library for the [Device GmbH DvLIR IP network readout adapter](https://www.device.de/index.php/produkte/smart-metering/dvlir) REST API.
This go client is an open-source library to communicate with the Device GmbH DvLIR IP network readout adapter REST API.

## Code Style

This project was written according to the **[uber-go](https://github.com/uber-go/guide/blob/master/style.md)** coding style.

## Features

- Retrieve a csv file with the latest data
- Retrieve the momentary data
- Retrieve general information about the device (including information about the connected electric meter, network information, system informatrion)
- Retrieve and change network information (including if a dhcp server is used, the IP address, subnet mask, gateway and dns-server of the adapter, if a ntp-server is used and the name of the ntp-server)
- Retrieve and change system information (including the saving interval, the reset safety code, the deletion safety code and if a reset with the default password is allowed)
- Reset the adapter to factory settings
- Deleting all saved data
- Upload a firmware file
- Restart the adapter

## Installation

```
go get github.com/inexio/dvlir-restapi-go-client
```

or

```
git clone https://github.com/inexio/dvlir-restapi-go-client.git
```

## Usage

```go
    //Create a new DvLIR api client
    dvlirClient, err := NewDvLIRClient(ip, pw)
    
    //Login into the adapter to get a valid session id
    err = dvlirClient.Login()
    
    //Retrieve the data file
    file, err := dvlirClient.GetDataFile(10)
    
    //Logout of the adapter again
    err = dvlirClient.Logout()
```

### Tests

Our library provides a few unit and intergration tests. To use these tests, the yaml config file in the config directory must be adapted to your setup.

In order to run a test, run the following command inside of the root directory of this repository:

```
go test --run <Name of the test you want to run>
```

If you want to check if your setup works, run:

```
go test --run TestDvLIRClient_Setup
```

It must be noted, that you have to run the tests one at a time, because quite a few of the tests cause the adapter to restart, which would cause the other tests to fail.

If you want to upload a firmware file to your adapter, the firmware file needs to be in the project folder.

## Getting Help

If there are any problems or something does not work as intended, open an issue on GitHub.

## Contribution

Contributions to the project are welcome.

We are looking forward to your bug reports, suggestions and fixes.

If you want to make any contributions make sure your go report does match up with our projects score of **A+**.

When you contribute make sure your code is conform to the **uber-go** coding style.

Happy Coding!

