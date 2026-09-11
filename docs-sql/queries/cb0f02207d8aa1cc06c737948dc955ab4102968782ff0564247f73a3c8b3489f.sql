SELECT
    labels AS percentile,
    round(quantiles) AS views
FROM
(
    SELECT
        quantiles(0.999, 0.99, 0.95, 0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1)(view_count) AS quantiles,
        ['99.9th', '99th', '95th', '90th', '80th', '70th','60th', '50th', '40th', '30th', '20th', '10th'] AS labels
    FROM youtube
)
ARRAY JOIN
    quantiles,
    labels;
