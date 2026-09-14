  CREATE TABLE table(
      s String,
      -- 此索引使用 'bitpacking'，覆盖下面的表级默认值：
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- 此索引从表设置继承 'none'：
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
