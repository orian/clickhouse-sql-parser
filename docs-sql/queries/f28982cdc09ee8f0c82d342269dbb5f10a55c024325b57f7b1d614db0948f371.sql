WITH buckets AS (
  SELECT
    $__timeInterval(Timestamp) AS ts,
    count() AS bucket_count
  FROM $__sourceTable
  WHERE TimestampTime >= fromUnixTimestamp64Milli({startDateMilliseconds:Int64})
        - toIntervalSecond($__interval_s * 30) -- Buscar 30 intervalos anteriores
    AND TimestampTime < fromUnixTimestamp64Milli({endDateMilliseconds:Int64})
    AND SeverityText = 'error'
  GROUP BY ts
  ORDER BY ts
  WITH FILL
    FROM toDateTime(fromUnixTimestamp64Milli({startDateMilliseconds:Int64}))
    TO toDateTime(fromUnixTimestamp64Milli({endDateMilliseconds:Int64}))
    STEP toIntervalSecond($__interval_s)
),

anomaly_detection AS (
  SELECT
    ts,
    bucket_count,
    avg(bucket_count) OVER (
      ORDER BY ts ROWS BETWEEN 30 PRECEDING AND 1 PRECEDING
    ) AS previous_30_avg,
    stddevPop(bucket_count) OVER (
      ORDER BY ts ROWS BETWEEN 30 PRECEDING AND 1 PRECEDING
    ) AS previous_30_stddev,
    greatest(
      bucket_count - (previous_30_avg + 2 * previous_30_stddev), 0
    ) AS excess_error_count
  FROM buckets
)

SELECT ts, excess_error_count
FROM anomaly_detection
WHERE ts >= fromUnixTimestamp64Milli({startDateMilliseconds:Int64})
  AND ts < fromUnixTimestamp64Milli({endDateMilliseconds:Int64})
