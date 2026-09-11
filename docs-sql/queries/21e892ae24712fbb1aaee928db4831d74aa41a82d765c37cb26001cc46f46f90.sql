--- events 테이블 생성. tenant_id를 기본 키의 일부로 사용
CREATE TABLE events
(
    tenant_id UInt32,                 -- 테넌트 식별자
    id UUID,                    -- 고유 이벤트 ID
    type LowCardinality(String), -- 이벤트 유형
    timestamp DateTime,          -- 이벤트 타임스탬프
    user_id UInt32,               -- 이벤트를 발생시킨 사용자 ID
    data String,                 -- 이벤트 데이터
)
ORDER BY (tenant_id, timestamp)
