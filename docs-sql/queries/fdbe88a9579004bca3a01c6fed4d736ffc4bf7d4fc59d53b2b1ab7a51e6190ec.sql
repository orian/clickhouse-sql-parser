        -- 過去1時間にnyc_taxiデータベースで実行された、実行時間の長いクエリ上位5件を検索
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
