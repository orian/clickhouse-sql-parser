-- tenant_1 のテーブルを作成する
CREATE TABLE events
(
    id UUID,                    -- 一意のイベントID
    type LowCardinality(String), -- イベントの種類
    timestamp DateTime,          -- イベントのタイムスタンプ
    user_id UInt32,               -- イベントをトリガーしたユーザーのID
    data String,                 -- イベントデータ
)
ORDER BY (timestamp, user_id);
