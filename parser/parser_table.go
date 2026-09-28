package parser

import (
	"errors"
	"fmt"
	"strings"
)

func (p *Parser) parseDDL(pos Pos) (DDL, error) {
	switch {
	case p.matchKeyword(KeywordCreate),
		p.matchKeyword(KeywordAttach):
		isAttach := p.matchKeyword(KeywordAttach)
		_ = p.lexer.consumeToken()
		orReplace := p.tryConsumeKeywords(KeywordOr, KeywordReplace)
		if orReplace && !p.matchOneOfKeywords(KeywordTemporary, KeywordTable, KeywordView, KeywordFunction, KeywordDictionary) {
			return nil, fmt.Errorf("expected keyword: TEMPORARY|TABLE|VIEW|FUNCTION|DICTIONARY, but got %q", p.lastTokenText())
		}
		if isAttach {
			return p.parseAttach(pos, orReplace)
		}
		switch {
		case p.matchKeyword(KeywordNamed):
			return p.parseCreateNamedCollection(pos)
		case p.matchKeyword(KeywordDatabase):
			return p.parseCreateDatabase(pos)
		case p.matchKeyword(KeywordDictionary):
			stmt, err := p.parseCreateDictionary(pos, orReplace)
			if err != nil {
				return nil, err
			}
			if stmt.Schema == nil {
				return nil, fmt.Errorf("expected dictionary definition after CREATE DICTIONARY %s", stmt.Name.String())
			}
			return stmt, nil
		case p.matchKeyword(KeywordTable),
			p.matchKeyword(KeywordTemporary):
			return p.parseCreateTable(pos, orReplace)
		case p.matchKeyword(KeywordFunction):
			return p.parseCreateFunction(pos, orReplace)
		case p.matchKeyword(KeywordMaterialized):
			return p.parseCreateMaterializedView(pos)
		case p.matchKeyword(KeywordLive):
			return p.parseCreateLiveView(pos)
		case p.matchKeyword(KeywordView):
			return p.parseCreateView(pos, orReplace)
		case p.matchKeyword(KeywordRole):
			return p.parseCreateRole(pos)
		case p.matchKeyword(KeywordUser):
			return p.parseCreateUser(pos)
		case p.matchKeyword(KeywordIndex), p.matchUnquotedIdent("UNIQUE"):
			return p.parseCreateIndex(pos)
		default:
			return nil, fmt.Errorf("expected keyword: NAMED|DATABASE|DICTIONARY|TABLE|VIEW|ROLE|USER|FUNCTION|MATERIALIZED|INDEX, but got %q",
				p.lastTokenKind())
		}
	case p.matchKeyword(KeywordAlter):
		_ = p.lexer.consumeToken()
		switch {
		case p.matchKeyword(KeywordRole):
			return p.parseAlterRole(pos)
		case p.matchKeyword(KeywordTable):
			return p.parseAlterTable(pos)
		default:
			return nil, fmt.Errorf("expected keyword: TABLE|ROLE, but got %q", p.lastTokenText())
		}
	case p.matchKeyword(KeywordDetach):
		_ = p.lexer.consumeToken()
		return p.parseDetach(pos)
	case p.matchKeyword(KeywordDrop):
		_ = p.lexer.consumeToken()
		switch {
		case p.matchKeyword(KeywordDatabase):
			stmt, err := p.parseDropDatabase(pos)
			if err != nil {
				return nil, err
			}
			if stmt.Permanently {
				return nil, errors.New("PERMANENTLY is only valid with DETACH")
			}
			return stmt, nil
		case p.matchKeyword(KeywordTemporary),
			p.matchKeyword(KeywordView),
			p.matchKeyword(KeywordDictionary),
			p.matchKeyword(KeywordTable):
			stmt, err := p.parseDropStmt(pos)
			if err != nil {
				return nil, err
			}
			if stmt.Permanently {
				return nil, errors.New("PERMANENTLY is only valid with DETACH")
			}
			return stmt, nil
		case p.matchKeyword(KeywordUser),
			p.matchKeyword(KeywordRole):
			return p.parserDropUserOrRole(pos)
		case p.matchKeyword(KeywordIndex):
			return p.parseDropIndex(pos)
		default:
			return nil, fmt.Errorf("expected keyword: DATABASE|TABLE|INDEX, but got %q", p.lastTokenText())
		}
	case p.matchKeyword(KeywordTruncate):
		return p.parseTruncateTable(pos)
	case p.matchKeyword(KeywordRename):
		return p.parseRenameStmt(pos)
	}
	return nil, nil // nolint
}

// parseAttach parses ATTACH TABLE|VIEW|MATERIALIZED VIEW|LIVE VIEW|DICTIONARY|
// DATABASE. It shares the CREATE parsers and marks the result IsAttach so
// that it prints back as ATTACH. ClickHouse accepts no OR REPLACE and no
// FUNCTION, ROLE, USER or NAMED COLLECTION after ATTACH.
func (p *Parser) parseAttach(pos Pos, orReplace bool) (DDL, error) {
	if orReplace {
		return nil, errors.New("ATTACH does not support OR REPLACE")
	}
	switch {
	case p.matchKeyword(KeywordDatabase):
		stmt, err := p.parseCreateDatabase(pos)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	case p.matchKeyword(KeywordDictionary):
		stmt, err := p.parseCreateDictionary(pos, false)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	case p.matchKeyword(KeywordTable), p.matchKeyword(KeywordTemporary):
		stmt, err := p.parseCreateTable(pos, false)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	case p.matchKeyword(KeywordMaterialized):
		stmt, err := p.parseCreateMaterializedView(pos)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	case p.matchKeyword(KeywordLive):
		stmt, err := p.parseCreateLiveView(pos)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	case p.matchKeyword(KeywordView):
		stmt, err := p.parseCreateView(pos, false)
		if err != nil {
			return nil, err
		}
		stmt.IsAttach = true
		return stmt, nil
	default:
		return nil, fmt.Errorf("expected keyword: DATABASE|DICTIONARY|TABLE|VIEW|MATERIALIZED|LIVE after ATTACH, but got %q",
			p.lastTokenText())
	}
}

// matchPermanently reports whether the current token is the PERMANENTLY
// modifier of DETACH. It is not a reserved keyword, so it is matched as an
// unquoted identifier.
func (p *Parser) matchPermanently() bool {
	return isUnquotedIdent(p.last()) && strings.EqualFold(p.last().String, "PERMANENTLY")
}

// parseDetach parses DETACH TABLE|VIEW|DICTIONARY|DATABASE ... [PERMANENTLY]
// [SYNC]. It shares the DROP parsers and marks the result IsDetach so that it
// prints back as DETACH rather than DROP.
func (p *Parser) parseDetach(pos Pos) (DDL, error) {
	switch {
	case p.matchKeyword(KeywordDatabase):
		stmt, err := p.parseDropDatabase(pos)
		if err != nil {
			return nil, err
		}
		stmt.IsDetach = true
		return stmt, nil
	case p.matchKeyword(KeywordTemporary),
		p.matchKeyword(KeywordView),
		p.matchKeyword(KeywordDictionary),
		p.matchKeyword(KeywordTable):
		stmt, err := p.parseDropStmt(pos)
		if err != nil {
			return nil, err
		}
		stmt.IsDetach = true
		return stmt, nil
	default:
		return nil, fmt.Errorf("expected keyword: DATABASE|TABLE|VIEW|DICTIONARY after DETACH, but got %q", p.lastTokenText())
	}
}

func (p *Parser) parseCreateDatabase(pos Pos) (*CreateDatabase, error) {
	if err := p.expectKeyword(KeywordDatabase); err != nil {
		return nil, err
	}

	// try to parse IF NOT EXISTS clause
	ifNotExists, err := p.tryParseIfNotExists()
	if err != nil {
		return nil, err
	}
	// parse database name
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	StatementEnd := name.End()
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if onCluster != nil {
		StatementEnd = onCluster.End()
	}
	engineExpr, err := p.tryParseEngineExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	if engineExpr != nil {
		StatementEnd = engineExpr.End()
	}
	commentExpr, err := p.tryParseComment()
	if err != nil {
		return nil, err
	}
	if commentExpr != nil {
		StatementEnd = commentExpr.End()
	}
	return &CreateDatabase{
		CreatePos:    pos,
		StatementEnd: StatementEnd,
		Name:         name,
		IfNotExists:  ifNotExists,
		OnCluster:    onCluster,
		Engine:       engineExpr,
		Comment:      commentExpr,
	}, nil
}

func (p *Parser) parseCreateDictionary(pos Pos, orReplace bool) (*CreateDictionary, error) {
	if err := p.expectKeyword(KeywordDictionary); err != nil {
		return nil, err
	}

	createDict := &CreateDictionary{
		CreatePos: pos,
		OrReplace: orReplace,
	}

	// parse IF NOT EXISTS clause if exists
	var err error
	createDict.IfNotExists, err = p.tryParseIfNotExists()
	if err != nil {
		return nil, err
	}

	// parse dictionary name
	name, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	createDict.Name = name

	// try parse UUID clause if exists
	uuid, err := p.tryParseUUID()
	if err != nil {
		return nil, err
	}
	createDict.UUID = uuid

	// parse ON CLUSTER clause if exists
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createDict.OnCluster = onCluster
	createDict.StatementEnd = p.End()
	switch {
	case onCluster != nil:
		createDict.StatementEnd = onCluster.End()
	case uuid != nil:
		createDict.StatementEnd = uuid.End()
	default:
		createDict.StatementEnd = name.End()
	}

	// `ATTACH DICTIONARY name` re-attaches a detached dictionary and has no
	// definition; parseDDL rejects the short form for CREATE.
	if !p.matchTokenKind(TokenKindLParen) {
		return createDict, nil
	}

	// parse dictionary schema clause (required)
	schema, err := p.parseDictionarySchemaClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createDict.Schema = schema

	// parse dictionary engine clause (required)
	engine, err := p.parseDictionaryEngineClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createDict.Engine = engine
	createDict.StatementEnd = engine.End()

	// parse COMMENT clause if exists
	comment, err := p.tryParseComment()
	if err != nil {
		return nil, err
	}
	createDict.Comment = comment
	if comment != nil {
		createDict.StatementEnd = comment.End()
	}

	return createDict, nil
}

