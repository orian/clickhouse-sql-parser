  SELECT
      URLDomain,
      groupArrayDistinct(OS) AS distinct_os_codes
  FROM metrica.hits_v1
  WHERE URLDomain != '' -- ضع في الاعتبار فقط السجلات التي تحتوي على نطاق مسجّل
  GROUP BY URLDomain
  ORDER BY URLDomain ASC
  LIMIT 20;
