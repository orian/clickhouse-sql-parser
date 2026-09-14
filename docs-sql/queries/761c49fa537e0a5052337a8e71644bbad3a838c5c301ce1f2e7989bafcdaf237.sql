SET allow_experimental_ai_functions = 1;
SELECT aiSimilarity('cat', 'kitten', 'text-embedding-3-small', map('credentials', 'ai_embedding_credentials'))
