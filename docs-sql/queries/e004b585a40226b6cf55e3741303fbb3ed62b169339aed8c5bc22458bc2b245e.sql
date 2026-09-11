  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- 仅考虑记录了域名的 hits
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
