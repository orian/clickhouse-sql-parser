    # إنشاء TABLE بثلاثة columns
    CREATE TABLE users_imported
    (
       `username` String,
       `firstname` String,
       `lastname` String
    )
    ENGINE = MergeTree
    ORDER BY firstname
