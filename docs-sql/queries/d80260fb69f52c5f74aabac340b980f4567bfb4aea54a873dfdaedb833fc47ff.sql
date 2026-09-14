        -- 查找过去 1 小时内 nyc_taxi 数据库中耗时最长的 5 个查询
        SELECT
            normalized_query_hash,
            type,
            event_time,
            query_duration_ms,
            query,
            read_rows,
            tables
        FROM merge('system', '^query_log')
        WHERE has(databases, 'nyc_taxi')
          AND event_time >= now() - INTERVAL 1 HOUR
          AND type = 'QueryFinish'
          AND is_initial_query = 1
          AND query_kind = 'Select'
        ORDER BY query_duration_ms DESC
        LIMIT 1 BY normalized_query_hash
        LIMIT 5
        FORMAT VERTICAL
