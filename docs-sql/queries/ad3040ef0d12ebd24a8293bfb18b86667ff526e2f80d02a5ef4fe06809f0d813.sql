SET skip_redundant_aliases_in_udf = 0;
CREATE FUNCTION IF NOT EXISTS test_03274 AS ( x ) -> ((x + 1 as y, y + 2));

EXPLAIN SYNTAX SELECT test_03274(4 + 2);
