package parser

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func (p *Parser) tryParseWithClause(pos Pos) (*WithClause, error) {
	if !p.matchKeyword(KeywordWith) {
		return nil, nil
	}
	return p.parseWithClause(pos)
}

func (p *Parser) parseWithClause(pos Pos) (*WithClause, error) {
	if err := p.expectKeyword(KeywordWith); err != nil {
		return nil, err
	}
	hasRecursive := p.tryConsumeKeywords(KeywordRecursive)

	cteExpr, err := p.parseCTEStmt(p.Pos())
	if err != nil {
		return nil, err
	}
	ctes := []*CTEStmt{cteExpr}
	for p.tryConsumeTokenKind(TokenKindComma) != nil {
		// ClickHouse allows a trailing comma immediately before the SELECT
		// governed by this WITH clause.
		if p.matchKeyword(KeywordSelect) {
			break
		}
		cteExpr, err := p.parseCTEStmt(p.Pos())
		if err != nil {
			return nil, err
		}
		ctes = append(ctes, cteExpr)
	}

	return &WithClause{
		WithPos:      pos,
		CTEs:         ctes,
		EndPos:       ctes[len(ctes)-1].End(),
		HasRecursive: hasRecursive,
	}, nil
}

func (p *Parser) tryParseTopClause(pos Pos) (*TopClause, error) {
	if !p.matchKeyword(KeywordTop) {
		return nil, nil
	}
	return p.parseTopClause(pos)
}

func (p *Parser) parseTopClause(pos Pos) (*TopClause, error) {
	if err := p.expectKeyword(KeywordTop); err != nil {
		return nil, err
	}

	number, err := p.parseNumber(p.Pos())
	if err != nil {
		return nil, err
	}
	topEnd := number.End()

	withTies := false
	if p.tryConsumeKeywords(KeywordWith) {
		topEnd = p.End()
		if err := p.expectKeyword(KeywordTies); err != nil {
			return nil, err
		}
		withTies = true
	}
	return &TopClause{
		TopPos:   pos,
		TopEnd:   topEnd,
		Number:   number,
		WithTies: withTies,
	}, nil
}

func (p *Parser) tryParseDistinctOn(pos Pos) (*DistinctOn, error) {
	if !p.matchKeyword(KeywordOn) {
		return nil, nil
	}
	return p.parseDistinctOn(pos)
}

func (p *Parser) parseDistinctOn(pos Pos) (*DistinctOn, error) {
	if err := p.expectKeyword(KeywordOn); err != nil {
		return nil, err
	}

	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}

	col, err := p.ParseNestedIdentifier(p.Pos())
	if err != nil {
		return nil, err
	}
	idents := []*NestedIdentifier{col}

	for p.matchTokenKind(TokenKindComma) {
		_ = p.lexer.consumeToken()

		col, err = p.ParseNestedIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		idents = append(idents, col)
	}

	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}

	return &DistinctOn{
		Idents:        idents,
		DistinctOnPos: pos,
		DistinctOnEnd: p.Pos(),
	}, nil
}

func (p *Parser) tryParseFromClause(pos Pos) (*FromClause, error) {
	if !p.matchKeyword(KeywordFrom) {
		return nil, nil
	}
	return p.parseFromClause(pos)
}

func (p *Parser) parseFromClause(pos Pos) (*FromClause, error) {
	if err := p.expectKeyword(KeywordFrom); err != nil {
		return nil, err
	}

	expr, err := p.parseJoinExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &FromClause{
		FromPos: pos,
		Expr:    expr,
	}, nil
}

func (p *Parser) tryParseJoinConstraints(pos Pos) (Expr, error) {
	switch {
	case p.tryConsumeKeywords(KeywordOn):
		columnExprList, err := p.parseColumnExprList(p.Pos())
		if err != nil {
			return nil, err
		}
		return &OnClause{
			OnPos: pos,
			On:    columnExprList,
		}, nil
	case p.tryConsumeKeywords(KeywordUsing):
		hasParen := p.tryConsumeTokenKind(TokenKindLParen) != nil
		columnExprList, err := p.parseColumnExprListWithLParen(p.Pos())
		if err != nil {
			return nil, err
		}
		if hasParen {
			if err := p.expectTokenKind(TokenKindRParen); err != nil {
				return nil, err
			}
		}
		return &UsingClause{
			UsingPos: pos,
			Using:    columnExprList,
			UsingEnd: p.prevEnd(),
		}, nil
	}
	return nil, nil
}

