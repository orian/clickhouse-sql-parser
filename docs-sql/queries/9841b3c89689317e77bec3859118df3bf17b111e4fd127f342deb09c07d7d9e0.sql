    SELECT count(*)
    FROM hackernews
    WHERE hasAllTokens(lower(comment), 'avx sve');
