CREATE TABLE electric_vehicle_state
(
    vin String, -- 车辆识别码
    last_update DateTime64 Materialized now64(), -- 可选（与 argMax 配合使用）
    battery_level Nullable(UInt8), -- 单位：%
    lat Nullable(Float64), -- 纬度（°）
    lon Nullable(Float64), -- 经度（°）
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- 单位：°C
    speed_kmh Nullable(Float32) -- 来自传感器
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