func (p *Parser) parseJoinOp(_ Pos) []string {
	var modifiers []string
	switch {
	case p.tryConsumeKeywords(KeywordCross): // cross join
		modifiers = append(modifiers, KeywordCross)
	case p.matchKeyword(KeywordAny), p.matchKeyword(KeywordAll):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordFull) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
		if p.matchKeyword(KeywordLeft) || p.matchKeyword(KeywordRight) || p.matchKeyword(KeywordInner) || p.matchKeyword(KeywordOuter) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordSemi), p.matchKeyword(KeywordAsof):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordLeft) || p.matchKeyword(KeywordRight) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
		if p.matchKeyword(KeywordOuter) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordInner):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordAll) || p.matchKeyword(KeywordAny) || p.matchKeyword(KeywordAsof) || p.matchKeyword(KeywordArray) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordLeft):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordOuter) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
		if p.matchKeyword(KeywordSemi) || p.matchKeyword(KeywordAnti) ||
			p.matchKeyword(KeywordAny) || p.matchKeyword(KeywordAll) ||
			p.matchKeyword(KeywordAsof) || p.matchKeyword(KeywordArray) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordRight):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordOuter) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
		if p.matchKeyword(KeywordSemi) || p.matchKeyword(KeywordAnti) ||
			p.matchKeyword(KeywordAny) || p.matchKeyword(KeywordAll) ||
			p.matchKeyword(KeywordAsof) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordFull):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
		if p.matchKeyword(KeywordOuter) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
		if p.matchKeyword(KeywordAll) || p.matchKeyword(KeywordAny) {
			modifiers = append(modifiers, p.last().String)
			_ = p.lexer.consumeToken()
		}
	case p.matchKeyword(KeywordArray):
		modifiers = append(modifiers, p.last().String)
		_ = p.lexer.consumeToken()
	}
	return modifiers
}

func (p *Parser) parseJoinTableExpr(_ Pos) (Expr, error) {
	switch {
	case p.matchTokenKind(TokenKindIdent), p.matchTokenKind(TokenKindString), p.matchTokenKind(TokenKindLParen):
		tableExpr, err := p.parseTableExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		statementEnd := tableExpr.End()

		hasFinal := p.matchKeyword(KeywordFinal)
		if hasFinal {
			if tableExpr.Stream != nil {
				return nil, errors.New("FINAL must come before STREAM")
			}
			statementEnd = p.End()
			_ = p.lexer.consumeToken()
		}

		sampleRatio, err := p.tryParseSampleClause(p.Pos())
		if err != nil {
			return nil, err
		}
		if sampleRatio != nil {
			if tableExpr.Stream != nil {
				return nil, errors.New("SAMPLE cannot follow STREAM")
			}
			statementEnd = sampleRatio.End()
		}
		return &JoinTableExpr{
			Table:        tableExpr,
			SampleRatio:  sampleRatio,
			HasFinal:     hasFinal,
			StatementEnd: statementEnd,
		}, nil
	default:
		return nil, fmt.Errorf("expected table name or subquery, got %s", fmt.Sprintf("%v", p.lastTokenKind()))
	}
}

func (p *Parser) parseJoinRightExpr(pos Pos) (expr Expr, err error) {
	var rightExpr Expr
	var modifiers []string
	if p.tryConsumeTokenKind(TokenKindComma) != nil {
		return p.parseJoinExpr(p.Pos())
	}
	// GLOBAL is kept as the first modifier: it changes how a distributed
	// join runs, so dropping it silently changes the query (#57).
	// ClickHouse has no LOCAL join; `a LOCAL JOIN b` aliases a as LOCAL,
	// which parseTableExpr handles.
	if p.tryConsumeKeywords(KeywordGlobal) {
		modifiers = append(modifiers, KeywordGlobal)
		if p.matchKeyword(KeywordArray) || p.matchKeyword(KeywordLeft) && p.peekKeyword(KeywordArray) {
			return nil, errors.New("GLOBAL cannot be used with ARRAY JOIN")
		}
		if !p.matchKeyword(KeywordJoin) {
			op := p.parseJoinOp(p.Pos())
			if len(op) == 0 {
				return nil, fmt.Errorf("expected JOIN after GLOBAL, got %s", p.lastTokenKind())
			}
			modifiers = append(modifiers, op...)
		}
	} else {
		modifiers = p.parseJoinOp(p.Pos())
	}

	if len(modifiers) != 0 && !p.matchKeyword(KeywordJoin) {
		return nil, fmt.Errorf("expected JOIN, got %s", p.lastTokenKind())
	}
	if !p.tryConsumeKeywords(KeywordJoin) {
		return nil, nil
	}

	modifiers = append(modifiers, KeywordJoin)

	// Check if this is an ARRAY JOIN
	if slices.Contains(modifiers, KeywordArray) {
		// For ARRAY JOIN, parse column expression list instead of table expression
		expr, err = p.parseColumnExprList(p.Pos())
		if err != nil {
			return nil, err
		}

		// ARRAY JOIN doesn't have constraints (ON/USING)
		// try parse next join
		rightExpr, err = p.parseJoinRightExpr(p.Pos())
		if err != nil {
			return nil, err
		}
		return &JoinExpr{
			JoinPos:     pos,
			Left:        expr,
			Right:       rightExpr,
			Modifiers:   modifiers,
			Constraints: nil,
		}, nil
	}

	expr, err = p.parseJoinTableExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	constrains, err := p.tryParseJoinConstraints(p.Pos())
	if err != nil {
		return nil, err
	}

	// try parse next join
	rightExpr, err = p.parseJoinRightExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &JoinExpr{
		JoinPos:     pos,
		Left:        expr,
		Right:       rightExpr,
		Modifiers:   modifiers,
		Constraints: constrains,
	}, nil
}

func (p *Parser) parseJoinExpr(pos Pos) (expr Expr, err error) {
	if expr, err = p.parseJoinTableExpr(p.Pos()); err != nil {
		return nil, err
	}
	rightExpr, err := p.parseJoinRightExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	if rightExpr == nil {
		return expr, nil
	}
	return &JoinExpr{
		JoinPos: pos,
		Left:    expr,
		Right:   rightExpr,
	}, nil
}

