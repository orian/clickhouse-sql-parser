
        CREATE TABLE t_text_index_merge_vocabulary_intersecting_{codec}
        (
            id UInt64,
            s String,
            INDEX idx_s (s) TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = '{codec}')
        )
        ENGINE = MergeTree ORDER BY id
    