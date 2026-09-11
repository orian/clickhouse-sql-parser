CREATE TABLE tab (city Nullable(String)) ENGINE=Memory;
INSERT INTO tab (city) VALUES (NULL), ('Amsterdam'), ('New York'), ('Tokyo'), ('Valencia'), (NULL);
SELECT any(city), anyRespectNulls(city) FROM tab;