func (p *Parser) parseCreateNamedCollection(pos Pos) (*CreateNamedCollection, error) {
	if err := p.expectKeyword(KeywordNamed); err != nil {
		return nil, err
	}

	if err := p.expectKeyword(KeywordCollection); err != nil {
		return nil, err
	}

	createCollection := &CreateNamedCollection{
		CreatePos: pos,
	}

	// parse IF NOT EXISTS clause if exists
	var err error
	createCollection.IfNotExists, err = p.tryParseIfNotExists()
	if err != nil {
		return nil, err
	}

	// parse collection name
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	createCollection.Name = name

	// parse ON CLUSTER clause if exists
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createCollection.OnCluster = onCluster

	// parse AS keyword
	if err := p.expectKeyword(KeywordAs); err != nil {
		return nil, err
	}

	// parse parameters
	params := make([]*NamedCollectionParam, 0)
	for !p.atEOF() {
		param, err := p.parseNamedCollectionParam(p.Pos())
		if err != nil {
			return nil, err
		}
		params = append(params, param)

		// Check if there's another parameter
		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}
	createCollection.Params = params

	if len(params) > 0 {
		createCollection.StatementEnd = params[len(params)-1].End()
	} else if onCluster != nil {
		createCollection.StatementEnd = onCluster.End()
	} else {
		createCollection.StatementEnd = name.End()
	}

	return createCollection, nil
}

func (p *Parser) parseNamedCollectionParam(pos Pos) (*NamedCollectionParam, error) {
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindSingleEQ); err != nil {
		return nil, err
	}

	// Parse the value - can be string, number, or identifier
	var value Expr
	switch {
	case p.matchTokenKind(TokenKindString):
		literal, err := p.parseLiteral(p.Pos())
		if err != nil {
			return nil, err
		}
		value = literal
	case p.matchTokenKind(TokenKindInt), p.matchTokenKind(TokenKindFloat):
		literal, err := p.parseLiteral(p.Pos())
		if err != nil {
			return nil, err
		}
		value = literal
	case p.matchTokenKind(TokenKindIdent):
		ident, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		value = ident
	default:
		return nil, fmt.Errorf("expected string, number or identifier in named collection parameter, got %s", p.lastTokenKind())
	}

	param := &NamedCollectionParam{
		ParamPos: pos,
		Name:     name,
		Value:    value,
	}

	// Parse optional [NOT] OVERRIDABLE clause
	if p.tryConsumeKeywords(KeywordNot) {
		param.NotOverridable = true
		if err := p.expectKeyword(KeywordOverridable); err != nil {
			return nil, err
		}
	} else if p.tryConsumeKeywords(KeywordOverridable) {
		param.Overridable = true
	}
	if param.Overridable || param.NotOverridable {
		param.ParamEnd = p.prevEnd()
	}

	return param, nil
}

func (p *Parser) parseCreateTable(pos Pos, orReplace bool) (*CreateTable, error) {
	createTable := &CreateTable{CreatePos: pos, OrReplace: orReplace}
	createTable.HasTemporary = p.tryConsumeKeywords(KeywordTemporary)

	if err := p.expectKeyword(KeywordTable); err != nil {
		return nil, err
	}

	// parse IF NOT EXISTS clause if exists
	var err error
	createTable.IfNotExists, err = p.tryParseIfNotExists()
	if err != nil {
		return nil, err
	}
	if orReplace && createTable.IfNotExists {
		return nil, errors.New("OR REPLACE cannot be combined with IF NOT EXISTS")
	}

	tableIdentifier, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	createTable.Name = tableIdentifier

	// try parse UUID clause if exists
	uuid, err := p.tryParseUUID()
	if err != nil {
		return nil, err
	}
	createTable.UUID = uuid
	// parse ON CLUSTER clause if exists
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createTable.OnCluster = onCluster

	tableSchema, err := p.parseTableSchemaClause(p.Pos())
	if err != nil {
		return nil, err
	}
	createTable.TableSchema = tableSchema
	// The statement ends after its last part so far; ENGINE, AS, COMMENT and
	// SETTINGS below extend it.
	createTable.StatementEnd = p.prevEnd()

	engineExpr, err := p.tryParseEngineExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	if engineExpr != nil {
		createTable.Engine = engineExpr
		createTable.StatementEnd = engineExpr.End()

		// The TimeSeries engine accepts a tail of SAMPLES/DATA, TAGS and
		// METRICS target clauses after the engine expression.
		if strings.EqualFold(engineExpr.Name, "TimeSeries") {
			targets, err := p.parseTimeSeriesTargets()
			if err != nil {
				return nil, err
			}
			createTable.TimeSeriesTargets = targets
			// Targets are ordered by first occurrence, so a later part of
			// an earlier target can end after the last target.
			for _, target := range targets {
				if target.End() > createTable.StatementEnd {
					createTable.StatementEnd = target.End()
				}
			}
		}
	}

	if p.tryConsumeKeywords(KeywordAs) {
		// After AS, we can have: SELECT/WITH (with or without parens), or table_function(...)
		// Check if it's a SELECT/WITH query (explicitly check keywords/paren before ident)
		if p.matchKeyword(KeywordSelect) || p.matchKeyword(KeywordWith) || p.matchTokenKind(TokenKindLParen) {
			// It's a SELECT or WITH query (with or without parentheses)
			subQuery, err := p.parseSubQuery(p.Pos())
			if err != nil {
				return nil, err
			}
			createTable.SubQuery = subQuery
			createTable.StatementEnd = subQuery.End()
		} else if p.matchTokenKind(TokenKindIdent) {
			// It's a table function: remote(...), remoteSecure(...), etc.
			ident, err := p.parseIdent()
			if err != nil {
				return nil, err
			}
			if p.matchTokenKind(TokenKindLParen) {
				argsExpr, err := p.parseTableArgList(p.Pos())
				if err != nil {
					return nil, err
				}
				tableFunc := &TableFunctionExpr{
					Name: ident,
					Args: argsExpr,
				}
				createTable.TableFunction = tableFunc
				createTable.StatementEnd = tableFunc.End()
			} else {
				return nil, fmt.Errorf("expected ( after identifier in AS clause, got %q", p.lastTokenKind())
			}
		} else {
			return nil, fmt.Errorf("expected SELECT, WITH or identifier after AS, got %q", p.lastTokenKind())
		}
	}

	comment, err := p.tryParseComment()
	if err != nil {
		return nil, err
	}
	createTable.Comment = comment
	if comment != nil {
		createTable.StatementEnd = comment.End()
	}

	// A trailing SETTINGS clause holds query-level settings for the CREATE
	// itself. It is kept apart from the engine's storage SETTINGS.
	if p.matchKeyword(KeywordSettings) {
		settings, err := p.tryParseSettingsClause(p.Pos())
		if err != nil {
			return nil, err
		}
		createTable.Settings = settings
		createTable.StatementEnd = settings.End()
	}
	return createTable, nil
}

func (p *Parser) parseIdentOrFunction(_ Pos) (Expr, error) {
	ident, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	switch {
	case p.matchTokenKind(TokenKindLBracket):
		params, err := p.parseArrayParams(p.Pos())
		if err != nil {
			return nil, err
		}
		return &ObjectParams{
			Object: ident,
			Params: params,
		}, nil
	case p.matchTokenKind(TokenKindLParen):
		params, err := p.parseFunctionParams(p.Pos())
		if err != nil {
			return nil, err
		}
		funcExpr := &FunctionExpr{
			Name:   ident,
			Params: params,
		}

		overPos := p.Pos()
		if p.tryConsumeKeywords(KeywordOver) {
			var overExpr Expr
			switch {
			case p.matchTokenKind(TokenKindIdent):
				overExpr, err = p.parseIdent()
			case p.matchTokenKind(TokenKindLParen):
				overExpr, err = p.parseWindowCondition(p.Pos())
				if err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("expected window name or (, but got %q", p.lastTokenKind())
			}

			if err != nil {
				return nil, err
			}
			return &WindowFunctionExpr{
				Function: funcExpr,
				OverPos:  overPos,
				OverExpr: overExpr,
			}, nil
		}
		return funcExpr, nil
	case p.tryConsumeTokenKind(TokenKindDot) != nil:
		switch {
		case p.matchTokenKind(TokenKindIdent):
			fields := []*Ident{ident}
			for {
				child, err := p.parseIdent()
				if err != nil {
					return nil, err
				}
				fields = append(fields, child)
				if p.tryConsumeTokenKind(TokenKindDot) == nil {
					break
				}
			}
			return &Path{Fields: fields}, nil
		case p.matchTokenKind("*"):
			nextIdent, err := p.parseColumnStar(p.Pos())
			if err != nil {
				return nil, err
			}
			return &NestedIdentifier{
				Ident:    ident,
				DotIdent: nextIdent,
			}, nil
		case p.matchTokenKind(TokenKindInt):
			i, err := p.parseNumber(p.Pos())
			if err != nil {
				return nil, err
			}
			return &IndexOperation{
				Object:    ident,
				Operation: TokenKindDot,
				Index:     i,
			}, nil
		default:
			return nil, fmt.Errorf("expected IDENT, NUMBER or *, but got %q", p.lastTokenKind())
		}
	}
	return ident, nil
}

func (p *Parser) parseTableIdentifier(_ Pos) (*TableIdentifier, error) {
	ident, err := p.parseIdentOrString()
	if err != nil {
		return nil, err
	}
	dotIdent, err := p.tryParseDotIdentOrString(p.Pos())
	if err != nil {
		return nil, err
	}
	if dotIdent != nil {
		return &TableIdentifier{
			Database: ident,
			Table:    dotIdent,
		}, nil
	}
	return &TableIdentifier{
		Table: ident,
	}, nil
}

func (p *Parser) parseTableSchemaClause(pos Pos) (*TableSchemaClause, error) {
	switch {
	case p.matchTokenKind(TokenKindLParen):
		// parse column definitions
		if err := p.expectTokenKind(TokenKindLParen); err != nil {
			return nil, err
		}

		columns, err := p.parseTableColumns()
		if err != nil {
			return nil, err
		}
		// ClickHouse rejects an empty column list `()`.
		if len(columns) == 0 {
			return nil, fmt.Errorf("expected column definition, got %s", p.lastTokenKind())
		}

		rightParenPos := p.Pos()
		if err := p.expectTokenKind(TokenKindRParen); err != nil {
			return nil, err
		}
		return &TableSchemaClause{
			SchemaPos: pos,
			SchemaEnd: rightParenPos,
			Columns:   columns,
		}, nil
	case p.matchKeyword(KeywordAs) && !p.peekKeyword(KeywordSelect) && !p.peekKeyword(KeywordWith) && !p.peekTokenKind(TokenKindLParen):
		// Handle AS only if followed by identifier (not SELECT/WITH/LPAREN)
		// This handles: AS ident, AS ident.ident, AS ident(...)
		// CREATE TABLE will handle: AS SELECT, AS WITH, AS (SELECT ...)
		p.tryConsumeKeywords(KeywordAs)

		ident, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		switch {
		case p.matchTokenKind(TokenKindDot):
			// it's a database.table
			dotIdent, err := p.tryParseDotIdent(p.Pos())
			if err != nil {
				return nil, err
			}
			return &TableSchemaClause{
				SchemaPos: pos,
				SchemaEnd: dotIdent.End(),
				AliasTable: &TableIdentifier{
					Database: ident,
					Table:    dotIdent,
				},
			}, nil
		case p.matchTokenKind(TokenKindLParen):
			// it's a table function
			argsExpr, err := p.parseTableArgList(pos)
			if err != nil {
				return nil, err
			}
			return &TableSchemaClause{
				SchemaPos: pos,
				SchemaEnd: p.prevEnd(),
				TableFunction: &TableFunctionExpr{
					Name: ident,
					Args: argsExpr,
				},
			}, nil
		default:
			return &TableSchemaClause{
				SchemaPos: pos,
				SchemaEnd: p.prevEnd(),
				AliasTable: &TableIdentifier{
					Table: ident,
				},
			}, nil
		}
	}
	// no schema is ok for MATERIALIZED VIEW
	return nil, nil
}

