  INSERT INTO infile_globs FROM INFILE 'input_*.csv' FORMAT CSV;
  INSERT INTO infile_globs FROM INFILE 'input_{1,2}.csv' FORMAT CSV;
  INSERT INTO infile_globs FROM INFILE 'input_?.csv' FORMAT CSV;
