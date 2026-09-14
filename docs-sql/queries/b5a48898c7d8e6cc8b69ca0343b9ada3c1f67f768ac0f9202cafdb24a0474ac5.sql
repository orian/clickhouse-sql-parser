    SELECT *
    FROM gcs(
        'https://storage.googleapis.com/clickhouse-docs-example-bucket/epidemiology.csv',
        'GOOG1E...YOUR_ACCESS_KEY',
        'YOUR_SECRET_KEY',
        'CSVWithNames'
    );
