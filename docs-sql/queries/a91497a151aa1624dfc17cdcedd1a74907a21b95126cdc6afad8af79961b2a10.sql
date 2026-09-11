   INSERT INTO imdb.roles
   SELECT now() as created_at, 845466 as actor_id, rand(number) % 412320 as movie_id, 'Himself' as role
   FROM system.numbers
   LIMIT 10;
