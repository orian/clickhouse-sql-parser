        CREATE MATERIALIZED VIEW uk_price_paid_by_town_mv
        TO uk_price_paid_by_town
        AS SELECT
            town,
            date,
            price,
            type
        FROM uk_price_paid;
