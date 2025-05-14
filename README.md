# go-markup

## Overview
`go-markup` is a lightweight Go library for parsing HTML-like markup strings into a flexible tree of nodes. It is designed for simplicity, custom markup processing, and easy integration into your Go projects.

`go-markup`은 HTML과 유사한 마크업 문자열을 트리 구조의 노드로 파싱하는 경량 Go 라이브러리입니다. 간편하게 커스텀 마크업을 처리하고 Go 프로젝트에 쉽게 통합할 수 있도록 설계되었습니다.

## Features
- Parses HTML-like markup to a node tree
- Supports tag attributes and nested elements
- Extracts both element text and child nodes
- Simple, dependency-free implementation

## How to Build
1. Make sure you have Go 1.24 or higher installed.
2. Clone the repository:
   ```sh
   git clone https://github.com/starpia-forge/gomarkup.git
   cd gomarkup
   ```
3. Build the library or use it as a module in your Go project.


## How to Use

```shell
go get -u github.com/starpia-forge/go-markup
```

```go
package main

import (
    "fmt"

    "github.com/starpia-forge/go-markup"
)

func main() {
    markup := `<div id="main"><b>Hello</b> World</div>`

    nodes, err := gomarkup.ParseMarkup(markup)
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }

    for _, node := range nodes {
        printTree(node, 0)
    }
}

// printTree displays the node tree hierarchy with tag, attributes, and text.
func printTree(n *gomarkup.Node, depth int) {
    prefix := ""
    for i := 0; i < depth; i++ {
        prefix += "  "
    }
    fmt.Printf("%s<Tag=%s Attrs=%v Text=%q>\n", prefix, n.Tag, n.Attributes, n.Text)
    for _, c := range n.Children {
        printTree(c, depth+1)
    }
}
```

## Cautions
- This library is intended for simple, trusted markup and is **not** a full HTML parser.
- It does not handle malformed or complex HTML, scripts, or styles.
- Use with untrusted or complicated markup is not recommended.

## License
MIT License
