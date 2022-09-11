package netgsm

import (
	"encoding/xml"
	"html"
	"net/http"
	"strings"
)

type Config struct {
	ApiUrl       string
	SmsCompany   string
	SmsMsgHeader string
	SmsUserCode  string
	SmsPassword  string
}

type API struct {
	Config Config
}

type Request struct {
	XMLName  xml.Name `xml:"xml,omitempty"`
	MainBody struct {
		Header struct {
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
	} `xml:"mainbody,omitempty"`
}

type Response struct {
	XMLName xml.Name `xml:"xml,omitempty"`
	XML     struct {
		Main struct {
			Code  string `xml:"code,omitempty"`
			JobID string `xml:"jobID,omitempty"`
		} `xml:"main,omitempty"`
	} `xml:"xml,omitempty"`
}

func (api *API) Sms(request Request) (res Response) {
	request.MainBody.Header.Company = api.Config.SmsCompany
	request.MainBody.Header.MsgHeader = api.Config.SmsMsgHeader
	request.MainBody.Header.UserCode = api.Config.SmsUserCode
	request.MainBody.Header.Password = api.Config.SmsPassword
	request.MainBody.Header.Type = "1:n"
	request.MainBody.Body.Msg = "<![CDATA[" + request.MainBody.Body.Msg + " - ]]>"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return res
	}
	response, err := http.Post(api.Config.ApiUrl, "text/xml; charset=utf-8", strings.NewReader(xml.Header+html.UnescapeString(string(postdata))))
	if err != nil {
		return res
	}
	defer response.Body.Close()
	decoder := xml.NewDecoder(response.Body)
	decoder.Decode(&res)
	return res
}
