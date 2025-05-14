package gomarkup

import (
	"testing"
)

func TestParseMarkup(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantErr    bool
		verifyFunc func(t *testing.T, result []*Node)
	}{
		{
			name:    "empty input",
			input:   "",
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 0 {
					t.Errorf("expected no nodes, got %d", len(result))
				}
			},
		},
		{
			name:    "single valid node",
			input:   "<tag></tag>",
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 1 {
					t.Errorf("expected 1 node, got %d", len(result))
				}
				if result[0].Tag != "tag" {
					t.Errorf("expected tag 'tag', got '%s'", result[0].Tag)
				}
			},
		},
		{
			name:    "multiple sibling nodes",
			input:   "<tag1></tag1><tag2></tag2>",
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 2 {
					t.Errorf("expected 2 nodes, got %d", len(result))
				}
				if result[0].Tag != "tag1" || result[1].Tag != "tag2" {
					t.Errorf("expected tag1 and tag2, got '%s' and '%s'", result[0].Tag, result[1].Tag)
				}
			},
		},
		{
			name:    "nested nodes",
			input:   "<parent><child></child></parent>",
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 1 {
					t.Errorf("expected 1 node, got %d", len(result))
				}
				if result[0].Tag != "parent" {
					t.Errorf("expected parent node, got '%s'", result[0].Tag)
				}
				if len(result[0].Children) != 1 || result[0].Children[0].Tag != "child" {
					t.Errorf("expected child node, got '%+v'", result[0].Children)
				}
			},
		},
		{
			name:    "text before tag",
			input:   "text <tag></tag>",
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 1 {
					t.Errorf("expected 1 node, got %d", len(result))
				}
				if result[0].Tag != "tag" {
					t.Errorf("expected tag 'tag', got '%s'", result[0].Tag)
				}
			},
		},
		{
			name:    "invalid markup - missing closing tag",
			input:   "<tag>",
			wantErr: true,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 0 {
					t.Errorf("expected no nodes on error, got %d", len(result))
				}
			},
		},
		{
			name:    "invalid markup - incorrect nested structure",
			input:   "<parent><child></child>",
			wantErr: true,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 0 {
					t.Errorf("expected no nodes on error, got %d", len(result))
				}
			},
		},
		{
			name:    "attributes in tag",
			input:   `<tag attribute="value"></tag>`,
			wantErr: false,
			verifyFunc: func(t *testing.T, result []*Node) {
				if len(result) != 1 {
					t.Errorf("expected 1 node, got %d", len(result))
				}
				if result[0].Tag != "tag" {
					t.Errorf("expected tag 'tag', got '%s'", result[0].Tag)
				}
				if val, ok := result[0].Attributes["attribute"]; !ok || val != "value" {
					t.Errorf("expected attribute 'attribute' with value 'value', got '%+v'", result[0].Attributes)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMarkup(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseMarkup() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.verifyFunc(t, result)
		})
	}
}
