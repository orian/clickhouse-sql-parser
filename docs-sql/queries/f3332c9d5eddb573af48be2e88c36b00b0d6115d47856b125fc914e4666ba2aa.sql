--- 创建表 events，将 tenant_id 作为主键的一部分
CREATE TABLE events
(
    tenant_id UInt32,                 -- 租户标识符
    id UUID,                    -- 唯一事件 ID
    type LowCardinality(String), -- 事件类型
    timestamp DateTime,          -- 事件时间戳
    user_id UInt32,               -- 触发事件的用户 ID
    data String,                 -- 事件数据
)
ORDER BY (tenant_id, timestamp)