func (p *Parser) parseTableColumns() ([]Expr, error) {
	columns := make([]Expr, 0)
	for !p.atEOF() {
		switch {
		case p.matchKeyword(KeywordIndex):
			indexPos := p.Pos()
			_ = p.lexer.consumeToken()
			index, err := p.parseTableIndex(indexPos)
			if err != nil {
				return nil, err
			}
			columns = append(columns, index)
		case p.matchKeyword(KeywordConstraint):
			constraintPos := p.Pos()
			_ = p.lexer.consumeToken()
			ident, err := p.parseIdent()
			if err != nil {
				return nil, err
			}
			if !p.matchOneOfKeywords(KeywordCheck, KeywordAssume) {
				return nil, fmt.Errorf("expected keyword: %s or %s, but got %s",
					KeywordCheck, KeywordAssume, p.lastTokenKind())
			}
			constraintTypeToken := p.last()
			constraintType := &Ident{
				NamePos: constraintTypeToken.Pos,
				NameEnd: constraintTypeToken.End,
				Name:    constraintTypeToken.String,
			}
			_ = p.lexer.consumeToken()
			expr, err := p.parseExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			columns = append(columns, &ConstraintClause{
				ConstraintPos: constraintPos,
				Constraint:    ident,
				Type:          constraintType,
				Expr:          expr,
			})
		case p.matchKeyword(KeywordProjection):
			projection, err := p.parseTableProjection(p.Pos(), true)
			if err != nil {
				return nil, err
			}
			columns = append(columns, projection)
		default:
			column, err := p.tryParseTableColumnExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			if column == nil {
				break
			}
			columns = append(columns, column)
		}
		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}
	// end of column definitions
	return columns, nil
}

func (p *Parser) tryParseTableColumnExpr(pos Pos) (*ColumnDef, error) {
	if !p.matchTokenKind(TokenKindIdent) {
		return nil, nil // nolint
	}
	return p.parseTableColumnExpr(pos)
}

func (p *Parser) parseTableColumnExpr(pos Pos) (*ColumnDef, error) {
	// Not a column definition, just return
	column := &ColumnDef{NamePos: pos}
	// parse column name
	name, err := p.ParseNestedIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	column.Name = name
	columnEnd := name.End()

	if p.matchTokenKind(TokenKindIdent) && !p.matchKeyword(KeywordRemove) {
		columnType, err := p.parseColumnType(p.Pos())
		if err != nil {
			return nil, err
		}
		column.Type = columnType
		columnEnd = columnType.End()
	}

	nullable := p.tryParseNull(p.Pos())
	if nullable != nil {
		columnEnd = nullable.End()
	}
	notNull, err := p.tryParseNotNull(p.Pos())
	if err != nil {
		return nil, err
	}
	if notNull != nil {
		columnEnd = notNull.End()
	}

	switch {
	case p.tryConsumeKeywords(KeywordDefault):
		column.DefaultExpr, err = p.parseExpr(p.Pos())
		if err == nil {
			columnEnd = column.DefaultExpr.End()
		}
	case p.tryConsumeKeywords(KeywordMaterialized):
		column.MaterializedExpr, err = p.parseExpr(p.Pos())
		if err == nil {
			columnEnd = column.MaterializedExpr.End()
		}
	case p.tryConsumeKeywords(KeywordEphemeral):
		// EPHEMERAL accepts an optional expression. Skip the expression when the
		// next token clearly ends the column definition (',' or ')') or starts
		// a subsequent column-def clause (COMMENT, CODEC, TTL).
		column.IsEphemeral = true
		if p.matchTokenKind(TokenKindComma) || p.matchTokenKind(")") ||
			p.matchKeyword(KeywordComment) || p.matchKeyword(KeywordCodec) || p.matchKeyword(KeywordTtl) {
			columnEnd = p.Pos()
		} else {
			column.EphemeralExpr, err = p.parseExpr(p.Pos())
			if err == nil {
				columnEnd = column.EphemeralExpr.End()
			}
		}
	case p.tryConsumeKeywords(KeywordAlias):
		column.AliasExpr, err = p.parseExpr(p.Pos())
		if err == nil {
			columnEnd = column.AliasExpr.End()
		}
	}
	if err != nil {
		return nil, err
	}

	comment, err := p.tryParseColumnComment(p.Pos())
	if err != nil {
		return nil, err
	}
	if comment != nil {
		columnEnd = comment.End()
	}

	codec, err := p.tryParseCompressionCodecs(p.Pos())
	if err != nil {
		return nil, err
	}
	if codec != nil {
		columnEnd = codec.End()
	}
	ttl, err := p.tryParseTTLClause(p.Pos(), false)
	if err != nil {
		return nil, err
	}
	if ttl != nil {
		columnEnd = ttl.End()
	}
	column.TTL = ttl

	column.ColumnEnd = columnEnd
	column.Comment = comment
	column.Codec = codec
	column.Nullable = nullable
	column.NotNull = notNull
	return column, nil
}

func (p *Parser) parseTableArgExpr(pos Pos) (Expr, error) {
	switch {
	case p.matchTokenKind(TokenKindIdent):
		ident, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		switch {
		// nest identifier
		case p.matchTokenKind(TokenKindDot):
			dotIdent, err := p.tryParseDotIdent(p.Pos())
			if err != nil {
				return nil, err
			}
			return &NestedIdentifier{
				Ident:    ident,
				DotIdent: dotIdent,
			}, nil
		case p.matchTokenKind(TokenKindLParen):
			argsExpr, err := p.parseTableArgList(pos)
			if err != nil {
				return nil, err
			}
			return &TableFunctionExpr{
				Name: ident,
				Args: argsExpr,
			}, nil
		default:
			return ident, nil
		}
	case p.matchTokenKind(TokenKindLParen):
		return p.parseSubQuery(p.Pos())
	case p.matchTokenKind(TokenKindInt), p.matchTokenKind(TokenKindString), p.matchKeyword(KeywordNull):
		return p.parseLiteral(p.Pos())
	default:
		return nil, fmt.Errorf("unexpected token: %q, expected <Name>, <literal>", p.lastTokenKind())
	}
}

// parseTableArg parses one table-function argument. The common forms
// (a name, db.table, a nested table function, a subquery or a literal) keep
// their table-argument AST; anything else, such as `currentDatabase() || 'x'`,
// `n + 1` or `-1`, is parsed as a general expression (#129).
func (p *Parser) parseTableArg(pos Pos) (Expr, error) {
	state := p.lexer.saveState()
	arg, err := p.parseTableArgExpr(pos)
	if err == nil && (p.matchTokenKind(TokenKindComma) || p.matchTokenKind(TokenKindRParen)) {
		return arg, nil
	}
	p.lexer.restoreState(state)
	return p.parseExpr(pos)
}

func (p *Parser) parseTableArgList(pos Pos) (*TableArgListExpr, error) {
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	args := make([]Expr, 0)
	// A table function may take no arguments, e.g. currentDatabase().
	for !p.atEOF() && !p.matchTokenKind(TokenKindRParen) {
		// Check if this is a named parameter (identifier followed by =)
		var arg Expr
		var err error

		// Try to detect named parameter pattern: last token is identifier, next token is =
		isNamedParam := false
		lastKind := p.lastTokenKind()
		if lastKind == TokenKindIdent || lastKind == TokenKindKeyword {
			// Last token is an identifier, peek at the next token
			nextToken, peekErr := p.lexer.peekToken()

			if peekErr == nil && nextToken != nil && nextToken.Kind == TokenKindSingleEQ {
				isNamedParam = true
			}
		}

		if isNamedParam {
			// Parse as named parameter - the identifier is already the last token
			// We need to get it, consume the =, and parse the value
			name := &Ident{
				NamePos: p.last().Pos,
				NameEnd: p.last().End,
				Name:    p.last().String,
			}
			// Consume the = token
			if err := p.lexer.consumeToken(); err != nil {
				return nil, err
			}
			if err := p.expectTokenKind(TokenKindSingleEQ); err != nil {
				return nil, err
			}
			// Parse the value
			value, err := p.parseTableArg(p.Pos())
			if err != nil {
				return nil, err
			}
			arg = &NamedParameterExpr{
				NamePos: name.NamePos,
				Name:    name,
				Value:   value,
			}
		} else {
			// Parse as regular table arg expression
			arg, err = p.parseTableArg(p.Pos())
		}

		if err != nil {
			return nil, err
		}
		args = append(args, arg)
		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}

	rightParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	return &TableArgListExpr{
		LeftParenPos:  pos,
		RightParenPos: rightParenPos,
		Args:          args,
	}, nil
}

func (p *Parser) tryParseClusterClause(pos Pos) (*ClusterClause, error) {
	if !p.tryConsumeKeywords(KeywordOn) {
		return nil, nil // nolint
	}
	if err := p.expectKeyword(KeywordCluster); err != nil {
		return nil, err
	}

	var expr Expr
	var err error
	switch {
	case p.matchTokenKind(TokenKindIdent):
		expr, err = p.parseIdent()
	case p.matchTokenKind(TokenKindString):
		expr, err = p.parseString(p.Pos())
	default:
		return nil, fmt.Errorf("unexpected token: %q, expected <IDENT> or <STRING>", p.lastTokenText())
	}
	if err != nil {
		return nil, err
	}
	return &ClusterClause{
		OnPos: pos,
		Expr:  expr,
	}, nil
}

func (p *Parser) tryParsePartitionByClause(pos Pos) (*PartitionByClause, error) {
	if !p.tryConsumeKeywords(KeywordPartition) {
		return nil, nil // nolint
	}

	if err := p.expectKeyword(KeywordBy); err != nil {
		return nil, err
	}

	// parse partition key list
	columnExpr, err := p.parseColumnExprListWithLParen(p.Pos())
	if err != nil {
		return nil, err
	}
	return &PartitionByClause{
		PartitionPos: pos,
		Expr:         columnExpr,
	}, nil
}