func (p *Parser) parseTableExpr(pos Pos) (*TableExpr, error) {
	var expr Expr
	var err error
	switch {
	case p.matchTokenKind(TokenKindString), p.matchTokenKind(TokenKindIdent):
		// table name
		tableIdentifier, err := p.parseTableIdentifier(p.Pos())
		if err != nil {
			return nil, err
		}
		// it's a table name
		if tableIdentifier.Database != nil || !p.matchTokenKind(TokenKindLParen) { // database.table
			expr = tableIdentifier
		} else {
			// table function expr
			tableArgs, err := p.parseTableArgList(p.Pos())
			if err != nil {
				return nil, err
			}
			expr = &TableFunctionExpr{
				Name: tableIdentifier.Table,
				Args: tableArgs,
			}
		}
	case p.matchTokenKind(TokenKindLParen):
		expr, err = p.parseSubQuery(p.Pos())
	default:
		return nil, errors.New("expect table name or subquery")
	}
	if err != nil {
		return nil, err
	}

	tableEnd := expr.End()
	if p.tryConsumeKeywords(KeywordAs) {
		alias, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		expr = &AliasExpr{
			Expr:     expr,
			AliasPos: alias.Pos(),
			Alias:    alias,
		}
		tableEnd = expr.End()
	} else if p.matchTokenKind(TokenKindIdent) && p.lastTokenKind() != TokenKindKeyword && !p.matchStreamKeyword() ||
		p.matchKeyword(KeywordLocal) {
		// LOCAL is not a join keyword in ClickHouse; like an identifier it
		// aliases the table (`FROM a LOCAL JOIN b` is `FROM a AS LOCAL JOIN b`).
		// STREAM is the streaming-query modifier, never an implicit alias.
		alias, err := p.parseIdent()
		if err != nil {
			return nil, err
		}
		expr = &AliasExpr{
			Expr:     expr,
			AliasPos: alias.Pos(),
			Alias:    alias,
		}
		tableEnd = expr.End()
	}

	isFinalExist := false
	if p.tryConsumeKeywords(KeywordFinal) {
		switch expr.(type) {
		case *TableFunctionExpr:
			return nil, errors.New("table function doesn't support FINAL")
		case *SelectQuery:
			return nil, errors.New("subquery doesn't support FINAL")
		}
		isFinalExist = true
		tableEnd = p.prevEnd() // include FINAL
	}

	// `STREAM [BOUNDED] [UNORDERED]` turns the read into a streaming query.
	// It must not be read as an implicit alias (#60).
	var stream *StreamClause
	if p.matchStreamKeyword() {
		stream = &StreamClause{StreamPos: p.Pos(), StreamEnd: p.End()}
		_ = p.lexer.consumeToken()
		for {
			var modifier string
			switch {
			case p.matchUnquotedIdent("BOUNDED"):
				modifier = "BOUNDED"
			case p.matchUnquotedIdent("UNORDERED"):
				modifier = "UNORDERED"
			}
			if modifier == "" {
				break
			}
			stream.Modifiers = append(stream.Modifiers, modifier)
			stream.StreamEnd = p.End()
			_ = p.lexer.consumeToken()
		}
		tableEnd = stream.StreamEnd
	}

	return &TableExpr{
		TablePos: pos,
		TableEnd: tableEnd,
		Expr:     expr,
		HasFinal: isFinalExist,
		Stream:   stream,
	}, nil
}

// matchUnquotedIdent reports whether the current token is the unquoted
// identifier word (case-insensitive). ClickHouse does not reserve words such
// as STREAM, BOUNDED or UNORDERED, so they are matched this way.
func (p *Parser) matchUnquotedIdent(word string) bool {
	return isUnquotedIdent(p.last()) && strings.EqualFold(p.last().String, word)
}

// matchStreamKeyword reports whether the current token is the STREAM
// modifier of a table expression.
func (p *Parser) matchStreamKeyword() bool {
	return p.matchUnquotedIdent("STREAM")
}

func (p *Parser) tryParsePrewhereClause(pos Pos) (*PrewhereClause, error) {
	if !p.matchKeyword(KeywordPrewhere) {
		return nil, nil
	}
	return p.parsePrewhereClause(pos)
}
func (p *Parser) parsePrewhereClause(pos Pos) (*PrewhereClause, error) {
	if err := p.expectKeyword(KeywordPrewhere); err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &PrewhereClause{
		PrewherePos: pos,
		Expr:        expr,
	}, nil
}

func (p *Parser) tryParseWhereClause(pos Pos) (*WhereClause, error) {
	if !p.matchKeyword(KeywordWhere) {
		return nil, nil
	}
	return p.parseWhereClause(pos)
}

func (p *Parser) parseWhereClause(pos Pos) (*WhereClause, error) {
	if err := p.expectKeyword(KeywordWhere); err != nil {
		return nil, err
	}

	expr, err := p.parseExpr(p.Pos())
	if err != nil {
		return nil, err
	}
	return &WhereClause{
		WherePos: pos,
		Expr:     expr,
	}, nil
}

func (p *Parser) tryParseGroupByClause(pos Pos) (*GroupByClause, error) {
	if !p.matchKeyword(KeywordGroup) {
		return nil, nil
	}
	return p.parseGroupByClause(pos)
}

