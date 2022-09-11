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
	"encoding/xml"
	"fmt"

	netgsm "github.com/ozgur-soft/netgsm.go/src"
)

func main() {
	api, req := netgsm.Api("header", "username", "password")
	req.Body.Msg = "test"
	req.Body.No = "905555555555"
	res := api.Sms(req) // Normal sms
	// res := api.Otp(req) // Hızlı sms
	pretty, _ := xml.MarshalIndent(res, " ", " ")
	fmt.Println(string(pretty))
}
```
