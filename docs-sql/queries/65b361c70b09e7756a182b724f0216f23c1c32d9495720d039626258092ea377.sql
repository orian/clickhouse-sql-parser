CREATE TABLE samples_for_my_table
(
    `id` UUID,
    `timestamp` DateTime64(3),
    `value` Float64
)
ENGINE = MergeTree
ORDER BY (id, timestamp);

CREATE TABLE tags_for_my_table ...

CREATE TABLE metrics_for_my_table ...

CREATE TABLE my_table ENGINE=TimeSeries SAMPLES samples_for_my_table TAGS tags_for_my_table METRICS metrics_for_my_table;
