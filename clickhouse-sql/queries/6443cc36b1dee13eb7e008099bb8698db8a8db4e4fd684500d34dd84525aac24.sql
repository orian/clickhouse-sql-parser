
        SELECT * FROM file('test_pq', Parquet, 'map_tuple_column Map(String, Tuple(a Nullable(String)))') Format Null
    