// syntax: groupByClause? (WITH (CUBE | ROLLUP))? (WITH TOTALS)?
func (p *Parser) parseGroupByClause(pos Pos) (*GroupByClause, error) {
	if err := p.expectKeyword(KeywordGroup); err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordBy); err != nil {
		return nil, err
	}

	var expr Expr
	var err error
	aggregateType := ""
	switch {
	case p.matchKeyword(KeywordCube) || p.matchKeyword(KeywordRollup):
		aggregateType = p.last().String
		_ = p.lexer.consumeToken()
		expr, err = p.parseFunctionParams(p.Pos())
	case p.tryConsumeKeywords(KeywordGrouping, KeywordSets):
		aggregateType = "GROUPING SETS"
		expr, err = p.parseFunctionParams(p.Pos())
	case p.tryConsumeKeywords(KeywordAll):
		aggregateType = "ALL"
	default:
		expr, err = p.parseColumnExprListWithLParen(p.Pos())
	}
	if err != nil {
		return nil, err
	}
	groupBy := &GroupByClause{
		GroupByPos:    pos,
		AggregateType: aggregateType,
		Expr:          expr,
	}

	// parse [WITH ROLLUP | WITH CUBE] [WITH TOTALS]; as in ClickHouse, each
	// modifier appears at most once and TOTALS comes last.
	for p.tryConsumeKeywords(KeywordWith) {
		switch {
		case p.matchKeyword(KeywordCube), p.matchKeyword(KeywordRollup):
			if groupBy.WithCube || groupBy.WithRollup || groupBy.WithTotals {
				return nil, fmt.Errorf("unexpected WITH %s after GROUP BY modifiers", p.last().String)
			}
			if p.tryConsumeKeywords(KeywordCube) {
				groupBy.WithCube = true
			} else {
				_ = p.tryConsumeKeywords(KeywordRollup)
				groupBy.WithRollup = true
			}
		case p.matchKeyword(KeywordTotals):
			if groupBy.WithTotals {
				return nil, fmt.Errorf("duplicate WITH TOTALS modifier")
			}
			_ = p.tryConsumeKeywords(KeywordTotals)
			groupBy.WithTotals = true
		default:
			return nil, fmt.Errorf("expected CUBE, ROLLUP or TOTALS, got %s", p.lastTokenKind())
		}
	}
	groupBy.GroupByEnd = p.prevEnd()

	return groupBy, nil
}

func (p *Parser) tryParseLimitAfterLimitByClause(pos Pos) (*LimitClause, error) {
	if !p.matchKeyword(KeywordLimit) {
		return nil, nil
	}

	return p.parseLimitClause(pos)
}

func (p *Parser) tryParseLimitClause(pos Pos) (*LimitClause, error) {
	if !p.matchKeyword(KeywordLimit) && !p.matchKeyword(KeywordOffset) {
		return nil, nil
	}

	return p.parseLimitClause(pos)
}

func (p *Parser) parseLimitClause(pos Pos) (*LimitClause, error) {
	var limit Expr
	var offset Expr
	var err error
	if p.tryConsumeKeywords(KeywordLimit) {
		limit, err = p.parseExpr(p.Pos())
		if err != nil {
			return nil, err
		}

		if p.tryConsumeKeywords(KeywordOffset) {
			offset, err = p.parseExpr(p.Pos())
		} else if p.tryConsumeTokenKind(TokenKindComma) != nil {
			offset = limit
			limit, err = p.parseExpr(p.Pos())
		}
	} else if p.tryConsumeKeywords(KeywordOffset) {
		offset, err = p.parseExpr(p.Pos())
	}

	if err != nil {
		return nil, err
	}

	limitClause := &LimitClause{
		LimitPos: pos,
		Limit:    limit,
		Offset:   offset,
	}
	// LIMIT n [OFFSET m] WITH TIES (#48). ClickHouse only checks at analysis
	// time that ORDER BY is present, so that is not enforced here.
	if limit != nil && p.matchKeyword(KeywordWith) && p.peekKeyword(KeywordTies) {
		_ = p.lexer.consumeToken() // WITH
		limitClause.WithTiesEnd = p.End()
		_ = p.lexer.consumeToken() // TIES
		limitClause.WithTies = true
	}
	return limitClause, nil
}

func (p *Parser) tryParseLimitByClause(pos Pos) (Expr, error) {
	if !p.matchKeyword(KeywordLimit) {
		return nil, nil
	}
	return p.parseLimitByClause(pos)
}

func (p *Parser) parseBetweenClause(expr Expr) (*BetweenClause, error) {
	if err := p.expectKeyword(KeywordBetween); err != nil {
		return nil, err
	}

	betweenExpr, err := p.parseSubExpr(p.Pos(), PrecedenceBetweenLike)
	if err != nil {
		return nil, err
	}

	andPos := p.Pos()
	if err := p.expectKeyword(KeywordAnd); err != nil {
		return nil, err
	}

	andExpr, err := p.parseSubExpr(p.Pos(), PrecedenceBetweenLike)
	if err != nil {
		return nil, err
	}

	return &BetweenClause{
		Expr:    expr,
		Between: betweenExpr,
		AndPos:  andPos,
		And:     andExpr,
	}, nil
}

