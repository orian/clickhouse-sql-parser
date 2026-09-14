    ALTER TABLE uk_price_paid
        ADD PROJECTION uk_price_paid_by_town
        (
            SELECT town, date, price, type
            ORDER BY (town, date)
        );
