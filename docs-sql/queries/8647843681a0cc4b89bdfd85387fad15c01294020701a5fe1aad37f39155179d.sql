CREATE TABLE docs
(
    id UInt64,
    body String,
    INDEX body_idx body TYPE text(tokenizer = 'japanese')
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO docs VALUES (1, '日本語の形態素解析エンジン'), (2, 'これはテストの文章です');

SELECT id FROM docs WHERE hasAllTokens(body, '形態 解析', 'japanese') ORDER BY id;
