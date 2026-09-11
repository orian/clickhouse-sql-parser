
      SELECT count() FROM tab
      WHERE str LIKE {prefix_pattern}
      SETTINGS optimize_rewrite_like_perfect_affix=1
  