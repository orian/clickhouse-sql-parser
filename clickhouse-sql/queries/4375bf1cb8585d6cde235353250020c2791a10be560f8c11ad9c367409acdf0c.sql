
      SELECT count() FROM tab
      WHERE str LIKE {suffix_pattern}
      SETTINGS optimize_rewrite_like_perfect_affix=1
  