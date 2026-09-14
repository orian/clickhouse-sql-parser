
    CREATE TABLE tab
    (
        id     UInt64,
        message String,
        INDEX  idx(message) TYPE text(tokenizer = 'splitByNonAlpha')
    )
    ENGINE = MergeTree
    ORDER BY id
  