func (p *Parser) tryParsePrimaryKeyClause(pos Pos) (*PrimaryKeyClause, error) {
	if !p.tryConsumeKeywords(KeywordPrimary) {
		return nil, nil // nolint
	}

	if err := p.expectKeyword(KeywordKey); err != nil {
		return nil, err
	}

	// parse partition key list
	columnExpr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &PrimaryKeyClause{
		PrimaryPos: pos,
		Expr:       columnExpr,
	}, nil
}

func (p *Parser) tryParseOrderByClause(pos Pos) (*OrderByClause, error) {
	if !p.tryConsumeKeywords(KeywordOrder) {
		return nil, nil // nolint
	}

	if err := p.expectKeyword(KeywordBy); err != nil {
		return nil, err
	}
	return p.parseOrderByClause(pos)
}

func (p *Parser) parseOrderByClause(pos Pos) (*OrderByClause, error) {
	orderByListExpr := &OrderByClause{OrderPos: pos, ListEnd: pos}
	items := make([]Expr, 0)
	for {
		expr, err := p.parseOrderExpr(pos)
		if err != nil {
			return nil, err
		}
		if expr == nil {
			break
		}
		items = append(items, expr)

		if p.atEOF() || p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}
	if len(items) > 0 {
		orderByListExpr.ListEnd = items[len(items)-1].End()
	}
	orderByListExpr.Items = items

	// Parse optional INTERPOLATE clause
	if p.matchKeyword(KeywordInterpolate) {
		interpolatePos := p.Pos()
		_ = p.lexer.consumeToken()
		interpolate, err := p.parseInterpolateClause(interpolatePos)
		if err != nil {
			return nil, err
		}
		orderByListExpr.Interpolate = interpolate
		orderByListExpr.ListEnd = interpolate.End()
	}

	return orderByListExpr, nil
}

func (p *Parser) parseOrderExpr(pos Pos) (*OrderExpr, error) {
	// parse column expr
	columnExpr, err := p.parseExpr(pos)
	if err != nil {
		return nil, err
	}

	var alias *Ident
	if p.matchKeyword(KeywordAs) {
		// It should be a subquery instead of an order by alias if the `AS` is followed by `SELECT` keyword.
		nextToken, err := p.lexer.peekToken()
		if err != nil {
			return nil, err
		}
		if nextToken != nil && nextToken.ToString() == KeywordSelect {
			return &OrderExpr{
				OrderPos: pos,
				Expr:     columnExpr,
			}, nil
		}
		// consume the `AS` keyword
		_ = p.lexer.consumeToken()
		alias, err = p.parseIdent()
		if err != nil {
			return nil, err
		}
	} else if p.matchKeyword(KeywordTtl) {
		return &OrderExpr{
			OrderPos: pos,
			Expr:     columnExpr,
		}, nil
	}

	direction := OrderDirectionNone
	switch {
	case p.matchKeyword(KeywordAsc), p.matchKeyword(KeywordAscending):
		direction = OrderDirectionAsc
		_ = p.lexer.consumeToken()
	case p.matchKeyword(KeywordDesc), p.matchKeyword(KeywordDescending):
		direction = OrderDirectionDesc
		_ = p.lexer.consumeToken()
	}

	// [NULLS FIRST|LAST] [COLLATE 'locale'], in this order as in ClickHouse.
	nulls := ""
	if p.tryConsumeKeywords(KeywordNulls) {
		switch {
		case p.tryConsumeKeywords(KeywordFirst):
			nulls = KeywordFirst
		case p.tryConsumeKeywords(KeywordLast):
			nulls = KeywordLast
		default:
			return nil, fmt.Errorf("expected FIRST or LAST after NULLS, got %s", p.lastTokenKind())
		}
	}
	var collate *StringLiteral
	if p.tryConsumeKeywords(KeywordCollate) {
		collate, err = p.parseString(p.Pos())
		if err != nil {
			return nil, err
		}
	}

	// Parse optional WITH FILL clause
	var fill *Fill
	if p.matchKeyword(KeywordWith) && p.peekKeyword(KeywordFill) {
		_ = p.lexer.consumeToken() // WITH
		fillPos := p.Pos()         // FILL
		_ = p.lexer.consumeToken()
		fill, err = p.parseFillClause(fillPos)
		if err != nil {
			return nil, err
		}
	}

	return &OrderExpr{
		OrderPos:  pos,
		Alias:     alias,
		Expr:      columnExpr,
		Direction: direction,
		Nulls:     nulls,
		Collate:   collate,
		Fill:      fill,
		OrderEnd:  p.prevEnd(),
	}, nil
}

func (p *Parser) parseFillClause(fillPos Pos) (*Fill, error) {
	fill := &Fill{FillPos: fillPos}

	// Parse optional FROM clause
	if p.tryConsumeKeywords(KeywordFrom) {
		fromExpr, err := p.parseExpr(fillPos)
		if err != nil {
			return nil, err
		}
		fill.From = fromExpr
	}

	// Parse optional TO clause
	if p.tryConsumeKeywords(KeywordTo) {
		toExpr, err := p.parseExpr(fillPos)
		if err != nil {
			return nil, err
		}
		fill.To = toExpr
	}

	// Parse optional STEP clause
	if p.tryConsumeKeywords(KeywordStep) {
		stepExpr, err := p.parseExpr(fillPos)
		if err != nil {
			return nil, err
		}
		fill.Step = stepExpr
	}

	// Parse optional STALENESS clause
	if p.tryConsumeKeywords(KeywordStaleness) {
		stalenessExpr, err := p.parseExpr(fillPos)
		if err != nil {
			return nil, err
		}
		fill.Staleness = stalenessExpr
	}

	return fill, nil
}

func (p *Parser) parseInterpolateClause(interpolatePos Pos) (*InterpolateClause, error) {
	interpolate := &InterpolateClause{
		InterpolatePos: interpolatePos,
		ListEnd:        interpolatePos + Pos(len("INTERPOLATE")),
	}

	if p.tryConsumeTokenKind(TokenKindLParen) == nil {
		// INTERPOLATE without columns is valid
		return interpolate, nil
	}

	items := make([]*InterpolateItem, 0)
	for {
		column, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		item := &InterpolateItem{Column: column}

		if p.tryConsumeKeywords(KeywordAs) {
			expr, err := p.parseExpr(interpolatePos)
			if err != nil {
				return nil, err
			}
			item.Expr = expr
		}

		items = append(items, item)

		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}

	rparen := p.tryConsumeTokenKind(TokenKindRParen)
	if rparen == nil {
		return nil, fmt.Errorf("expected ')' after INTERPOLATE column list")
	}

	interpolate.Items = items
	interpolate.ListEnd = rparen.End

	return interpolate, nil
}

func (p *Parser) tryParseTTLClause(pos Pos, allowMultiValues bool) (*TTLClause, error) {
	if !p.tryConsumeKeywords(KeywordTtl) {
		return nil, nil // nolint
	}
	ttlExprList := &TTLClause{TTLPos: pos, ListEnd: pos}
	// accept the TTL keyword
	items, err := p.parseTTLClause(pos, allowMultiValues)
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		ttlExprList.ListEnd = items[len(items)-1].End()
	}
	ttlExprList.Items = items
	return ttlExprList, nil
}

// parseTTLClause parses the TTL clause.
// allowMultiValues is used to determine whether to allow multiple TTL values.
func (p *Parser) parseTTLClause(pos Pos, allowMultiValues bool) ([]*TTLExpr, error) {
	items := make([]*TTLExpr, 0)
	expr, err := p.parseTTLExpr(pos)
	if err != nil {
		return nil, err
	}
	items = append(items, expr)
	for allowMultiValues && !p.atEOF() && p.tryConsumeTokenKind(TokenKindComma) != nil {
		expr, err = p.parseTTLExpr(pos)
		if err != nil {
			return nil, err
		}
		items = append(items, expr)
	}
	return items, nil
}

func (p *Parser) tryParseTTLPolicy(pos Pos) (*TTLPolicy, error) {
	var rule *TTLPolicyRule
	switch {
	case p.tryConsumeKeywords(KeywordTo):
		if p.tryConsumeKeywords(KeywordDisk) {
			value, err := p.parseString(p.Pos())
			if err != nil {
				return nil, err
			}
			rule = &TTLPolicyRule{RulePos: pos, ToDisk: value}
		} else if p.tryConsumeKeywords(KeywordVolume) {
			value, err := p.parseString(p.Pos())
			if err != nil {
				return nil, err
			}
			rule = &TTLPolicyRule{RulePos: pos, ToVolume: value}
		} else {
			return nil, fmt.Errorf("unexpected token: %q, expected DISK or VOLUME", p.lastTokenKind())
		}
	case p.matchKeyword(KeywordDelete), p.matchKeyword(KeywordRecompress):
		token := p.last()
		_ = p.lexer.consumeToken()
		action := &TTLPolicyRuleAction{
			ActionPos: token.Pos,
			ActionEnd: token.End,
			Action:    token.ToString(),
		}
		codec, err := p.tryParseCompressionCodecs(p.Pos())
		if err != nil {
			return nil, err
		}
		action.Codec = codec
		rule = &TTLPolicyRule{RulePos: pos, Action: action}
	}

	where, err := p.tryParseWhereClause(p.Pos())
	if err != nil {
		return nil, err
	}

	groupBy, err := p.tryParseGroupByClause(p.Pos())
	if err != nil {
		return nil, err
	}

	if rule == nil && where == nil && groupBy == nil {
		return nil, nil // nolint
	}
	policy := &TTLPolicy{Item: rule, Where: where, GroupBy: groupBy}
	if groupBy != nil && p.tryConsumeKeywords(KeywordSet) {
		for {
			pos := p.Pos()
			column, err := p.ParseNestedIdentifier(pos)
			if err != nil {
				return nil, err
			}
			if err := p.expectTokenKind(TokenKindSingleEQ); err != nil {
				return nil, err
			}
			expr, err := p.parseExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			policy.Assignments = append(policy.Assignments, &UpdateAssignment{
				AssignmentPos: pos, Column: column, Expr: expr,
			})
			if !p.matchTokenKind(TokenKindComma) {
				break
			}
			// A comma may begin another TTL rule instead of an assignment.
			saved := p.lexer.saveState()
			_ = p.tryConsumeTokenKind(TokenKindComma)
			if p.last() == nil || p.matchTokenKind(TokenKindEOF) || p.matchTokenKind(";") {
				return nil, fmt.Errorf("expected assignment or TTL expression after comma")
			}
			afterComma := p.lexer.saveState()
			_, err = p.ParseNestedIdentifier(p.Pos())
			isAssignment := err == nil && p.matchTokenKind(TokenKindSingleEQ)
			p.lexer.restoreState(saved)
			if !isAssignment {
				break
			}
			p.lexer.restoreState(afterComma)
		}
	}
	return policy, nil
}

