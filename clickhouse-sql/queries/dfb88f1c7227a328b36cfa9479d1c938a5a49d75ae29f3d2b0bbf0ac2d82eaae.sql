
        SELECT * FROM left JOIN right ON left.id = right.id
        SETTINGS query_plan_join_swap_table = 0, joined_block_split_single_row = 0
    