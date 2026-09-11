          -- Consultas con mayor uso de memoria
          SELECT
              type,
              event_time,
              query_id,
              formatReadableSize(memory_usage) AS memory,
              ProfileEvents.Values[indexOf(ProfileEvents.Names, 'UserTimeMicroseconds')] AS userCPU,
              ProfileEvents.Values[indexOf(ProfileEvents.Names, 'SystemTimeMicroseconds')] AS systemCPU,
              (ProfileEvents['CachedReadBufferReadFromCacheMicroseconds']) / 1000000 AS FromCacheSeconds,
              (ProfileEvents['CachedReadBufferReadFromSourceMicroseconds']) / 1000000 AS FromSourceSeconds,
              normalized_query_hash
          FROM clusterAllReplicas('default', merge('system', '^query_log'))
          WHERE has(databases, 'nyc_taxi')
            AND type = 'QueryFinish'
            AND is_initial_query = 1
            AND query_kind = 'Select'
            AND event_time >= now() - INTERVAL 2 DAY
            AND user NOT ILIKE '%internal%'
          ORDER BY memory_usage DESC
          LIMIT 30
          SETTINGS skip_unavailable_shards = 1
