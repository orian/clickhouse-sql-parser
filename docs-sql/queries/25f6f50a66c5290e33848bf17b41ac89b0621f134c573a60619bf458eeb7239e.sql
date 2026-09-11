  CREATE TABLE table(
      s String,
      -- Este índice usa 'bitpacking', anulando el valor predeterminado a nivel de tabla que aparece abajo:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- Este índice hereda 'none' del ajuste de la tabla:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
