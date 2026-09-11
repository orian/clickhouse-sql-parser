SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group_of_3,
       timeSeriesRemoveTags(group_of_3, ['env', 'region']) AS group_of_1,
       timeSeriesGroupToTags(group_of_1),
       timeSeriesRemoveTags(group_of_1, ['__name__', 'nonexistent']) AS empty_group,
       timeSeriesGroupToTags(empty_group)
