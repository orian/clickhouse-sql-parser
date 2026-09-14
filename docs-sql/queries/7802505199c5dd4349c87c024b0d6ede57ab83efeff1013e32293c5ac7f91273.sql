SELECT
'é' AS original, -- e + 결합 악센트 부호 (U+0065 + U+0301)
length(original),
normalizeUTF8NFC('é') AS nfc_normalized, -- é (U+00E9)
length(nfc_normalized);
