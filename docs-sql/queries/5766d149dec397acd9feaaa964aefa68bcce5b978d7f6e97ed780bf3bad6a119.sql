CREATE TABLE london
(
    id           String,
    geometry     Geometry,
    properties   Nullable(JSON),
    name         String MATERIALIZED properties.name,
    feature_type String MATERIALIZED properties.feature_type
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO london
SELECT id, geometry, properties
FROM file('london.geojson', GeoJSON);
