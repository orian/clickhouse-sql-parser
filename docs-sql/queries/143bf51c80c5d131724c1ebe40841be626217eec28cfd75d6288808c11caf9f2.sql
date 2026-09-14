/* ORIGINAL QUERY */
SELECT category, sum(value) FROM t GROUP BY category HAVING service = 'svc1';

/* FIXED QUERY */
SELECT category, sum(value) FROM t WHERE service = 'svc1' GROUP BY category;
