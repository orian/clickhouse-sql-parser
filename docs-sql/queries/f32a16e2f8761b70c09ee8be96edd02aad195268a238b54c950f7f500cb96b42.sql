CREATE QUOTA my_quota KEYED BY normalized_query_hash FOR INTERVAL 1 hour MAX queries = 100 TO my_user;
