CREATE TABLE mww_ttest (sample_data Float64, sample_index UInt8) ENGINE = Memory;
INSERT INTO mww_ttest VALUES (10, 0), (11, 0), (12, 0), (1, 1), (2, 1), (3, 1);

SELECT mannWhitneyUTest('greater')(sample_data, sample_index) FROM mww_ttest;
