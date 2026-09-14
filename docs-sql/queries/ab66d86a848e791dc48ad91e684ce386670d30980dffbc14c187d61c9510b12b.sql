SELECT
    'é' AS original, -- é (U+00E9)
    length(original),
    normalizeUTF8NFD('é') AS nfd_normalized, -- e + 결합 악센트 (U+0065 + U+0301)
    length(nfd_normalized);
