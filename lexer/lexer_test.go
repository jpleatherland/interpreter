package lexer

import (
	"testing"

	"github.com/jpleatherland/interpreter/token"
)

func TestNextToken(t *testing.T) {
	type tokenTest struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	type testCase struct {
		name           string
		input          string
		expectedTokens []tokenTest
	}

	var tests []testCase

	tests = append(tests, testCase{
		name:  "Simple symbols",
		input: `=+(){},;`,
		expectedTokens: []tokenTest{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		},
	})

	tests = append(tests, testCase{
		name: "assignments and functions",
		input: `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`,
		expectedTokens: []tokenTest{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lexer := New(tt.input)
			for i, expected := range tt.expectedTokens {
				tok := lexer.NextToken()

				if tok.Type != expected.expectedType {
					t.Fatalf("test %q, token[%d] - expected type: %q, got: %q",
						tt.name, i, expected.expectedType, tok.Type)
				}

				if tok.Literal != expected.expectedLiteral {
					t.Fatalf("test %q, token[%d] - expected literal: %q, got: %q",
						tt.name, i, expected.expectedLiteral, tok.Literal)
				}

			}
		})
	}
}
