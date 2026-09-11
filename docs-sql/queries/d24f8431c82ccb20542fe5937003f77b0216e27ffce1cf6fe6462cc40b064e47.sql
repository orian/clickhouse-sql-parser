SELECT
    timeslot,
    toTypeName(sumMap(statusMapTuple)),
    toTypeName(sumMapWithOverflow(statusMapTuple))
FROM sum_map
GROUP BY timeslot;
