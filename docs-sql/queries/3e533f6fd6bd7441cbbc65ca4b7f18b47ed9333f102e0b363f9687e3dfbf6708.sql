        -- 지난 1시간 동안 nyc_taxi 데이터베이스에서 실행 시간이 긴 쿼리 상위 5개 찾기
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
