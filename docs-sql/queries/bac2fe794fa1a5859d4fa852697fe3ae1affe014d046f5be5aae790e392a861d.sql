SELECT radio, count() AS c FROM cell_towers GROUP BY radio ORDER BY c DESC
