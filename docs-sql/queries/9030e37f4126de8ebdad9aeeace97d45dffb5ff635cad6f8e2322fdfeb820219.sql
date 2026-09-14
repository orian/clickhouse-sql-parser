-- tenant_1용 테이블 생성
CREATE TABLE tenant_1.events
(
    id UUID,                    -- 고유한 이벤트 ID
    type LowCardinality(String), -- 이벤트 유형
    timestamp DateTime,          -- 이벤트 타임스탬프
    user_id UInt32,               -- 이벤트를 발생시킨 사용자 ID
    data String,                 -- 이벤트 데이터
)
ORDER BY (timestamp, user_id);

-- tenant_2용 테이블 생성
CREATE TABLE tenant_2.events
(
    id UUID,                    -- 고유한 이벤트 ID
    type LowCardinality(String), -- 이벤트 유형
    timestamp DateTime,          -- 이벤트 타임스탬프
    user_id UInt32,               -- 이벤트를 발생시킨 사용자 ID
    data String,                 -- 이벤트 데이터
)
ORDER BY (timestamp, user_id);
