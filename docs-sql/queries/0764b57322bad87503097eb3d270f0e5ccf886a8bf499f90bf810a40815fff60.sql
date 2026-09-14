CREATE TABLE poly (id UInt32, shape Polygon) ENGINE = Memory;
INSERT INTO poly VALUES (1, [[(0, 0), (10, 0), (10, 10), (0, 10)], [(4, 4), (6, 4), (6, 6), (4, 6)]]);
SELECT id, pointInPolygon((2., 2.), shape) AS res FROM poly;
