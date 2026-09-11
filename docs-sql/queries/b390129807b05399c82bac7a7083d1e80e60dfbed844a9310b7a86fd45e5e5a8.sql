CREATE TABLE covid19 (
    date Date,
    location_key LowCardinality(String),
    new_confirmed Int32,
    new_deceased Int32,
    new_recovered Int32,
    new_tested Int32,
    cumulative_confirmed Int32,
    cumulative_deceased Int32,
    cumulative_recovered Int32,
    cumulative_tested Int32
)
ENGINE = MergeTree
ORDER BY (location_key, date);
