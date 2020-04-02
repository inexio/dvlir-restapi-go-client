package dvlirclient

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
)

type client struct {
	*clientData
}

/*
clientData - Contains data of a client
*/
type clientData struct {
	ipAddress string
	password  string
	sessionID string
	resty     *resty.Client
}

/*
NotValidError - Is returned when the client was not initialized properly
*/
type NotValidError struct{}

func (m *NotValidError) Error() string {
	return "client was not created properly with the func NewDvLIRClient(baseUrl string)"
}

/*
isValid - returns true if a client is valid and false if a client is invalid
*/
func (c *client) isValid() bool {
	return c.clientData != nil
}

func (d *DvLIRClient) get(path string, body string, pwd string) (*resty.Response, error) {
	request := d.resty.R()
	if body != "" {
		request.SetBody(body)
	}
	response, err := request.Get("http://" + d.ipAddress + path + pwd)
	if err != nil {
		return nil, errors.Wrap(err, "error during http request")
	}

	if response.StatusCode() != 200 {
		return nil, errors.Wrap(getHTTPError(response), "http status code != 200")
	}

	return response, nil
}

func (d *DvLIRClient) post(path string, body string, header, queryParams map[string]string, file bool) (*resty.Response, error) {
	request := d.resty.R()
	request.SetHeader("Content-Type", "application/json")

	if header != nil {
		request.SetHeaders(header)
	}

	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	if file {
		request.SetFile("firmware", body)
	} else if body != "" {
		request.SetBody(body)
	}

	var response *resty.Response
	var err error
	response, err = request.Post(path)
	if err != nil {
		return nil, errors.Wrap(err, "error during http request")
	}
	return response, nil
}

//Http error handling

/*
HTTPError - Represents an http error returned by the api.
*/
type HTTPError struct {
	StatusCode int
	Status     string
	Body       *ErrorResponse
}

/*
ErrorResponse - Contains error information.
*/
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (h HTTPError) Error() string {
	msg := "http error: status code: " + strconv.Itoa(h.StatusCode) + " // status: " + h.Status
	if h.Body != nil {
		msg += " // message: " + h.Body.Message
	}
	return msg
}

func getHTTPError(response *resty.Response) error {
	httpError := HTTPError{
		StatusCode: response.StatusCode(),
		Status:     response.Status(),
	}
	var errorResponse ErrorResponse
	err := json.Unmarshal(response.Body(), &errorResponse)
	if err != nil {
		return httpError
	}
	httpError.Body = &errorResponse
	return httpError
}

/*
init is used for initialising the config file
*/
func init() {
	// Search config in home directory with name "eve-ng-api" (without extension).
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("dvlir-api")

	//Set env var prefix to only match certain vars
	viper.SetEnvPrefix("DvLIR_API")

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Couldn't read config")
	}
}

func switchCase(e string) bool {
	switch e {
	case "Yes", "yes", "No", "no":
		return true
	default:
		return false
	}
}

/*
DataLineConversion converts a string into a DataLine struct
*/
func (d *DvLIRClient) DataLineConversion(input string) DataLine {
	var line DataLine

	inputSlice := strings.Split(input, ";")

	line.Index = inputSlice[0]
	line.Date = inputSlice[1]
	line.Time = inputSlice[2]
	line.DvLIRSn = inputSlice[3]
	line.MeterNumber = inputSlice[4]
	line.OneEightZero = inputSlice[5]
	line.OneEightOne = inputSlice[6]
	line.OneEightTwo = inputSlice[7]
	line.TwoEightZero = inputSlice[8]
	line.TwoEightOne = inputSlice[9]
	line.TwoEightTwo = inputSlice[10]
	line.Power = inputSlice[11]
	line.Status = inputSlice[12]

	return line
}

/*
LineSplitter splits a string at every instance of \r\n
*/
func (d *DvLIRClient) LineSplitter(input string) []string {
	var lines []string
	lines = strings.Split(input, "\r\n")
	return lines
}

/*
HashTagSplitter splits a string at every instance of a #
*/
func (d *DvLIRClient) HashTagSplitter(input string) []string {
	var info []string
	info = strings.Split(input, "#")
	return info
}
