CREATE TABLE my_executable_table (
   x UInt32,
   y String
)
ENGINE = Executable('my_script.py', TabSeparated, (SELECT 10))
