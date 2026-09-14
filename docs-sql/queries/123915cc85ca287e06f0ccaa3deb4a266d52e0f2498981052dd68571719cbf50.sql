CREATE TABLE multi_metrics(
    date Date,
    browser_metrics Nested(
        browser String,
        impressions UInt32,
        clicks UInt32
    )
)
ENGINE = Memory;

INSERT INTO multi_metrics VALUES
    ('2000-01-01', ['Firefox', 'Chrome'], [100, 200], [10, 25]),
    ('2000-01-01', ['Chrome', 'Safari'], [150, 50], [20, 5]),
    ('2000-01-01', ['Firefox', 'Edge'], [80, 40], [8, 4]);

SELECT
    sumMappedArrays(browser_metrics.browser, browser_metrics.impressions, browser_metrics.clicks) AS result
FROM multi_metrics;
