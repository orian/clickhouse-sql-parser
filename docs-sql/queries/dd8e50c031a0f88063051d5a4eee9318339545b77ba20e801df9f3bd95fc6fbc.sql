    INSERT INTO otel_traces
    SELECT * FROM file('~/data/data-otel-traces.csv', CSVWithNames);
