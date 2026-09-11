CREATE TABLE electric_vehicle_state
(
    vin String, -- идентификационный номер транспортного средства
    last_update DateTime64 Materialized now64(), -- необязательно (используется с argMax)
    battery_level Nullable(UInt8), -- в %
    lat Nullable(Float64), -- широта (°)
    lon Nullable(Float64), -- долгота (°)
    firmware_version Nullable(String),
    cabin_temperature Nullable(Float32), -- в °C
    speed_kmh Nullable(Float32) -- с датчика
)
ENGINE = CoalescingMergeTree
ORDER BY vin;
