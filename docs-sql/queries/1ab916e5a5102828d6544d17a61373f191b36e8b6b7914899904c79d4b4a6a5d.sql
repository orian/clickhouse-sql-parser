SHOW CREATE TABLE cookies;

SHOW CREATE TABLE cookies

Query id: 248ec8e2-5bce-45b3-97d9-ed68edf445a5

┌─statement────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ CREATE TABLE default.cookies
(
    `id` String,
    `timestamp` DateTime
)
ENGINE = MergeTree
ORDER BY id
SETTINGS index_granularity = 8192 │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

