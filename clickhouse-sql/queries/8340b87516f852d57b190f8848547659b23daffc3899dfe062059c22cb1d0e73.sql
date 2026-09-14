
        CREATE TABLE text_index_lazy_wide
        (
            id UInt64,
            s String,
            INDEX idx s TYPE text(tokenizer = splitByNonAlpha, posting_list_codec = 'bitpacking') GRANULARITY 100000000
        )
        ENGINE = MergeTree ORDER BY id
    