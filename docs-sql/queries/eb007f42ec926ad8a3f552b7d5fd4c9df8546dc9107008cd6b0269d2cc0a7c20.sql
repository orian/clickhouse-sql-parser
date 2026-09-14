SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group,
       timeSeriesRemoveAllTagsExcept(group, ['env']) AS result_group,
       timeSeriesGroupToTags(result_group)
