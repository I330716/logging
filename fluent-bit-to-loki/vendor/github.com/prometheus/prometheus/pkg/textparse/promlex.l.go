// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textparse

import (
	"github.com/pkg/errors"
)

const (
	sInit = iota
	sComment
	sMeta1
	sMeta2
	sLabels
	sLValue
	sValue
	sTimestamp
	sExemplar
	sEValue
	sETimestamp
)

// Lex is called by the parser generated by "go tool yacc" to obtain each
// token. The method is opened before the matching rules block and closed at
// the end of the file.
func (l *promlexer) Lex() token {
	if l.i >= len(l.b) {
		return tEOF
	}
	c := l.b[l.i]
	l.start = l.i

yystate0:

	switch yyt := l.state; yyt {
	default:
		panic(errors.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: sComment
		goto yystart8
	case 2: // start condition: sMeta1
		goto yystart19
	case 3: // start condition: sMeta2
		goto yystart21
	case 4: // start condition: sLabels
		goto yystart24
	case 5: // start condition: sLValue
		goto yystart29
	case 6: // start condition: sValue
		goto yystart33
	case 7: // start condition: sTimestamp
		goto yystart36
	}

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate5
	case c == ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate7
	case c == '\n':
		goto yystate4
	case c == '\t' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate4:
	c = l.next()
	goto yyrule2

yystate5:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '\t' || c == ' ':
		goto yystate6
	}

yystate6:
	c = l.next()
	switch {
	default:
		goto yyrule4
	case c == '\t' || c == ' ':
		goto yystate6
	}

yystate7:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate7
	}

	goto yystate8 // silence unused label error
yystate8:
	c = l.next()
yystart8:
	switch {
	default:
		goto yyabort
	case c == 'H':
		goto yystate9
	case c == 'T':
		goto yystate14
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate9:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate10
	}

yystate10:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'L':
		goto yystate11
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'P':
		goto yystate12
	}

yystate12:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\t' || c == ' ':
		goto yystate13
	}

yystate13:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c == '\t' || c == ' ':
		goto yystate13
	}

yystate14:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'Y':
		goto yystate15
	}

yystate15:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'P':
		goto yystate16
	}

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate17
	}

yystate17:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\t' || c == ' ':
		goto yystate18
	}

yystate18:
	c = l.next()
	switch {
	default:
		goto yyrule7
	case c == '\t' || c == ' ':
		goto yystate18
	}

	goto yystate19 // silence unused label error
yystate19:
	c = l.next()
yystart19:
	switch {
	default:
		goto yyabort
	case c == ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate20:
	c = l.next()
	switch {
	default:
		goto yyrule8
	case c >= '0' && c <= ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

	goto yystate21 // silence unused label error
yystate21:
	c = l.next()
yystart21:
	switch {
	default:
		goto yyrule9
	case c == '\t' || c == ' ':
		goto yystate23
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate22
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate22
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == ' ':
		goto yystate23
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate22
	}

	goto yystate24 // silence unused label error
yystate24:
	c = l.next()
yystart24:
	switch {
	default:
		goto yyabort
	case c == ',':
		goto yystate25
	case c == '=':
		goto yystate26
	case c == '\t' || c == ' ':
		goto yystate3
	case c == '}':
		goto yystate28
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate27
	}

yystate25:
	c = l.next()
	goto yyrule15

yystate26:
	c = l.next()
	goto yyrule14

yystate27:
	c = l.next()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate27
	}

yystate28:
	c = l.next()
	goto yyrule13

	goto yystate29 // silence unused label error
yystate29:
	c = l.next()
yystart29:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate30
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate30:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate31
	case c == '\\':
		goto yystate32
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate30
	}

yystate31:
	c = l.next()
	goto yyrule16

yystate32:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate30
	}

	goto yystate33 // silence unused label error
yystate33:
	c = l.next()
yystart33:
	switch {
	default:
		goto yyabort
	case c == '\t' || c == ' ':
		goto yystate3
	case c == '{':
		goto yystate35
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'z' || c >= '|' && c <= 'ÿ':
		goto yystate34
	}

yystate34:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'z' || c >= '|' && c <= 'ÿ':
		goto yystate34
	}

yystate35:
	c = l.next()
	goto yyrule11

	goto yystate36 // silence unused label error
yystate36:
	c = l.next()
yystart36:
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate37
	case c == '\t' || c == ' ':
		goto yystate3
	case c >= '0' && c <= '9':
		goto yystate38
	}

yystate37:
	c = l.next()
	goto yyrule19

yystate38:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9':
		goto yystate38
	}

yyrule1: // \0
	{
		return tEOF
	}
yyrule2: // \n
	{
		l.state = sInit
		return tLinebreak
		goto yystate0
	}
yyrule3: // [ \t]+
	{
		return tWhitespace
	}
yyrule4: // #[ \t]+
	{
		l.state = sComment
		goto yystate0
	}
yyrule5: // #
	{
		return l.consumeComment()
	}
yyrule6: // HELP[\t ]+
	{
		l.state = sMeta1
		return tHelp
		goto yystate0
	}
yyrule7: // TYPE[\t ]+
	{
		l.state = sMeta1
		return tType
		goto yystate0
	}
yyrule8: // {M}({M}|{D})*
	{
		l.state = sMeta2
		return tMName
		goto yystate0
	}
yyrule9: // {C}*
	{
		l.state = sInit
		return tText
		goto yystate0
	}
yyrule10: // {M}({M}|{D})*
	{
		l.state = sValue
		return tMName
		goto yystate0
	}
yyrule11: // \{
	{
		l.state = sLabels
		return tBraceOpen
		goto yystate0
	}
yyrule12: // {L}({L}|{D})*
	{
		return tLName
	}
yyrule13: // \}
	{
		l.state = sValue
		return tBraceClose
		goto yystate0
	}
yyrule14: // =
	{
		l.state = sLValue
		return tEqual
		goto yystate0
	}
yyrule15: // ,
	{
		return tComma
	}
yyrule16: // \"(\\.|[^\\"])*\"
	{
		l.state = sLabels
		return tLValue
		goto yystate0
	}
yyrule17: // [^{ \t\n]+
	{
		l.state = sTimestamp
		return tValue
		goto yystate0
	}
yyrule18: // {D}+
	{
		return tTimestamp
	}
yyrule19: // \n
	{
		l.state = sInit
		return tLinebreak
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	// Workaround to gobble up comments that started with a HELP or TYPE
	// prefix. We just consume all characters until we reach a newline.
	// This saves us from adding disproportionate complexity to the parser.
	if l.state == sComment {
		return l.consumeComment()
	}
	return tInvalid
}

func (l *promlexer) consumeComment() token {
	for c := l.cur(); ; c = l.next() {
		switch c {
		case 0:
			return tEOF
		case '\n':
			l.state = sInit
			return tComment
		}
	}
}
