  CREATE TABLE table(
      s String,
      -- この索引は 'bitpacking' を使用し、以下のテーブルレベルのデフォルトを上書きします:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- この索引はテーブル設定から 'none' を継承します:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
