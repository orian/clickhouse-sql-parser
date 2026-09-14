SELECT MobilePhoneModel, count() AS c
FROM metrica.hits
WHERE
      RegionID = 229
  AND EventDate >= '2013-07-01'
  AND EventDate <= '2013-07-31'
  AND MobilePhone != 0
  AND MobilePhoneModel NOT IN ('', 'iPad')
GROUP BY MobilePhoneModel
ORDER BY c DESC
LIMIT 8;
