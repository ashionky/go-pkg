/**
 * @Author pibing
 * @create 2020/12/27 11:40 AM
 */

package http

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	ips, e := GetServerIP()
	fmt.Println(ips)
	fmt.Println(e)
}
