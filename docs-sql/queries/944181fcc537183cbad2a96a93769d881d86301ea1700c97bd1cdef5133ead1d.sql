SELECT key_values['default'] AS default_disk_total
FROM system.asynchronous_metrics
WHERE metric = 'DiskTotal'
