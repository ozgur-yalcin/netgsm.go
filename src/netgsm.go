package netgsm

import (
	"encoding/xml"
	"net/http"
	"netgsm/config"
	"strings"
	"time"
)

type SmsData struct {
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

func Sms(xmlrequest SmsData) bool {
	loc, _ := time.LoadLocation("Europe/Istanbul")
	xmlrequest.MainBody.Header.Company = config.SmsCompany
	xmlrequest.MainBody.Header.MsgHeader = config.SmsMsgHeader
	xmlrequest.MainBody.Header.UserCode = config.SmsUserCode
	xmlrequest.MainBody.Header.Password = config.SmsPassword
	xmlrequest.MainBody.Header.Type = "1:n"
	xmlrequest.MainBody.Header.StartDate = time.Now().In(loc).Format("020120061504")
	xmlrequest.MainBody.Header.StopDate = time.Now().In(loc).Add(24 * time.Hour).Format("020120061504")
	xmlrequest.MainBody.Body.Msg = "<![CDATA[" + xmlrequest.MainBody.Body.Msg + " - ]]>"
	postdata, _ := xml.Marshal(xmlrequest)
	repl := strings.NewReplacer("&lt;!", "<!", "]&gt;", "]>", "<xml>", "", "</xml>", "")
	resp, err := http.Post(config.APIURL, "text/xml; charset=utf-8", strings.NewReader(xml.Header+repl.Replace(string(postdata))))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}
