SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS dest_group,
       timeSeriesTagsToGroup([('code', '404'), ('message', 'Page not found')], '__name__', 'http_codes') AS src_group,
       timeSeriesCopyTags(dest_group, src_group, ['__name__', 'code', 'env']) AS result_group,
       timeSeriesGroupToTags(result_group)
