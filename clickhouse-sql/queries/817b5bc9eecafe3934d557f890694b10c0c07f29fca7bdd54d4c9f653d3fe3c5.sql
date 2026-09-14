
    CREATE TABLE tab_filter_not_in (
      id UInt64,
      text String,
      INDEX idx (text) TYPE text(tokenizer = splitByNonAlpha, postprocessor = if(text NOT IN ('fox', 'cat', 'dog', 'bird', 'tree', 'river', 'stone'), '', text))
    )
    ENGINE = MergeTree
    ORDER BY id
  