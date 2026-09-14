
        CREATE TABLE json_to_string_shared_data (id UInt64, payload JSON(max_dynamic_paths=0)) ENGINE = MergeTree ORDER BY id
    