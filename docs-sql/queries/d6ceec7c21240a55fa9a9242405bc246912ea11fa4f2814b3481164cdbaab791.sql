CREATE TABLE issues (body String) ENGINE = Memory;
INSERT INTO issues VALUES ('The application exits unexpectedly after login.');
SELECT body, aiFilter(body, 'describes a bug', map('credentials', 'ai_text_credentials')) AS is_bug FROM issues LIMIT 5
