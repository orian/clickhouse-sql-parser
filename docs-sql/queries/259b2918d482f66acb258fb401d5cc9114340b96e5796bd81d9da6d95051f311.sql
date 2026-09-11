SELECT MVTEncode('points')(geom, tuple(cluster_count)::Tuple(cluster_count UInt64)) AS tile
FROM
(
    SELECT MVTEncodeGeom((lon, lat)::Point, 10, 550, 335) AS geom, count() AS cluster_count
    FROM points
    GROUP BY geom
)
SETTINGS allow_suspicious_types_in_group_by = 1;
