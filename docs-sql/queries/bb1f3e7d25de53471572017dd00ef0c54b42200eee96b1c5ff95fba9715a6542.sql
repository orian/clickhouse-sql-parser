CREATE TABLE foursquare_mercator
(
    fsq_place_id Nullable(String),
    name Nullable(String),
    latitude Float64,
    longitude Float64,
    address Nullable(String),
    locality Nullable(String),
    region LowCardinality(Nullable(String)),
    postcode LowCardinality(Nullable(String)),
    admin_region LowCardinality(Nullable(String)),
    post_town LowCardinality(Nullable(String)),
    po_box LowCardinality(Nullable(String)),
    country LowCardinality(Nullable(String)),
    date_created Nullable(Date),
    date_refreshed Nullable(Date),
    date_closed Nullable(Date),
    tel Nullable(String),
    website Nullable(String),
    email Nullable(String),
    facebook_id Nullable(Int64),
    instagram Nullable(String),
    twitter Nullable(String),
    fsq_category_ids Array(Nullable(String)),
    fsq_category_labels Array(Nullable(String)),
    placemaker_url Nullable(String),
    geom Nullable(String),
    bbox Tuple(
        xmin Nullable(Float64),
        ymin Nullable(Float64),
        xmax Nullable(Float64),
        ymax Nullable(Float64)
    ),
    category LowCardinality(Nullable(String)) ALIAS fsq_category_labels[1],
    mercator_x UInt32 MATERIALIZED 0xFFFFFFFF * ((longitude + 180) / 360),
    mercator_y UInt32 MATERIALIZED 0xFFFFFFFF * ((1 / 2) - ((log(tan(((latitude + 90) / 360) * pi())) / 2) / pi())),
    INDEX idx_x mercator_x TYPE minmax,
    INDEX idx_y mercator_y TYPE minmax
)
ENGINE = MergeTree
ORDER BY mortonEncode(mercator_x, mercator_y);
