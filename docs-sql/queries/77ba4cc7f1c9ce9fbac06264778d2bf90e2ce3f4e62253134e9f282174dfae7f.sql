-- 对于 Map 类型的列：
INDEX map_key_index mapKeys(map_column) TYPE bloom_filter
INDEX map_value_index mapValues(map_column) TYPE bloom_filter

-- 对于 JSON 类型的列：
INDEX json_paths_index JSONAllPaths(json_column) TYPE bloom_filter

-- 对于 Tuple 类型的列：
INDEX tuple_1_index tuple_column.1 TYPE bloom_filter
INDEX tuple_2_index tuple_column.2 TYPE bloom_filter

-- 对于 Nested 类型的列：
INDEX nested_1_index col.nested_col1 TYPE bloom_filter
INDEX nested_2_index col.nested_col2 TYPE bloom_filter
