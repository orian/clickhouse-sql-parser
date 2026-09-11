-- 为 tenant_1 创建表
CREATE TABLE events
(
    id UUID,                    -- 唯一事件 ID
    type LowCardinality(String), -- 事件类型
    timestamp DateTime,          -- 事件时间戳
    user_id UInt32,               -- 触发事件的用户 ID
    data String,                 -- 事件数据
)
ORDER BY (timestamp, user_id);
