
        create table json_advanced_shared_data_top (json JSON(max_dynamic_paths=0))
        engine=MergeTree order by tuple()
        settings min_bytes_for_wide_part=0, min_rows_for_wide_part=0,
                 object_shared_data_serialization_version='advanced',
                 object_shared_data_serialization_version_for_zero_level_parts='advanced',
                 object_shared_data_buckets_for_wide_part=1
    