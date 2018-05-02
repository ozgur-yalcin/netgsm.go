# netgsm
An easy-to-use netgsm.com.tr API with golang

```go
package main

import (
	"fmt"
	"netgsm/config"
	"netgsm/src"
)

func init() {
	config.SmsCompany = ""   // Firma Adı
	config.SmsMsgHeader = "" // Mesaj başlığı
	config.SmsUserCode = ""  // Kullanıcı Adı
	config.SmsPassword = ""  // Şifre
}

func main() {
	smsdata := netgsm.SmsData{}
	smsdata.MainBody.Body.Msg = "test"
	smsdata.MainBody.Body.No = "905555555555"
	send := netgsm.Sms(smsdata)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
```

# Security
If you discover any security related issues, please email ozguryalcin@outlook.com instead of using the issue tracker.

# License
The MIT License (MIT). Please see License File for more information.
