        CREATE TABLE uk_price_paid_by_town
        (
            town       LowCardinality(String),
            date       Date,
            price      UInt32,
            type       Enum8('terraced' = 1, 'semi-detached' = 2, 'detached' = 3, 'flat' = 4, 'other' = 0)
        )
        ENGINE = MergeTree
        PARTITION BY toYYYYMM(date)
        ORDER BY (town, date);
