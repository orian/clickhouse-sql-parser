SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group1,
       timeSeriesTagsToGroup([], '__name__', 'http_failures') AS group2,
       timeSeriesTagsToGroup([]) AS empty_group,
       timeSeriesTagsToGroup([], '__name__', 'http_failures') AS same_group2,
       throwIf(same_group2 != group2),
       timeSeriesGroupToTags(group2)
