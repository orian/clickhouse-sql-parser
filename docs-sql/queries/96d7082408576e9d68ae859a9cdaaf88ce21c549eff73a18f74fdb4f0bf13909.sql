--- テーブル events を作成する。tenant_id を主キーの一部として使用する
CREATE TABLE events
(
    tenant_id UInt32,                 -- テナント識別子
    id UUID,                    -- 一意のイベントID
    type LowCardinality(String), -- イベントの種別
    timestamp DateTime,          -- イベントのタイムスタンプ
    user_id UInt32,               -- イベントをトリガーしたユーザーのID
    data String,                 -- イベントデータ
)
ORDER BY (tenant_id, timestamp)
