package parser

import "flag"

// updateBeautifyFixtures controls whether TestBeautifyVisitor_Fixtures
// overwrites the .expected.sql files with the visitor's current output
// instead of asserting against them.
//
// Use it to re-pin fixtures after an intentional beautify change:
//
//	go test ./parser -run TestBeautifyVisitor_Fixtures -update-beautify
var updateBeautifyFixtures bool

func init() {
	flag.BoolVar(&updateBeautifyFixtures, "update-beautify", false,
		"overwrite .expected.sql files in parser/testdata/beautify/ with current BeautifyVisitor output")
}
