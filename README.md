# netgsm
An easy-to-use netgsm.com.tr API with golang

```go
package main

import (
	"netgsm/src"
)

func main() {
	smsdata := netgsm.SmsData{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := netgsm.Sms(smsdata)
	if send {
		// your code
	}
}
