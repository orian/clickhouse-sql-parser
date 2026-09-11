WITH
    1 AS buffer,
    4096 AS extent,
    MVTBoundingBox({z:UInt8}, {x:UInt32}, {y:UInt32}, buffer / extent) AS bounding_box   -- margin matches the clip buffer
SELECT MVTEncode('points')(geom, tuple(cluster_count)::Tuple(cluster_count UInt64))
FROM
(
    SELECT MVTEncodeGeom((lon, lat)::Point, {z:UInt8}, {x:UInt32}, {y:UInt32}) AS geom, count() AS cluster_count
    FROM points
    WHERE lon BETWEEN bounding_box.1 AND bounding_box.3 AND lat BETWEEN bounding_box.2 AND bounding_box.4   -- index-using prefilter
    GROUP BY geom
)
SETTINGS allow_suspicious_types_in_group_by = 1
