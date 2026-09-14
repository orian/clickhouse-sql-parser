CREATE TABLE posts
(
    Id UInt32,
    Title String,
    ViewCount UInt32,
    Version UInt32
)
ENGINE = ReplacingMergeTree(Version)
ORDER BY Id
