CREATE TABLE tab(t Tuple(a UInt32, b Tuple(c String, d UInt32))) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO tab VALUES ((3, ('c', 4)));

SELECT flattenTuple(t) FROM tab;
