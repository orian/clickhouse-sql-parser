CREATE TABLE electric_vehicle_state
(
    vin String, -- 차량 식별 번호
    last_update DateTime64 Materialized now64(), -- 선택 사항 (argMax와 함께 사용)
    battery_level Nullable(UInt8), -- % 단위
    lat Nullable(Float64), -- 위도 (°)
    lon Nullable(Float64), -- 경도 (°)
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- °C 단위
    speed_kmh Nullable(Float32) -- 센서 측정값
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
