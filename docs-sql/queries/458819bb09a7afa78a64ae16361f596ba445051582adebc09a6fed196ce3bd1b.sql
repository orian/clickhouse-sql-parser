SELECT * FROM remote(`127.0.0.1`, view(SELECT a, b, c FROM table_name));
