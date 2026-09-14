
        CREATE TABLE t_subcolumns (a Array(UInt64), s Nullable(String), m Map(String, UInt64)) ENGINE = MergeTree ORDER BY tuple()
    