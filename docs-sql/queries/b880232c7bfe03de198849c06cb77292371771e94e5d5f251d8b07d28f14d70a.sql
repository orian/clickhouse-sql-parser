-- 스케줄을 변경하고 기존 설정 및 의존성을 삭제합니다.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE;

-- 스케줄을 변경하고 재시도 동작을 조정합니다.
ALTER TABLE rmv MODIFY REFRESH EVERY 30 MINUTE
SETTINGS refresh_retries = 5,
         refresh_retry_initial_backoff_ms = 500,
         refresh_retry_max_backoff_ms = 60000;

-- 주기를 변경하면서 의존성을 유지합니다.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR DEPENDS ON other_rmv;

-- `DEPENDS ON`을 생략하여 의존성을 삭제합니다.
ALTER TABLE rmv MODIFY REFRESH EVERY 6 HOUR;
