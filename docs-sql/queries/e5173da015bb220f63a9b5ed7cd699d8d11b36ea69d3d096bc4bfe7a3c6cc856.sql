SET allow_experimental_ai_functions = 1;
SELECT aiEmbed('Hello world', 'text-embedding-3-small', map('credentials', 'ai_embedding_credentials', 'dimensions', '256'))
