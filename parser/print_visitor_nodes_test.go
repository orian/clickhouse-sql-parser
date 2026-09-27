package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// printNode prints node with a fresh PrintVisitor, turning a panic into an
// error so one broken node type cannot abort the whole check.
func printNode(node Expr) (out string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	printer := NewPrintVisitor()
	if err := node.Accept(printer); err != nil {
		return "", err
	}
	return printer.String(), nil
}

// TestPrintVisitor_EveryNodeMatchesString checks that printing any node of any
// fixture statement with PrintVisitor gives the same SQL as the node's
// String() (#98). PrintVisitor can be applied to a subtree, so every VisitX
// must be correct on its own, not only when reached from a statement.
func TestPrintVisitor_EveryNodeMatchesString(t *testing.T) {
	for _, dir := range sqlFixtureDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
		require.NoError(t, err)
		for _, file := range files {
			t.Run(file, func(t *testing.T) {
				src, err := os.ReadFile(file)
				require.NoError(t, err)
				stmts, err := NewParser(string(src)).ParseStmts()
				if err != nil {
					t.Skip("fixture does not parse")
				}
				for _, stmt := range stmts {
					Walk(stmt, func(node Expr) bool {
						got, err := printNode(node)
						require.NoError(t, err, "%T: %s", node, node.String())
						require.Equal(t, node.String(), got, "%T", node)
						return true
					})
				}
			})
		}
	}
}

// countingVisitor counts how often each column identifier is visited.
type countingVisitor struct {
	DefaultASTVisitor
	seen map[string]int
}

func (v *countingVisitor) VisitIdent(i *Ident) error {
	v.seen[i.Name]++
	return nil
}

// TestAccept_VisitsChildrenOnce checks that IntervalFrom and WindowFrameParam
// no longer walk their child in Accept (refactor-visitor.md invariant 1), so a
// visitor embedding DefaultASTVisitor sees the child exactly once.
func TestAccept_VisitsChildrenOnce(t *testing.T) {
	for _, tc := range []struct {
		sql   string
		ident string
	}{
		{"SELECT EXTRACT(HOUR FROM ts_col) FROM t", "ts_col"},
		{"SELECT sum(x) OVER (ORDER BY y ROWS BETWEEN {start:UInt32} PRECEDING AND CURRENT ROW) FROM t", "start"},
	} {
		t.Run(tc.sql, func(t *testing.T) {
			stmts, err := NewParser(tc.sql).ParseStmts()
			require.NoError(t, err)
			visitor := &countingVisitor{seen: map[string]int{}}
			visitor.Self = visitor
			require.NoError(t, stmts[0].Accept(visitor))
			require.Equal(t, 1, visitor.seen[tc.ident], "visits of %s: %v", tc.ident, visitor.seen)
		})
	}
}
