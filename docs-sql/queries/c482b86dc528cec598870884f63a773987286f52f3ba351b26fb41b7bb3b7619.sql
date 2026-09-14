  CREATE TABLE table(
      s String,
      -- Cet index utilise 'bitpacking', en remplaçant la valeur par défaut définie au niveau de la table ci-dessous :
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- Cet index hérite de 'none' à partir du paramètre de table :
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