func (p *Parser) parseLimitByClause(pos Pos) (Expr, error) {
	limit, err := p.parseLimitClause(pos)
	if err != nil {
		return nil, err
	}

	var by *ColumnExprList
	if !p.tryConsumeKeywords(KeywordBy) {
		return limit, nil
	}
	if by, err = p.parseColumnExprListWithLParen(p.Pos()); err != nil {
		return nil, err
	}
	return &LimitByClause{
		Limit:  limit,
		ByExpr: by,
	}, nil
}

func (p *Parser) tryParseWindowFrameClause(pos Pos) (*WindowFrameClause, error) {
	if !p.matchKeyword(KeywordRows) && !p.matchKeyword(KeywordRange) {
		return nil, nil
	}
	return p.parseWindowFrameClause(pos)
}

func (p *Parser) parseWindowFrameClause(pos Pos) (*WindowFrameClause, error) {
	var windowFrameType string
	if p.matchKeyword(KeywordRows) || p.matchKeyword(KeywordRange) {
		windowFrameType = p.last().String
		_ = p.lexer.consumeToken()
	} else {
		return nil, fmt.Errorf("expected ROWS or RANGE for window frame")
	}

	var expr Expr
	if p.tryConsumeKeywords(KeywordBetween) {
		left, err := p.parseFrameExtent()
		if err != nil {
			return nil, err
		}
		andPos := p.Pos()
		if err := p.expectKeyword(KeywordAnd); err != nil {
			return nil, err
		}
		right, err := p.parseFrameExtent()
		if err != nil {
			return nil, err
		}
		expr = &BetweenClause{
			Between: left,
			AndPos:  andPos,
			And:     right,
		}
	} else {
		// single extent
		extent, err := p.parseFrameExtent()
		if err != nil {
			return nil, err
		}
		expr = extent
	}

	return &WindowFrameClause{
		FramePos: pos,
		Type:     windowFrameType,
		Extend:   expr,
	}, nil
}

// parseFrameExtent parses a single frame extent
func (p *Parser) parseFrameExtent() (Expr, error) {
	switch {
	case p.matchKeyword(KeywordCurrent):
		return p.parseFrameCurrentRow()
	case p.matchKeyword(KeywordUnbounded):
		return p.parseFrameUnbounded()
	case p.matchTokenKind(TokenKindInt):
		return p.parseFrameNumber()
	case p.matchTokenKind(TokenKindLBrace):
		return p.parseFrameParam()
	case p.matchKeyword(KeywordInterval):
		return p.parseFrameInterval()
	default:
		return nil, fmt.Errorf("expected UNBOUNDED, CURRENT ROW, integer, parameter, or interval")
	}
}

func (p *Parser) parseFrameCurrentRow() (Expr, error) {
	currentPos := p.Pos()
	_ = p.lexer.consumeToken()
	if err := p.expectKeyword(KeywordRow); err != nil {
		return nil, err
	}
	rowEnd := p.End()
	return &WindowFrameCurrentRow{
		CurrentPos: currentPos,
		RowEnd:     rowEnd,
	}, nil
}

func (p *Parser) parseFrameUnbounded() (Expr, error) {
	unboundedPos := p.Pos()
	_ = p.lexer.consumeToken()

	direction, err := p.parseFrameDirection()
	if err != nil {
		return nil, err
	}
	return &WindowFrameUnbounded{
		UnboundedPos: unboundedPos,
		Direction:    direction,
	}, nil
}

func (p *Parser) parseFrameNumber() (Expr, error) {
	number, err := p.parseNumber(p.Pos())
	if err != nil {
		return nil, err
	}

	direction, endPos, err := p.parseFrameDirectionWithEnd()
	if err != nil {
		return nil, err
	}
	return &WindowFrameNumber{
		EndPos:    endPos,
		Number:    number,
		Direction: direction,
	}, nil
}

func (p *Parser) parseFrameParam() (Expr, error) {
	queryParam, err := p.parseQueryParam(p.Pos())
	if err != nil {
		return nil, err
	}

	direction, endPos, err := p.parseFrameDirectionWithEnd()
	if err != nil {
		return nil, err
	}
	return &WindowFrameParam{
		Param:     queryParam,
		EndPos:    endPos,
		Direction: direction,
	}, nil
}

func (p *Parser) parseFrameInterval() (Expr, error) {
	intervalExpr, err := p.parseInterval(true)
	if err != nil {
		return nil, err
	}

	direction, endPos, err := p.parseFrameDirectionWithEnd()
	if err != nil {
		return nil, err
	}
	return &WindowFrameExtendExpr{
		Expr:      intervalExpr,
		Direction: direction,
		EndPos:    endPos,
	}, nil
}

func (p *Parser) parseFrameDirection() (string, error) {
	switch {
	case p.matchKeyword(KeywordPreceding), p.matchKeyword(KeywordFollowing):
		direction := p.last().String
		_ = p.lexer.consumeToken()
		return direction, nil
	default:
		return "", fmt.Errorf("expected PRECEDING or FOLLOWING, got %s", p.lastTokenKind())
	}
}

func (p *Parser) parseFrameDirectionWithEnd() (string, Pos, error) {
	if !p.matchKeyword(KeywordPreceding) && !p.matchKeyword(KeywordFollowing) {
		return "", 0, fmt.Errorf("expected PRECEDING or FOLLOWING, got %s", p.lastTokenKind())
	}
	endPos := p.End()
	direction := p.last().String
	_ = p.lexer.consumeToken()
	return direction, endPos, nil
}

