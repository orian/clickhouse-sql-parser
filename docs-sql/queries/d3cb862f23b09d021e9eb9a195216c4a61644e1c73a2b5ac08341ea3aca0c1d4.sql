SELECT timeSeriesTagsToGroup([('region', 'eu'), ('env', 'dev')], '__name__', 'http_requests_count') AS group,
       timeSeriesExtractTag(group, '__name__'),
       timeSeriesExtractTag(group, 'env'),
       timeSeriesExtractTag(group, 'instance')
