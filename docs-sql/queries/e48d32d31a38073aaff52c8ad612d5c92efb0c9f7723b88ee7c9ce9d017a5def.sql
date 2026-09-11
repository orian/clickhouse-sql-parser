server_id = 1               -- 1以上であれば何でも可（0以外）
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- 10.5.0で導入
expire_logs_days = 1        -- 1以上を指定（0の場合、ログは永久に保持される）
