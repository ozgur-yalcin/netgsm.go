package main

import (
	NetGsm "./src/netgsm"
)

func main() {
	smsdata := NetGsm.SmsData{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := NetGsm.Sms(smsdata)
	if send {
		// your code
	}
}