func (p *Parser) tryParseWindowClause(pos Pos) (*WindowClause, error) {
	if !p.matchKeyword(KeywordWindow) {
		return nil, nil
	}
	return p.parseWindowClause(pos)
}

func (p *Parser) parseWindowCondition(pos Pos) (*WindowExpr, error) {
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}
	var windowName *Ident
	if p.canParseWindowNameInParens() {
		var err error
		windowName, err = p.parseIdent()
		if err != nil {
			return nil, err
		}
	}
	partitionBy, err := p.tryParsePartitionByClause(pos)
	if err != nil {
		return nil, err
	}
	orderBy, err := p.tryParseOrderByClause(p.Pos())
	if err != nil {
		return nil, err
	}
	frame, err := p.tryParseWindowFrameClause(p.Pos())
	if err != nil {
		return nil, err
	}
	rightParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	return &WindowExpr{
		LeftParenPos:  pos,
		RightParenPos: rightParenPos,
		WindowName:    windowName,
		PartitionBy:   partitionBy,
		OrderBy:       orderBy,
		Frame:         frame,
	}, nil
}

func (p *Parser) canParseWindowNameInParens() bool {
	if !p.matchTokenKind(TokenKindIdent) {
		return false
	}
	if !p.matchTokenKind(TokenKindKeyword) {
		return true
	}

	savedState := p.lexer.saveState()
	defer p.lexer.restoreState(savedState)

	switch {
	case p.matchKeyword(KeywordPartition), p.matchKeyword(KeywordOrder):
		_ = p.lexer.consumeToken()
		return !p.matchKeyword(KeywordBy)
	case p.matchKeyword(KeywordRows), p.matchKeyword(KeywordRange):
		_ = p.lexer.consumeToken()
		return !p.matchKeyword(KeywordBetween) &&
			!p.matchKeyword(KeywordCurrent) &&
			!p.matchKeyword(KeywordUnbounded) &&
			!p.matchTokenKind(TokenKindInt) &&
			!p.matchTokenKind(TokenKindLBrace) &&
			!p.matchKeyword(KeywordInterval)
	default:
		return true
	}
}

func (p *Parser) parseWindowClause(pos Pos) (*WindowClause, error) {
	if err := p.expectKeyword(KeywordWindow); err != nil {
		return nil, err
	}

	windows := make([]*WindowDefinition, 0, 1)
	for {
		windowName, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		asPos := p.Pos()
		if err := p.expectKeyword(KeywordAs); err != nil {
			return nil, err
		}

		condition, err := p.parseWindowCondition(p.Pos())
		if err != nil {
			return nil, err
		}

		windows = append(windows, &WindowDefinition{
			Name:  windowName,
			AsPos: asPos,
			Expr:  condition,
		})

		if p.tryConsumeTokenKind(TokenKindComma) == nil {
			break
		}
	}

	var endPos Pos
	if len(windows) > 0 {
		endPos = windows[len(windows)-1].End()
	}

	return &WindowClause{
		WindowPos: pos,
		EndPos:    endPos,
		Windows:   windows,
	}, nil
}

func (p *Parser) tryParseHavingClause(pos Pos) (*HavingClause, error) {
	if !p.matchKeyword(KeywordHaving) {
		return nil, nil
	}
	return p.parseHavingClause(pos)
}

func (p *Parser) parseHavingClause(pos Pos) (*HavingClause, error) {
	if err := p.expectKeyword(KeywordHaving); err != nil {
		return nil, err
	}

	expr, err := p.parseColumnsExpr(p.Pos())
	if err != nil {
		return nil, err
	}

	return &HavingClause{
		HavingPos: pos,
		Expr:      expr,
	}, nil
}

func (p *Parser) parseSubQuery(_ Pos) (*SubQuery, error) {
	pos := p.Pos()

	hasParen := p.tryConsumeTokenKind(TokenKindLParen) != nil

	selectQuery, err := p.parseSelectQuery(p.Pos())
	if err != nil {
		return nil, err
	}
	var rightParenPos Pos
	if hasParen {
		rightParenPos = p.Pos()
		if err := p.expectTokenKind(TokenKindRParen); err != nil {
			return nil, err
		}
		// `(query) UNION ALL ...`: the parentheses group only the first
		// operand, so they become a group node and the chain continues.
		if p.matchKeyword(KeywordUnion) || p.matchKeyword(KeywordExcept) {
			group := &SelectQuery{
				SelectPos:    pos,
				StatementEnd: p.prevEnd(),
				HasParen:     true,
				Group:        selectQuery,
			}
			if err := p.parseSetOperation(group); err != nil {
				return nil, err
			}
			return &SubQuery{Select: group}, nil
		}
	}

	return &SubQuery{
		HasParen:      hasParen,
		Select:        selectQuery,
		RightParenPos: rightParenPos,
	}, nil
}

