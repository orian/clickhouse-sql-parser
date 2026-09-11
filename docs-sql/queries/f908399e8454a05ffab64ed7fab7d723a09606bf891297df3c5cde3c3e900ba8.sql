CREATE NAMED COLLECTION ai_text_credentials AS
    provider = 'openai',
    endpoint = 'https://api.openai.com/v1/chat/completions',
    model = 'gpt-4o-mini',
    api_key = 'sk-...';

-- The embedding functions (`aiEmbed`, `aiSimilarity`) do not read `model` from the named collection,
-- pass it as a positional argument instead. Defining `model` in an embedding collection is an error,
-- not silently ignored.
CREATE NAMED COLLECTION ai_embedding_credentials AS
    provider = 'openai',
    endpoint = 'https://api.openai.com/v1/embeddings',
    api_key = 'sk-...';
