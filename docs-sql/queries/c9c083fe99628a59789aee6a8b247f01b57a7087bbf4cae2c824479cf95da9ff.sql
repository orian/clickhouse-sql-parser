SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group,
       timeSeriesGroupToTags(group) AS sorted_tags,
       timeSeriesTagsToGroup(sorted_tags) AS same_group,
       throwIf(same_group != group)
