DROP TABLE IF EXISTS test;
CREATE TABLE test (s LowCardinality(String)) ENGINE = Memory;

-- create two parts:

INSERT INTO test VALUES ('ab'), ('cd'), ('ab'), ('ab'), ('df');
INSERT INTO test VALUES ('ef'), ('cd'), ('ab'), ('cd'), ('ef');

-- the order the blocks are read in is not defined, so sort the result to make it reproducible:

SELECT s, lowCardinalityIndices(s) AS index FROM test ORDER BY s, index;
