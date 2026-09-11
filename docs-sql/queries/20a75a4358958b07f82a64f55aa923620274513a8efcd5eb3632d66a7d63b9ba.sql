SELECT properties.name AS name, geometry, variantType(geometry)
FROM file('london.geojson', GeoJSON);
