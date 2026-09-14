
        SELECT tuple_column.a FROM file('test_pq_index', Parquet, 'tuple_column Tuple(a Nullable(String))') Format Null
    