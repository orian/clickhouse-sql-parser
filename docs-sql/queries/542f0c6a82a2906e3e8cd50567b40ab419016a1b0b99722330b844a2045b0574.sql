CREATE TABLE my_table ENGINE=TimeSeries
SAMPLES INNER COLUMNS (timestamp UInt32 CODEC(DoubleDelta, ZSTD(1)), value Float32 CODEC(ZSTD(3)))
