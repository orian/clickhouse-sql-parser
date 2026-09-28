package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// BenchmarkParseSQLFiles benchmarks parsing all SQL files in the testdata/query directory
func BenchmarkParseSQLFiles(b *testing.B) {
	testFiles, err := filepath.Glob("testdata/query/*.sql")
	if err != nil {
		b.Fatalf("Failed to glob test files: %v", err)
	}

	for _, file := range testFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			b.Fatalf("Failed to read file %s: %v", file, err)
		}

		b.Run(filepath.Base(file), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser := NewParser(string(content))
				_, err := parser.ParseStmts()
				if err != nil {
					b.Fatalf("Failed to parse SQL from %s: %v", file, err)
				}
			}
		})
	}
}

// BenchmarkParseComplexQueries benchmarks parsing specifically complex SQL queries
func BenchmarkParseComplexQueries(b *testing.B) {
	complexQueries := []string{
		"testdata/query/select_with_multi_join.sql",
		"testdata/query/select_with_window_function.sql",
		"testdata/query/select_simple_with_with_clause.sql",
		"testdata/query/select_with_left_join.sql",
		"testdata/benchdata/posthog_huge_0.sql",
		"testdata/benchdata/posthog_huge_1.sql",
	}

	for _, queryFile := range complexQueries {
		content, err := os.ReadFile(queryFile)
		if err != nil {
			b.Fatalf("Failed to read file %s: %v", queryFile, err)
		}

		b.Run(queryFile, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				parser := NewParser(string(content))
				_, err := parser.ParseStmts()
				if err != nil {
					b.Fatalf("Failed to parse SQL from %s: %v", queryFile, err)
				}
			}
		})
	}
}

// benchInputs returns the SQL inputs of the corpus benchmarks: the files
// listed one per line in the file named by BENCH_INPUTS, or by default every
// testdata/query/*.sql and testdata/benchdata/*.sql fixture. huge holds the
// large PostHog queries among them. scripts/bench_compare.sh sets
// BENCH_INPUTS so that two versions are measured on identical inputs.
func benchInputs(b *testing.B) (all, huge []string) {
	b.Helper()
	var paths []string
	if list := os.Getenv("BENCH_INPUTS"); list != "" {
		data, err := os.ReadFile(list)
		if err != nil {
			b.Fatalf("read BENCH_INPUTS: %v", err)
		}
		paths = strings.Fields(string(data))
	} else {
		for _, pattern := range []string{"testdata/query/*.sql", "testdata/benchdata/*.sql"} {
			matches, err := filepath.Glob(pattern)
			if err != nil {
				b.Fatal(err)
			}
			paths = append(paths, matches...)
		}
	}
	for _, path := range paths {
		src, err := os.ReadFile(path)
		if err != nil {
			b.Fatal(err)
		}
		all = append(all, string(src))
		if strings.HasPrefix(filepath.Base(path), "posthog_huge") {
			huge = append(huge, string(src))
		}
	}
	if len(all) == 0 || len(huge) == 0 {
		b.Fatal("no benchmark inputs")
	}
	return all, huge
}

func benchParseAll(b *testing.B, inputs []string) [][]Expr {
	b.Helper()
	parsed := make([][]Expr, len(inputs))
	for i, src := range inputs {
		stmts, err := NewParser(src).ParseStmts()
		if err != nil {
			b.Fatalf("parse: %v", err)
		}
		parsed[i] = stmts
	}
	return parsed
}

func benchPrintVisitor(parsed [][]Expr) {
	for _, stmts := range parsed {
		for _, stmt := range stmts {
			printer := NewPrintVisitor()
			_ = stmt.Accept(printer)
			_ = printer.String()
		}
	}
}

// BenchmarkParseCorpus parses every benchmark input once per iteration.
func BenchmarkParseCorpus(b *testing.B) {
	all, _ := benchInputs(b)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchParseAll(b, all)
	}
}

// BenchmarkParseHuge parses the large PostHog queries.
func BenchmarkParseHuge(b *testing.B) {
	_, huge := benchInputs(b)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchParseAll(b, huge)
	}
}

// BenchmarkStringCorpus prints every parsed input with String().
func BenchmarkStringCorpus(b *testing.B) {
	all, _ := benchInputs(b)
	parsed := benchParseAll(b, all)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, stmts := range parsed {
			for _, stmt := range stmts {
				_ = stmt.String()
			}
		}
	}
}

// BenchmarkPrintVisitorCorpus prints every parsed input with PrintVisitor,
// the formatter behind the CLI's -format.
func BenchmarkPrintVisitorCorpus(b *testing.B) {
	all, _ := benchInputs(b)
	parsed := benchParseAll(b, all)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchPrintVisitor(parsed)
	}
}

// BenchmarkPrintVisitorHuge prints the large PostHog queries with
// PrintVisitor.
func BenchmarkPrintVisitorHuge(b *testing.B) {
	_, huge := benchInputs(b)
	parsed := benchParseAll(b, huge)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchPrintVisitor(parsed)
	}
}
