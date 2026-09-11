CREATE TABLE trips (pickup_datetime DateTime('UTC'), trip_distance Float64, total_amount Float64) ENGINE = Memory;

-- A fare of 3, plus 2.5 for every unit of distance.
INSERT INTO trips
SELECT toDateTime('2020-01-01 00:00:00', 'UTC') + number * 60, number % 10 + 1, 2.5 * (number % 10 + 1) + 3
FROM numbers(1000);

-- One model per year of the data.
CREATE TABLE models ENGINE = Memory AS
SELECT
    toYear(pickup_datetime) AS year,
    stochasticLinearRegressionState(0.01, 0.0, 10, 'SGD')(total_amount, trip_distance) AS model
FROM trips
GROUP BY year;

SELECT
    trip_distance,
    round(evalMLMethod(model, trip_distance), 2) AS predicted,
    total_amount
FROM trips
LEFT JOIN models ON year = toYear(pickup_datetime)
ORDER BY pickup_datetime
LIMIT 5