func (p *Parser) parseSelectQuery(_ Pos) (*SelectQuery, error) {
	if !p.matchKeyword(KeywordSelect) && !p.matchKeyword(KeywordWith) && !p.matchTokenKind(TokenKindLParen) {
		return nil, fmt.Errorf("expected SELECT, WITH or (, got %s", p.lastTokenKind())
	}

	var selectStmt *SelectQuery
	var err error
	if p.matchTokenKind(TokenKindLParen) {
		selectStmt, err = p.parseSelectGroup()
	} else {
		selectStmt, err = p.parseSelectStmt(p.Pos())
	}
	if err != nil {
		return nil, err
	}
	if err := p.parseSetOperation(selectStmt); err != nil {
		return nil, err
	}
	return selectStmt, nil
}

// parseSetOperation parses an optional `UNION ALL|DISTINCT query` or
// `EXCEPT query` continuation of selectStmt, or the clauses after a final
// parenthesised group, and extends selectStmt's end to cover them.
func (p *Parser) parseSetOperation(selectStmt *SelectQuery) error {
	defer func() {
		// The statement ends with the last query of its UNION/EXCEPT chain.
		for _, next := range []*SelectQuery{selectStmt.UnionAll, selectStmt.UnionDistinct, selectStmt.Except} {
			if next != nil && next.End() > selectStmt.StatementEnd {
				selectStmt.StatementEnd = next.End()
			}
		}
	}()
	switch {
	case p.tryConsumeKeywords(KeywordUnion):
		switch {
		case p.tryConsumeKeywords(KeywordAll):
			unionAllExpr, err := p.parseSelectQuery(p.Pos())
			if err != nil {
				return err
			}
			selectStmt.UnionAll = unionAllExpr
		case p.tryConsumeKeywords(KeywordDistinct):
			unionDistinctExpr, err := p.parseSelectQuery(p.Pos())
			if err != nil {
				return err
			}
			selectStmt.UnionDistinct = unionDistinctExpr
		default:
			return fmt.Errorf("expected ALL or DISTINCT, got %s", p.lastTokenKind())
		}
	case p.tryConsumeKeywords(KeywordExcept):
		exceptExpr, err := p.parseSelectQuery(p.Pos())
		if err != nil {
			return err
		}
		selectStmt.Except = exceptExpr
	default:
		if selectStmt.Group != nil {
			return p.parseGroupTail(selectStmt)
		}
	}
	return nil
}

// parseSelectGroup parses a parenthesised set-operation operand `( query )`
// into a group node: the query inside the parentheses goes to Group, and the
// caller attaches any UNION/EXCEPT that follows the closing parenthesis to
// the group node itself (#78).
func (p *Parser) parseSelectGroup() (*SelectQuery, error) {
	lParenPos := p.Pos()
	if err := p.expectTokenKind(TokenKindLParen); err != nil {
		return nil, err
	}
	inner, err := p.parseSelectQuery(p.Pos())
	if err != nil {
		return nil, err
	}
	if err := p.expectTokenKind(TokenKindRParen); err != nil {
		return nil, err
	}
	return &SelectQuery{
		SelectPos:    lParenPos,
		StatementEnd: p.prevEnd(),
		HasParen:     true,
		Group:        inner,
	}, nil
}

// parseGroupTail parses the clauses ClickHouse accepts after a final
// parenthesised group: `[SETTINGS ...]` or `FORMAT fmt [SETTINGS ...]`. The
// SETTINGS before and after FORMAT are the same query-level clause here, so
// only one of them may be given.
func (p *Parser) parseGroupTail(group *SelectQuery) error {
	settings, err := p.tryParseSettingsClause(p.Pos())
	if err != nil {
		return err
	}
	format, err := p.tryParseFormat(p.Pos())
	if err != nil {
		return err
	}
	var outputSettings *SettingsClause
	if format != nil {
		outputSettings, err = p.tryParseSettingsClause(p.Pos())
		if err != nil {
			return err
		}
	}
	if settings != nil && outputSettings != nil {
		return errors.New("duplicate query-level SETTINGS clause")
	}
	group.Settings = settings
	group.Format = format
	group.OutputSettings = outputSettings
	if settings != nil || format != nil {
		group.StatementEnd = p.prevEnd()
	}
	return nil
}

