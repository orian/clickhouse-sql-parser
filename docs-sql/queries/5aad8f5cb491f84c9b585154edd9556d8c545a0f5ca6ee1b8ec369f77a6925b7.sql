SELECT
    toLastDayOfWeek(toDateTime('2023-04-21 10:20:30')), /* a Friday */
    toLastDayOfWeek(toDateTime('2023-04-21 10:20:30'), 1), /* a Friday */
    toLastDayOfWeek(toDate('2023-04-23')), /* a Sunday */
    toLastDayOfWeek(toDate('2023-04-23'), 1) /* a Sunday */
FORMAT Vertical
