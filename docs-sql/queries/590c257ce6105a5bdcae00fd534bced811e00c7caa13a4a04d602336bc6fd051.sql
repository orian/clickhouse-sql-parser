WITH toDateTime64('2020-01-01 10:20:30.999999999', 9) AS dt64
SELECT toStartOfNanosecond(dt64, 'Asia/Istanbul');
