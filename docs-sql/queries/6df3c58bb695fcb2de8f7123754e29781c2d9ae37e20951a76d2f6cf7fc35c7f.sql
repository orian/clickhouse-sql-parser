-- Increment the count
ALTER TABLE posts UPDATE ViewCount = ViewCount + 1 WHERE Id = 404346

-- The updated value is immediately visible
SELECT ViewCount FROM posts WHERE Id = 404346
