pub struct Lexer {
    input: String,
    curr_pos: usize,
    next_pos: usize,
    ch: char,
}

pub fn init_lexer(input: &String) -> Lexer {
    Lexer {
        input: input.to_string(),
        curr_pos: 0,
        next_pos: 1,
        ch: input.chars().next().unwrap_or('0'),
    }
}

impl Lexer {
    fn advance_char(&mut self) -> char {
        self.ch = self.peek_char();
        self.advance_pointers();
        self.ch
    }

    fn peek_char(&mut self) -> char {
        if self.next_pos >= self.input.len() {
            return '0';
        }
        self.input.chars().nth(self.next_pos).unwrap()
    }

    fn read_ident(&mut self) -> &str {
        let curr_pos = self.curr_pos;
        while is_letter(self.ch) {
            self.advance_char();
        }
        &mut self.input[curr_pos..self.curr_pos]
    }

    fn read_number(&mut self) -> &str {
        let curr_pos = self.curr_pos;
        while is_digit(self.ch) {
            self.advance_char();
        }
        &mut self.input[curr_pos..self.curr_pos]
    }

    fn skip_white_space(&mut self) {
        while is_whitespace(self.ch) {
            self.advance_char();
        }
    }

    fn advance_pointers(&mut self) {
        self.curr_pos = self.next_pos;
        self.next_pos += 1;
    }
}

fn is_whitespace(ch: char) -> bool {
    return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r';
}

fn is_letter(ch: char) -> bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z';
}

fn is_digit(ch: char) -> bool {
    return '0' <= ch && ch <= '9';
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_init_lexer() {
        let input = String::from("let foo = 123;");
        let l = init_lexer(&input);
        assert_eq!(l.input, input);
        assert_eq!(l.curr_pos, 0);
        assert_eq!(l.next_pos, 1);
        assert_eq!(l.ch, 'l');
    }

    #[test]
    fn test_advance_char() {
        let l = &mut fake_lexer("let foo = 123;");
        assert_eq!(l.advance_char(), 'e');
        assert_eq!(l.curr_pos, 1);
        assert_eq!(l.next_pos, 2);
        assert_eq!(l.advance_char(), 't');
    }

    #[test]
    fn test_peek_char() {
        let l = &mut fake_lexer("let foo = 123;");
        assert_eq!(l.peek_char(), 'e');
    }

    #[test]
    fn test_peek_char_empty() {
        let l = &mut fake_lexer(&String::from(""));
        assert_eq!(l.peek_char(), '0');
    }

    #[test]
    fn test_read_ident() {
        let l = &mut fake_lexer(&String::from("foobar 123"));
        assert_eq!(l.read_ident(), "foobar");
    }

    #[test]
    fn test_read_number() {
        let l = &mut fake_lexer(&String::from("123 foobar"));
        assert_eq!(l.read_number(), "123");
    }

    #[test]
    fn test_skip_white_space() {
        let l = &mut fake_lexer(&String::from("a  \n \t \r b"));
        l.advance_char();
        l.skip_white_space();
        assert_eq!(l.ch, 'b');
    }

    #[test]
    fn test_is_letter() {
        assert_eq!(is_letter('e'), true);
        assert_eq!(is_letter('4'), false);
    }

    #[test]
    fn test_is_digit() {
        assert_eq!(is_digit('e'), false);
        assert_eq!(is_digit('4'), true);
    }

    fn fake_lexer(input: &str) -> Lexer {
        init_lexer(&String::from(input))
    }
}