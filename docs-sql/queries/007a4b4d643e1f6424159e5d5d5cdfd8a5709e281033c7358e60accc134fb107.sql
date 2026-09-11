SELECT JSONAllPathsWithTypes(tags)
FROM arxiv
FORMAT PrettyJSONEachRow

{
  "JSONAllPathsWithTypes(tags)": {
    "28_03_2025.comment": "String",
    "28_03_2025.name": "String",
    "28_03_2025.score": "Int64",
    "28_03_2025.updates": "Array(JSON(max_dynamic_types=16, max_dynamic_paths=256))",
    "tag_1.comment": "String",
    "tag_1.name": "String",
    "tag_1.score": "String"
  }
}
