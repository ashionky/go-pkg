/**
 * @Author pibing
 * @create 2020/11/18 4:41 PM
 */

package http

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var contentType="application/json"
func TestPost(t *testing.T) {
	url  := "http://47.108.203.220:8555"
	bc_no := strconv.FormatInt(11941, 16)
	bc_no="0x"+bc_no
	params:=[]string{bc_no}
	body :=map[string]interface{}{
		"jsonrpc":"2.0",
		"method":"eth_getBlockTransactionCountByNumber",
		"params":params,
		"id":time.Now().Unix(),
	}
	data,err:=Post(url,body,contentType)
	if err !=nil {
		fmt.Printf("eth_getBlockTransactionCountByNumber err:%v",err)
		return
	}
	fmt.Println(string(data))

}
