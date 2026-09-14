SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- 설정이 비활성화된 상태에서도 이 구문은 정상적으로 실행됩니다.
-- 테이블이 이미 존재하므로 삽입이 차단되지 않습니다.
INSERT INTO events VALUES (1, '14:30:25');

-- 단, Time 유형으로 새 테이블을 생성하면 실패합니다.
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
