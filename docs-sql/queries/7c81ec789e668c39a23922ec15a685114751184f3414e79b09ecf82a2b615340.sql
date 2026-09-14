/*
 * 计算两个日期之间的天数。
 * /* 如果任一参数为 NULL，则返回 NULL */
 */
SELECT
    dateDiff('day', toDate('2024-01-01'), toDate('2024-12-31')) AS days_in_year, -- 365
    dateDiff('day', toDate('2020-01-01'), today()) AS days_since  #! since 2020
    ///////////////////////////////////////////////////////////////////
    # TODO: 添加小时/分钟版本
