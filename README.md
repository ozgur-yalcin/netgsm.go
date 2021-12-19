[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-soft/netgsm.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-soft/netgsm.go)](https://pkg.go.dev/github.com/ozgur-soft/netgsm.go/src)

# Netgsm.go
An easy-to-use netgsm.com.tr API with golang

# Installation
```bash
go get github.com/ozgur-soft/netgsm.go
```

# Usage
```go
package main

import (
	"fmt"

	netgsm "github.com/ozgur-soft/netgsm.go/src"
)

func main() {
	config := netgsm.Config{SmsCompany: "NETGSM", SmsMsgHeader: "", SmsUserCode: "", SmsPassword: "", ApiUrl: "https://api.netgsm.com.tr/sms/send/xml"}
	api := &netgsm.API{config}
	request := netgsm.Request{}
	request.MainBody.Body.Msg = "test"
	request.MainBody.Body.No = "905555555555"
	send := api.Sms(request)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
```
