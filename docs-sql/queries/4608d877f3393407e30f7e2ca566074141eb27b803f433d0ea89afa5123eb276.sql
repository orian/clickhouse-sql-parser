SELECT
        toStartOfWeek(toDateTime('2023-04-21 10:20:30')), /* a Friday */
        toStartOfWeek(toDateTime('2023-04-21 10:20:30'), 1), /* a Friday */
        toStartOfWeek(toDate('2023-04-24')), /* a Monday */
        toStartOfWeek(toDate('2023-04-24'), 1) /* a Monday */
    FORMAT Vertical
