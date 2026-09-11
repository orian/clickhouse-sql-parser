SELECT quantilePrometheusHistogram(bucket_upper_bound, cumulative_bucket_value)
FROM VALUES('bucket_upper_bound Float64, cumulative_bucket_value UInt64', (0, 6), (0.5, 11), (1, 14), (inf, 19));
