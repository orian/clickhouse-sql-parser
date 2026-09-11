server_id = 1               -- или больше; любое значение, кроме 0
log_bin = ON
binlog_format = ROW
binlog_row_image = FULL
binlog_row_metadata = FULL  -- добавлено в 10.5.0
expire_logs_days = 1        -- или больше; значение 0 означает бессрочное хранение журналов
