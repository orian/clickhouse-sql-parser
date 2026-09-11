
        CREATE TABLE IF NOT EXISTS codec_{seq_type}_{type}_{codec}
        (n {type} CODEC({codec}))
        ENGINE = MergeTree
        ORDER BY tuple()
        SETTINGS merge_selector_base=1000
    