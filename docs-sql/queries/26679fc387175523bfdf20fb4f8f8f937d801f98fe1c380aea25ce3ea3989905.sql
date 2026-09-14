CREATE TABLE table(
    s String,
    -- This index uses 'bitpacking', overriding the table-level default below:
    INDEX idx_a s TYPE text(tokenizer = 'splitByNonAlpha', posting_list_codec = 'bitpacking'),
    -- This index inherits 'none' from the table setting:
    INDEX idx_b lower(s) TYPE text(tokenizer = 'splitByNonAlpha'))
ENGINE = MergeTree()
ORDER BY tuple()
SETTINGS text_index_posting_list_codec = 'none';
