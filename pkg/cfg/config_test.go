/**
 * @Author pibing
 * @create 2020/11/16 2:29 PM
 */

package cfg

import (
	"fmt"
	"testing"
)

func TestInitcfg(t *testing.T) {
	configFile:="./dev.yml"
	_ = Initcfg(configFile)
	fmt.Print(config.Mysql.Dbname)
}
