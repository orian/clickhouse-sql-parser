package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParser_RejectsInvalidAccessControl pins role and GRANT statements that
// ClickHouse 26.8 rejects and the parser used to accept (#50), next to the
// valid forms that still parse.
func TestParser_RejectsInvalidAccessControl(t *testing.T) {
	for _, sql := range []string{
		"ALTER ROLE r1 RENAME TO r2, r3 RENAME TO r4",
		"CREATE ROLE r1 ON CLUSTER c1, r2",
		"CREATE ROLE r1 ON CLUSTER c1, r2 ON CLUSTER c2",
		"GRANT SELECT(x, y) ON db.t TO john WITH GRANT OPTION WITH ADMIN OPTION",
		"GRANT SELECT ON db.t TO john WITH ADMIN OPTION",
		"GRANT SELECT(x, y) ON db.* TO john",
		"GRANT SELECT(x, y) ON *.* TO john",
		"GRANT SELECT ON *.table TO john",
		"GRANT ADMIN OPTION ON *.* TO r",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.Error(t, err)
		})
	}
	for _, sql := range []string{
		"ALTER ROLE r1 ON CLUSTER c RENAME TO r2",
		"ALTER ROLE r1, r2 SETTINGS max_memory_usage = 1",
		"CREATE ROLE r1, r2 ON CLUSTER c",
		"GRANT SELECT(x, y) ON db.t TO john WITH GRANT OPTION WITH REPLACE OPTION",
		"GRANT SELECT(x) ON t TO john",
		"GRANT SELECT ON db.* TO john",
		"GRANT SELECT ON *.* TO john",
	} {
		t.Run(sql, func(t *testing.T) {
			_, err := NewParser(sql).ParseStmts()
			require.NoError(t, err)
		})
	}
}
