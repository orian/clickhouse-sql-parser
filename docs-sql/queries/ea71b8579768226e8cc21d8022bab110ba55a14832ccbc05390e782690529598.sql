ALTER TABLE system.session_log_1
    MODIFY COLUMN interface Enum8('TCP' = 1, 'HTTP' = 2, 'gRPC' = 3, 'MySQL' = 4, 'PostgreSQL' = 5, 'Local' = 6, 'TCP_Interserver' = 7, 'Prometheus' = 8, 'Background' = 9, 'ArrowFlight' = 10);
