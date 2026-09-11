INSERT INTO shorttable FROM INFILE 'list.json' FORMAT JSONEachRow;
SELECT * FROM shorttable
