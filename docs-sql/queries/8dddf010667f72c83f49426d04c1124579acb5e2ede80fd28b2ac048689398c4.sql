ALTER TABLE otel_v2.otel_traces
ADD PROJECTION prj_traceid_time
(
    SELECT *
    ORDER BY (TraceId, toDateTime(Timestamp))
);
