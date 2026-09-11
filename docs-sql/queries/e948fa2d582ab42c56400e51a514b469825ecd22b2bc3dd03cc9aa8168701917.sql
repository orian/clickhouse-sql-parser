CREATE TABLE regexp_os (id UInt64, parent_id UInt64, regexp String, keys Array(String), values Array(String)) ENGINE = Memory;
INSERT INTO regexp_os VALUES
    (1, 0, 'Linux',      ['os_replacement'], ['Linux']),
    (2, 1, 'Android',    ['os_replacement'], ['Android']),
    (3, 2, 'Android 1',  ['os_replacement'], ['Android']),
    (4, 3, 'Android 12', ['os_replacement'], ['Android']);
CREATE DICTIONARY regexp_tree (regexp String, os_replacement String DEFAULT 'Other') PRIMARY KEY regexp SOURCE(CLICKHOUSE(TABLE 'regexp_os' DB currentDatabase())) LIFETIME(MIN 0 MAX 0) LAYOUT(REGEXP_TREE);
SELECT
    'Mozilla/5.0 (Linux; Android 12; SM-G998B) Mobile Safari/537.36' AS user_agent,

    -- This will match ALL applicable patterns
    dictGetAll('regexp_tree', 'os_replacement', 'Mozilla/5.0 (Linux; Android 12; SM-G998B) Mobile Safari/537.36') AS all_matches,

    -- This returns only the first match
    dictGet('regexp_tree', 'os_replacement', 'Mozilla/5.0 (Linux; Android 12; SM-G998B) Mobile Safari/537.36') AS first_match;
