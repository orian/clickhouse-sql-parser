
        CREATE TABLE text_index_trivial_count (id UInt64, text String,
            INDEX idx text TYPE text(tokenizer = splitByNonAlpha) GRANULARITY 1)
        ENGINE = MergeTree ORDER BY id
    