func (p *Parser) parseTTLExpr(pos Pos) (*TTLExpr, error) {
	columnExpr, err := p.parseExpr(pos)
	if err != nil {
		return nil, err
	}
	policy, err := p.tryParseTTLPolicy(p.Pos())
	if err != nil {
		return nil, err
	}
	return &TTLExpr{
		TTLPos: pos,
		Expr:   columnExpr,
		Policy: policy,
	}, nil
}

func (p *Parser) tryParseSampleByClause(pos Pos) (*SampleByClause, error) {
	if !p.tryConsumeKeywords(KeywordSample) {
		return nil, nil // nolint
	}

	if err := p.expectKeyword(KeywordBy); err != nil {
		return nil, err
	}

	// parse sample by expr
	columnExpr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &SampleByClause{
		SamplePos: pos,
		Expr:      columnExpr,
	}, nil
}

func (p *Parser) tryParseSettingsClause(pos Pos) (*SettingsClause, error) {
	if !p.tryConsumeKeywords(KeywordSettings) {
		return nil, nil // nolint
	}
	return p.parseSettingsClause(pos)
}

func (p *Parser) parseSettingsClause(pos Pos) (*SettingsClause, error) {
	settings := &SettingsClause{SettingsPos: pos, ListEnd: pos}
	items, err := p.parseSettingsList(p.Pos())
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("settings list is empty")
	}
	settings.ListEnd = items[len(items)-1].End()
	settings.Items = items
	return settings, nil
}

func (p *Parser) parseSettingsList(pos Pos) ([]*SettingExpr, error) {
	items := make([]*SettingExpr, 0)
	expr, err := p.parseSettingsExpr(pos)
	if err != nil {
		return nil, err
	}
	items = append(items, expr)
	for p.tryConsumeTokenKind(TokenKindComma) != nil {
		expr, err = p.parseSettingsExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		items = append(items, expr)
	}
	return items, nil
}

func (p *Parser) parseSettingsExpr(pos Pos) (*SettingExpr, error) {
	ident, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindSingleEQ); err != nil {
		return nil, err
	}

	var expr Expr
	switch {
	case p.matchTokenKind(TokenKindInt):
		number, err := p.parseNumber(p.Pos())
		if err != nil {
			return nil, err
		}
		expr = number
	case p.matchTokenKind(TokenKindString):
		str, err := p.parseString(p.Pos())
		expr = str
		if err != nil {
			return nil, err
		}
	case p.matchTokenKind(TokenKindLBrace):
		m, err := p.parseMapLiteral(p.Pos())
		if err != nil {
			return nil, err
		}
		expr = m
	case p.matchKeyword(KeywordTrue), p.matchKeyword(KeywordFalse):
		// Handle TRUE/FALSE keywords as boolean literals
		lastToken := p.last()
		_ = p.lexer.consumeToken()
		expr = &BoolLiteral{
			LiteralPos: lastToken.Pos,
			LiteralEnd: lastToken.End,
			Literal:    lastToken.String,
		}
	default:
		return nil, fmt.Errorf("unexpected token: %q, expected <number>, <bool> or <string>", p.lastTokenKind())
	}

	return &SettingExpr{
		SettingsPos: pos,
		Name:        ident,
		Expr:        expr,
	}, nil
}

func (p *Parser) parseDestinationClause(pos Pos) (*DestinationClause, error) {
	if err := p.expectKeyword(KeywordTo); err != nil {
		return nil, err
	}

	tableIdentifier, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	return &DestinationClause{
		ToPos:           pos,
		TableIdentifier: tableIdentifier,
	}, nil
}

func (p *Parser) tryParseEngineExpr(pos Pos) (*EngineExpr, error) {
	if !p.matchKeyword(KeywordEngine) {
		return nil, nil // nolint
	}
	return p.parseEngineExpr(pos)
}

