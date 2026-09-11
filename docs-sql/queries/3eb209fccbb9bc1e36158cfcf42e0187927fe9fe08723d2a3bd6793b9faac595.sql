CREATE DICTIONARY cache_dict (id UInt64, data String) PRIMARY KEY id
SOURCE(CLICKHOUSE(TABLE 'cache_src')) LIFETIME(MIN 0 MAX 900) LAYOUT(CACHE(SIZE_IN_CELLS 1000));

-- nothing is cached yet, so nothing comes back and the source is not queried
SELECT count() FROM cache_dict WHERE id IN (1, 2, 3);
0

-- looking the keys up populates the cache
SELECT dictGet('cache_dict', 'data', toUInt64(number + 1)) FROM numbers(3);

-- and now the same read sees them
SELECT count() FROM cache_dict WHERE id IN (1, 2, 3);
3
