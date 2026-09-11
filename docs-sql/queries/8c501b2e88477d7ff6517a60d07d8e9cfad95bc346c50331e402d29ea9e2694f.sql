SELECT
    arrayJoin([100, 12345, 432546534]) AS elapsed,
    formatReadableTimeDelta(elapsed, 'minutes') AS time_delta
