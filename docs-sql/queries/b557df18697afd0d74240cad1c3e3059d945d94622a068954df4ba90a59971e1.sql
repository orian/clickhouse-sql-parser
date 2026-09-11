SET allow_experimental_ai_functions = 1;
CREATE TABLE issues (body String) ENGINE = Memory;
INSERT INTO issues VALUES ('The application exits unexpectedly after login.');
SELECT body, aiClassify(body, ['bug', 'question', 'feature'], map('credentials', 'ai_text_credentials')) AS kind FROM issues LIMIT 5
