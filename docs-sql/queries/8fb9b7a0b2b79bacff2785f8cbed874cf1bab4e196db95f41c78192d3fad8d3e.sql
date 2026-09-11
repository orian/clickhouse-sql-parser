CREATE TABLE test (
  a Array(Int32),
  b Array(String),
  c Array(DateTime)
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO test VALUES ([], [], []);

SELECT emptyArrayToSingle(a), emptyArrayToSingle(b), emptyArrayToSingle(c) FROM test;
