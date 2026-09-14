INSERT INTO imdb_dbt.actor_summary ("id", "name", "num_movies", "avg_rank", "genres", "directors", "updated_at")
WITH actor_summary AS (
   SELECT id,
      any(actor_name) AS name,
      uniqExact(movie_id)    AS num_movies,
      avg(rank)                AS avg_rank,
      uniqExact(genre)         AS genres,
      uniqExact(director_name) AS directors,
      max(created_at) AS updated_at
   FROM (
      SELECT imdb.actors.id AS id,
         concat(imdb.actors.first_name, ' ', imdb.actors.last_name) AS actor_name,
         imdb.movies.id AS movie_id,
         imdb.movies.rank AS rank,
         genre,
         concat(imdb.directors.first_name, ' ', imdb.directors.last_name) AS director_name,
         created_at
      FROM imdb.actors
         JOIN imdb.roles ON imdb.roles.actor_id = imdb.actors.id
         LEFT OUTER JOIN imdb.movies ON imdb.movies.id = imdb.roles.movie_id
         LEFT OUTER JOIN imdb.genres ON imdb.genres.movie_id = imdb.movies.id
         LEFT OUTER JOIN imdb.movie_directors ON imdb.movie_directors.movie_id = imdb.movies.id
         LEFT OUTER JOIN imdb.directors ON imdb.directors.id = imdb.movie_directors.director_id
   )
   GROUP BY id
)

SELECT *
FROM actor_summary
-- 이 필터는 증분 실행 시에만 적용됩니다
WHERE id > (SELECT max(id) FROM imdb_dbt.actor_summary) OR updated_at > (SELECT max(updated_at) FROM imdb_dbt.actor_summary)
