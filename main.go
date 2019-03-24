package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	var u = flag.String("u", "", "Your userid.")
	var p = flag.String("p", "", "Your password.")

	flag.Parse()
	fmt.Println(*u, *p)

	url := "http://192.168.254.220/a70.htm"

	res, _ := http.Post(url, "text/html; charset=gbk", strings.NewReader("DDDDD="+*u+"&upass="+*p+"&R1=0&R2=&R6=0&para=00&0MKKey=123456"))

	body, _ := ioutil.ReadAll(res.Body)
	text := string(body)

	fmt.Println(text[5840:5885])
	time.Sleep(1 * time.Second)
}
