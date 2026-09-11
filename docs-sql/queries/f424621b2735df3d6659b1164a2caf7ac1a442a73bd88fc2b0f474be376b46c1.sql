SELECT timeSeriesTagsToGroup([('__name__', 'up'), ('job', 'api-server'), ('src1', 'a'), ('src2', 'b'), ('src3', 'c')]) AS group,
       timeSeriesJoinTags(group, 'foo', ',', ['src1', 'src2', 'src3']) AS result_group,
       timeSeriesGroupToTags(result_group)
