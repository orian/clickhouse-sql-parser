        CREATE TABLE uk_price_paid
        (
        price      UInt32,
        date       Date,
        postcode   LowCardinality(String),
        type       Enum8('terraced' = 1, 'semi-detached' = 2, 'detached' = 3, 'flat' = 4, 'other' = 0),
        is_new     UInt8,
        duration   Enum8('freehold' = 1, 'leasehold' = 2, 'unknown' = 0),
        addr1      String,
        addr2      String,
        street     LowCardinality(String),
        locality   LowCardinality(String),
        town       LowCardinality(String),
        district   LowCardinality(String),
        county     LowCardinality(String)
        )
        ENGINE = MergeTree
        PARTITION BY toYYYYMM(date)
        ORDER BY (postcode, addr1, addr2);
