-- Inserts all 10 rows: `limit` applies to the INSERT statement, not to its SELECT
INSERT INTO t SETTINGS limit = 2 SELECT number FROM numbers(10);

-- Inserts 2 rows: `limit` applies to the SELECT itself
INSERT INTO t SELECT number FROM numbers(10) SETTINGS limit = 2;
