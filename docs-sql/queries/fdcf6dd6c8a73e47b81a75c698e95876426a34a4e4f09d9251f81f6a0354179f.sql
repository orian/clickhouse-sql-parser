        INSERT INTO uk_price_paid_by_town
        SELECT
            town,
            date,
            price,
            type
        FROM uk_price_paid;