func (p *Parser) parseEngineExpr(pos Pos) (*EngineExpr, error) {
	if err := p.expectKeyword(KeywordEngine); err != nil {
		return nil, err
	}
	_ = p.tryConsumeTokenKind(TokenKindSingleEQ)

	engineExpr := &EngineExpr{EnginePos: pos}
	var engineEnd Pos
	switch {
	case p.matchTokenKind(TokenKindIdent):
		ident, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		engineExpr.Name = ident.Name
		engineEnd = ident.End()
		if p.matchTokenKind(TokenKindLParen) {
			params, err := p.parseFunctionParams(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.Params = params
			engineEnd = params.End()
		}
	default:
		return nil, fmt.Errorf("unexpected token: %s", p.lastTokenKind())
	}

	for !p.atEOF() {
		switch {
		case p.matchKeyword(KeywordOrder):
			if engineExpr.OrderBy != nil {
				return nil, fmt.Errorf("duplicate ORDER BY clause")
			}
			orderBy, err := p.tryParseOrderByClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.OrderBy = orderBy
			engineEnd = orderBy.End()
		case p.matchKeyword(KeywordPartition):
			if engineExpr.PartitionBy != nil {
				return nil, fmt.Errorf("duplicate PARTITION BY clause")
			}
			partitionBy, err := p.tryParsePartitionByClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.PartitionBy = partitionBy
			engineEnd = partitionBy.End()
		case p.matchKeyword(KeywordPrimary):
			if engineExpr.PrimaryKey != nil {
				return nil, fmt.Errorf("duplicate PRIMARY KEY clause")
			}
			primaryKey, err := p.tryParsePrimaryKeyClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.PrimaryKey = primaryKey
			engineEnd = primaryKey.End()
		case p.matchKeyword(KeywordSample):
			if engineExpr.SampleBy != nil {
				return nil, fmt.Errorf("duplicate SAMPLE BY clause")
			}
			sampleBy, err := p.tryParseSampleByClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.SampleBy = sampleBy
			engineEnd = sampleBy.End()
		case p.matchKeyword(KeywordTtl):
			if engineExpr.TTL != nil {
				return nil, fmt.Errorf("duplicate TTL clause")
			}
			ttl, err := p.tryParseTTLClause(p.Pos(), true)
			if err != nil {
				return nil, err
			}
			engineExpr.TTL = ttl
			engineEnd = ttl.End()
		case p.matchKeyword(KeywordSettings):
			if engineExpr.Settings != nil {
				// A second SETTINGS clause ends the storage definition; the
				// caller parses it as the statement's query-level settings.
				engineExpr.EngineEnd = engineEnd
				return engineExpr, nil
			}
			settingsClause, err := p.tryParseSettingsClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engineExpr.Settings = settingsClause
			engineEnd = settingsClause.End()
		default:
			engineExpr.EngineEnd = engineEnd
			return engineExpr, nil
		}
	}
	engineExpr.EngineEnd = engineEnd
	return engineExpr, nil
}

// timeSeriesTargetKind normalises a TimeSeries target keyword to its slot.
// SAMPLES and its backwards-compat alias DATA both map to "samples". The
// two-word RECENT SAMPLES target is matched by matchTimeSeriesTarget.
func timeSeriesTargetKind(s string) (kind string, ok bool) {
	switch strings.ToUpper(s) {
	case "SAMPLES", "DATA":
		return "samples", true
	case "TAGS":
		return "tags", true
	case "METRICS":
		return "metrics", true
	}
	return "", false
}

// isUnquotedIdent reports whether token is an identifier written without
// quotes. A quoted identifier (e.g. `DATA`) is never treated as a keyword.
func isUnquotedIdent(token *Token) bool {
	return token != nil && token.Kind == TokenKindIdent &&
		token.QuoteType != DoubleQuote && token.QuoteType != BackTicks
}

// matchTimeSeriesTarget reports the normalised slot for the current token when
// it starts a TimeSeries target keyword: SAMPLES, DATA, TAGS, METRICS, or
// RECENT followed by SAMPLES.
func (p *Parser) matchTimeSeriesTarget() (kind string, ok bool) {
	if !isUnquotedIdent(p.last()) {
		return "", false
	}
	if strings.EqualFold(p.last().String, "RECENT") {
		next, err := p.lexer.peekToken()
		if err != nil || !isUnquotedIdent(next) || !strings.EqualFold(next.String, "SAMPLES") {
			return "", false
		}
		return "recent_samples", true
	}
	return timeSeriesTargetKind(p.last().String)
}

// parseTimeSeriesTargets parses the optional tail of SAMPLES/DATA, TAGS,
// METRICS and RECENT SAMPLES target clauses that may follow an
// `ENGINE = TimeSeries` expression. Each occurrence of a target keyword
// introduces one part of that target:
//
//	<KEYWORD> [db.]table
//	<KEYWORD> INNER UUID 'uuid'
//	<KEYWORD> INNER COLUMNS (...)
//	<KEYWORD> [INNER] ENGINE = engine ...
//
// As in ClickHouse, the parts of one target may be spread over several
// occurrences in any order, and each part may be given at most once. The parts
// are collected into a single clause per target (DATA and SAMPLES share the
// "samples" target), ordered by the first occurrence of each target.
func (p *Parser) parseTimeSeriesTargets() ([]*TimeSeriesTargetClause, error) {
	var targets []*TimeSeriesTargetClause
	byKind := make(map[string]*TimeSeriesTargetClause)
	for {
		kind, ok := p.matchTimeSeriesTarget()
		if !ok {
			break
		}
		kwToken := p.last()
		keyword := kwToken.String
		kwEnd := kwToken.End
		_ = p.lexer.consumeToken() // consume the target keyword identifier
		if kind == "recent_samples" {
			keyword += " " + p.last().String
			kwEnd = p.last().End
			_ = p.lexer.consumeToken() // consume SAMPLES
		}

		clause := byKind[kind]
		if clause == nil {
			clause = &TimeSeriesTargetClause{
				KindPos: kwToken.Pos,
				KindEnd: kwEnd,
				Kind:    kind,
				Keyword: keyword,
			}
			byKind[kind] = clause
			targets = append(targets, clause)
		}
		duplicate := func(part string) error {
			return fmt.Errorf("duplicate TimeSeries target clause %s%s", keyword, part)
		}

		var partEnd Pos
		switch {
		case p.tryConsumeKeywords(KeywordInner):
			switch {
			case p.matchKeyword(KeywordUuid):
				if clause.InnerUUID != nil {
					return nil, duplicate(" INNER UUID")
				}
				uuid, err := p.parseUUID()
				if err != nil {
					return nil, err
				}
				clause.InnerUUID = uuid
				partEnd = uuid.End()
			case p.tryConsumeKeywords(KeywordColumns):
				if clause.InnerColumns != nil {
					return nil, duplicate(" INNER COLUMNS")
				}
				columns, err := p.parseTableSchemaClause(p.Pos())
				if err != nil {
					return nil, err
				}
				if columns == nil {
					return nil, fmt.Errorf("expected ( after %s INNER COLUMNS", keyword)
				}
				clause.InnerColumns = columns
				partEnd = columns.End()
			case p.matchKeyword(KeywordEngine):
				if clause.InnerEngine != nil {
					return nil, duplicate(" INNER ENGINE")
				}
				innerEngine, err := p.parseEngineExpr(p.Pos())
				if err != nil {
					return nil, err
				}
				clause.InnerEngine = innerEngine
				clause.EngineShorthand = false
				partEnd = innerEngine.End()
			default:
				return nil, fmt.Errorf("expected UUID, COLUMNS or ENGINE after %s INNER, got %s", keyword, p.lastTokenKind())
			}
		case p.matchKeyword(KeywordEngine):
			// ClickHouse emits `<KEYWORD> ENGINE = ...` for auto-generated
			// TimeSeries targets in SHOW CREATE TABLE, omitting INNER.
			if clause.InnerEngine != nil {
				return nil, duplicate(" INNER ENGINE")
			}
			innerEngine, err := p.parseEngineExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			clause.InnerEngine = innerEngine
			clause.EngineShorthand = true
			partEnd = innerEngine.End()
		default:
			if clause.External != nil {
				return nil, duplicate(" table")
			}
			external, err := p.parseTableIdentifier(p.Pos())
			if err != nil {
				return nil, err
			}
			clause.External = external
			partEnd = external.End()
		}
		if partEnd > clause.KindEnd {
			clause.KindEnd = partEnd
		}
	}
	return targets, nil
}

func (p *Parser) parseStmt(pos Pos) (Expr, error) {
	expr, err := p.parseStmtBody(pos)
	if err != nil {
		return nil, err
	}
	// Statement can be terminated by ';' or EOF
	if p.last() != nil && !p.matchTokenKind(";") {
		return nil, fmt.Errorf("<EOF> or ';' was expected, but got: %q", p.lastTokenText())
	}
	return expr, nil
}

// parseStmtBody parses one statement with its trailing output clauses, but not
// the terminating ';' or end of input. EXPLAIN uses it for the explained
// statement.
func (p *Parser) parseStmtBody(pos Pos) (Expr, error) {
	var err error
	var expr Expr
	switch {
	case p.matchKeyword(KeywordCreate),
		p.matchKeyword(KeywordAttach),
		p.matchKeyword(KeywordAlter),
		p.matchKeyword(KeywordDrop),
		p.matchKeyword(KeywordDetach),
		p.matchKeyword(KeywordTruncate),
		p.matchKeyword(KeywordRename):
		expr, err = p.parseDDL(pos)
	case p.matchKeyword(KeywordSelect), p.matchKeyword(KeywordWith), p.matchTokenKind(TokenKindLParen):
		expr, err = p.parseSelectQuery(pos)
	case p.matchKeyword(KeywordDelete):
		expr, err = p.parseDeleteClause(pos)
	case p.matchKeyword(KeywordInsert):
		expr, err = p.parseInsertStmt(p.Pos())
	case p.matchKeyword(KeywordUse):
		expr, err = p.parseUseStmt(pos)
	case p.matchKeyword(KeywordSet):
		expr, err = p.parseSetStmt(pos)
	case p.matchKeyword(KeywordSettings):
		expr, err = p.parseSettingsStmt(pos)
	case p.matchKeyword(KeywordSystem):
		expr, err = p.parseSystemStmt(pos)
	case p.matchKeyword(KeywordOptimize):
		expr, err = p.parseOptimizeStmt(pos)
	case p.matchKeyword(KeywordCheck):
		expr, err = p.parseCheckStmt(pos)
	case p.matchKeyword(KeywordExplain):
		expr, err = p.parseExplainStmt(pos)
	case p.matchKeyword(KeywordGrant):
		expr, err = p.parseGrantPrivilegeStmt(pos)
	case p.matchKeyword(KeywordShow):
		expr, err = p.parseShowStmt(pos)
	case p.matchKeyword(KeywordDesc), p.matchKeyword(KeywordDescribe):
		expr, err = p.parseDescribeStmt(pos)
	default:
		if p.last() == nil {
			return nil, errors.New("unexpected end of input")
		}
		return nil, fmt.Errorf("unexpected token: %q", p.lastTokenText())
	}
	if err != nil {
		return nil, err
	}
	if err := p.tryParseQueryOutput(expr); err != nil {
		return nil, err
	}
	return expr, nil
}

// tryParseQueryOutput parses the query output clauses `[FORMAT fmt]
// [SETTINGS ...]` that may end a statement into its embedded OutputClauses
// (#41). SELECT and INSERT handle FORMAT themselves. A statement that does
// not embed OutputClauses is one ClickHouse accepts no output clauses for
// (GRANT, SYSTEM, USE, SET, DELETE, USER/ROLE statements, CREATE FUNCTION,
// CREATE NAMED COLLECTION), so the clause is rejected instead of silently
// dropped.
func (p *Parser) tryParseQueryOutput(stmt Expr) error {
	switch stmt.(type) {
	case *SelectQuery, *InsertStmt:
		return nil
	}
	if !p.matchKeyword(KeywordFormat) && !p.matchKeyword(KeywordSettings) {
		return nil
	}
	holder, ok := stmt.(outputClausesHolder)
	if !ok {
		return fmt.Errorf("%s is not supported after this statement", p.lastTokenText())
	}
	format, err := p.tryParseFormat(p.Pos())
	if err != nil {
		return err
	}
	settings, err := p.tryParseSettingsClause(p.Pos())
	if err != nil {
		return err
	}
	// CREATE TABLE's trailing query-level SETTINGS (CreateTable.Settings) is
	// this same clause written before FORMAT; ClickHouse allows it only once.
	if ct, ok := stmt.(*CreateTable); ok && ct.Settings != nil && settings != nil {
		return errors.New("duplicate query-level SETTINGS clause")
	}
	output := holder.outputClauses()
	output.Format = format
	output.OutputSettings = settings
	return nil
}

func (p *Parser) ParseStmts() ([]Expr, error) {
	var stmts []Expr
	for {
		if err := p.lexer.consumeToken(); err != nil {
			return nil, p.wrapError(err)
		}
		if p.atEOF() {
			break
		}
		if p.matchTokenKind(";") {
			continue
		}
		stmt, err := p.parseStmt(p.Pos())
		if err != nil {
			return nil, p.wrapError(err)
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}

func (p *Parser) parseUseStmt(pos Pos) (*UseStmt, error) {
	if err := p.expectKeyword(KeywordUse); err != nil {
		return nil, err
	}

	database, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	return &UseStmt{
		UsePos:       pos,
		Database:     database,
		StatementEnd: database.End(),
	}, nil
}

func (p *Parser) parseShowStmt(pos Pos) (*ShowStmt, error) {
	if err := p.expectKeyword(KeywordShow); err != nil {
		return nil, err
	}

	var showType string
	var target *TableIdentifier

	// Parse the type of SHOW statement
	switch {
	case p.matchKeyword(KeywordCreate):
		// SHOW CREATE TABLE table_name
		showType = "CREATE"
		_ = p.lexer.consumeToken()

		if err := p.expectKeyword(KeywordTable); err != nil {
			return nil, err
		}
		showType += " TABLE"

		tableIdent, err := p.parseTableIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		target = tableIdent

	case p.matchKeyword(KeywordDatabases):
		// SHOW DATABASES [optional clauses]
		showType = "DATABASES"
		_ = p.lexer.consumeToken()

	case p.matchKeyword(KeywordTables):
		// SHOW TABLES
		showType = "TABLES"
		_ = p.lexer.consumeToken()

	default:
		return nil, fmt.Errorf("expected CREATE, DATABASES, or TABLES after SHOW, got %q", p.lastTokenText())
	}

	stmt := &ShowStmt{
		ShowPos:  pos,
		ShowType: showType,
		Target:   target,
	}

	// Parse optional clauses for SHOW DATABASES
	if showType == "DATABASES" {
		// Parse [[NOT] LIKE | ILIKE '<pattern>']
		if p.matchKeyword(KeywordNot) {
			stmt.NotLike = true
			_ = p.lexer.consumeToken()
		}

		if p.matchKeyword(KeywordLike) || p.matchKeyword(KeywordIlike) {
			if p.matchKeyword(KeywordLike) {
				stmt.LikeType = "LIKE"
			} else {
				stmt.LikeType = "ILIKE"
			}
			_ = p.lexer.consumeToken()

			// Parse pattern expression
			pattern, err := p.parseExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			stmt.LikePattern = pattern
		}

		// Parse [LIMIT <N>]
		if p.matchKeyword(KeywordLimit) {
			_ = p.lexer.consumeToken()
			limit, err := p.parseExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			stmt.Limit = limit
		}

		// Parse [INTO OUTFILE filename]
		if p.matchKeyword(KeywordInto) {
			_ = p.lexer.consumeToken()
			if err := p.expectKeyword(KeywordOutfile); err != nil {
				return nil, err
			}

			// Parse filename as a string literal
			outFile, err := p.parseString(p.Pos())
			if err != nil {
				return nil, err
			}
			stmt.OutFile = outFile
		}

		// A trailing FORMAT is parsed by parseStmt as a query output clause.
	}

	// Set statement end position
	stmt.StatementEnd = p.prevEnd()

	return stmt, nil
}

func (p *Parser) parseDescribeStmt(pos Pos) (*DescribeStmt, error) {
	// DESC and DESCRIBE are both supported
	if !p.matchKeyword(KeywordDesc) && !p.matchKeyword(KeywordDescribe) {
		return nil, fmt.Errorf("expected DESC or DESCRIBE")
	}
	_ = p.lexer.consumeToken()

	// TABLE keyword is optional after DESC/DESCRIBE
	var describeType string
	if p.matchKeyword(KeywordTable) {
		_ = p.lexer.consumeToken()
		describeType = "TABLE"
	}

	tableIdent, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}

	return &DescribeStmt{
		DescribePos:  pos,
		StatementEnd: tableIdent.End(),
		DescribeType: describeType,
		Target:       tableIdent,
	}, nil
}

// syntax: TRUNCATE TEMPORARY? TABLE (IF EXISTS)? tableIdentifier clusterClause?;
func (p *Parser) parseTruncateTable(pos Pos) (*TruncateTable, error) {
	if err := p.expectKeyword(KeywordTruncate); err != nil {
		return nil, err
	}

	isTemporary := p.tryConsumeKeywords(KeywordTemporary)

	if err := p.expectKeyword(KeywordTable); err != nil {
		return nil, err
	}

	ifExists, err := p.tryParseIfExists()
	if err != nil {
		return nil, err
	}

	tableName, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}

	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}

	truncateTable := &TruncateTable{
		TruncatePos:  pos,
		IsTemporary:  isTemporary,
		IfExists:     ifExists,
		Name:         tableName,
		OnCluster:    onCluster,
		StatementEnd: tableName.End(),
	}

	if onCluster != nil {
		truncateTable.StatementEnd = onCluster.End()
	}

	return truncateTable, nil
}

