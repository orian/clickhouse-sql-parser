
        CREATE TABLE t_text_index_merge_multiblock_intersecting_{codec}
        (
            id UInt64,
            s String,
            INDEX idx_s (s) TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = '{codec}', posting_list_block_size = 8192)
        )
        ENGINE = MergeTree ORDER BY id
    