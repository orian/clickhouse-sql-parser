    SELECT * FROM imdb_dbt.actor_summary WHERE num_movies > 5 ORDER BY avg_rank  DESC LIMIT 10;
