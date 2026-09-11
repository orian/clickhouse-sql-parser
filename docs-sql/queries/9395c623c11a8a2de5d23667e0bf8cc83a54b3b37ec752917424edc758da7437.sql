DROP TABLE IF EXISTS doc;

CREATE TABLE doc (
    id UInt32,
    msg String,
    INDEX idx(msg) TYPE text(tokenizer = splitByString(['()', '\\']))
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO doc VALUES (1, '()a,\\bc()d'), (2, '()\\a()bc\\d'), (3, ',()a\\,bc,(),d,');

SELECT count() FROM doc WHERE hasAnyTokens(msg, ['a', 'd']);
