DROP TABLE IF EXISTS log;

CREATE TABLE log (
    id UInt32,
    tags Array(String),
    attributes Map(String, String),
    INDEX idx_tags (tags) TYPE text(tokenizer = splitByNonAlpha),
    INDEX idx_attributes_keys mapKeys(attributes) TYPE text(tokenizer = array),
    INDEX idx_attributes_vals mapValues(attributes) TYPE text(tokenizer = array)
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO log VALUES
    (1, ['clickhouse', 'clickhouse cloud'], {'address': '192.0.0.1', 'log_level': 'INFO'}),
    (2, ['chdb'], {'embedded': 'true', 'log_level': 'DEBUG'});

SELECT count() FROM log WHERE hasAnyTokens(mapKeys(attributes), ['address', 'log_level']);
