    CREATE TABLE IF NOT EXISTS  nginxdb.access_logs (
      message String
    )
    ENGINE = MergeTree()
    ORDER BY tuple()
