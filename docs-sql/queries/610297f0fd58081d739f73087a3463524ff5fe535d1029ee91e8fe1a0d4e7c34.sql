CREATE TABLE data
ENGINE = Memory AS
SELECT
    10 AS value,
    toDateTime('2020-01-01') + (3600 * number) AS time
FROM numbers_mt(10);

-- Calculate timeunit using intDiv
SELECT
    value,
    time,
    exponentialMovingAverage(1)(value, intDiv(toUInt32(time), 3600)) OVER (ORDER BY time ASC) AS res,
    intDiv(toUInt32(time), 3600) AS timeunit
FROM data
ORDER BY time ASC
