    SELECT *
    FROM s3(
    'https://learn-clickhouse.s3.us-east-2.amazonaws.com/uk_property_prices/uk_prices.csv.zst'
    )
    LIMIT 5;
