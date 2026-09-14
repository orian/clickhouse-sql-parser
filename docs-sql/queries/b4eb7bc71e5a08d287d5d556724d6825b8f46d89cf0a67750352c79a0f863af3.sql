SELECT description
FROM system.documentation
WHERE type = 'Table Engine' AND name = 'MergeTree'
FORMAT TSVRaw;
