// Copyright 2021 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

type SemanticTokenTypes string

const (
	SemanticTokenNamespace SemanticTokenTypes = "namespace"

	// Represents a generic type. Acts as a fallback for types which
	// can't be mapped to a specific type like class or enum.
	SemanticTokenType          SemanticTokenTypes = "type"
	SemanticTokenClass         SemanticTokenTypes = "class"
	SemanticTokenEnum          SemanticTokenTypes = "enum"
	SemanticTokenInterface     SemanticTokenTypes = "interface"
	SemanticTokenStruct        SemanticTokenTypes = "struct"
	SemanticTokenTypeParameter SemanticTokenTypes = "typeParameter"
	SemanticTokenParameter     SemanticTokenTypes = "parameter"
	SemanticTokenVariable      SemanticTokenTypes = "variable"
	SemanticTokenProperty      SemanticTokenTypes = "property"
	SemanticTokenEnumMember    SemanticTokenTypes = "enumMember"
	SemanticTokenEvent         SemanticTokenTypes = "event"
	SemanticTokenFunction      SemanticTokenTypes = "function"
	SemanticTokenMethod        SemanticTokenTypes = "method"
	SemanticTokenMacro         SemanticTokenTypes = "macro"
	SemanticTokenKeyword       SemanticTokenTypes = "keyword"
	SemanticTokenModifier      SemanticTokenTypes = "modifier"
	SemanticTokenComment       SemanticTokenTypes = "comment"
	SemanticTokenString        SemanticTokenTypes = "string"
	SemanticTokenNumber        SemanticTokenTypes = "number"
	SemanticTokenRegexp        SemanticTokenTypes = "regexp"
	SemanticTokenOperator      SemanticTokenTypes = "operator"
)

type SemanticTokenModifiers string

const (
	SemanticTokenModifierDeclaration    SemanticTokenModifiers = "declaration"
	SemanticTokenModifierDefinition     SemanticTokenModifiers = "definition"
	SemanticTokenModifierReadonly       SemanticTokenModifiers = "readonly"
	SemanticTokenModifierStatic         SemanticTokenModifiers = "static"
	SemanticTokenModifierDeprecated     SemanticTokenModifiers = "deprecated"
	SemanticTokenModifierAbstract       SemanticTokenModifiers = "abstract"
	SemanticTokenModifierAsync          SemanticTokenModifiers = "async"
	SemanticTokenModifierModification   SemanticTokenModifiers = "modification"
	SemanticTokenModifierDocumentation  SemanticTokenModifiers = "documentation"
	SemanticTokenModifierDefaultLibrary SemanticTokenModifiers = "defaultLibrary"
)

// TokenFormat is an additional token format capability to allow future extensions of the format.
//
// @since 3.16.0.
type TokenFormat string

// TokenFormatRelative described using relative positions.
const TokenFormatRelative TokenFormat = "relative"

type SemanticTokensLegend struct {
	// TokenTypes is the token types a server uses.
	TokenTypes []SemanticTokenTypes `json:"tokenTypes"`

	// TokenModifiers is the token modifiers a server uses.
	TokenModifiers []SemanticTokenModifiers `json:"tokenModifiers"`
}
