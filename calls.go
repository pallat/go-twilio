package twilio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL      = "https://api.twilio.com/2010-04-01"
	makeCallPath = "/Accounts/{AccountSid}/Calls.json"
)

var (
	accountSID string
	authToken  string
)

func SetAccount(sid, token string) {
	accountSID = sid
	authToken = token
}

func CreateCall(from, to, hookURL string) (response CallResponse, err error) {
	apiURL := baseURL + strings.Replace(makeCallPath, "{AccountSid}", accountSID, 1)

	v := url.Values{}
	v.Set("To", to)
	v.Set("From", from)
	v.Set("Url", hookURL)

	rb := *strings.NewReader(v.Encode())

	c := http.Client{}
	req, err := http.NewRequest("POST", apiURL, &rb)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth(accountSID, authToken)

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return
	}

	return
}

type CallResponse struct {
	Sid             string          `json:"sid"`
	DateCreated     string          `json:"date_created"`
	DateUpdated     string          `json:"date_updated"`
	ParentCall_sid  string          `json:"parent_call_sid"`
	AccountSid      string          `json:"account_sid"`
	To              string          `json:"to"`
	ToFormatted     string          `json:"to_formatted"`
	From            string          `json:"from"`
	FromFormatted   string          `json:"from_formatted"`
	PhoneNumberSid  string          `json:"phone_number_sid"`
	Status          string          `json:"status"`
	StartTime       string          `json:"start_time"`
	EndTime         string          `json:"end_time"`
	Duration        string          `json:"duration"`
	Price           string          `json:"price"`
	PriceUnit       string          `json:"price_unit"`
	Direction       string          `json:"direction"`
	AnsweredBy      string          `json:"answered_by"`
	APIVersion      string          `json:"api_version"`
	Annotation      string          `json:"annotation"`
	ForwardedFrom   string          `json:"forwarded_from"`
	GroupSid        string          `json:"group_sid"`
	CallerName      string          `json:"caller_name"`
	URI             string          `json:"uri"`
	SubresourceUris SubresourceURIS `json:"subresource_uris"`
}

type SubresourceURIS struct {
	Notifications string `json:"notifications"`
	Recordings    string `json:"recordings"`
}
