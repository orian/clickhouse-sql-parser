SELECT
    hostName(),
    database,
    table,
    round(elapsed, 0) AS elapsed_seconds,
    round(progress, 4) AS progress_ratio,
    formatReadableTimeDelta((elapsed / progress) - elapsed) AS estimated_time_remaining,
    num_parts,
    result_part_name
FROM clusterAllReplicas(default, merge(system, '^merges'))
ORDER BY (elapsed / progress) - elapsed ASC
