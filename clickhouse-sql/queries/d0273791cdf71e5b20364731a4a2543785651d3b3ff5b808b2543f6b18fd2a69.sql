
        CREATE TABLE permute_cache_src
        (
            c1 String,
            c2 String,
            c3 String,
            c4 String,
            c5 String,
            ts DateTime64(6),
            payload String,
            value UInt64
        )
        ENGINE = Memory
    