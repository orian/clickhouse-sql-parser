SELECT
    'Ä ① Hello' AS original,
    normalizeUTF8NFKCCasefold('Ä ① Hello') AS nfkc_cf_normalized;
