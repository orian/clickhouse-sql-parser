server_id = 1            -- 或更大；不能为 0
log_bin = ON
binlog_format = ROW      -- 默认值
binlog_row_image = FULL  -- 默认值
expire_logs_days = 1     -- 或更大；0 表示日志永久保留
