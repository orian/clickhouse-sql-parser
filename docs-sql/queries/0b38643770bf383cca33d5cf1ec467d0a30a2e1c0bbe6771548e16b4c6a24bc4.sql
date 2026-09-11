SELECT _table,
   multiIf(
     walkover IS NOT NULL,
     walkover,
     variantType(score) = 'Array(String)',
     toBool(arrayExists(
        x -> position(x, 'W/O') > 0,
        variantElement(score, 'Array(String)')
     )),
     variantElement(score, 'String') LIKE '%W/O%'
   ),
   count()
FROM merge('atp_matches*')
GROUP BY ALL
ORDER BY _table;
