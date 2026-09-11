CREATE TABLE memory (i UInt32) ENGINE = Memory SETTINGS min_bytes_to_keep = 4096, max_bytes_to_keep = 16384;

/* 1. testing oldest block doesn't get deleted due to min-threshold - 3000 rows */
INSERT INTO memory SELECT * FROM numbers(0, 1600); -- 8'192 bytes

/* 2. adding block that doesn't get deleted */
INSERT INTO memory SELECT * FROM numbers(1000, 100); -- 1'024 bytes

/* 3. testing oldest block gets deleted - 9216 bytes - 1100 */
INSERT INTO memory SELECT * FROM numbers(9000, 1000); -- 8'192 bytes

/* 4. checking a very large block overrides all */
INSERT INTO memory SELECT * FROM numbers(9000, 10000); -- 65'536 bytes

SELECT total_bytes, total_rows FROM system.tables WHERE name = 'memory' AND database = currentDatabase();
