  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- 기록된 도메인이 있는 hits만 고려
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
