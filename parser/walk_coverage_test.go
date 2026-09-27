package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// enteredNodes records every node a DefaultASTVisitor-based traversal enters.
type enteredNodes struct {
	DefaultASTVisitor
	seen map[Expr]bool
	list []Expr
}

func (v *enteredNodes) Enter(e Expr) {
	if isNilExpr(e) || v.seen[e] {
		return
	}
	v.seen[e] = true
	v.list = append(v.list, e)
}

// TestWalk_ReachesEveryNodeTheVisitorReaches checks that Walk (and so Find,
// FindAll and Transform) visits every node that Accept with a
// DefaultASTVisitor reaches, for every statement of every fixture (#105).
func TestWalk_ReachesEveryNodeTheVisitorReaches(t *testing.T) {
	files := fixtureSQLFiles(t)

	for _, file := range files {
		src, err := os.ReadFile(file)
		require.NoError(t, err)
		stmts, err := NewParser(string(src)).ParseStmts()
		if err != nil {
			continue
		}
		for _, stmt := range stmts {
			walked := map[Expr]bool{}
			Walk(stmt, func(e Expr) bool {
				walked[e] = true
				return true
			})
			visitor := &enteredNodes{seen: map[Expr]bool{}}
			visitor.Self = visitor
			require.NoError(t, stmt.Accept(visitor), "%s: %s", file, stmt.String())
			for _, e := range visitor.list {
				if !walked[e] {
					t.Errorf("%s: Walk misses %T %q in %s", file, e, e.String(), clip(stmt.String()))
				}
			}
		}
	}
}

// TestVisitor_ReachesEveryNodeWalkReaches is the reverse of
// TestWalk_ReachesEveryNodeTheVisitorReaches: a visitor embedding
// DefaultASTVisitor enters every node Walk reaches (#106).
func TestVisitor_ReachesEveryNodeWalkReaches(t *testing.T) {
	for _, file := range fixtureSQLFiles(t) {
		src, err := os.ReadFile(file)
		require.NoError(t, err)
		stmts, err := NewParser(string(src)).ParseStmts()
		if err != nil {
			continue
		}
		for _, stmt := range stmts {
			visitor := &enteredNodes{seen: map[Expr]bool{}}
			visitor.Self = visitor
			require.NoError(t, stmt.Accept(visitor), "%s: %s", file, stmt.String())
			Walk(stmt, func(e Expr) bool {
				if !isNilExpr(e) && !visitor.seen[e] {
					t.Errorf("%s: visitor misses %T %q in %s", file, e, e.String(), clip(stmt.String()))
				}
				return true
			})
		}
	}
}

// TestTraversal_ReachesEveryChildField checks, for every node of every
// fixture statement, that each AST node held in one of its exported fields
// (directly, in a slice or in an embedded struct) is reached by both Walk and
// the DefaultASTVisitor recursion. Unlike the two tests above, it also catches
// a child that both traversals skip (#106).
func TestTraversal_ReachesEveryChildField(t *testing.T) {
	for _, file := range fixtureSQLFiles(t) {
		src, err := os.ReadFile(file)
		require.NoError(t, err)
		stmts, err := NewParser(string(src)).ParseStmts()
		if err != nil {
			continue
		}
		for _, stmt := range stmts {
			visitor := &enteredNodes{seen: map[Expr]bool{}}
			visitor.Self = visitor
			require.NoError(t, stmt.Accept(visitor), "%s: %s", file, stmt.String())
			walked := map[Expr]bool{}
			Walk(stmt, func(e Expr) bool {
				walked[e] = true
				return true
			})
			for node := range walked {
				for _, child := range childFields(node) {
					if !walked[child] || !visitor.seen[child] {
						t.Errorf("%s: %T child %T %q: walked=%v visited=%v in %s",
							file, node, child, child.String(), walked[child], visitor.seen[child], clip(stmt.String()))
					}
				}
			}
		}
	}
}

// childFields returns the non-nil AST nodes held in node's exported fields.
func childFields(node Expr) []Expr {
	exprType := reflect.TypeOf((*Expr)(nil)).Elem()
	var children []Expr
	var add func(f reflect.Value)
	add = func(f reflect.Value) {
		switch f.Kind() {
		case reflect.Ptr, reflect.Interface:
			if !f.IsNil() && f.Type().Implements(exprType) && !isNilExpr(f.Interface().(Expr)) {
				children = append(children, f.Interface().(Expr))
			}
		case reflect.Slice:
			for i := 0; i < f.Len(); i++ {
				add(f.Index(i))
			}
		case reflect.Struct:
			if f.CanAddr() && f.Addr().Type().Implements(exprType) {
				children = append(children, f.Addr().Interface().(Expr))
				return
			}
			for i := 0; i < f.NumField(); i++ {
				if f.Type().Field(i).IsExported() {
					add(f.Field(i))
				}
			}
		}
	}
	add(reflect.ValueOf(node).Elem())
	return children
}

