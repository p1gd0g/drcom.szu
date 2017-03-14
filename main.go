package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	url := "http://192.168.254.220/a70.htm"

	// xxx 帐号 ??? 密码
	res, _ := http.Post(url, "text/html; charset=gbk", strings.NewReader("DDDDD=xxx&upass=???&R1=0&R2=&R6=0&para=00&0MKKey=123456"))

	body, _ := ioutil.ReadAll(res.Body)
	text := string(body)

	fmt.Println(text[5840:5885])
	time.Sleep(3 * time.Second)
}
