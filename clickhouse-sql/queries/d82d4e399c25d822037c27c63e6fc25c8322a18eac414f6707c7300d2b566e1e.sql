
    CREATE TABLE tab_filter_{drop_pct} (
      id UInt64,
      text String,
      INDEX idx (text) TYPE text(tokenizer = splitByNonAlpha, postprocessor = if(text IN ('a', 'an', 'and', 'are', 'as', 'at', 'be', 'but', 'by', 'for', 'if', 'in', 'into', 'is', 'it', 'no', 'not', 'of', 'on', 'or', 'such', 'that', 'the', 'their', 'then', 'there', 'these', 'they', 'this', 'to', 'was', 'will', 'with'), '', text))
    )
    ENGINE = MergeTree
    ORDER BY id
  