SET enable_json_type = 1;

CREATE TABLE people
(
    `json` JSON(username String)
)
ENGINE = MergeTree
ORDER BY json.username;
