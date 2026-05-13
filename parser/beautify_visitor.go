package parser

import "strings"

// BeautifyVisitor emits indented, line-broken SQL — the "beautify" counterpart
// to PrintVisitor. Like every visitor on this branch, it embeds DefaultASTVisitor
// so that any unimplemented Visit method falls back to default descent.
//
// Architectural note: formatting belongs to the visitor, not to the nodes.
// Upstream introduced a Formatter type with a FormatSQL(*Formatter) method on
// every AST node (commit 9275c63 and the beautify follow-ups). That breaks our
// branch's invariant — see refactor-visitor.md — so we keep formatting logic
// inside the visitor and leave the nodes untouched.
//
// This file is a stub. To port a specific upstream beautify improvement (e.g.
// FROM/JOIN indentation from #244/50c0046, INSERT column-list wrapping from
// #246, ON CLUSTER line breaks from #251, etc.), translate the relevant
// FormatSQL body from upstream's parser/format.go into a VisitX method here,
// writing into b.builder while tracking depth with Indent/Outdent.
type BeautifyVisitor struct {
	DefaultASTVisitor
	builder strings.Builder
	depth   int
	indent  string // characters per level; defaults to two spaces
}

func NewBeautifyVisitor() *BeautifyVisitor {
	v := &BeautifyVisitor{indent: "  "}
	v.Self = v
	return v
}

func (b *BeautifyVisitor) String() string {
	return b.builder.String()
}

// WriteIndent writes the current indentation prefix to the buffer.
func (b *BeautifyVisitor) WriteIndent() {
	for i := 0; i < b.depth; i++ {
		b.builder.WriteString(b.indent)
	}
}

// Indent increases the indentation level by one.
func (b *BeautifyVisitor) Indent() { b.depth++ }

// Outdent decreases the indentation level by one.
func (b *BeautifyVisitor) Outdent() {
	if b.depth > 0 {
		b.depth--
	}
}
