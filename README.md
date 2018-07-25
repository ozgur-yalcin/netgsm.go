[![Linux Build Status](https://travis-ci.org/OzqurYalcin/netgsm.svg?branch=master)](https://travis-ci.org/OzqurYalcin/netgsm) [![Windows Build Status](https://ci.appveyor.com/api/projects/status/sm4un9iwhqg9sdna?svg=true)](https://ci.appveyor.com/project/OzqurYalcin/netgsm) [![Build Status](https://circleci.com/gh/OzqurYalcin/netgsm.svg?style=svg)](https://circleci.com/gh/OzqurYalcin/netgsm)

# Netgsm
An easy-to-use netgsm.com.tr API with golang

# Security
If you discover any security related issues, please email ozguryalcin@outlook.com instead of using the issue tracker.

# License
The MIT License (MIT). Please see License File for more information.

# Installation
```bash
go get github.com/OzqurYalcin/netgsm
```

# Usage
```go
package main

import (
	"fmt"

	"github.com/OzqurYalcin/netgsm/config"
	"github.com/OzqurYalcin/netgsm/src"
)

func init() {
	config.SmsCompany = "NETGSM"
	config.SmsMsgHeader = "" // Başlık
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
```
