    -- from clickhouse-01
    SELECT count(*)
    FROM uk.uk_price_paid_local
    --    ┌──count()─┐
    -- 1. │ 15107353 │ -- 15.11 million
    --    └──────────┘

    --from clickhouse-02
    SELECT count(*)
    FROM uk.uk_price_paid_local
    --    ┌──count()─┐
    -- 1. │ 15105202 │ -- 15.11 million
    --    └──────────┘
