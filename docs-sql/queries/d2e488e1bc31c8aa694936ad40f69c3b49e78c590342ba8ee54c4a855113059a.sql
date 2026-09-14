    ALTER TABLE otel_traces ADD INDEX idx_duration Duration TYPE minmax GRANULARITY 1;
    ALTER TABLE otel_traces MATERIALIZE INDEX idx_duration;
