///**
// * @Author pibing
// * @create 2021/12/29 5:53 PM
// */
//
package dirtyfilter

import (
	"go-pkg/pkg/dirtyfilter/lang"
	"strings"
	"sync"
)

// Trie 短语组成的Trie树.
type Trie struct {
	Root *Node
}

var LangMapTrie map[string]*Trie

var once = sync.Once{}

//加载语种的脏词树
func init() {
	once.Do(func() {
		LangMapTrie = make(map[string]*Trie)
		for key, v := range lang.LangMap {
			trie := NewTrie()
			trie.Add(v...)
			LangMapTrie[key] = trie
		}
	})
}

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool             //是否是根节点
	isPathEnd  bool             //是否是末节点
	Character  string           //当前节点的单词
	Children   map[string]*Node //子节点
}

// NewTrie 新建一棵Trie
func NewTrie() *Trie {
	return &Trie{
		Root: NewRootNode(""),
	}
}

// NewNode 新建子节点
func NewNode(character string) *Node {
	return &Node{
		Character: character,
		Children:  make(map[string]*Node, 0),
	}
}

// NewRootNode 新建根节点
func NewRootNode(character string) *Node {
	return &Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[string]*Node, 0),
	}
}

// Add 添加脏词
func (tree *Trie) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

/*如
i like to eat apple， i like to drink water, fuck you
tree树
root->i->like->to->eat->apple
                 ->drink->water
    ->fuck->you*/
func (tree *Trie) add(word string) {
	var current = tree.Root //从根节点开始建树
	//var runes = []rune(word)     //以字符作节点，由于产品的过滤需求，暂废弃
	var runes = strings.Split(word, " ") //分隔空格，以单词作节点
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		r = strings.ToLower(r) //转换成小写
		if next, ok := current.Children[r]; ok {
			current = next //存在，则把该子节点作为当前节点
		} else {
			newNode := NewNode(r) //不存在，则追加当前单词作节点
			current.Children[r] = newNode
			current = newNode
		}
		if position == len(runes)-1 {
			current.isPathEnd = true //树的终端
		}
	}
}

// Replace 词语替换
func (tree *Trie) Replace(text string, replace string) (string, bool) {
	if replace == "" {
		replace = "***"
	}

	var (
		parent  = tree.Root //开始检索的起始节点
		current *Node       //检索出的当前节点
		//runes   = []rune(text)
		runes  = strings.Split(text, " ") //以空格分隔文本为字符串数组
		length = len(runes)               //字符串数组长度,及需要检索的单词树
		left   = 0                        //下一个要检索单词的角标
		ok     bool
		c      = 0
		flag   = false //是否有脏词
	)

	for position := 0; position < len(runes); position++ {
		s := runes[position]
		s = strings.ToLower(s) //转换成小写比对
		c++
		current, ok = parent.Children[s] //查询树中是否存在该单词

		/*若不存在节点
		  或者
		  存在节点，但是该节点不是末节点且当前文本没有下个单词 如 文本是fuck，词库中是fuck you
		  则重新从根节点开始建树下一个单词
		*/
		if !ok || (!current.IsPathEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++ //下个单词的坐标
			continue
		}
		//存在节点，且节点是末节点，如文本 fuck you 词库中是fuck，需替换fuck
		if current.IsPathEnd() && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = replace //替换该位置的单词为***
				flag = true
			}
		}
		/*存在，且不是末节点，如文本fuck you   词库中是fuck you ，
		 *fuck存在的当前节点current将作为you的查询的起始节点
		 */
		parent = current //当前节点作为下个单词查询的根节点
	}

	return strings.Join(runes, " "), flag //把替换后的数组
}

// IsPathEnd 判断是否为某个路径的结束
func (node *Node) IsPathEnd() bool {
	return node.isPathEnd
}
