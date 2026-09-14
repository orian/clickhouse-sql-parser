CREATE TABLE fully_random_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(generateRandomStructure(7, 123), 1, 10)
LIMIT 1000;

DESCRIBE TABLE fully_random_table;
