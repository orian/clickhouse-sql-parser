ALTER TABLE otel_v2.otel_traces
MATERIALIZE PROJECTION prj_traceid_time;
