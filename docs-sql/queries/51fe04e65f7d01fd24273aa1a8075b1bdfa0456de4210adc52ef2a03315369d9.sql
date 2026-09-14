CREATE TABLE electric_vehicle_state
(
    vin String, -- número de identificação do veículo
    last_update DateTime64 Materialized now64(), -- opcional (usado com argMax)
    battery_level Nullable(UInt8), -- em %
    lat Nullable(Float64), -- latitude (°)
    lon Nullable(Float64), -- longitude (°)
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- em °C
    speed_kmh Nullable(Float32) -- do sensor
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
