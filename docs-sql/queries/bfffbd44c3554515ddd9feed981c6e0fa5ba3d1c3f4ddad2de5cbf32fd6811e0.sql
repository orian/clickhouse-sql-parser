CREATE TABLE tab(city Nullable(String)) ENGINE=Memory;
INSERT INTO tab (city) VALUES ('Amsterdam'), (NULL), ('New York'), ('Tokyo'), ('Valencia'), (NULL);
SELECT anyLast(city), anyLastRespectNulls(city) FROM tab;
