SELECT
    'H₂O²' AS original,                           -- H + subscript 2 + O + superscript 2
    normalizeUTF8NFKD('H₂O²') AS nfkd_normalized  -- Converts to H 2 O 2
