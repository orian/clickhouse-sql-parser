    INSERT INTO cell_towers SELECT * FROM s3('https://datasets-documentation.s3.amazonaws.com/cell_towers/cell_towers.csv.xz', NOSIGN, 'CSVWithNames')
