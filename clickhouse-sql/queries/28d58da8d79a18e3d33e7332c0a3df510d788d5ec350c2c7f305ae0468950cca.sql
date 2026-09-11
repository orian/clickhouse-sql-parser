
    CREATE TABLE tab
    (
        id   UInt64,
        name String,
        INDEX idx(name) TYPE text(tokenizer = array)
    )
    ENGINE = MergeTree
    ORDER BY id
  