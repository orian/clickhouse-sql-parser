  SELECT
      toHour(EventTime) AS hour_of_day,
      -- sumForEach を使って、デスクトップとモバイルの訪問数を 1 回で集計
      sumForEach([
          IsMobile = 0, -- デスクトップからの訪問 (IsMobile = 0)
          IsMobile = 1  -- モバイルからの訪問 (IsMobile = 1)
      ]) AS device_counts
  FROM metrica.hits
  GROUP BY hour_of_day
  ORDER BY hour_of_day;
