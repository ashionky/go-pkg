package dirtyfilter

import (
	"fmt"
	"testing"
)

//todo 也可以不按语言，但需要把所有词组加载到一棵树上
func TestName(t *testing.T) {
	s := "fuck you mm" //fuck you 需要过滤的词
	if trie, ok := LangMapTrie["en"]; ok {
		s2, ok2 := trie.Replace(s, "")
		fmt.Println(s2, ok2) //*** you mm true
	}

}
