  SELECT
      change_type,
      path,
      old_path,
      time,
      commit_hash
  FROM git.file_changes
  WHERE (path = 'src/Functions/geometryFromColumn.h') OR (old_path = 'src/Functions/geometryFromColumn.h')
