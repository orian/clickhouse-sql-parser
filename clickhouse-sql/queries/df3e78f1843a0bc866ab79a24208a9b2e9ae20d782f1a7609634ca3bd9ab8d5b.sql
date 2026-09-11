
        CREATE TABLE t_in_unclustered (k UInt64, u UUID, v UInt64) ENGINE = MergeTree ORDER BY tuple()
    