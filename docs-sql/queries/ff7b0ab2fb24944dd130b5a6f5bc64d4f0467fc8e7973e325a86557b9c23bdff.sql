CREATE TABLE arxiv
(
  `doc` JSON(update_date Date)
)
ENGINE = MergeTree
ORDER BY doc.update_date
