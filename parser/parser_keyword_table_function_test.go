package parser

import "testing"

func TestParser_KeywordNamedTableFunctions(t *testing.T) {
	tests := []struct {
		name         string
		sql          string
		wantFrom     bool
		wantDescribe bool
		wantFunction bool
	}{
		{
			name:         "format in FROM",
			sql:          `SELECT * FROM format(JSONEachRow, 'a String', '{"a":"x"}')`,
			wantFrom:     true,
			wantFunction: true,
		},
		{
			name:         "format in DESCRIBE",
			sql:          `DESC format(JSONEachRow, '{"x" : 1}')`,
			wantDescribe: true,
			wantFunction: true,
		},
		{
			name: "keyword column before FROM",
			sql:  `SELECT from FROM t`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmts, err := NewParser(tt.sql).ParseStmts()
			if err != nil {
				t.Fatalf("ParseStmts() error = %v", err)
			}
			if len(stmts) != 1 {
				t.Fatalf("ParseStmts() returned %d statements, want 1", len(stmts))
			}
			beautifyRoundTrip(t, stmts[0])
			if tt.wantFrom {
				query, ok := stmts[0].(*SelectQuery)
				if !ok || query.From == nil {
					t.Errorf("expected a SelectQuery with a FROM clause, got %#v", stmts[0])
				}
			}
			if tt.wantDescribe {
				describe, ok := stmts[0].(*DescribeStmt)
				if !ok || describe.TargetExpr == nil || describe.Target != nil {
					t.Errorf("expected a DescribeStmt with a function TargetExpr, got %#v", stmts[0])
				}
			}
			if tt.wantFunction {
				found := false
				Walk(stmts[0], func(expr Expr) bool {
					if function, ok := expr.(*TableFunctionExpr); ok {
						if name, ok := function.Name.(*Ident); ok && name.Name == "format" {
							found = true
						}
					}
					return true
				})
				if !found {
					t.Error("AST traversal did not reach the format table function")
				}
			}
		})
	}
}
