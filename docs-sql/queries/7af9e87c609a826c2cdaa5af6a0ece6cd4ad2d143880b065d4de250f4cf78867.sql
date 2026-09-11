server_id = 1               -- 1 이상; 0이 아닌 값
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- 10.5.0에서 도입됨
expire_logs_days = 1        -- 1 이상; 0으로 설정하면 로그가 영구 보존됨
