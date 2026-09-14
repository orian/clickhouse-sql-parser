package parser

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"testing"
)

var docsSQL = flag.Bool("docs-sql", false, "parse the extracted documentation corpus and compare its baseline")
var docsSQLUpdate = flag.Bool("docs-sql-update", false, "parse documentation SQL and replace its baseline")
var docsSQLStrict = flag.Bool("docs-sql-strict", false, "fail for every documentation example rejected by the parser")

type docsSource struct {
	Path             string            `json:"path"`
	Line             int               `json:"line"`
	EndLine          int               `json:"end_line"`
	Info             string            `json:"info"`
	Kind             string            `json:"kind"`
	ExpectedErrors   []string          `json:"expected_errors,omitempty"`
	Attributes       map[string]string `json:"attributes,omitempty"`
	HasSubstitutions bool              `json:"has_substitutions,omitempty"`
}
type docsExample struct {
	ID      string       `json:"id"`
	File    string       `json:"file"`
	Sources []docsSource `json:"sources"`
}
type docsResult struct {
	ID         string       `json:"id"`
	File       string       `json:"file"`
	Status     string       `json:"status"`
	Statements int          `json:"statements"`
	Error      string       `json:"error,omitempty"`
	Sources    []docsSource `json:"sources"`
}
type docsReport struct {
	Examples    int            `json:"examples"`
	Occurrences int            `json:"occurrences"`
	Counts      map[string]int `json:"counts"`
	Results     []docsResult   `json:"results"`
}

func parseDocsExample(sql string) (status string, statements int, message string) {
	status = "accepted"
	defer func() {
		if value := recover(); value != nil {
			status, message = "panic", fmt.Sprintf("%v\n%s", value, debug.Stack())
		}
	}()
	stmts, err := NewParser(sql).ParseStmts()
	if err != nil {
		return "rejected", 0, err.Error()
	}
	if len(stmts) == 0 {
		return "empty", 0, "no statements parsed"
	}
	return status, len(stmts), ""
}

// TestDocumentationSQL is opt-in: this corpus includes intentionally incomplete
// grammar templates, foreign dialects, and mislabeled output as well as queries.
// The baseline records observations, not assertions of ClickHouse SQL validity.
func TestDocumentationSQL(t *testing.T) {
	if !*docsSQL && !*docsSQLUpdate && !*docsSQLStrict {
		t.Skip("use -docs-sql to parse the documentation corpus")
	}
	runSQLCorpus(t, "docs-sql", *docsSQLUpdate, *docsSQLStrict)
}

// Both corpora use the same parser, integrity checks, and baseline semantics.
func runSQLCorpus(t *testing.T, directory string, update, strict bool) {
	t.Helper()
	root := filepath.Join("..", directory)
	readJSON := func(name string, target any) {
		t.Helper()
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(data, target); err != nil {
			t.Fatal(err)
		}
	}
	var manifest struct {
		Version     int           `json:"version"`
		Occurrences int           `json:"occurrences"`
		Examples    []docsExample `json:"examples"`
	}
	readJSON("manifest.json", &manifest)
	if manifest.Version != 1 || len(manifest.Examples) == 0 {
		t.Fatal("missing or unsupported corpus manifest")
	}
	baseline := map[string]string{}
	if !update && !strict {
		readJSON("baseline.json", &baseline)
	}
	observed := map[string]string{}
	report := docsReport{Examples: len(manifest.Examples), Occurrences: manifest.Occurrences, Counts: map[string]int{}}
	for _, example := range manifest.Examples {
		t.Run(example.ID, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(root, example.File))
			if err != nil {
				t.Fatal(err)
			}
			if fmt.Sprintf("%x", sha256.Sum256(data)) != example.ID {
				t.Fatal("SQL content does not match manifest hash; rerun extractor")
			}
			if _, exists := observed[example.ID]; exists {
				t.Fatal("duplicate manifest ID")
			}
			status, statements, message := parseDocsExample(string(data))
			observed[example.ID] = status
			report.Counts[status]++
			report.Results = append(report.Results, docsResult{example.ID, example.File, status, statements, message, example.Sources})
			if strict && status != "accepted" {
				t.Errorf("%s (%s): %s", example.File, status, message)
			} else if !update && !strict && baseline[example.ID] != status {
				t.Errorf("%s: baseline %q -> %q: %s", example.File, baseline[example.ID], status, message)
			}
		})
	}
	writeJSON := func(name string, value any) {
		t.Helper()
		data, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, name), append(data, '\n'), 0644); err != nil {
			t.Fatal(err)
		}
	}
	// Always write the complete report before reporting baseline differences.
	writeJSON("report.json", report)
	t.Logf("%d unique examples / %d occurrences: %v; details in %s/report.json", report.Examples, report.Occurrences, report.Counts, directory)
	if update {
		if len(observed) != len(manifest.Examples) {
			t.Fatal("refusing to update baseline from a filtered or incomplete corpus run")
		}
		if t.Failed() {
			t.Fatal("refusing to update baseline after corpus integrity failures")
		}
		writeJSON("baseline.json", observed)
	} else if !strict && !reflect.DeepEqual(baseline, observed) {
		t.Errorf("%s baseline changed (including added/removed examples); inspect report.json, then use -%s-update to accept", directory, directory)
	}
}
