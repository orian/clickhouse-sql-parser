CREATE [OR REPLACE] DICTIONARY [IF NOT EXISTS] [db.]dictionary_name [ON CLUSTER cluster]
(
    key1  type1  [DEFAULT | EXPRESSION expr1] [IS_OBJECT_ID],
    key2  type2  [DEFAULT | EXPRESSION expr2],
    attr1 type2  [DEFAULT | EXPRESSION expr3] [HIERARCHICAL|INJECTIVE],
    attr2 type2  [DEFAULT | EXPRESSION expr4] [HIERARCHICAL|INJECTIVE]
)
PRIMARY KEY key1, key2
SOURCE(SOURCE_NAME([param1 value1 ... paramN valueN]))
LAYOUT(LAYOUT_NAME([param_name param_value]))
LIFETIME({MIN min_val MAX max_val | max_val})
SETTINGS(setting_name = setting_value, setting_name = setting_value, ...)
COMMENT 'Comment'