// fixtureSQLFiles lists every .sql fixture under testdata, skipping the
// generated format and beautify outputs.
func fixtureSQLFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	require.NoError(t, filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && (info.Name() == "format" || info.Name() == "beautify") {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.HasSuffix(path, ".sql") {
			files = append(files, path)
		}
		return nil
	}))
	require.NotEmpty(t, files)
	return files
}

func clip(s string) string {
	if len(s) > 120 {
		return fmt.Sprintf("%s…", s[:120])
	}
	return s
}

// TestWalk_Issue105Regressions pins the children Walk used to skip (#105).
func TestWalk_Issue105Regressions(t *testing.T) {
	names := func(sql string, match func(Expr) (string, bool)) []string {
		stmts, err := NewParser(sql).ParseStmts()
		require.NoError(t, err)
		var out []string
		Walk(stmts[0], func(e Expr) bool {
			if name, ok := match(e); ok {
				out = append(out, name)
			}
			return true
		})
		return out
	}
	tables := func(e Expr) (string, bool) {
		if table, ok := e.(*TableIdentifier); ok {
			return table.String(), true
		}
		return "", false
	}
	numbers := func(e Expr) (string, bool) {
		if number, ok := e.(*NumberLiteral); ok {
			return number.String(), true
		}
		return "", false
	}
	idents := func(e Expr) (string, bool) {
		if ident, ok := e.(*Ident); ok {
			return ident.String(), true
		}
		return "", false
	}

	require.Equal(t, []string{"a", "b"}, names("SELECT * FROM a EXCEPT SELECT * FROM b", tables))
	require.Equal(t, []string{"a", "b", "c"}, names("SELECT * FROM a EXCEPT SELECT * FROM b EXCEPT SELECT * FROM c", tables))
	require.Equal(t, []string{"1", "2", "3", "4"}, names("INSERT INTO t (a, b) VALUES (1, 2), (3, 4)", numbers))
	require.Equal(t, []string{"c", "k", "v"}, names("CREATE NAMED COLLECTION c AS k = v", idents))
	require.Contains(t, names("SELECT sum(x) OVER (ORDER BY y ROWS BETWEEN {start:UInt32} PRECEDING AND CURRENT ROW) FROM t", idents), "start")

	var codecs []string
	stmts, err := NewParser("CREATE TABLE t (d DateTime) ENGINE = MergeTree ORDER BY d TTL d + INTERVAL 1 MONTH RECOMPRESS CODEC(ZSTD(17))").ParseStmts()
	require.NoError(t, err)
	Walk(stmts[0], func(e Expr) bool {
		if codec, ok := e.(*CompressionCodec); ok {
			codecs = append(codecs, codec.String())
		}
		return true
	})
	require.Len(t, codecs, 1)

	// Map keys are walked in place, so rewriting one changes the AST.
	stmts, err = NewParser("SELECT * FROM t SETTINGS additional_table_filters = {'t': 'x = 1'}").ParseStmts()
	require.NoError(t, err)
	Walk(stmts[0], func(e Expr) bool {
		if s, ok := e.(*StringLiteral); ok && s.Literal == "t" {
			s.Literal = "renamed"
		}
		return true
	})
	require.Contains(t, stmts[0].String(), "{'renamed': 'x = 1'}")
}

// TestVisitor_Issue106Regressions pins the children the visitor used to skip
// (#106): a visitor embedding DefaultASTVisitor sees them.
func TestVisitor_Issue106Regressions(t *testing.T) {
	for _, tc := range []struct {
		sql  string
		want string // String() of a node the visitor must enter
	}{
		{"SELECT DISTINCT ON (a, b) a, c FROM t", "b"},
		{"ALTER TABLE t ADD PROJECTION p (SELECT * ORDER BY y)", "y"},
		{"CREATE TABLE t UUID '11111111-1111-1111-1111-111111111111' (a Int32) ENGINE = Memory", "'11111111-1111-1111-1111-111111111111'"},
		{"CREATE DATABASE db", "db"},
		{"CREATE TABLE t (a Int32) ENGINE = Memory COMMENT 'c'", "'c'"},
		{"ALTER TABLE t DROP DETACHED PARTITION 'p' SETTINGS allow_drop_detached = 1", "allow_drop_detached=1"},
		{"RENAME TABLE a TO b", "a TO b"},
		{"SELECT * FROM t STREAM", "STREAM"},
	} {
		t.Run(tc.sql, func(t *testing.T) {
			stmts, err := NewParser(tc.sql).ParseStmts()
			require.NoError(t, err)
			visitor := &enteredNodes{seen: map[Expr]bool{}}
			visitor.Self = visitor
			require.NoError(t, stmts[0].Accept(visitor))
			var entered []string
			for _, e := range visitor.list {
				entered = append(entered, e.String())
			}
			require.Contains(t, entered, tc.want)
		})
	}
}
