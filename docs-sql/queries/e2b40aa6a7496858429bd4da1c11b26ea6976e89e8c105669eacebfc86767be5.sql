server_id = 1            -- или больше; любое значение, кроме 0
log_bin = ON
binlog_format = ROW      -- значение по умолчанию
binlog_row_image = FULL  -- значение по умолчанию
expire_logs_days = 1     -- или больше; 0 означает, что журналы хранятся бессрочно
