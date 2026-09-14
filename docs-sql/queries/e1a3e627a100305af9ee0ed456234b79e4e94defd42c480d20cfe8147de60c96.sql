    CREATE MATERIALIZED VIEW otel_traces_1m_mv_backfill
    TO otel_traces_1m
    AS
    SELECT
        toStartOfMinute(Timestamp) AS Timestamp,
        ServiceName,
        StatusCode,
        count() AS count,
        avgState(Duration) AS avg__Duration,
        maxSimpleState(Duration) AS max__Duration,
        quantilesState(0.95, 0.99)(Duration) AS quantiles__Duration
    FROM otel_traces_backfill
    GROUP BY
        Timestamp,
        ServiceName,
        StatusCode;
