SELECT metric, value, key_values
FROM system.asynchronous_metrics
WHERE metric IN ('MemoryResident', 'OSUserTime', 'DiskTotal', 'BlockReadBytes', 'NetworkReceiveBytes')
ORDER BY metric
