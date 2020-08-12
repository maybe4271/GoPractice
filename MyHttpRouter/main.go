package main

import (
	"fmt"
	"net/http"
	"strings"
)

type nodeType uint8

//节点类型： 静态节点：eg:static status step中的a/根节点/除另外三个之外的节点/通配符*匹配*后的所有字符
const (
	static nodeType = iota
	root
	param
	catchAll
)

type Handle func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
	name string
}

func (r *Router) Handle(method, path string, handle Handle) {
	fmt.Println("%s %s", method, path)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type node struct {
	path      string
	indices   string
	wildChild bool
	nType     nodeType
	priority  uint32
	children  []*node
	handle    Handle
}

func longestCommonPrefix(str1, str2 string) int {
	i := 0
	len_ := min(len(str1), len(str2))
	for i < len_ && str1[i] == str2[i] {
		i++
	}
	return i
}

func (n *node) addRoute(path string, handle Handle) {
	fullPath := path
	n.priority++

	//if empty
	if len(n.path) == 0 && len(n.indices) == 0 {
		n.insertChild(path, fullPath, handle)
		n.nType = root
		return
	}

walk:
	for {
		i := longestCommonPrefix(path, n.path)
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],
				indices:   n.indices,
				wildChild: n.wildChild,
				nType:     static,
				priority:  n.priority - 1,
				children:  n.children,
				handle:    n.handle,
			}
			n.path = path[:i]
			//child节点的path字段的首字符合集
			n.indices = string([]byte{n.path[i]})
			n.wildChild = false
			n.children = []*node{&child}
			n.handle = nil
		}

		if i < len(path) {
			path = path[i:]

			//only one child
			if n.wildChild {
				n = n.children[0]
				n.priority++

				if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
					n.nType != catchAll && (len(n.path) >= len(path) || path[len(n.path)] == '/') {
					continue walk
				} else {
					pathSeg := path
					if n.nType != catchAll {
						pathSeg = strings.SplitN(pathSeg, '/', 2)[0]
					}
					prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				}
			}
		}
	}
}

func (n *node) insertChild(path, fullPath string, handle Handle) {
	fmt.Println("insertChild running")
}

func main() {
	s := strings.Split(",", "")
	fmt.Println(s, len(s))
	fmt.Println("Hello World")
}
