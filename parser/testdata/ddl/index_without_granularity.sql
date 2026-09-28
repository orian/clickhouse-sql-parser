CREATE TABLE t (id UInt64, s String, INDEX idx s TYPE text(tokenizer = 'splitByNonAlpha')) ENGINE = MergeTree ORDER BY id;
CREATE TABLE t (a Int32, INDEX i a TYPE minmax, INDEX j (a, a * 2) TYPE set(100) GRANULARITY 4) ENGINE = MergeTree ORDER BY a;
ALTER TABLE t ADD INDEX i a TYPE minmax;
ALTER TABLE t ADD INDEX IF NOT EXISTS i a TYPE bloom_filter(0.01) AFTER j;
