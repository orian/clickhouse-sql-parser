INSERT INTO FUNCTION file('london_export.geojson', GeoJSON)
SELECT id, geometry, properties FROM london;
