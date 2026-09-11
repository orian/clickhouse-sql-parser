SET input_format_skip_unknown_fields = 0;
INSERT INTO shorttable FROM INFILE 'list.json' FORMAT JSONEachRow;
