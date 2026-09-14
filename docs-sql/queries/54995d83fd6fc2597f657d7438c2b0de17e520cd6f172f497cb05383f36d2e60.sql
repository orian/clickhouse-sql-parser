SELECT
    count() AS num_rows,
    uniqExact(_file) AS num_files
FROM s3('https://coiled-datasets-rp.s3.us-east-1.amazonaws.com/1trc/measurements-100*.parquet', 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', headers('x-amz-request-payer' = 'requester'))

┌───num_rows─┬─num_files─┐
│ 1110000000 │       111 │
└────────────┴───────────┘

1 row in set. Elapsed: 3.089 sec. Processed 1.09 billion rows, 0.00 B (353.55 million rows/s., 0.00 B/s.)
Peak memory usage: 192.27 KiB.
