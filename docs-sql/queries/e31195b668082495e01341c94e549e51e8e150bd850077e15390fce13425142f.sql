CREATE TABLE sentiment_pooled (
id UInt64,
sentiment Float32
)
ENGINE = ExecutablePool(
    'sentiment.py',
    TabSeparated,
    (SELECT id, comment FROM hackernews WHERE id > 0 AND comment != '' LIMIT 20000)
)
SETTINGS
    pool_size = 4;
