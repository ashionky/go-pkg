/**
 * @Author pibing
 * @create 2020/12/11 4:44 PM
 */

package file

import (
	"fmt"

	"testing"
)

type data struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Arr  []string `json:"arr"`
}

func TestDownJson(t *testing.T) {
	mp := map[string]interface{}{}
	path := "./test.json"
	_ = ReadJson(mp, path)
	fmt.Println(mp)

}
