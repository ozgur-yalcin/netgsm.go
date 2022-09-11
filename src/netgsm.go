package netgsm

import (
	"encoding/xml"
	"html"
	"net/http"
	"strings"
)

type API struct {
	Endpoint string
}

type Request struct {
	XMLName xml.Name `xml:"mainbody,omitempty"`
	Header  struct {
		Company   string `xml:"company,omitempty"`
		UserCode  string `xml:"usercode,omitempty"`
		Password  string `xml:"password,omitempty"`
		StartDate string `xml:"startdate,omitempty"`
		StopDate  string `xml:"stopdate,omitempty"`
		Type      string `xml:"type,omitempty"`
		MsgHeader string `xml:"msgheader,omitempty"`
	} `xml:"header,omitempty"`
	Body struct {
		Msg string `xml:"msg,omitempty"`
		No  string `xml:"no,omitempty"`
	} `xml:"body,omitempty"`
}

type Response struct {
	XMLName xml.Name `xml:"xml,omitempty"`
	Main    struct {
		Code  string `xml:"code,omitempty"`
		JobID string `xml:"jobID,omitempty"`
		Error string `xml:"error,omitempty"`
	} `xml:"main,omitempty"`
}

func Api(header, usercode, password string) (*API, *Request) {
	api := new(API)
	api.Endpoint = "https://api.netgsm.com.tr"
	request := new(Request)
	request.Header.Company = "NETGSM"
	request.Header.MsgHeader = header
	request.Header.UserCode = usercode
	request.Header.Password = password
	request.Header.Type = "1:n"
	return api, request
}

func (api *API) Sms(request *Request) bool {
	request.Body.Msg = "<![CDATA[" + request.Body.Msg + " - ]]>"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return false
	}
	response, err := http.Post(api.Endpoint+"/sms/send/xml", "text/xml", strings.NewReader(xml.Header+html.UnescapeString(string(postdata))))
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return true
}

func (api *API) Otp(request *Request) (res Response) {
	request.Body.Msg = "<![CDATA[" + request.Body.Msg + " - ]]>"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return res
	}
	response, err := http.Post(api.Endpoint+"/sms/send/otp", "text/xml", strings.NewReader(xml.Header+html.UnescapeString(string(postdata))))
	if err != nil {
		return res
	}
	defer response.Body.Close()
	decoder := xml.NewDecoder(response.Body)
	decoder.Decode(&res)
	return res
}
