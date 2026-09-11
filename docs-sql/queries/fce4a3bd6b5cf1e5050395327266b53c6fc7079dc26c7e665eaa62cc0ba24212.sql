  CREATE TABLE table(
      s String,
      -- يستخدم هذا الفهرس 'bitpacking'، متجاوزًا الإعداد الافتراضي على مستوى الجدول أدناه:
      INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
      -- يرث هذا الفهرس 'none' من إعداد الجدول:
      INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
  ENGINE = MergeTree()
  ORDER BY tuple()
  SETTINGS text_index_posting_list_codec = 'none';
