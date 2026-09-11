SELECT
    '① ② ③' AS original,                           -- Circled number characters
    normalizeUTF8NFKC('① ② ③') AS nfkc_normalized  -- Converts to 1 2 3
