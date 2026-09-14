SELECT version
  FROM pg_get_loaded_modules()
 WHERE module_name = 'chdb';
