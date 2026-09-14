        CREATE TABLE otel_http_status_counts
        (
            Timestamp    DateTime,
            ServiceName  LowCardinality(String),
            StatusCounts Map(String, UInt64)
        )
        ENGINE = MergeTree()
        ORDER BY (ServiceName, Timestamp);

        INSERT INTO otel_http_status_counts VALUES
            ('2025-12-26 10:00:00', 'cart-service',      {'2xx': 150, '4xx': 12, '5xx': 3}),
            ('2025-12-26 10:01:00', 'cart-service',      {'2xx': 200, '4xx': 8,  '5xx': 1}),
            ('2025-12-26 10:00:00', 'inventory-service', {'2xx': 90,  '4xx': 5}),
            ('2025-12-26 10:01:00', 'inventory-service', {'2xx': 110, '4xx': 3,  '5xx': 2}),
            ('2025-12-26 10:00:00', 'payment-service',   {'2xx': 50,  '5xx': 10}),
            ('2025-12-26 10:01:00', 'payment-service',   {'2xx': 45,  '4xx': 2,  '5xx': 15});
