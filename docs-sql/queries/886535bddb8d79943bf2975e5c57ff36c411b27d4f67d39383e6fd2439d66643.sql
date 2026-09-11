SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group_of_3,
       timeSeriesRemoveTag(group_of_3, '__name__') AS group_of_2,
       timeSeriesGroupToTags(group_of_2),
       timeSeriesRemoveTag(group_of_2, 'env') AS group_of_1,
       timeSeriesGroupToTags(group_of_1),
       timeSeriesRemoveTag(group_of_1, 'region') AS empty_group,
       timeSeriesGroupToTags(empty_group)
