BACKUP TABLE trips
TO S3('http://seaweedfs:8333/clickhouse/backups/trips1', 'your_access_key_id', 'your_secret_access_key');

--- DROP TABLE trips;

RESTORE TABLE trips
FROM S3('http://seaweedfs:8333/clickhouse/backups/trips1', 'your_access_key_id', 'your_secret_access_key');
