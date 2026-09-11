ALTER TABLE stackoverflow.posts
  (ADD INDEX view_count_idx ViewCount TYPE minmax GRANULARITY 1);

ALTER TABLE stackoverflow.posts MATERIALIZE INDEX view_count_idx;
