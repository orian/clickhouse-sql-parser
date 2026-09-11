
        CREATE TABLE codec_{seq_type}_{type}_{codec} (n {type} CODEC({codec}))
        ENGINE = MergeTree PARTITION BY tuple() ORDER BY tuple()
        SETTINGS parts_to_delay_insert = 5000, parts_to_throw_insert = 5000, auto_statistics_types='';
    