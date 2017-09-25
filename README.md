# netgsm-go
An easy-to-use netgsm.com.tr API with golang

`package main

import "fmt"

func main() {
	smsdata := src.SmsData{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := src.Sms(smsdata)
	if send {
		fmt.Println("ok")
	}
}
`
