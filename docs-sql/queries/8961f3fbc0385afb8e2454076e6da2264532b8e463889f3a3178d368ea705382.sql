SET PERSIST log_bin = ON;
SET PERSIST binlog_format = ROW;
SET PERSIST binlog_row_image = FULL;
SET PERSIST binlog_row_metadata = FULL;
SET PERSIST binlog_expire_logs_seconds = 86400;
