
        SELECT throwIf(
            count() != 1 OR countIf(part_type = 'Wide') != 1,
            'Expected one active Wide part')
        FROM system.parts
        WHERE database = currentDatabase()
          AND table = 'group_by_low_cardinality_array_join'
          AND active
    