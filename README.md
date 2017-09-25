# netgsm-go
An easy-to-use netgsm.com.tr API with golang

```go
package main

import (
	"fmt"
	NetGsm "./src/netgsm"
)

func main() {
	smsdata := NetGsm.SmsRequest{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := NetGsm.Sms(smsdata)
	if send {
		fmt.Println("ok")
	}
}
