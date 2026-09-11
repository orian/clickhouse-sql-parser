  CREATE TABLE table(
      s String,
      -- Este índice usa 'bitpacking', substituindo o padrão no nível da tabela abaixo:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- Este índice herda 'none' da configuração da tabela:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
