CREATE TABLE jsons
(
    `json` String
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO jsons VALUES ('{"foo":"true","qux":1}');

SELECT simpleJSONHas(json, 'foo') FROM jsons;
SELECT simpleJSONHas(json, 'bar') FROM jsons;
