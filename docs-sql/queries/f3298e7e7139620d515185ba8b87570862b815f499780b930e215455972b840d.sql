CREATE TABLE dynamic_events
(
    `id`   UInt64,
    `ts`   DateTime DEFAULT now(),
    `data` JSON(
        max_dynamic_paths = 512,
        `event_type` LowCardinality(String),
        `version` UInt8
    )
)
ENGINE = MergeTree
ORDER BY (data.event_type, ts)
