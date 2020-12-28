/**
 * @Author pibing
 * @create 2020/12/27 12:17 PM
 */

package jwt_util

import (
	"fmt"
	"testing"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	Name string
	Age  int
}

func Test( t *testing.T)  {
	user:=User{
		Name: "张三",
		Age:  11,
	}

	var result map[string]interface{}
	_=mapstructure.Decode(user,&result)
	s, e := MakeToken(result)
	fmt.Println(s)
	fmt.Println(e)

	token := ParseToken(s)
	u:=User{}
	_=mapstructure.Decode(token,&u)
	fmt.Println(u.Age)
}
