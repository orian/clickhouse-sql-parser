SELECT id, properties.name AS name, variantType(geometry) AS geo_type
FROM file('london.geojson', GeoJSON);
