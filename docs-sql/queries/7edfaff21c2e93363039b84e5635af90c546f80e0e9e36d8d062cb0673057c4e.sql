SELECT
    ip,
    count() AS total
FROM error_log GROUP BY ip ORDER BY total DESC LIMIT 10
FORMAT Template SETTINGS format_template_resultset = 'output.results',
                         format_template_row = 'output.rows';

== Top 10 IPs ==

9.8.4.6 generated 3 requests
9.5.1.1 generated 3 requests
2.4.8.9 generated 3 requests
4.8.8.2 generated 3 requests
4.5.4.4 generated 3 requests
3.3.6.4 generated 2 requests
8.9.5.9 generated 2 requests
2.5.1.8 generated 2 requests
6.8.3.6 generated 2 requests
6.6.3.5 generated 2 requests

--- 1000 rows read in 0.001380604 ---
