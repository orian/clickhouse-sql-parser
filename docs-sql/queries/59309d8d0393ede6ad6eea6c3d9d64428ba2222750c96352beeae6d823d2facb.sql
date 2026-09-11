SELECT
    'é' AS original, -- é (U+00E9)
    length(original),
    normalizeUTF8NFD('é') AS nfd_normalized, -- e + combining acute (U+0065 + U+0301)
    length(nfd_normalized);