func (p *Parser) parseDeleteClause(pos Pos) (*DeleteClause, error) {
	if err := p.expectKeyword(KeywordDelete); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordFrom); err != nil {
		return nil, err
	}
	tableIdentifier, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}

	if err := p.expectKeyword(KeywordWhere); err != nil {
		return nil, err
	}
	whereExpr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}

	return &DeleteClause{
		DeletePos: pos,
		Table:     tableIdentifier,
		OnCluster: onCluster,
		WhereExpr: whereExpr,
	}, nil
}

func (p *Parser) parseColumnNamesExpr(pos Pos) (*ColumnNamesExpr, error) {
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	var columnNames []NestedIdentifier
	for !p.atEOF() && p.tryConsumeTokenKind(TokenKindRParen) == nil {
		name, err := p.ParseNestedIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}

		columnNames = append(columnNames, *name)
		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}
	rightParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	return &ColumnNamesExpr{
		LeftParenPos:  pos,
		RightParenPos: rightParenPos,
		ColumnNames:   columnNames,
	}, nil
}

func (p *Parser) parseTypedPlaceholder(pos Pos) (Expr, error) {
	if err := p.expectTokenKind(TokenKindLBrace); err != nil {
		return nil, err
	}

	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	if err := p.expectTokenKind(TokenKindColon); err != nil {
		return nil, err
	}
	columnType, err := p.parseColumnType(p.Pos())
	if err != nil {
		return nil, err
	}

	rightBracePos := p.Pos()
	if err := p.expectTokenKind(TokenKindRBrace); err != nil {
		return nil, err
	}
	return &TypedPlaceholder{
		LeftBracePos:  pos,
		RightBracePos: rightBracePos,
		Name:          name,
		Type:          columnType,
	}, nil
}

func (p *Parser) parseAssignmentValues(pos Pos) (*AssignmentValues, error) {
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}
	p.inValues++
	defer func() { p.inValues-- }()

	var value Expr
	var err error
	values := make([]Expr, 0)
	for !p.atEOF() && p.tryConsumeTokenKind(TokenKindRParen) == nil {
		switch {
		case p.matchTokenKind(TokenKindLParen):
			value, err = p.parseAssignmentValues(p.Pos())
		case p.matchTokenKind(TokenKindLBrace) && p.peekTokenKind(TokenKindString):
			// a map, e.g. {'a': 1}, read by the Values format (#50)
			value, err = p.parseMapLiteral(p.Pos())
		case p.matchTokenKind(TokenKindLBrace):
			// placeholder with type, e.g. {a :Int32}, {b :DateTime(6)}
			value, err = p.parseTypedPlaceholder(p.Pos())
		default:
			value, err = p.parseExpr(p.Pos())
		}
		if err != nil {
			return nil, err
		}
		values = append(values, value)
		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}
	rightParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	return &AssignmentValues{
		LeftParenPos:  pos,
		RightParenPos: rightParenPos,
		Values:        values,
	}, nil
}

// parseInsertTableFunction parses the table function of
// `INSERT INTO [TABLE] FUNCTION f(args) [(columns)]`. Unlike an aggregate
// call, a second parenthesised list here is the insert's column list, not a
// parametric argument list, so only the function's own arguments are parsed.
func (p *Parser) parseInsertTableFunction() (*FunctionExpr, error) {
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	leftParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}
	args, err := p.parseColumnExprListWithLParen(p.Pos())
	if err != nil {
		return nil, err
	}
	rightParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	return &FunctionExpr{
		Name: name,
		Params: &ParamExprList{
			LeftParenPos:  leftParenPos,
			RightParenPos: rightParenPos,
			Items:         args,
		},
	}, nil
}

func (p *Parser) parseInsertStmt(pos Pos) (*InsertStmt, error) {
	if err := p.expectKeyword(KeywordInsert); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordInto); err != nil {
		return nil, err
	}

	insertExpr := &InsertStmt{InsertPos: pos}
	insertExpr.HasTableKeyword = p.tryConsumeKeywords(KeywordTable)

	var table Expr
	var err error
	if p.tryConsumeKeywords(KeywordFunction) {
		table, err = p.parseInsertTableFunction()
	} else {
		table, err = p.parseTableIdentifier(p.Pos())
	}
	if err != nil {
		return nil, err
	}
	insertExpr.Table = table

	if p.matchTokenKind(TokenKindLParen) && !p.peekKeyword(KeywordSelect) && !p.peekKeyword(KeywordWith) {
		// parse column names
		insertExpr.ColumnNames, err = p.parseColumnNamesExpr(p.Pos())
		if err != nil {
			return nil, err
		}
	}

	switch {
	case p.matchKeyword(KeywordFormat):
		insertExpr.Format, err = p.parseFormat(p.Pos())
	case p.matchKeyword(KeywordValues):
		// consume VALUES keyword
		_ = p.lexer.consumeToken()
		values := make([]*AssignmentValues, 0)
		for !p.atEOF() {
			value, err := p.parseAssignmentValues(p.Pos())
			if err != nil {
				return nil, err
			}
			values = append(values, value)
			if p.tryConsumeTokenKind(TokenKindComma) == nil {
				break
			}
		}
		insertExpr.Values = values
	case p.matchKeyword(KeywordSelect), p.matchKeyword(KeywordWith), p.matchTokenKind(TokenKindLParen):
		insertExpr.SelectExpr, err = p.parseSelectQuery(p.Pos())
	default:
		// do nothing
	}

	if err != nil {
		return nil, err
	}
	return insertExpr, nil
}

func (p *Parser) parseRenameStmt(pos Pos) (*RenameStmt, error) {
	if err := p.expectKeyword(KeywordRename); err != nil {
		return nil, err
	}

	renameTarget := KeywordTable
	switch {
	case p.tryConsumeKeywords(KeywordDictionary):
		renameTarget = KeywordDictionary
	case p.tryConsumeKeywords(KeywordDatabase):
		renameTarget = KeywordDatabase
	default:
		if err := p.expectKeyword(KeywordTable); err != nil {
			return nil, err
		}
	}

	targetPair, err := p.parseTargetPair(p.Pos())
	if err != nil {
		return nil, err
	}
	tablePairList := []*TargetPair{targetPair}
	for p.tryConsumeTokenKind(TokenKindComma) != nil {
		tablePair, err := p.parseTargetPair(p.Pos())
		if err != nil {
			return nil, err
		}
		tablePairList = append(tablePairList, tablePair)
	}
	// ClickHouse renames one database per statement (#50).
	if renameTarget == KeywordDatabase && len(tablePairList) > 1 {
		return nil, errors.New("RENAME DATABASE takes a single pair")
	}

	renameStmt := &RenameStmt{
		RenamePos:    pos,
		StatementEnd: tablePairList[len(tablePairList)-1].End(),

		RenameTarget:   renameTarget,
		TargetPairList: tablePairList,
	}

	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if onCluster != nil {
		renameStmt.OnCluster = onCluster
		renameStmt.StatementEnd = onCluster.End()
	}

	return renameStmt, nil
}

func (p *Parser) parseTargetPair(_ Pos) (*TargetPair, error) {
	oldTable, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	if err = p.expectKeyword(KeywordTo); err != nil {
		return nil, err
	}
	newTable, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}

	return &TargetPair{
		Old: oldTable,
		New: newTable,
	}, nil
}

func (p *Parser) parseCreateFunction(pos Pos, orReplace bool) (*CreateFunction, error) {
	if err := p.expectKeyword(KeywordFunction); err != nil {
		return nil, err
	}
	ifNotExists, err := p.tryParseIfNotExists()
	if err != nil {
		return nil, err
	}
	functionName, err := p.parseIdent()
	if err != nil {
		return nil, err
	}
	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordAs); err != nil {
		return nil, err
	}
	params, err := p.parseFunctionParams(p.Pos())
	if err != nil {
		return nil, err
	}
	if err := p.expectTokenKind(TokenKindArrow); err != nil {
		return nil, err
	}
	expr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &CreateFunction{
		CreatePos:    pos,
		OrReplace:    orReplace,
		IfNotExists:  ifNotExists,
		FunctionName: functionName,
		OnCluster:    onCluster,
		Params:       params,
		Expr:         expr,
	}, nil
}

// Dictionary parsing functions

func (p *Parser) parseDictionarySchemaClause(pos Pos) (*DictionarySchemaClause, error) {
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	schema := &DictionarySchemaClause{
		SchemaPos: pos,
	}

	// Parse first attribute
	attr, err := p.parseDictionaryAttribute(p.Pos())
	if err != nil {
		return nil, err
	}
	schema.Attributes = append(schema.Attributes, attr)

	// Parse additional attributes
	for p.tryConsumeTokenKind(TokenKindComma) != nil {
		attr, err := p.parseDictionaryAttribute(p.Pos())
		if err != nil {
			return nil, err
		}
		schema.Attributes = append(schema.Attributes, attr)
	}

	rParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	schema.RParenPos = rParenPos

	return schema, nil
}

func (p *Parser) parseDictionaryAttribute(pos Pos) (*DictionaryAttribute, error) {
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	columnType, err := p.parseColumnType(p.Pos())
	if err != nil {
		return nil, err
	}

	attr := &DictionaryAttribute{
		NamePos: pos,
		Name:    name,
		Type:    columnType,
	}

	// Parse optional attribute properties
	for {
		switch {
		case p.tryConsumeKeywords(KeywordDefault):
			if attr.Default != nil {
				return nil, fmt.Errorf("duplicate DEFAULT clause")
			}
			literal, err := p.parseLiteral(p.Pos())
			if err != nil {
				return nil, err
			}
			attr.Default = literal
		case p.tryConsumeKeywords(KeywordExpression):
			if attr.Expression != nil {
				return nil, fmt.Errorf("duplicate EXPRESSION clause")
			}
			expr, err := p.parseExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			attr.Expression = expr
		case p.tryConsumeKeywords(KeywordHierarchical):
			if attr.Hierarchical {
				return nil, fmt.Errorf("duplicate HIERARCHICAL clause")
			}
			attr.Hierarchical = true
		case p.tryConsumeKeywords(KeywordInjective):
			if attr.Injective {
				return nil, fmt.Errorf("duplicate INJECTIVE clause")
			}
			attr.Injective = true
		case p.tryConsumeKeywords(KeywordIs_object_id):
			if attr.IsObjectId {
				return nil, fmt.Errorf("duplicate IS_OBJECT_ID clause")
			}
			attr.IsObjectId = true
		default:
			// No more attribute properties
			return attr, nil
		}
	}
}

