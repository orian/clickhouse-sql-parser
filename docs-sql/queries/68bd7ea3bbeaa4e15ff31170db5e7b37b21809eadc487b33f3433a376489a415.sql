CREATE TABLE student_ttest (sample_data Float64, sample_index UInt8) ENGINE = Memory;
INSERT INTO student_ttest VALUES (20.3, 0), (21.1, 0), (21.9, 1), (21.7, 0), (19.9, 1), (21.8, 1);

SELECT studentTTest(sample_data, sample_index) FROM student_ttest;
