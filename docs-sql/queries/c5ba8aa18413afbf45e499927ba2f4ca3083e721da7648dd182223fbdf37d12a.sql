SELECT count()
FROM [...]
WHERE hasToken(column_with_text_index, 'token')
-- or: hasAllTokens(column_with_text_index, ['token', ...])
-- or: hasAnyTokens(column_with_text_index, ['token', ...])
