        ALTER TABLE otel_traces
            ADD COLUMN HostName String
            MATERIALIZED ResourceAttributes['host.name'];
