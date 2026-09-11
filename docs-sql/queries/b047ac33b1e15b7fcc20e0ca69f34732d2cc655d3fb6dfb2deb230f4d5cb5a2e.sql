CREATE TABLE geo_multipoint (mp MultiPoint) ENGINE = Memory();
INSERT INTO geo_multipoint VALUES([(0, 0), (10, 0), (10, 10), (0, 10)]);
SELECT mp, toTypeName(mp) FROM geo_multipoint;
