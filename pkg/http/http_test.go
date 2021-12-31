/**
 * @Author pibing
 * @create 2020/11/18 4:41 PM
 */

package http

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

var contentType = "application/json"

func TestPost(t *testing.T) {

}



func TestPost2(t *testing.T) {
	url := "http://sppc.svinsight.com/login_register"

	body := map[string]interface{}{
		"mobile": "13786769225",
		"type":  "1",
		"code":  "123121",
		"password":  "ffd121314",
	}
	data, err := Post(url, body, contentType)
	if err != nil {
		fmt.Println( err)
		return
	}

	fmt.Println( data)
}

const ascstr = "0123456789"
func GenNToken(size int) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	length := byte(len(ascstr))
	for k, v := range bytes {
		bytes[k] = ascstr[v%length]
	}
	return string(bytes)
}

func TestPost3(t *testing.T) {
	wg:=&sync.WaitGroup{}
	for i:=0;i<10000;i++ {
		wg.Add(1)
		//go dustesslogin(wg)
		go getcode(wg)
	}
	wg.Wait()
	fmt.Println(111)

}

func getcode(wg *sync.WaitGroup)  {
	defer wg.Done()
	url := "http://sppc.svinsight.com/get_verification_code"

	body := map[string]interface{}{
		"mobile": "13"+GenNToken(9),
		"type":  "1",
	}
	data, err := PostHeader(url, body)
	if err != nil {
		fmt.Println( err)
		return
	}
	fmt.Println( data)
	//fmt.Println( body["mobile"])
}

func dustesslogin(wg *sync.WaitGroup)  {
	defer wg.Done()
	url := "https://wp.dustess.com/login"

	 body:=map[string]string{
		"loginName":GenNToken(9),
		"pwd":string(GenNToken(32)),
	}
	_, err := Post(url, body, contentType)
	if err != nil {
		return
	}

	//fmt.Println( data)
}
