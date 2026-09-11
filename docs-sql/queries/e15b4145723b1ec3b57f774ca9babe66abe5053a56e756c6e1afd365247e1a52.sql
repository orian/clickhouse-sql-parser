SELECT sparkbar(9, toDate('2020-01-01'), toDate('2020-01-10'))(event_date, cnt) FROM (SELECT sum(value) AS cnt, event_date FROM spark_bar_data GROUP BY event_date);
