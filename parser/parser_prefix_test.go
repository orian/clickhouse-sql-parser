package parser

import "testing"

// TestParser_FixturePrefixes deterministically exercises EOF in every byte of
// small fixtures (including strings and comments), and at token boundaries in
// larger fixtures to bound the quadratic cost of reparsing prefixes.
func TestParser_FixturePrefixes(t *testing.T) {
	for _, dir := range sqlFixtureDirs {
		for _, seed := range readSQLSeeds(t, dir, 0) {
			t.Run(seed.name, func(t *testing.T) {
				check := func(end int) {
					t.Helper()
					defer func() {
						if value := recover(); value != nil {
							t.Fatalf("panic at byte %d for SQL %q: %v", end, seed.sql[:end], value)
						}
					}()
					_, _ = NewParser(seed.sql[:end]).ParseStmts()
				}
				if len(seed.sql) <= 1024 {
					for end := 0; end <= len(seed.sql); end++ {
						check(end)
					}
					return
				}
				check(0)
				lexer := NewLexer(seed.sql)
				for {
					if err := lexer.consumeToken(); err != nil {
						t.Fatal(err)
					}
					if lexer.lastToken == nil {
						break
					}
					check(int(lexer.lastToken.Pos))
					check(int(lexer.lastToken.End))
				}
				check(len(seed.sql))
			})
		}
	}
}
