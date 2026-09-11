SELECT catboostEvaluate('/var/lib/clickhouse/user_files/occupy.bin', Temperature, Humidity, Light, CO2, HumidityRatio) AS prediction FROM occupancy LIMIT 1
