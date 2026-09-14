    -- clickhouse-01
    SELECT count(*)
    FROM uk.uk_price_paid_local

    --   ┌─count()─┐
    -- 1.│   10000 │
    --   └─────────┘

    -- clickhouse-02
    SELECT count(*)
    FROM uk.uk_price_paid_local

    --   ┌─count()─┐
    -- 1.│   10000 │
    --   └─────────┘
