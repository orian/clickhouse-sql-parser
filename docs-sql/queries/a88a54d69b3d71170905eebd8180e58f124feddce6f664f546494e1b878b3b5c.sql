server_id = 1            -- 1以上であれば何でも可（0以外）
log_bin = ON
binlog_format = ROW      -- デフォルト値
binlog_row_image = FULL  -- デフォルト値
expire_logs_days = 1     -- 1以上；0を指定するとログが永久に保持される
