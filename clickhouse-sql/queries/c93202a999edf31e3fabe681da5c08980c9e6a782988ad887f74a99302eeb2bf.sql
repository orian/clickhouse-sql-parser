
    CREATE TABLE t_text_index_analysis
    (
        s String,
        INDEX idx_s (s) TYPE text(tokenizer = 'splitByNonAlpha')
    )
    ENGINE = MergeTree
    ORDER BY tuple()
    SETTINGS index_granularity = 1
  