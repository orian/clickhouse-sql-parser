SELECT properties.name AS name, variantType(geometry), geometry.Point, geometry.LineString, geometry.Polygon
FROM file('london.geojson', GeoJSON);
