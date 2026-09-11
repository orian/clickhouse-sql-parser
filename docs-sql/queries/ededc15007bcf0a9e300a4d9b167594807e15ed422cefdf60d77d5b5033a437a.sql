  CREATE TABLE table(
      s String,
      -- Этот индекс использует 'bitpacking', переопределяя заданное ниже значение по умолчанию на уровне таблицы:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- Этот индекс наследует 'none' из настройки таблицы:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
