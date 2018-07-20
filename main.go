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
	api := new(netgsm.API)
	api.Lock()
	defer api.Unlock()
	request := new(netgsm.Request)
	request.MainBody.Body.Msg = "test"
	request.MainBody.Body.No = "905555555555"
	send := api.Sms(request)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
