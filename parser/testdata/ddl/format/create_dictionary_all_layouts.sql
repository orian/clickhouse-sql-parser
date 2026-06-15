-- Origin SQL:
-- Coverage for every dictionary LAYOUT documented by ClickHouse.
-- See clickhouse-docs/en/sql-reference/statements/create/dictionary/layouts/.

CREATE DICTIONARY test.l_flat (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(FLAT());
CREATE DICTIONARY test.l_flat_args (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(FLAT(INITIAL_ARRAY_SIZE 50000 MAX_ARRAY_SIZE 5000000));

CREATE DICTIONARY test.l_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED());
CREATE DICTIONARY test.l_hashed_args (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED(SHARDS 1 SHARD_LOAD_QUEUE_BACKLOG 10000 MAX_LOAD_FACTOR 0.5));
CREATE DICTIONARY test.l_sparse_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(SPARSE_HASHED(SHARDS 1));
CREATE DICTIONARY test.l_ck_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_HASHED(SHARDS 1));
CREATE DICTIONARY test.l_ck_sparse_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_SPARSE_HASHED());

CREATE DICTIONARY test.l_hashed_array (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED_ARRAY(SHARDS 1));
CREATE DICTIONARY test.l_ck_hashed_array (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_HASHED_ARRAY());

CREATE DICTIONARY test.l_range_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(RANGE_HASHED(range_lookup_strategy 'max'));
CREATE DICTIONARY test.l_range_hashed_empty (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(RANGE_HASHED());
CREATE DICTIONARY test.l_ck_range_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_RANGE_HASHED());

CREATE DICTIONARY test.l_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(CACHE(SIZE_IN_CELLS 1000000000));
CREATE DICTIONARY test.l_ck_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_CACHE(SIZE_IN_CELLS 1000));
CREATE DICTIONARY test.l_ssd_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(SSD_CACHE(BLOCK_SIZE 4096 FILE_SIZE 16777216 READ_BUFFER_SIZE 1048576 PATH '/var/lib/clickhouse/user_files/test_dict'));
CREATE DICTIONARY test.l_ck_ssd_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_SSD_CACHE(BLOCK_SIZE 4096 PATH '/var/lib/clickhouse/user_files/ck_dict'));

CREATE DICTIONARY test.l_direct (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(DIRECT());
CREATE DICTIONARY test.l_ck_direct (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_DIRECT());

CREATE DICTIONARY test.l_ip_trie (prefix String, asn UInt32) PRIMARY KEY prefix SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(IP_TRIE);
CREATE DICTIONARY test.l_ip_trie_args (prefix String, asn UInt32) PRIMARY KEY prefix SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(IP_TRIE(ACCESS_TO_KEY_FROM_ATTRIBUTES 1));

CREATE DICTIONARY test.l_regexp_tree (regexp String, v String) PRIMARY KEY regexp SOURCE(YAMLRegExpTree(PATH '/var/lib/clickhouse/user_files/regexp_tree.yaml')) LIFETIME(0) LAYOUT(REGEXP_TREE);

CREATE DICTIONARY test.l_polygon (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON(STORE_POLYGON_KEY_COLUMN 1));
CREATE DICTIONARY test.l_polygon_simple (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_SIMPLE());
CREATE DICTIONARY test.l_polygon_index_each (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_INDEX_EACH(MAX_DEPTH 5 MIN_INTERSECTIONS 1));
CREATE DICTIONARY test.l_polygon_index_cell (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_INDEX_CELL());


-- Format SQL:
CREATE DICTIONARY test.l_flat (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(FLAT());
CREATE DICTIONARY test.l_flat_args (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(FLAT(INITIAL_ARRAY_SIZE 50000 MAX_ARRAY_SIZE 5000000));
CREATE DICTIONARY test.l_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED());
CREATE DICTIONARY test.l_hashed_args (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED(SHARDS 1 SHARD_LOAD_QUEUE_BACKLOG 10000 MAX_LOAD_FACTOR 0.5));
CREATE DICTIONARY test.l_sparse_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(SPARSE_HASHED(SHARDS 1));
CREATE DICTIONARY test.l_ck_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_HASHED(SHARDS 1));
CREATE DICTIONARY test.l_ck_sparse_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_SPARSE_HASHED());
CREATE DICTIONARY test.l_hashed_array (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(HASHED_ARRAY(SHARDS 1));
CREATE DICTIONARY test.l_ck_hashed_array (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_HASHED_ARRAY());
CREATE DICTIONARY test.l_range_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(RANGE_HASHED(range_lookup_strategy 'max'));
CREATE DICTIONARY test.l_range_hashed_empty (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(RANGE_HASHED());
CREATE DICTIONARY test.l_ck_range_hashed (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_RANGE_HASHED());
CREATE DICTIONARY test.l_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(CACHE(SIZE_IN_CELLS 1000000000));
CREATE DICTIONARY test.l_ck_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_CACHE(SIZE_IN_CELLS 1000));
CREATE DICTIONARY test.l_ssd_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(SSD_CACHE(BLOCK_SIZE 4096 FILE_SIZE 16777216 READ_BUFFER_SIZE 1048576 PATH '/var/lib/clickhouse/user_files/test_dict'));
CREATE DICTIONARY test.l_ck_ssd_cache (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_SSD_CACHE(BLOCK_SIZE 4096 PATH '/var/lib/clickhouse/user_files/ck_dict'));
CREATE DICTIONARY test.l_direct (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(DIRECT());
CREATE DICTIONARY test.l_ck_direct (id UInt64, v String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(COMPLEX_KEY_DIRECT());
CREATE DICTIONARY test.l_ip_trie (prefix String, asn UInt32) PRIMARY KEY prefix SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(IP_TRIE());
CREATE DICTIONARY test.l_ip_trie_args (prefix String, asn UInt32) PRIMARY KEY prefix SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(IP_TRIE(ACCESS_TO_KEY_FROM_ATTRIBUTES 1));
CREATE DICTIONARY test.l_regexp_tree (regexp String, v String) PRIMARY KEY regexp SOURCE(YAMLRegExpTree(PATH '/var/lib/clickhouse/user_files/regexp_tree.yaml')) LIFETIME(0) LAYOUT(REGEXP_TREE());
CREATE DICTIONARY test.l_polygon (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON(STORE_POLYGON_KEY_COLUMN 1));
CREATE DICTIONARY test.l_polygon_simple (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_SIMPLE());
CREATE DICTIONARY test.l_polygon_index_each (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_INDEX_EACH(MAX_DEPTH 5 MIN_INTERSECTIONS 1));
CREATE DICTIONARY test.l_polygon_index_cell (key String, v String) PRIMARY KEY key SOURCE(CLICKHOUSE(host 'x' port 9000 table 't')) LIFETIME(0) LAYOUT(POLYGON_INDEX_CELL());
