CREATE TABLE electric_vehicle_state
(
    vin String, -- vehicle identification number
    last_update DateTime64 Materialized now64(), -- optional (used with argMax)
    battery_level Nullable(UInt8), -- in %
    lat Nullable(Float64), -- latitude (°)
    lon Nullable(Float64), -- longitude (°)
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- in °C
    speed_kmh Nullable(Float32) -- from sensor
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
