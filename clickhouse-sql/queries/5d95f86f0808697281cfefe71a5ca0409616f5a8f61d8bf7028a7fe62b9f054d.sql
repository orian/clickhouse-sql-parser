
    CREATE TABLE text_index_build (text String, INDEX idx_text (text) TYPE text(tokenizer = splitByNonAlpha)) ENGINE = MergeTree ORDER BY tuple();
  