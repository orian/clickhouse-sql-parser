SELECT id, geometry, properties
FROM london
ORDER BY id
INTO OUTFILE 'london_export.geojson'
FORMAT GeoJSON;
