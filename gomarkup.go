package gomarkup

import (
	"fmt"
	"regexp"
	"strings"
)

// Node represents a single node in a tree structure,
// containing tag data, attributes, child nodes, and text content.
type Node struct {
	Tag        string
	Attributes map[string]interface{}
	Children   []*Node
	Text       string
}

// ParseMarkup parses a markup language string and
// returns a slice of Nodes representing the parsed structure or an error.
func ParseMarkup(input string) ([]*Node, error) {
	input = strings.TrimSpace(input)
	var nodes []*Node
	for len(input) > 0 {
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			break
		}
		if input[0] != '<' {
			i := strings.Index(input, "<")
			if i == -1 {
				break
			}
			input = input[i:]
			continue
		}
		node, consumed, err := parseNode(input)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
		input = input[consumed:]
	}
	return nodes, nil
}

// parseNode parses an input string representing a single HTML-like element
// into a Node structure, consuming its content.
//
// Returns the parsed Node, the number of characters consumed,
// or an error if parsing fails.
func parseNode(input string) (*Node, int, error) {
	tagOpen := regexp.MustCompile(`^<([a-zA-Z0-9_\-]+)((?:\s+[^>]+)?)\s*>`)
	loc := tagOpen.FindStringSubmatchIndex(input)
	if loc == nil {
		return nil, 0, fmt.Errorf("invalid tag start: %s", input)
	}
	tagName := input[loc[2]:loc[3]]
	rawAttrs := input[loc[4]:loc[5]]
	attrs := parseAttributes(rawAttrs)
	offset := loc[1]
	rem := input[offset:]

	endTag := "</" + tagName + ">"
	endIdx := strings.Index(rem, endTag)
	if endIdx == -1 {
		return nil, 0, fmt.Errorf("no closing tag for <%s>", tagName)
	}
	body := rem[:endIdx]
	children, text := parseChildren(body)

	node := &Node{
		Tag:        tagName,
		Attributes: attrs,
		Children:   children,
		Text:       text,
	}
	consumed := offset + endIdx + len(endTag)
	return node, consumed, nil
}

// parseAttributes extracts key-value pairs from a string formatted as
// HTML-like attributes and returns them as a map.
func parseAttributes(s string) map[string]interface{} {
	res := map[string]interface{}{}
	attrRx := regexp.MustCompile(`([a-zA-Z0-9_\-]+)\s*=\s*"([^"]*)"`)
	matches := attrRx.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		res[m[1]] = m[2]
	}
	return res
}

// parseChildren extracts child nodes and text from an HTML-like body string,
// returning parsed nodes and leftover text.
func parseChildren(body string) ([]*Node, string) {
	var res []*Node
	txt := ""
	for {
		body = strings.TrimSpace(body)
		if len(body) == 0 {
			break
		}
		if body[0] != '<' {
			i := strings.Index(body, "<")
			if i == -1 {
				txt += body
				break
			}
			txt += body[:i]
			body = body[i:]
			continue
		}
		node, consumed, err := parseNode(body)
		if err != nil {
			break
		}
		res = append(res, node)
		body = body[consumed:]
	}
	txt = strings.TrimSpace(txt)
	return res, txt
}

// printTree prints a hierarchical representation of
// a tree structure defined by Node, starting at the specified depth.
//
// Each node is printed with its tag, attributes, and text, prefixed by
// an indentation proportional to its depth level.
func printTree(n *Node, depth int) {
	prefix := strings.Repeat("  ", depth)
	fmt.Printf("%s<Tag=%s Attrs=%v Text=%q>\n", prefix, n.Tag, n.Attributes, n.Text)
	for _, c := range n.Children {
		printTree(c, depth+1)
	}
}
