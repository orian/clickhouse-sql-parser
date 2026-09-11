-- ✗ Column names in SELECT (use Identifier carefully)
SELECT {column: Identifier} FROM users;  -- Limited support

-- ✗ Arbitrary SQL fragments
SELECT * FROM users {where_clause: String};  -- NOT SUPPORTED

-- ✗ ALTER TABLE statements
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- NOT SUPPORTED

-- ✗ Multiple statements
{statements: String};  -- NOT SUPPORTED
