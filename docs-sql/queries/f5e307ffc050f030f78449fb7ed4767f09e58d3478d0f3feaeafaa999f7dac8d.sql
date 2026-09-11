    INSERT INTO uk_price_paid
    SELECT
        toUInt32(price),
        date,
        postcode,
        transform(type, ['T', 'S', 'D', 'F', 'O'],
            ['terraced', 'semi-detached', 'detached', 'flat', 'other'], 'other') AS type,
        if(is_new = 'Y', 1, 0) AS is_new,
        transform(duration, ['F', 'L', 'U'],
            ['freehold', 'leasehold', 'unknown'], 'unknown') AS duration,
        addr1,
        addr2,
        street,
        locality,
        town,
        district,
        county
    FROM s3(
    'https://learn-clickhouse.s3.us-east-2.amazonaws.com/uk_property_prices/uk_prices.csv.zst'
    );
