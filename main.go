package main

import (
	"fmt"
	"netgsm/config"
	"netgsm/src"
)

func init() {
	config.SmsCompany = "NETGSM"
	config.SmsMsgHeader = "" // Mesaj başlığı
	config.SmsUserCode = ""  // Kullanıcı Adı
	config.SmsPassword = ""  // Şifre
}

func main() {
	smsdata := new(netgsm.SmsData)
	smsdata.Lock()
	defer smsdata.Lock()
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := netgsm.Sms(smsdata)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
