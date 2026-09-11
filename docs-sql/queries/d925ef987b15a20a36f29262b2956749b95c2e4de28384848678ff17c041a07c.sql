CREATE TABLE IF NOT EXISTS geo (geom Geometry) ENGINE = Memory();
INSERT INTO geo VALUES ((1, 2));
SELECT * FROM geo;
