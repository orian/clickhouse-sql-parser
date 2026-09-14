SET allow_experimental_ai_functions = 1;
SELECT aiGenerate('Explain ClickHouse', map('credentials', 'ai_text_credentials', 'system_prompt', 'You are a database expert. Be concise.'))
