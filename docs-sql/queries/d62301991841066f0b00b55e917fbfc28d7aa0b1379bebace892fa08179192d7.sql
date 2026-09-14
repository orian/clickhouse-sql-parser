SELECT
    domainWithoutWWW(URL) AS domain,
    count(),
    any(Title) AS title -- 각 도메인에서 처음 나타난 페이지 헤더를 가져옵니다.
FROM hits
GROUP BY domain
