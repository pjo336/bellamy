use std::collections::HashMap;

use super::token;

static ILLEGAL: &'static str = "ILLEGAL";
static EOF: &'static str = "EOF";

static IDENT: &'static str = "IDENT";
static INT: &'static str = "INT";

static ASSIGN: &'static str = "=";
static PLUS: &'static str = "+";
static MINUS: &'static str = "-";
static BANG: &'static str = "!";
static ASTERISK: &'static str = "*";
static SLASH: &'static str = "/";
static LT: &'static str = "<";
static GT: &'static str = ">";
static EQ: &'static str = "==";
static NE: &'static str = "!=";

static COMMA: &'static str = ",";
static SEMICOLON: &'static str = ";";

static LPAREN: &'static str = "(";
static RPAREN: &'static str = ")";
static LBRACE: &'static str = "{";
static RBRACE: &'static str = "}";

// Keywords
static FUNCTION: &'static str = "FUNCTION";
static LET: &'static str = "LET";
static TRUE: &'static str = "TRUE";
static FALSE: &'static str = "FALSE";
static RETURN: &'static str = "RETURN";
static IF: &'static str = "IF";
static ELSE: &'static str = "ELSE";

pub fn lookup_ident(ident: &str) -> Option<&token::TokenType> {
    let map = keyword_map();
    let res = &map.get(ident);
}

// TODO: Make static initialized for performance
fn keyword_map() -> HashMap<String, token::TokenType> {
    [
        ( String::from("fn"), String::from(FUNCTION)),
//        ("let", String::from(LET)),
//        ("true", String::from(TRUE)),
//        ("false", String::from(FALSE)),
//        ("return", String::from(RETURN)),
//        ("if", String::from(IF)),
//        ("else", String::from(ELSE)),
    ].iter().cloned().collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_lookup_ident() {
        assert_eq!(lookup_ident(&"foo"), None);
    }
}