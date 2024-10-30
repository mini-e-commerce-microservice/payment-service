package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("SB-Mid-server-LR5SazLSkB3wdU2MRwIpkXFp:")))
	r, _ := http.NewRequest()
	r.SetBasicAuth()
}
