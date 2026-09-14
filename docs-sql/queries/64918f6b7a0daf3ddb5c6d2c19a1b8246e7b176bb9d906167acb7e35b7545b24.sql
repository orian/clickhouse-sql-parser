    -- امنح أذونات لقراءة معلومات المخطط
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- امنح صلاحية القراءة لقاعدة البيانات والجداول الخاصة بك
    GRANT SELECT ON your_database.* TO your_databrain_user;
