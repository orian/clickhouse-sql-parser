SELECT name, database FROM system.tables WHERE match(name, {tableNameRegex: String}) AND database = {database: String};
