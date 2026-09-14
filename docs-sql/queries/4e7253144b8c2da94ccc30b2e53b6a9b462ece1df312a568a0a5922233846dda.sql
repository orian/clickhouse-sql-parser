-- Изменить расписание, удалить существующие настройки и зависимости.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- Изменить расписание и настроить поведение повторных попыток.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- Сохранить зависимость при изменении периода.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- Удалить зависимость, опустив `DEPENDS ON`.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
