
        INSERT INTO hierarchical_dictionary_source_table
        WITH 5000 AS first_level_start, 50000 as second_level_start
        SELECT
            (number + 1) as id,
            multiIf(id > second_level_start, id % (second_level_start - first_level_start) + first_level_start + 1,
                    id > first_level_start, (id % first_level_start) + 1,
                    0) as parent_id
        FROM system.numbers
        LIMIT 1000000;
    