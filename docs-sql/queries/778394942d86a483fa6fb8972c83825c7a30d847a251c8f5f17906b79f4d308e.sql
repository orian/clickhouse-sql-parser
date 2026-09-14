CREATE TABLE images
(
	`_file` LowCardinality(String),
	`caption` String,
	`image_embedding` Array(Float32)
)
ENGINE = MergeTree;
