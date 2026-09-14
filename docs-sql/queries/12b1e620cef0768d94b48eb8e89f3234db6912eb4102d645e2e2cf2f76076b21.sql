INSERT INTO db1.target_table_population (city, country, continent, population)
SELECT city, country, continent, 0 FROM db1.source_table;
