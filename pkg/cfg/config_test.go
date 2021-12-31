/**
 * @Author pibing
 * @create 2020/11/16 2:29 PM
 */

package cfg

import (
	"fmt"
	"github.com/jinzhu/configor"
	"testing"
)

type Cfg struct {
	Sl []string
}
func TestInitcfg(t *testing.T) {
	configFile := "./en.json"
	config:=Cfg{}
	err := configor.Load(&config, configFile)
	if err!=nil {

	}
	fmt.Println(config.Sl[0])

}
