package netgsm

import (
	"encoding/xml"
	"net/http"
	"strings"
	"time"
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

func (api API) Sms(request Request) bool {
	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		return false
	}
	request.MainBody.Header.Company = api.Config.SmsCompany
	request.MainBody.Header.MsgHeader = api.Config.SmsMsgHeader
	request.MainBody.Header.UserCode = api.Config.SmsUserCode
	request.MainBody.Header.Password = api.Config.SmsPassword
	request.MainBody.Header.Type = "1:n"
	request.MainBody.Header.StartDate = time.Now().In(loc).Format("020120061504")
	request.MainBody.Header.StopDate = time.Now().In(loc).Add(24 * time.Hour).Format("020120061504")
	request.MainBody.Body.Msg = "<![CDATA[" + request.MainBody.Body.Msg + " - ]]>"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return false
	}
	rpl := strings.NewReplacer("&lt;!", "<!", "]&gt;", "]>", "<xml>", "", "</xml>", "")
	res, err := http.Post(api.Config.ApiUrl, "text/xml; charset=utf-8", strings.NewReader(xml.Header+rpl.Replace(string(postdata))))
	if err != nil {
		return false
	}
	defer res.Body.Close()
	return true
}
