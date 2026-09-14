SELECT sum(PageViews * _sample_factor)
FROM visits
SAMPLE 10000000
