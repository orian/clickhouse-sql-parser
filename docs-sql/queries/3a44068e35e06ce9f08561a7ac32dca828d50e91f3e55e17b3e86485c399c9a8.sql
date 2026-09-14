CREATE TABLE welch_ttest (sample_data Float64, sample_index UInt8) ENGINE = Memory;
INSERT INTO welch_ttest VALUES (20.3, 0), (22.1, 0), (21.9, 0), (18.9, 1), (20.3, 1), (19, 1);

SELECT welchTTest(sample_data, sample_index) FROM welch_ttest;
