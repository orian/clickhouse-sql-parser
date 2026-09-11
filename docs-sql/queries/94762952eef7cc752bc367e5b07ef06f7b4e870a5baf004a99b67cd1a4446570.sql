CREATE TABLE mean_ztest (sample_data Float64, sample_index UInt8) ENGINE = Memory;
INSERT INTO mean_ztest VALUES (20.3, 0), (21.9, 0), (22.1, 0), (18.9, 1), (19, 1), (20.3, 1);

SELECT meanZTest(0.7, 0.45, 0.95)(sample_data, sample_index) FROM mean_ztest;
