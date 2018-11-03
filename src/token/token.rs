pub type TokenType = String;

pub struct Token {
    pub token_type: TokenType,
    pub literal: String,
}

impl Token {
    fn from_char(token_type: TokenType, ch: char) -> Token {
        Token {
            token_type,
            literal: ch.to_string(),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_from_char() {
        let t = Token::from_char(String::from("foo"), 'a');
        assert_eq!(t.token_type, String::from("foo"));
        assert_eq!(t.literal, String::from("a"))
    }
}