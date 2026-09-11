CREATE TABLE electric_vehicle_state
(
    vin String, -- 車両識別番号
    last_update DateTime64 Materialized now64(), -- 任意（argMaxと併用）
    battery_level Nullable(UInt8), -- %単位
    lat Nullable(Float64), -- 緯度（°）
    lon Nullable(Float64), -- 経度（°）
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- °C単位
    speed_kmh Nullable(Float32) -- センサーから取得
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
