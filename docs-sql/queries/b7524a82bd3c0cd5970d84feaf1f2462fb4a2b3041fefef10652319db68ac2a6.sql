CREATE DICTIONARY regions_dict
(
    region_id UInt64,
    parent_region UInt64 DEFAULT 0 HIERARCHICAL,
    region_name String DEFAULT ''
)
PRIMARY KEY region_id
SOURCE(...)
LAYOUT(HASHED())
LIFETIME(3600);
