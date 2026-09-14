MODIFY COLUMN [IF EXISTS] name
    [type] [default_expr] [COMMENT 'comment for column'] [codec] [STATISTICS] [TTL] [settings] [AFTER name_after | FIRST]
    | ADD ENUM VALUES ( 'name' [= number] [, ...] )
ALTER COLUMN [IF EXISTS] name
    TYPE [type] [default_expr] [COMMENT 'comment for column'] [codec] [STATISTICS] [TTL] [settings] [AFTER name_after | FIRST]
    | ADD ENUM VALUES ( 'name' [= number] [, ...] )
