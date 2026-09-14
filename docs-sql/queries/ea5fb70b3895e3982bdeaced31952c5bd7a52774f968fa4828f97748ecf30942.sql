SELECT
'é' AS original, -- e + acento agudo combinado (U+0065 + U+0301)
length(original),
normalizeUTF8NFC('é') AS nfc_normalized, -- é (U+00E9)
length(nfc_normalized);
