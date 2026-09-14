SELECT replaceRegexpOne(d, '(\\d{4})-(\\d{2})-(\\d{2})', '\\2/\\3/\\1') AS res
FROM values('d String', '2014-03-17', '2014-03-18', '2014-03-19')
