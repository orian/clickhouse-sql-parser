
    CREATE TABLE test (a UInt64, b UInt64) ENGINE = MergeTree() ORDER BY () AS SELECT number, intDiv(number, 10) FROM numbers_mt(5e7);
  