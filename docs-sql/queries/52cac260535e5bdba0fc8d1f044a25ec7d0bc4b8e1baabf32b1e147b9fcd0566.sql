        CREATE TABLE otel_traces
        (
            Timestamp          DateTime64(9),
            TraceId            String,
            SpanId             String,
            ParentSpanId       String,
            SpanName           LowCardinality(String),
            SpanKind           LowCardinality(String),
            ServiceName        LowCardinality(String),
            Duration           UInt64,
            StatusCode         LowCardinality(String),
            ResourceAttributes Map(LowCardinality(String), String),
            SpanAttributes     Map(LowCardinality(String), String)
        )
        ENGINE = MergeTree()
        ORDER BY (ServiceName, SpanName, toUnixTimestamp(Timestamp));
