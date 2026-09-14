
        SELECT tuple_column.a FROM file('test_pq', Parquet, 'tuple_column Tuple(a Nullable(String))') Format Null
    