
        SELECT * FROM file('test_pq_index', Parquet, 'array_tuple_column Array (Tuple(a Nullable(String)))') Format Null
    