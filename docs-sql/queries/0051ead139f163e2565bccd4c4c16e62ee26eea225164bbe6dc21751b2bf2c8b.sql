SELECT count() FROM cell_towers
WHERE pointInPolygon((lon, lat), (SELECT * FROM moscow))
