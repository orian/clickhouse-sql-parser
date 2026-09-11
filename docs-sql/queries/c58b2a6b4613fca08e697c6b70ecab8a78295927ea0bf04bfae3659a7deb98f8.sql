server_id = 1               -- 或更大；不能为 0
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- 10.5.0 中引入
expire_logs_days = 1        -- 或更大；0 表示日志将永久保留
