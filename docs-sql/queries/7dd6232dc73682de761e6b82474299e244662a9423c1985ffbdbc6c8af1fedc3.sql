SELECT timeSeriesTagsToGroup([('__name__', 'up'), ('job', 'api-server'), ('service', 'a:c')]) AS group,
       timeSeriesReplaceTag(group, 'foo', '$1', 'service', '(.*):.*') AS result_group,
       timeSeriesGroupToTags(result_group)
