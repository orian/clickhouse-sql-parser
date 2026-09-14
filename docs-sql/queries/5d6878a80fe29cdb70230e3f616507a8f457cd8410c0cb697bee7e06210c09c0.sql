    ALTER TABLE otel_logs
      ADD COLUMN ServiceVersion LowCardinality(String)
      MATERIALIZED LogAttributes['service.version'];
