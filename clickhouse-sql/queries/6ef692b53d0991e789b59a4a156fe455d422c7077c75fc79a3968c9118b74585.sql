
        CREATE TABLE IF NOT EXISTS delta_decode_{pattern}_{type}
        (n {type} CODEC(Delta))
        ENGINE = MergeTree
        ORDER BY tuple()
    