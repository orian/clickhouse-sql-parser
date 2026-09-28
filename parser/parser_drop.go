package parser

func (p *Parser) parseDropDatabase(pos Pos) (*DropDatabase, error) {
	if err := p.expectKeyword(KeywordDatabase); err != nil {
		return nil, err
	}

	isExists, err := p.tryParseIfExists()
	if err != nil {
		return nil, err
	}

	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	statementEnd := name.End()

	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if onCluster != nil {
		statementEnd = onCluster.End()
	}

	// PERMANENTLY is only valid for DETACH; the DROP caller rejects it.
	permanently := p.matchPermanently()
	if permanently {
		statementEnd = p.End()
		_ = p.lexer.consumeToken()
	}

	modifier, err := p.tryParseModifier()
	if err != nil {
		return nil, err
	}
	if modifier != "" {
		statementEnd = p.Pos()
	}

	return &DropDatabase{
		DropPos:      pos,
		Name:         name,
		IfExists:     isExists,
		OnCluster:    onCluster,
		Permanently:  permanently,
		Modifier:     modifier,
		StatementEnd: statementEnd,
	}, nil
}

func (p *Parser) parseDropStmt(pos Pos) (*DropStmt, error) {
	var isTemporary bool
	dropTarget := KeywordTable
	switch {
	case p.tryConsumeKeywords(KeywordDictionary):
		dropTarget = KeywordDictionary
	case p.tryConsumeKeywords(KeywordView):
		dropTarget = KeywordView
	default:
		isTemporary = p.tryConsumeKeywords(KeywordTemporary)
		if err := p.expectKeyword(KeywordTable); err != nil {
			return nil, err
		}
	}

	isExists, err := p.tryParseIfExists()
	if err != nil {
		return nil, err
	}

	name, err := p.parseTableIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}

	onCluster, err := p.tryParseClusterClause(p.Pos())
	if err != nil {
		return nil, err
	}

	// PERMANENTLY is only valid for DETACH; the DROP caller rejects it.
	permanently := p.matchPermanently()
	if permanently {
		_ = p.lexer.consumeToken()
	}

	modifier, err := p.tryParseModifier()
	if err != nil {
		return nil, err
	}

	return &DropStmt{
		DropPos:      pos,
		DropTarget:   dropTarget,
		Name:         name,
		IfExists:     isExists,
		OnCluster:    onCluster,
		IsTemporary:  isTemporary,
		Modifier:     modifier,
		Permanently:  permanently,
		StatementEnd: p.prevEnd(),
	}, nil
}

func (p *Parser) tryParseModifier() (string, error) {
	switch {
	case p.tryConsumeKeywords(KeywordSync):
		return "SYNC", nil
	case p.tryConsumeKeywords(KeywordNo):
		if err := p.expectKeyword(KeywordDelay); err != nil {
			return "", err
		}
		return "NO DELAY", nil
	}
	return "", nil
}

// parseDropIndex parses `INDEX [IF EXISTS] name ON [db.]table [ON CLUSTER c]`
// after DROP (#135).
func (p *Parser) parseDropIndex(pos Pos) (*DropIndex, error) {
	if err := p.expectKeyword(KeywordIndex); err != nil {
		return nil, err
	}
	drop := &DropIndex{DropPos: pos}
	var err error
	if drop.IfExists, err = p.tryParseIfExists(); err != nil {
		return nil, err
	}
	if drop.Name, err = p.parseIdent(); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordOn); err != nil {
		return nil, err
	}
	if drop.Table, err = p.parseTableIdentifier(p.Pos()); err != nil {
		return nil, err
	}
	if drop.OnCluster, err = p.tryParseClusterClause(p.Pos()); err != nil {
		return nil, err
	}
	drop.StatementEnd = p.prevEnd()
	return drop, nil
}
