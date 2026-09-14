WITH mergeTreePartInfo('all_12_25_7_4') AS info
SELECT info.partition_id, info.min_block, info.max_block, info.level, info.mutation;
