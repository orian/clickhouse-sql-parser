SELECT * FROM db1.tags_table WHERE arrayExists(x -> x > 3, tags_int);
