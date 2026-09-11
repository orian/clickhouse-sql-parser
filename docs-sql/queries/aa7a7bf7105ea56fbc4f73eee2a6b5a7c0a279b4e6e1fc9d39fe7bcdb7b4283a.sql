-- スケジュールを変更し、既存の設定と依存関係を削除する。
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- スケジュールを変更し、再試行の動作を調整する。
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- 期間を変更しながら依存関係を維持する。
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- `DEPENDS ON` を省略して依存関係を削除する。
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
