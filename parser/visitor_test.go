package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
)

func TestVisitor_Identical(t *testing.T) {
	visitor := DeprecatedDefaultASTVisitor{}

	for _, dir := range []string{"./testdata/dml", "./testdata/ddl", "./testdata/query", "./testdata/basic"} {
		outputDir := dir + "/format"

		entries, err := os.ReadDir(dir)
		require.NoError(t, err)
		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}
			t.Run(entry.Name(), func(t *testing.T) {
				fileBytes, err := os.ReadFile(filepath.Join(dir, entry.Name()))
				require.NoError(t, err)
				parser := Parser{
					lexer: NewLexer(string(fileBytes)),
				}
				stmts, err := parser.ParseStmts()
				require.NoError(t, err)
				var builder strings.Builder
				builder.WriteString("-- Origin SQL:\n")
				builder.Write(fileBytes)
				builder.WriteString("\n\n-- Format SQL:\n")
				var formatSQLBuilder strings.Builder
				for _, stmt := range stmts {
					err := stmt.Accept(&visitor)
					require.NoError(t, err)

					formatSQLBuilder.WriteString(stmt.String())
					formatSQLBuilder.WriteByte(';')
					formatSQLBuilder.WriteByte('\n')
				}
				formatSQL := formatSQLBuilder.String()
				builder.WriteString(formatSQL)
				validFormatSQL(t, formatSQL)
				g := goldie.New(t,
					goldie.WithNameSuffix(""),
					goldie.WithDiffEngine(goldie.ColoredDiff),
					goldie.WithFixtureDir(outputDir))
				g.Assert(t, entry.Name(), []byte(builder.String()))
			})
		}
	}
}

type simpleRewriteVisitor struct {
	DefaultVisitor
}

func (v *simpleRewriteVisitor) VisitTableIdentifier(expr *TableIdentifier) error {
	if err := v.DefaultVisitor.VisitTableIdentifier(expr); err != nil {
		return err
	}
	if expr.Table.String() == "group_by_all" {
		expr.Table = &Ident{Name: "hack"}
	}
	return nil
}

func (v *simpleRewriteVisitor) VisitOrderByExpr(expr *OrderExpr) error {
	if err := v.DefaultVisitor.VisitOrderByExpr(expr); err != nil {
		return err
	}
	expr.Direction = OrderDirectionDesc
	return nil
}

func TestVisitor_SimpleRewrite(t *testing.T) {
	visitor := &simpleRewriteVisitor{}
	visitor.self = visitor

	sql := `SELECT a, COUNT(b) FROM group_by_all GROUP BY CUBE(a) WITH CUBE WITH TOTALS ORDER BY a;`
	parser := NewParser(sql)
	stmts, err := parser.ParseStmts()
	require.NoError(t, err)

	require.Equal(t, 1, len(stmts))
	stmt := stmts[0]

	err = stmt.Accept(visitor)
	require.NoError(t, err)
	newSql := stmt.String()

	require.NotSame(t, sql, newSql)
	require.Contains(t, newSql, "hack")
	require.Contains(t, newSql, string(OrderDirectionDesc))
}

type nestedRewriteVisitor struct {
	DefaultVisitor
	stack []Expr
}

func (v *nestedRewriteVisitor) VisitTableIdentifier(expr *TableIdentifier) error {
	expr.Table = &Ident{Name: fmt.Sprintf("table%d", len(v.stack))}
	return nil
}

func (v *nestedRewriteVisitor) enter(expr Expr) {
	if s, ok := expr.(*SelectQuery); ok {
		v.stack = append(v.stack, s)
	}
}

func (v *nestedRewriteVisitor) leave(expr Expr) {
	if _, ok := expr.(*SelectQuery); ok {
		v.stack = v.stack[1:]
	}
}

func TestVisitor_NestRewrite(t *testing.T) {
	visitor := &nestedRewriteVisitor{}
	visitor.self = visitor

	sql := `SELECT replica_name FROM system.ha_replicas UNION DISTINCT SELECT replica_name FROM system.ha_unique_replicas format JSON`
	parser := NewParser(sql)
	stmts, err := parser.ParseStmts()
	require.NoError(t, err)

	require.Equal(t, 1, len(stmts))
	stmt := stmts[0]

	err = stmt.Accept(visitor)
	require.NoError(t, err)
	newSql := stmt.String()

	require.NotSame(t, sql, newSql)
	require.Less(t, strings.Index(newSql, "table1"), strings.Index(newSql, "table2"))
}
