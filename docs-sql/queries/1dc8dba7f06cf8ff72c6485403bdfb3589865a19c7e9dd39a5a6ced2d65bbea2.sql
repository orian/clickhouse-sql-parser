-- `index_granularity` is lowered to 8 here only to keep the example small enough to follow.
-- Do not change it in production: the default of 8192 is what makes the index sparse and cheap,
-- and a small value makes the index large and slows queries down.
CREATE TABLE ontime (FlightDate Date, Carrier String)
ENGINE = MergeTree ORDER BY FlightDate
SETTINGS index_granularity = 8;

-- Ten flights on each of four days, so a granule of eight rows spans more than one day.
INSERT INTO ontime SELECT toDate('2025-09-14') + intDiv(number, 10), 'AA' FROM numbers(40);

-- The granules that hold the rows of 2025-09-15 also hold rows of the neighbouring days,
-- and `indexHint` returns all of them.
SELECT FlightDate AS k, count() FROM ontime WHERE indexHint(k = '2025-09-15') GROUP BY k ORDER BY k ASC;
