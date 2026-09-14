  SELECT
      toHour(EventTime) AS hour_of_day,
      -- Используем sumForEach, чтобы за один проход подсчитать визиты с настольных и мобильных устройств
      sumForEach([
          IsMobile = 0, -- Визиты с настольных устройств (IsMobile = 0)
          IsMobile = 1  -- Визиты с мобильных устройств (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
