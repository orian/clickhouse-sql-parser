CREATE TABLE electric_vehicle_state
(
    vin String, -- número de identificación del vehículo
    last_update DateTime64 Materialized now64(), -- opcional (usado con argMax)
    battery_level Nullable(UInt8), -- en %
    lat Nullable(Float64), -- latitud (°)
    lon Nullable(Float64), -- longitud (°)
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- en °C
    speed_kmh Nullable(Float32) -- del sensor
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
