  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- ドメインが記録されている hits のみを対象とする
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
