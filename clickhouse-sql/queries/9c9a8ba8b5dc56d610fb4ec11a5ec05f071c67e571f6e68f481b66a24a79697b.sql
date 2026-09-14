
        CREATE TABLE t_in_clustered (k UInt64, u UUID, v UInt64) ENGINE = MergeTree ORDER BY (k, u)
    