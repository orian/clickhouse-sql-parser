  SELECT
      toHour(EventTime) AS hour_of_day,
      -- sumForEach를 사용하여 한 번에 데스크톱 방문과 모바일 방문을 집계합니다
      sumForEach([
          IsMobile = 0, -- 데스크톱 방문 (IsMobile = 0)
          IsMobile = 1  -- 모바일 방문 (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
