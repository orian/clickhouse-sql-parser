       -- الوصول إلى قاعدة البيانات الهدف: CREATE TABLE وDROP TABLE وSELECT وINSERT وTRUNCATE وغيرها.
       GRANT ALL ON <database>.* TO <user>;

       -- الوصول إلى جداول النظام لاكتشاف البيانات الوصفية وإدارة الأقسام.
       -- لا تشملها صلاحيات قاعدة البيانات الممنوحة أعلاه.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
