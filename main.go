package main

import (
	"netgsm/config"
	"netgsm/src"
)

func init() {
	config.SmsCompany = ""   // Firma Adı
	config.SmsMsgHeader = "" // Mesaj başlığı
	config.SmsUserCode = ""  // Kullanıcı Kodu
	config.SmsPassword = ""  // Kullanıcı Şifresi
}

func main() {
	smsdata := netgsm.SmsData{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := netgsm.Sms(smsdata)
	if send {
		// your code
	}
}