func (p *Parser) parseDictionaryEngineClause(pos Pos) (*DictionaryEngineClause, error) {
	engine := &DictionaryEngineClause{
		EnginePos: pos,
	}

	// Parse PRIMARY KEY clause (optional)
	if p.matchKeyword(KeywordPrimary) {
		primaryKey, err := p.parseDictionaryPrimaryKeyClause(p.Pos())
		if err != nil {
			return nil, err
		}
		engine.PrimaryKey = primaryKey
	}

	// Parse engine clauses (SOURCE, LIFETIME, LAYOUT, RANGE, SETTINGS)
	for {
		switch {
		case p.matchKeyword(KeywordSource):
			if engine.Source != nil {
				return nil, fmt.Errorf("duplicate SOURCE clause")
			}
			source, err := p.parseDictionarySourceClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engine.Source = source
		case p.matchKeyword(KeywordLifetime):
			if engine.Lifetime != nil {
				return nil, fmt.Errorf("duplicate LIFETIME clause")
			}
			lifetime, err := p.parseDictionaryLifetimeClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engine.Lifetime = lifetime
		case p.matchKeyword(KeywordLayout):
			if engine.Layout != nil {
				return nil, fmt.Errorf("duplicate LAYOUT clause")
			}
			layout, err := p.parseDictionaryLayoutClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engine.Layout = layout
		case p.matchKeyword(KeywordRange):
			if engine.Range != nil {
				return nil, fmt.Errorf("duplicate RANGE clause")
			}
			rangeClause, err := p.parseDictionaryRangeClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engine.Range = rangeClause
		case p.matchKeyword(KeywordSettings):
			if engine.Settings != nil {
				return nil, fmt.Errorf("duplicate SETTINGS clause")
			}
			settings, err := p.parseDictionarySettingsClause(p.Pos())
			if err != nil {
				return nil, err
			}
			engine.Settings = settings
		default:
			// No more engine clauses
			if engine.Source == nil {
				return nil, fmt.Errorf("SOURCE clause is required for dictionary")
			}
			return engine, nil
		}
	}
}

func (p *Parser) parseDictionaryPrimaryKeyClause(pos Pos) (*DictionaryPrimaryKeyClause, error) {
	if err := p.expectKeyword(KeywordPrimary); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordKey); err != nil {
		return nil, err
	}

	keys, err := p.parseColumnExprList(p.Pos())
	if err != nil {
		return nil, err
	}

	return &DictionaryPrimaryKeyClause{
		PrimaryKeyPos: pos,
		Keys:          keys,
		RParenPos:     keys.End() - 1,
	}, nil
}

func (p *Parser) parseDictionarySourceClause(pos Pos) (*DictionarySourceClause, error) {
	if err := p.expectKeyword(KeywordSource); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	sourceName, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	var args []*DictionaryArgExpr
	// Parse optional arguments
	for !p.matchTokenKind(TokenKindRParen) {
		arg, err := p.parseDictionaryArgExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		args = append(args, arg)

		// If there's no right paren, we expect another arg (no comma needed)
		if p.matchTokenKind(TokenKindRParen) {
			break
		}
	}

	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	rParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	return &DictionarySourceClause{
		SourcePos: pos,
		Source:    sourceName,
		Args:      args,
		RParenPos: rParenPos,
	}, nil
}

func (p *Parser) parseDictionaryArgExpr(pos Pos) (*DictionaryArgExpr, error) {
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	var value Expr
	// Parse the value part
	switch {
	case p.matchTokenKind(TokenKindString):
		literal, err := p.parseLiteral(p.Pos())
		if err != nil {
			return nil, err
		}
		value = literal
	case p.matchTokenKind(TokenKindInt), p.matchTokenKind(TokenKindFloat):
		literal, err := p.parseLiteral(p.Pos())
		if err != nil {
			return nil, err
		}
		value = literal
	case p.matchTokenKind(TokenKindIdent):
		ident, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		// Check if it's followed by optional parentheses
		if p.matchTokenKind(TokenKindLParen) {
			leftParenPos := p.Pos()
			_ = p.lexer.consumeToken() // consume (
			rightParenPos := p.Pos()
			if err := p.expectTokenKind(TokenKindRParen); err != nil {
				return nil, err
			}
			value = &FunctionExpr{
				Name: ident,
				Params: &ParamExprList{
					LeftParenPos:  leftParenPos,
					RightParenPos: rightParenPos,
					Items:         &ColumnExprList{ListPos: rightParenPos, ListEnd: rightParenPos},
				},
			}
		} else {
			value = ident
		}
	default:
		return nil, fmt.Errorf("expected identifier, string or number in dictionary argument, got %s", p.lastTokenKind())
	}

	return &DictionaryArgExpr{
		ArgPos: pos,
		Name:   name,
		Value:  value,
	}, nil
}

func (p *Parser) parseDictionaryLifetimeClause(pos Pos) (*DictionaryLifetimeClause, error) {
	if err := p.expectKeyword(KeywordLifetime); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	lifetime := &DictionaryLifetimeClause{
		LifetimePos: pos,
	}

	// Check for MIN/MAX form
	if p.matchKeyword(KeywordMin) || p.matchKeyword(KeywordMax) {
		isMinFirst := p.matchKeyword(KeywordMin)
		_ = p.lexer.consumeToken() // consume MIN or MAX

		first, err := p.parseNumber(p.Pos())
		if err != nil {
			return nil, err
		}

		// Expect the other keyword
		if isMinFirst {
			if err := p.expectKeyword(KeywordMax); err != nil {
				return nil, err
			}
		} else {
			if err := p.expectKeyword(KeywordMin); err != nil {
				return nil, err
			}
		}

		second, err := p.parseNumber(p.Pos())
		if err != nil {
			return nil, err
		}

		if isMinFirst {
			lifetime.Min = first
			lifetime.Max = second
		} else {
			lifetime.Min = second
			lifetime.Max = first
		}
	} else {
		// Simple form: LIFETIME(value)
		value, err := p.parseNumber(p.Pos())
		if err != nil {
			return nil, err
		}
		lifetime.Value = value
	}

	rParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	lifetime.RParenPos = rParenPos

	return lifetime, nil
}

func (p *Parser) parseDictionaryLayoutClause(pos Pos) (*DictionaryLayoutClause, error) {
	if err := p.expectKeyword(KeywordLayout); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	layoutName, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	var args []*DictionaryArgExpr
	// The inner argument list is optional: ClickHouse accepts both
	// LAYOUT(HASHED()) and the parameterless LAYOUT(FLAT) / LAYOUT(IP_TRIE).
	if p.tryConsumeTokenKind(TokenKindLParen) != nil {
		// Parse optional arguments
		for !p.matchTokenKind(TokenKindRParen) {
			arg, err := p.parseDictionaryArgExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			args = append(args, arg)

			// If there's no right paren, we expect another arg (no comma needed)
			if p.matchTokenKind(TokenKindRParen) {
				break
			}
		}

		if err := p.expectTokenKind(TokenKindRParen); err != nil {
			return nil, err
		}
	}

	rParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	return &DictionaryLayoutClause{
		LayoutPos: pos,
		Layout:    layoutName,
		Args:      args,
		RParenPos: rParenPos,
	}, nil
}

func (p *Parser) parseDictionaryRangeClause(pos Pos) (*DictionaryRangeClause, error) {
	if err := p.expectKeyword(KeywordRange); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	rangeClause := &DictionaryRangeClause{
		RangePos: pos,
	}

	// Parse MIN identifier MAX identifier or MAX identifier MIN identifier
	if p.matchKeyword(KeywordMin) {
		if err := p.expectKeyword(KeywordMin); err != nil {
			return nil, err
		}
		min, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		if err := p.expectKeyword(KeywordMax); err != nil {
			return nil, err
		}
		max, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		rangeClause.Min = min
		rangeClause.Max = max
	} else if p.matchKeyword(KeywordMax) {
		if err := p.expectKeyword(KeywordMax); err != nil {
			return nil, err
		}
		max, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		if err := p.expectKeyword(KeywordMin); err != nil {
			return nil, err
		}
		min, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		rangeClause.Min = min
		rangeClause.Max = max
	} else {
		return nil, fmt.Errorf("expected MIN or MAX in RANGE clause")
	}

	rParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	rangeClause.RParenPos = rParenPos

	return rangeClause, nil
}

func (p *Parser) parseDictionarySettingsClause(pos Pos) (*SettingsClause, error) {
	if err := p.expectKeyword(KeywordSettings); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	settings := &SettingsClause{SettingsPos: pos, ListEnd: pos}
	items := make([]*SettingExpr, 0)
	// Parse first setting
	expr, err := p.parseSettingsExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	items = append(items, expr)

	// Parse additional settings
	for p.tryConsumeTokenKind(TokenKindComma) != nil {
		expr, err := p.parseSettingsExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		items = append(items, expr)
	}

	if len(items) > 0 {
		settings.ListEnd = items[len(items)-1].End()
	}
	settings.Items = items

	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	settings.ListEnd = p.prevEnd() // include the closing )

	return settings, nil
}

// parseCreateIndex parses `[UNIQUE] INDEX [IF NOT EXISTS] name ON [db.]table
// [ON CLUSTER c] columns [TYPE type] [GRANULARITY n]` after CREATE (#135).
func (p *Parser) parseCreateIndex(pos Pos) (*CreateIndex, error) {
	create := &CreateIndex{CreatePos: pos}
	if p.matchUnquotedIdent("UNIQUE") {
		create.Unique = true
		_ = p.lexer.consumeToken()
	}
	if err := p.expectKeyword(KeywordIndex); err != nil {
		return nil, err
	}
	var err error
	if create.IfNotExists, err = p.tryParseIfNotExists(); err != nil {
		return nil, err
	}
	if create.Name, err = p.parseIdent(); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordOn); err != nil {
		return nil, err
	}
	if create.Table, err = p.parseTableIdentifier(p.Pos()); err != nil {
		return nil, err
	}
	if create.OnCluster, err = p.tryParseClusterClause(p.Pos()); err != nil {
		return nil, err
	}
	if p.tryConsumeTokenKind(TokenKindLParen) != nil {
		create.HasParen = true
		for {
			column, err := p.parseOrderExpr(p.Pos())
			if err != nil {
				return nil, err
			}
			create.Columns = append(create.Columns, column)
			if p.tryConsumeTokenKind(TokenKindComma) == nil {
				break
			}
		}
		if err := p.expectTokenKind(TokenKindRParen); err != nil {
			return nil, err
		}
	} else {
		expr, err := p.parseExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		create.Columns = []*OrderExpr{{OrderPos: expr.Pos(), Expr: expr, OrderEnd: expr.End()}}
	}
	if p.tryConsumeKeywords(KeywordType) {
		if create.IndexType, err = p.parseIndexType(); err != nil {
			return nil, err
		}
	}
	if p.tryConsumeKeywords(KeywordGranularity) {
		if create.Granularity, err = p.parseDecimal(p.Pos()); err != nil {
			return nil, err
		}
	}
	create.StatementEnd = p.prevEnd()
	return create, nil
}
