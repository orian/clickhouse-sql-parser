
        CREATE TABLE IF NOT EXISTS codec_{seq_type}_{type}_{codec} (n {type} CODEC({codec}))
        ENGINE = MergeTree PARTITION BY tuple()
        ORDER BY tuple();
    