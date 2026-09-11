-- テナント1のテーブルを作成 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- イベントの一意なID
    type LowCardinality(String), -- イベントの種別
    timestamp DateTime,          -- イベントのタイムスタンプ
    user_id UInt32,               -- イベントをトリガーしたユーザーのID
    data String,                 -- イベントデータ
)
ORDER BY (timestamp, user_id) -- 主キーは他の属性に集中させることができる

-- テナント2のテーブルを作成 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- イベントの一意なID
    type LowCardinality(String), -- イベントの種別
    timestamp DateTime,          -- イベントのタイムスタンプ
    user_id UInt32,               -- イベントをトリガーしたユーザーのID
    data String,                 -- イベントデータ
)
ORDER BY (timestamp, user_id) -- 主キーは他の属性に集中させることができる
