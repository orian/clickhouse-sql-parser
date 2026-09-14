INSERT INTO FUNCTION file('a.txt', 'RawBLOB') SELECT 'Hello' SETTINGS engine_file_truncate_on_insert = 1;
INSERT INTO FUNCTION file('b.txt', 'RawBLOB') SELECT 'World!' SETTINGS engine_file_truncate_on_insert = 1;

CREATE TABLE data (a String, b String) ENGINE = Memory;
INSERT INTO data SELECT file('a.txt'), file('b.txt');

SELECT * FROM data;
