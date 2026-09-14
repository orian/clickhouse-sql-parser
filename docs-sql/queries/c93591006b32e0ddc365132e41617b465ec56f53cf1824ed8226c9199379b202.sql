        SELECT
            'uk_price_paid' AS table,
            count() AS rows
        FROM uk_price_paid
        UNION ALL
        SELECT
            'uk_price_paid_by_town' AS table,
            count() AS rows
        FROM uk_price_paid_by_town;