func (p *Parser) parseSelectStmt(pos Pos) (*SelectQuery, error) { // nolint: funlen
	withClause, err := p.tryParseWithClause(pos)
	if err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordSelect); err != nil {
		return nil, err
	}
	// DISTINCT?
	hasDistinct := p.tryConsumeKeywords(KeywordDistinct)
	distinctOn, err := p.tryParseDistinctOn(p.Pos())
	if err != nil {
		return nil, err
	}

	top, err := p.tryParseTopClause(p.Pos())
	if err != nil {
		return nil, err
	}
	selectItems, err := p.parseSelectItems()
	if err != nil {
		return nil, err
	}

	statementEnd := pos
	if len(selectItems) > 0 {
		statementEnd = selectItems[len(selectItems)-1].End()
	}
	from, err := p.tryParseFromClause(p.Pos())
	if err != nil {
		return nil, err
	}

	if from != nil {
		statementEnd = from.End()
	}
	prewhere, err := p.tryParsePrewhereClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if prewhere != nil {
		statementEnd = prewhere.End()
	}
	where, err := p.tryParseWhereClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if where != nil {
		statementEnd = where.End()
	}
	groupBy, err := p.tryParseGroupByClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if groupBy != nil {
		statementEnd = groupBy.End()
	}
	withTotal := false
	if p.tryConsumeKeywords(KeywordWith) {
		totalsEnd := p.End()
		if err := p.expectKeyword(KeywordTotals); err != nil {
			return nil, err
		}
		withTotal = true
		statementEnd = totalsEnd
	}
	having, err := p.tryParseHavingClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if having != nil {
		statementEnd = having.End()
	}
	window, err := p.tryParseWindowClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if window != nil {
		statementEnd = window.End()
	}
	orderBy, err := p.tryParseOrderByClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if orderBy != nil {
		statementEnd = orderBy.End()
	}

	var limitBy *LimitByClause
	var limit *LimitClause
	parsedLimitBy, err := p.tryParseLimitByClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if parsedLimitBy != nil {
		statementEnd = parsedLimitBy.End()
		switch e := parsedLimitBy.(type) {
		case *LimitByClause:
			limitBy = e
			limit, err = p.tryParseLimitAfterLimitByClause(p.Pos())
			if err != nil {
				return nil, err
			}
			if limit != nil {
				statementEnd = limit.End()
			}
		case *LimitClause:
			limit = e
		}
	} else {
		limit, err = p.tryParseLimitClause(p.Pos())
		if err != nil {
			return nil, err
		}
		if limit != nil {
			statementEnd = limit.End()
		}
	}

	settings, err := p.tryParseSettingsClause(p.Pos())
	if err != nil {
		return nil, err
	}
	if settings != nil {
		statementEnd = settings.End()
	}

	format, err := p.tryParseFormat(p.Pos())
	if err != nil {
		return nil, err
	}
	var outputSettings *SettingsClause
	if format != nil {
		statementEnd = format.End()
		// SETTINGS after FORMAT is the query-level output clause, separate
		// from the SELECT's own SETTINGS parsed above.
		outputSettings, err = p.tryParseSettingsClause(p.Pos())
		if err != nil {
			return nil, err
		}
		if outputSettings != nil {
			statementEnd = outputSettings.End()
		}
	}

	return &SelectQuery{
		With:           withClause,
		SelectPos:      pos,
		StatementEnd:   statementEnd,
		Top:            top,
		HasDistinct:    hasDistinct,
		DistinctOn:     distinctOn,
		SelectItems:    selectItems,
		From:           from,
		Window:         window,
		Prewhere:       prewhere,
		Where:          where,
		GroupBy:        groupBy,
		Having:         having,
		OrderBy:        orderBy,
		LimitBy:        limitBy,
		Limit:          limit,
		Settings:       settings,
		Format:         format,
		OutputSettings: outputSettings,
		WithTotal:      withTotal,
	}, nil
}

func (p *Parser) parseCTEStmt(pos Pos) (*CTEStmt, error) {
	expr, err := p.parseExpr(pos)
	if err != nil {
		return nil, err
	}
	if err := p.expectKeyword(KeywordAs); err != nil {
		return nil, err
	}
	if p.matchTokenKind(TokenKindLParen) {
		selectQuery, err := p.parseSelectQuery(p.Pos())
		if err != nil {
			return nil, err
		}
		// CTEStmt prints the parentheses of `name AS (SELECT ...)` itself, so
		// unwrap the group they formed.
		if selectQuery.Group != nil && selectQuery.UnionAll == nil && selectQuery.UnionDistinct == nil &&
			selectQuery.Except == nil && selectQuery.Settings == nil && selectQuery.Format == nil {
			selectQuery = selectQuery.Group
		}
		return &CTEStmt{
			CTEPos: pos,
			Expr:   expr,
			Alias:  selectQuery,
		}, nil
	}
	name, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	return &CTEStmt{
		CTEPos: pos,
		Expr:   expr,
		Alias:  name,
	}, nil
}

func (p *Parser) tryParseSampleClause(pos Pos) (*SampleClause, error) {
	if !p.matchKeyword(KeywordSample) {
		return nil, nil
	}
	return p.parseSampleClause(pos)
}

func (p *Parser) parseSampleClause(pos Pos) (*SampleClause, error) {
	if err := p.expectKeyword(KeywordSample); err != nil {
		return nil, err
	}
	ratio, err := p.parseRatioExpr(p.Pos())
	if err != nil {
		return nil, err
	}

	var offset *RatioExpr
	if p.matchKeyword(KeywordOffset) {
		_ = p.lexer.consumeToken()
		offset, err = p.parseRatioExpr(p.Pos())
		if err != nil {
			return nil, err
		}
	}

	return &SampleClause{
		SamplePos: pos,
		Ratio:     ratio,
		Offset:    offset,
	}, nil
}

func (p *Parser) parseExplainStmt(pos Pos) (*ExplainStmt, error) {
	if err := p.expectKeyword(KeywordExplain); err != nil {
		return nil, err
	}

	var explainType string
	switch {
	case p.matchKeyword(KeywordSyntax),
		p.matchKeyword(KeywordPipeline),
		p.matchKeyword(KeywordEstimate),
		p.matchKeyword(KeywordAst):
		explainType = p.last().String
		_ = p.lexer.consumeToken()
	default:
		return nil, fmt.Errorf("expected SYNTAX, PIPELINE, ESTIMATE or AST, got %s", p.lastTokenKind())
	}
	stmt, err := p.parseSelectQuery(p.Pos())
	if err != nil {
		return nil, err
	}
	return &ExplainStmt{
		ExplainPos: pos,
		Type:       explainType,
		Statement:  stmt,
	}, nil
}
