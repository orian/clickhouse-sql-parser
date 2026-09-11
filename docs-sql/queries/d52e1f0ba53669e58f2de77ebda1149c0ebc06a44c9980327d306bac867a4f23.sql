-- Map 유형의 컬럼에 대해:
INDEX map_key_index mapKeys(map_column) TYPE bloom_filter
INDEX map_value_index mapValues(map_column) TYPE bloom_filter

-- JSON 유형의 컬럼에 대해:
INDEX json_paths_index JSONAllPaths(json_column) TYPE bloom_filter

-- Tuple 유형의 컬럼에 대해:
INDEX tuple_1_index tuple_column.1 TYPE bloom_filter
INDEX tuple_2_index tuple_column.2 TYPE bloom_filter

-- Nested 유형의 컬럼에 대해:
INDEX nested_1_index col.nested_col1 TYPE bloom_filter
INDEX nested_2_index col.nested_col2 TYPE bloom_filter
