    # إنشاء جدول يحتوي على 3 أعمدة
    CREATE TABLE users_imported
    (
       `username` String,
       `firstname` String,
       `lastname` String
    )
    ENGINE = MergeTree
    ORDER BY firstname
