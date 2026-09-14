BACKUP TABLE system.workloads, TABLE system.resources ON CLUSTER '{cluster}' TO S3('https://...', 'key', 'secret');
RESTORE TABLE system.workloads, TABLE system.resources ON CLUSTER '{cluster}' FROM S3('https://...', 'key', 'secret');
