WITH '< Server: nginx
< Date: Tue, 22 Jan 2019 00:26:14 GMT
< Content-Type: text/html; charset=UTF-8
< Connection: keep-alive
' AS s
SELECT extractAllGroupsVertical(s, '< ([\\w\\-]+): ([^\\r\\n]+)');
