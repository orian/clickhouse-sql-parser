-- 更改调度计划，删除现有设置和依赖项。
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- 更改调度计划并调整重试行为。
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- 在更改周期的同时保留依赖项。
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- 通过省略 `DEPENDS ON` 来删除依赖项。
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
