  CREATE TABLE table(
      s String,
      -- 이 인덱스는 'bitpacking'을 사용하며, 아래의 테이블 수준 기본값을 재정의합니다:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- 이 인덱스는 테이블 설정에서 'none'을 상속합니다:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
