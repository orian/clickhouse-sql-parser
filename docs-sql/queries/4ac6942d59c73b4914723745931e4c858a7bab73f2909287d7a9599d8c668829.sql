-- 테넌트 1을 위한 테이블 생성 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- 고유 이벤트 ID
    type LowCardinality(String), -- 이벤트 유형
    timestamp DateTime,          -- 이벤트 발생 타임스탬프
    user_id UInt32,               -- 이벤트를 트리거한 사용자 ID
    data String,                 -- 이벤트 데이터
)
ORDER BY (timestamp, user_id) -- 프라이머리 키는 다른 속성에 집중할 수 있음

-- 테넌트 2를 위한 테이블 생성 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- 고유 이벤트 ID
    type LowCardinality(String), -- 이벤트 유형
    timestamp DateTime,          -- 이벤트 발생 타임스탬프
    user_id UInt32,               -- 이벤트를 트리거한 사용자 ID
    data String,                 -- 이벤트 데이터
)
ORDER BY (timestamp, user_id) -- 프라이머리 키는 다른 속성에 집중할 수 있음
