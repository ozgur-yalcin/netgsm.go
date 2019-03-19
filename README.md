[![Build Status](https://travis-ci.org/OzqurYalcin/netgsm.svg?branch=master)](https://travis-ci.org/OzqurYalcin/netgsm) [![Build Status](https://circleci.com/gh/OzqurYalcin/netgsm.svg?style=svg)](https://circleci.com/gh/OzqurYalcin/netgsm) [![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/OzqurYalcin/netgsm/blob/master/LICENSE.md)

# Netgsm
An easy-to-use netgsm.com.tr API with golang

# Installation
```bash
go get github.com/OzqurYalcin/netgsm
```

# Usage
```go
package main

import (
	"fmt"

	netgsm "github.com/OzqurYalcin/netgsm/src"
)

func main() {
	config := netgsm.Config{SmsCompany: "NETGSM", SmsMsgHeader: "", SmsUserCode: "", SmsPassword: "", ApiUrl: "https://api.netgsm.com.tr/sms/send/xml"}
	api := &netgsm.API{config}
	request := &netgsm.Request{}
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
