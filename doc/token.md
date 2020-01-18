# Token
## TOKEN and separator
Specify lexical units (tokens and separators) that participate in SQL language.

## BNF
```
<token> ::= <nondelimiter token> | <delimiter token>
<nondelimiter token> ::=
  <REGULAR IDENTIFIER>
  | <key word>
  | <UNSIGED NUMERIC LITERAL>
  | <NATIONAL CHARACTER STRING LITERAL>
  | <BIT STRING LITERAL>
  | <HEX STRING LITERAL>
<REGULAR IDENTIFIER> ::= <identifier body>
<identifier body> ::= <identifier start> [ {<UNDERSCORE>|<identifier part>}...]
<identifier start> ::= !!! See rule
<identifier part> ::= <identifier start> | <DIGIT>
<key word> ::= <reserved word> | <non-reserved word>
<UNSIGNED NUMERIC LITERAL> ::= <exact numeric literal> | <approximate numeric literal>
<exact numeric literal> ::= <UNSIGNED INTEGER> [ <PERIOD> [ <UNSIGNED INTEGER> ] ] | <PERIOD> <UNSIGNED INTEGER>
<UNSIGNED INTEGER> ::= <DIGIT> ...
<approximate numeric literal> ::= <mantissa> E <exponent>
<mantissa> ::= <exact numeric literal>
<exponent> ::= <signed integer>
<signed integer> ::= [ <sign> ] <unsigned integer>
<sign> ::= <PLUS SIGN> | <MINUS SIGN>
<NATIONAL CHARACTER STRING LITERAL> ::= N <QUOTE> [ <character representation> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <character representation> ... ] <QUOTE> }... ]
<character representation> ::= <nonquote character> | <quote symbol>
<nonquote character> ::= !!! See rule
<quote symbol> ::= <QUOTE> <QUOTE>
<separator> ::= { <COMMENT> | <SPACE> | <NEWLINE> }...
<BIT STRING LITERAL> ::= B <QUOTE> [ <BIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <BIT> ... ] <QUOTE> }... ]
<BIT> ::= 0 | 1
<HEX STRING LITERAL> ::= X <QUOTE> [ <HEXIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <HEXIT> ... ] <QUOTE> }... ]
<HEXIT> ::= <DIGIT> | A | B | C | D | E | F | a | b | c | d | e | f
<delimiter token> ::=
  <CHARACTER STRING LITERAL>
  | <DATE STRING>
  | <TIME STRING>
  | <TIMESTAMP STRING>
  | <DELIMITED IDENTIFIER>
  | <SQL special character>
  | <NOT EQUALS OPERATOR>
  | <GREATER THAN OR EQUALS OPERATOR>
  | <LESS THAN OR EQUALS OPERATOR>
  | <CONCATENATION OPERATOR>
  | <DOUBLE PERIOD>
  | <LEFT BRACKET>
  | <RIGHT BRACKET>
<CHARACTER STRING LITERAL> ::= [ <introducer> <character set specification> ] <QUOTE> [ <character representation> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <character representation> ... ] <QUOTE> }... ]
<introducer> ::= <UNDERSCORE>
<character set specification> ::=
  <standard character repertoire name>
  | <implementation-defined character repertoire name>
  | <user-defined character repertoire name>
  | <standard universal character form-of-use name>
  | <implementation-defined universal character form-of-use name>
<standard character repertoire name> ::= <character set name>
<character set name> ::= [ <schema name> <PERIOD> ] <SQL language identifier>
<schema name> ::= [ <catalog name> <PERIOD> ] <unqualified schema name>
<catalog name> ::= <IDENTIFIER>
<unqualified schema name> ::= <IDENTIFIER>
<SQL language identifier> ::= <SQL language identifier start> [ { <UNDERSCORE> | <SQL language identifier part> }... ]
<SQL language identifier start> ::= <simple Latin letter>
<SQL language identifier part> ::= <simple Latin letter> | <DIGIT>
<implementation-defined character repertoire name> ::= <character set name>
<user-defined character repertoire name> ::= <character set name>
<standard universal character form-of-use name> ::= <character set name>
<implementation-defined universal character form-of-use name> ::= <character set name>
<DATE STRING> ::= <QUOTE> <date value> <QUOTE>
<date value> ::= <years value> <MINUS SIGN> <months value> <MINUS SIGN> <days value>
<years value> ::= <datetime value>
<datetime value> ::= <UNSIGNED INTEGER>
<months value> ::= <datetime value>
<days value> ::= <datetime value>
<TIME STRING> ::= <QUOTE> <time value> [ <time zone interval> ] <QUOTE>
<time value> ::= <hours value> <COLON> <minutes value> <COLON> <seconds value>
<hours value> ::= <datetime value>
<minutes value> ::= <datetime value>
<seconds value> ::= <seconds integer value> [ <PERIOD> [ <seconds fraction> ] ]
<seconds integer value> ::= <UNSIGNED INTEGER>
<seconds fraction> ::= <UNSIGNED INTEGER>
<time zone interval> ::= <sign> <hours value> <COLON> <minutes value>
<TIMESTAMP STRING> ::= <QUOTE> <date value> <SPACE> <time value> [ <time zone interval> ] <QUOTE>
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE> <delimited identifier body> <DOUBLE QUOTE>
<delimited identifier body> ::= <delimited identifier part> ...
<delimited identifier part> ::= <NONDOUBLEQUOTE CHARACTER> | <DOUBLEQUOTE SYMBOL>
<DOUBLEQUOTE SYMBOL> ::= <DOUBLE QUOTE> <DOUBLE QUOTE>
<IDENTIFIER> ::= [ <introducer> <character set specification> ] <actual identifier>
<actual identifier> ::= <REGULAR IDENTIFIER> | <DELIMITED IDENTIFIER>
<simple Latin letter> ::= <SIMPLE LATIN UPPER CASE LETTER> | <SIMPLE LATIN LOWER CASE LETTER>
<COMMENT> ::= <comment introducer> [ <comment character> ... ] <NEWLINE>
<comment introducer> ::= <MINUS SIGN> <MINUS SIGN> [<MINUS SIGN>...]
<comment character> ::= <nonquote character> | <QUOTE>
<SQL language character> ::= <simple Latin letter> | <DIGIT> | <SQL special character>
<SQL special character> ::=
  <SPACE>
  | <DOUBLE QUOTE>
  | <PERCENT>
  | <AMPERSAND>
  | <QUOTE>
  | <LEFT PAREN>
  | <RIGHT PAREN>
  | <ASTERISK>
  | <PLUS SIGN>
  | <COMMA>
  | <MINUS SIGN>
  | <PERIOD>
  | <SOLIDAS>
  | <COLON>
  | <SEMICOLON>
  | <LESS THAN OPERATOR>
  | <GREATER THAN OPERATOR>
  | <EQUALS OPERATOR>
  | <QUESTION MARK>
  | <UNDERSCORE>
  | <VERTICAL BAR>
<NOT EQUALS OPERATOR> ::= <LESS THAN OPERATOR><GREATER THAN OPERATOR>
<GREATER THAN OR EQUALS OPERATOR> ::= <GREATER THAN OPERATOR><EQUALS OPERATOR>
<LESS THAN OR EQUALS OPERATOR> ::= <LESS THAN OPERATOR><EQUALS OPERATOR>
<CONCATENATION OPERATOR> ::= <VERTICAL BAR><VERTICAL BAR>
<DOUBLE PERIOD> ::= <PERIOD><PERIOD>
<SPACE> ::= 0x20
<DOUBLE QUOTE> ::= "
<PERCENT> ::= %
<AMPERSAND> ::= &
<QUOTE> ::= '
<LEFT PAREN> ::= (
<RIGHT PAREN> ::= )
<ASTERISK> ::= *
<PLUS SIGN> ::= +
<COMMA> ::= ,
<MINUS SIGN> ::= -
<PERIOD> ::= .
<SOLIDAS> ::= /
<COLON> ::= :
<SEMICOLON> ::= ;
<LESS THAN OPERATOR> ::= <
<GREATER THAN OPERATOR> ::= >
<EQUALS OPERATOR> ::= =
<QUESTION MARK> ::= ?
<UNDERSCORE> ::= _
<VERTICAL BAR> ::= |
<LEFT BRACKET> ::= [
<RIGHT BRACKET> ::= ]
<SIMPLE LATIN UPPER CASE LETTER> ::= A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
<SIMPLE LATIN LOWER CASE LETTER> ::= a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
<DIGIT> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
<NEWLINE> ::= \n | \r
<reserved word> ::=
  ABSOLUTE | ACTION | ADD | ALL | ALLOCATE | ALTER | AND | ANY | ARE
  | AS | ASC | ASSERTION | AT | AUTHORIZATION | AVG
  | BEGIN | BETWEEN | BIT | BIT_LENGTH | BOTH | BY
  | CASCADE | CASCADED | CASE | CAST | CATALOG | CHAR | CHARACTER | CHARACTER_LENGTH
  | CHAR_LENGTH | CHECK | CLOSE | COALESCE | COLLATE | COLLATION | COLUMN | COMMIT
  | CONNECT | CONNECTION | CONSTRAINT | CONSTRAINTS | CONTINUE | CONVERT | CORRESPONDING
  | CREATE | CROSS | CURRENT | CURRENT_DATE | CURRENT_TIME | CURRENT_TIMESTAMP | CURRENT_USER | CURSOR
  | DATE | DAY | DEALLOCATE | DEC | DECIMAL | DECLARE | DEFAULT
  | DEFERRABLE | DEFERRED | DELETE | DESC | DESCRIBE | DESCRIPTOR | DIAGNOSTICS
  | DISCONNECT | DISTINCT | DOMAIN | DOUBLE | DROP
  | ELSE | END | END-EXEC | ESCAPE | EXCEPT | EXCEPTION | EXEC | EXECUTE | EXISTS | EXTERNAL | EXTRACT
  | FALSE | FETCH | FIRST | FLOAT | FOR | FOREIGN | FOUND | FROM | FULL
  | GET | GLOBAL | GO | GOTO | GRANT | GROUP
  | HAVING | HOUR
  | IDENTITY | IMMEDIATE | IN | INDICATOR | INITIALLY | INNER | INPUT | INSENSITIVE
  | INSERT | INT | INTEGER | INTERSECT | INTERVAL | INTO | IS | ISOLATION
  | JOIN
  | KEY
  | LANGUAGE | LAST | LEADING | LEFT | LEVEL | LIKE | LOCAL | LOWER
  | MATCH | MAX | MIN | MINUTE | MODULE | MONTH
  | NAMES | NATIONAL | NATURAL | NCHAR | NEXT | NO | NOT | NULL | NULLIF | NUMERIC
  | OCTET_LENGTH | OF | ON | ONLY | OPEN | OPTION | OR | ORDER | OUTER | OUTPUT | OVERLAPS
  | PAD | PARTIAL | POSITION | PRECISION | PREPARE | PRESERVE | PRIMARY | PRIOR | PRIVILEGES | PROCEDURE | PUBLIC
  | READ | REAL | REFERENCES | RELATIVE | RESTRICT | REVOKE | RIGHT | ROLLBACK | ROWS
  | SCHEMA | SCROLL | SECOND | SECTION | SELECT | SESSION | SESSION_USER | SET
  | SIZE | SMALLINT | SOME | SPACE | SQL | SQLCODE | SQLERROR | SQLSTATE | SUBSTRING | SUM | SYSTEM_USER
  | TABLE | TEMPORARY | THEN | TIME | TIMESTAMP | TIMEZONE_HOUR | TIMEZONE_MINUTE
  | TO | TRAILING | TRANSACTION | TRANSLATE | TRANSLATION | TRIM | TRUE
  | UNION | UNIQUE | UNKNOWN | UPDATE | UPPER | USAGE | USER | USING
  | VALUE | VALUES | VARCHAR | VARYING | VIEW
  | WHEN | WHENEVER | WHERE | WITH | WORK | WRITE
  | YEAR
  | ZONE
<non-reserved word> ::=
  ADA
  | C | CATALOG_NAME | CHARACTER_SET_CATALOG | CHARACTER_SET_NAME | CHARACTER_SET_SCHEMA
  | CLASS_ORIGIN | COBOL | COLLATION_CATALOG | COLLATION_NAME | COLLATION_SCHEMA
  | COLUMN_NAME | COMMAND_FUNCTION | COMMITTED | CONDITION_NUMBER | CONNECTION_NAME
  | CONSTRAINT_CATALOG | CONSTRAINT_NAME | CONSTRAINT_SCHEMA | CURSOR_NAME
  | DATA | DATETIME_INTERVAL_CODE | DATETIME_INTERVAL_PRECISION | DYNAMIC_FUNCTION
  | FORTRAN
  | LENGTH
  | MESSAGE_LENGTH | MESSAGE_OCTET_LENGTH | MESSAGE_TEXT | MORE | MUMPS
  | NAME | NULLABLE | NUMBER
  | PASCAL | PLI
  | REPEATABLE | RETURNED_LENGTH | RETURNED_OCTET_LENGTH | RETURNED_SQLSTATE | ROW_COUNT
  | SCALE | SCHEMA_NAME | SERIALIZABLE | SERVER_NAME | SUBCLASS_ORIGIN
  | TABLE_NAME | TYPE
  | UNCOMMITTED | UNNAMED
```

## Syntax Rules
### About `<SPACE>`
With the exception of the `<space>` character explicitly contained in `<timestamp string>` and `<interval string>` and the permitted `<separator>`s in `<bit string literal>`s and `<hex string literal>`s, a `<token>`, other than a `<character string literal>`, a `<national character string literal>`, or a `<delimited identifier>`, shall not include a `<space>` character or other `<separator>`.

### identifier start
1. A `<simple Latin letter>`
1. A character that is identified as a letter in the character repertoire identified by the `<module character set specification>` or by the `<character set specification>`
1. A character that is identified as a syllable in the character repertoire identified by the `<module character set specification>` or by the `<character set specification>`
1. A character that is identified as an ideograph in the character repertoire identified by the `<module character set specification>` or by the `<character set specification>`

### nonquote character
1. Any `<SQL language character>` other than a `<quote>`
1. Any character other than a `<quote>` in the character repertoire identified by the `<module character set specification>`
1. Any character other than a `<quote>` in the character repertoire identified by the `<character set specification>` or implied by "N".

### nondoublequote character
1. Any `<SQL language character>` other than a `<double quote>`
1. Any character other than a `<double quote>` in the character repertoire identified by the `<module character set specification>`
1. Any character other than a `<double quote>` in the character repertoire identified by the `<character set specification>`

## Character Sets
### SIMPLE LATIN LETTER
```
<simple Latin letter> ::=
     A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
   | a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
```

### DIGIT
```
<DIGIT> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
```

### SQL LANGUAGE CHARACTER
```
<SQL language character> ::=
     A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
   | a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
   | <SPACE> | <DOUBLE QUOTE> | <PERCENT> | <AMPERSAND> | <QUOTE> | <LEFT PAREN> | <RIGHT PAREN> | <ASTERISK>
   | <PLUS SIGN> | <COMMA> | <MINUS SIGN> | <PERIOD> | <SOLIDAS> | <COLON> | <SEMICOLON> | <LESS THAN OPERATOR>
   | <GREATER THAN OPERATOR> | <EQUALS OPERATOR> | <QUESTION MARK> | <UNDERSCORE> | <VERTICAL BAR>
```


### QUOTE ESCAPE CHARACTER
```
<QUOTE ESCAPE CHARACTER> ::=
     A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
   | a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
   | <SPACE> | <DOUBLE QUOTE> | <PERCENT> | <AMPERSAND> | <LEFT PAREN> | <RIGHT PAREN> | <ASTERISK> | <PLUS SIGN>
   | <COMMA> | <MINUS SIGN>| <PERIOD>| <SOLIDAS> | <COLON> | <SEMICOLON> | <LESS THAN OPERATOR>| <GREATER THAN OPERATOR>
   | <EQUALS OPERATOR> | <QUESTION MARK> | <UNDERSCORE> | <VERTICAL BAR> | <QUOTE><QUOTE>
```

### DOUBLEQUOTE ESCAPE CHARACTER
```
<DOUBLEQUOTE ESCAPE CHARACTER> ::=
     A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
   | a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
   | <SPACE> | <PERCENT> | <AMPERSAND> | <QUOTE> | <LEFT PAREN> | <RIGHT PAREN> | <ASTERISK>| <PLUS SIGN>
   | <COMMA> | <MINUS SIGN> | <PERIOD> | <SOLIDAS> | <COLON> | <SEMICOLON> | <LESS THAN OPERATOR>
   | <GREATER THAN OPERATOR> | <EQUALS OPERATOR> | <QUESTION MARK> | <UNDERSCORE> | <VERTICAL BAR>
   | <DOUBLE QUOTE><DOUBLE QUOTE>
```

## Define Token
### REGULAR IDENTIFIER
#### BNF
```
<REGULAR IDENTIFIER> ::= <identifier body>
<identifier body> ::= <identifier start> [ {<UNDERSCORE>|<identifier part>}...]
<identifier start> ::= !!! See rule
<identifier part> ::= <identifier start> | <DIGIT>
```
#### simplified BNF
```
<REGULAR IDENTIFIER> ::= <SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]
```

### KEYWORD
#### BNF
Omitted (See above)
### simplified BNF
```
<KEYWORD> ::= <SIMPLE LATIN LETTER>...
```

### UNSIGNED NUMERIC LITERAL
#### BNF
```
<UNSIGNED NUMERIC LITERAL> ::= <exact numeric literal> | <approximate numeric literal>
<exact numeric literal> ::= <UNSIGNED INTEGER> [ <PERIOD> [ <UNSIGNED INTEGER> ] ] | <PERIOD> <UNSIGNED INTEGER>
<UNSIGNED INTEGER> ::= <DIGIT> ...
<approximate numeric literal> ::= <mantissa> E <exponent>
<mantissa> ::= <exact numeric literal>
<exponent> ::= <signed integer>
<signed integer> ::= [ <sign> ] <unsigned integer>
<sign> ::= <PLUS SIGN> | <MINUS SIGN>
```
#### simplified BNF
```
<UNSIGNED NUMERIC LITERAL> ::= {<DIGIT>... [<PERIOD> [<DIGIT>...]] | <PERIOD> <DIGIT>...} [E [<PLUS SIGN>|<MINUS SIGN>] <DIGIT>...]
```

### NATIONAL CHARACTER STRING LITERAL
#### BNF
```
<NATIONAL CHARACTER STRING LITERAL> ::= N <QUOTE> [ <character representation> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <character representation> ... ] <QUOTE> }... ]
<character representation> ::= <nonquote character> | <quote symbol>
<nonquote character> ::= !!! See rule
<quote symbol> ::= <QUOTE> <QUOTE>
<separator> ::= { <COMMENT> | <SPACE> | <NEWLINE> }...
```
#### simplified BNF
```
<NATIONAL CHARACTER STRING LITERAL> ::= N <QUOTE> [<QUOTE ESCAPE CHARACTER>...]<QUOTE>[{<SEPARATOR>... <QUOTE>[<QUOTE ESCAPE CHARACTER>}...]<QUOTE>}...]
```

### SEPARATOR
#### BNF
```
<separator> ::= { <COMMENT> | <SPACE> | <NEWLINE> }...
```
#### simplified BNF
```
<SEPARATOR> ::= { <COMMENT> | <SPACE> | {\r | \n } }...
```

### COMMENT
#### BNF
```
<COMMENT> ::= <comment introducer> [ <comment character> ... ] <NEWLINE>
<comment introducer> ::= <MINUS SIGN> <MINUS SIGN> [<MINUS SIGN>...]
<comment character> ::= <nonquote character> | <QUOTE>
```
#### simplified BNF
```
<COMMENT> ::= <MINUS SIGN><MINUS SIGN>[<MINUS SIGN>...][<SQL LANGUAGE CHARACTER>...]{\r | \n}
```

### BIT STRING LITERAL
#### BNF
```
<BIT STRING LITERAL> ::= B <QUOTE> [ <BIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <BIT> ... ] <QUOTE> }... ]
<BIT> ::= 0 | 1
```
#### simplified BNF
```
<BIT STRING LITERAL> ::= B <QUOTE>[{0|1}...]<QUOTE>[{<SEPARATOR>... <QUOTE>[{0|1}...]<QUOTE>}...]
```

### HEX STRING LITERAL
#### BNF
```
<HEX STRING LITERAL> ::= X <QUOTE> [ <HEXIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <HEXIT> ... ] <QUOTE> }... ]
<HEXIT> ::= <DIGIT> | A | B | C | D | E | F | a | b | c | d | e | f
```
#### simplified BNF
```
<HEX STRING LITERAL> ::= X <QUOTE>[{<DIGIT>|A|B|C|D|E|F|a|b|c|d|e|f}...]<QUOTE>[{<SEPARATOR>... <QUOTE>[{<DIGIT>|A|B|C|D|E|F|a|b|c|d|e|f}...]<QUOTE>}...]
```

### CHARACTER STRING LITERAL
#### BNF
```
<CHARACTER STRING LITERAL> ::= [ <introducer> <character set specification> ] <QUOTE> [ <character representation> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <character representation> ... ] <QUOTE> }... ]
<introducer> ::= <UNDERSCORE>
<character set specification> ::=
  <standard character repertoire name>
  | <implementation-defined character repertoire name>
  | <user-defined character repertoire name>
  | <standard universal character form-of-use name>
  | <implementation-defined universal character form-of-use name>
<standard character repertoire name> ::= <character set name>
<character set name> ::= [ <schema name> <PERIOD> ] <SQL language identifier>
<schema name> ::= [ <catalog name> <PERIOD> ] <unqualified schema name>
<catalog name> ::= <IDENTIFIER>
<unqualified schema name> ::= <IDENTIFIER>
<SQL language identifier> ::= <SQL language identifier start> [ { <UNDERSCORE> | <SQL language identifier part> }... ]
<SQL language identifier start> ::= <simple Latin letter>
<SQL language identifier part> ::= <simple Latin letter> | <DIGIT>
<implementation-defined character repertoire name> ::= <character set name>
<user-defined character repertoire name> ::= <character set name>
<standard universal character form-of-use name> ::= <character set name>
<implementation-defined universal character form-of-use name> ::= <character set name>
<character representation> ::= <nonquote character> | <quote symbol>
<quote symbol> ::= <QUOTE> <QUOTE>
```
#### simplified BNF
```
<CHARACTER STRING LITERAL> ::= [<UNDERSCORE>[[<IDENTIFIER><PERIOD>]<IDENTIFIER><PERIOD>]<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]]<QUOTE>[<QUOTE ESCAPE CHARACTER>...]<QUOTE>[{<SEPARATOR>... <QUOTE>[<QUOTE ESCAPE CHARACTER>...]<QUOTE>}...]
```

### IDENTIFIER
#### BNF
```
<IDENTIFIER> ::= [ <introducer> <character set specification> ] <actual identifier>
<actual identifier> ::= <REGULAR IDENTIFIER> | <DELIMITED IDENTIFIER>
<introducer> ::= <UNDERSCORE>
<character set specification> ::=
  <standard character repertoire name>
  | <implementation-defined character repertoire name>
  | <user-defined character repertoire name>
  | <standard universal character form-of-use name>
  | <implementation-defined universal character form-of-use name>
<standard character repertoire name> ::= <character set name>
<character set name> ::= [ <schema name> <PERIOD> ] <SQL language identifier>
<schema name> ::= [ <catalog name> <PERIOD> ] <unqualified schema name>
<catalog name> ::= <IDENTIFIER>
<unqualified schema name> ::= <IDENTIFIER>
<SQL language identifier> ::= <SQL language identifier start> [ { <UNDERSCORE> | <SQL language identifier part> }... ]
<SQL language identifier start> ::= <simple Latin letter>
<SQL language identifier part> ::= <simple Latin letter> | <DIGIT>
<implementation-defined character repertoire name> ::= <character set name>
<user-defined character repertoire name> ::= <character set name>
<standard universal character form-of-use name> ::= <character set name>
<implementation-defined universal character form-of-use name> ::= <character set name>
<REGULAR IDENTIFIER> ::= <identifier body>
<identifier body> ::= <identifier start> [ {<UNDERSCORE>|<identifier part>}...]
<identifier start> ::= !!! See rule
<identifier part> ::= <identifier start> | <DIGIT>
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE> <delimited identifier body> <DOUBLE QUOTE>
<delimited identifier body> ::= <delimited identifier part> ...
<delimited identifier part> ::= <NONDOUBLEQUOTE CHARACTER> | <DOUBLEQUOTE SYMBOL>
<DOUBLEQUOTE SYMBOL> ::= <DOUBLE QUOTE> <DOUBLE QUOTE>
```
#### simplified BNF
```
<IDENTIFIER> ::= [<UNDERSCORE>[[<IDENTIFIER><PERIOD>]<IDENTIFIER><PERIOD>]<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]]{{<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]}|{<DOUBLE QUOTE><DOUBLEQUOTE ESCAPE CHARACTER>...<DOUBLE QUOTE>}}
```

### DELIMITED IDENTIFIER
#### BNF
```
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE> <delimited identifier body> <DOUBLE QUOTE>
<delimited identifier body> ::= <delimited identifier part> ...
<delimited identifier part> ::= <NONDOUBLEQUOTE CHARACTER> | <DOUBLEQUOTE SYMBOL>
<DOUBLEQUOTE SYMBOL> ::= <DOUBLE QUOTE> <DOUBLE QUOTE>
```
#### simplified BNF
```
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE><DOUBLEQUOTE ESCAPE CHARACTER>...<DOUBLE QUOTE>
```

### DATE STRING
#### BNF
```
<DATE STRING> ::= <QUOTE> <date value> <QUOTE>
<date value> ::= <years value> <MINUS SIGN> <months value> <MINUS SIGN> <days value>
<years value> ::= <datetime value>
<datetime value> ::= <UNSIGNED INTEGER>
<months value> ::= <datetime value>
<days value> ::= <datetime value>
<UNSIGNED INTEGER> ::= <DIGIT>...
```
#### simplified BNF
```
<DATE STRING> ::= <QUOTE><DIGIT>...<MINUS SIGN><DIGIT>...<MINUS SIGN><DIGIT>...<QUOTE>
```

### TIME STRING
#### BNF
```
<TIME STRING> ::= <QUOTE> <time value> [ <time zone interval> ] <QUOTE>
<time value> ::= <hours value> <COLON> <minutes value> <COLON> <seconds value>
<hours value> ::= <datetime value>
<minutes value> ::= <datetime value>
<seconds value> ::= <seconds integer value> [ <PERIOD> [ <seconds fraction> ] ]
<seconds integer value> ::= <UNSIGNED INTEGER>
<seconds fraction> ::= <UNSIGNED INTEGER>
<time zone interval> ::= <sign> <hours value> <COLON> <minutes value>
<datetime value> ::= <UNSIGNED INTEGER>
<UNSIGNED INTEGER> ::= <DIGIT>...
<sign> ::= <PLUS SIGN> | <MINUS SIGN>
```
#### simplified BNF
```
<TIME STRING> ::= <QUOTE><DIGIT>...<COLON><DIGIT>...<COLON><DIGIT>...[<PERIOD>[<DIGIT>...]][{<PLUS SUGN>|<MINUS SIGN>}<DIGIT>...<COLON><DIGIT>...]<QUOTE>
```

### TIMESTAMP STRING
#### BNF
```
<TIMESTAMP STRING> ::= <QUOTE> <date value> <SPACE> <time value> [ <time zone interval> ] <QUOTE>
<date value> ::= <years value> <MINUS SIGN> <months value> <MINUS SIGN> <days value>
<years value> ::= <datetime value>
<datetime value> ::= <UNSIGNED INTEGER>
<months value> ::= <datetime value>
<days value> ::= <datetime value>
<UNSIGNED INTEGER> ::= <DIGIT>...
<time value> ::= <hours value> <COLON> <minutes value> <COLON> <seconds value>
<hours value> ::= <datetime value>
<minutes value> ::= <datetime value>
<seconds value> ::= <seconds integer value> [ <PERIOD> [ <seconds fraction> ] ]
<seconds integer value> ::= <UNSIGNED INTEGER>
<seconds fraction> ::= <UNSIGNED INTEGER>
<time zone interval> ::= <sign> <hours value> <COLON> <minutes value>
<sign> ::= <PLUS SIGN> | <MINUS SIGN>
```
#### simplified BNF
```
<TIMESTAMP STRING> ::= <QUOTE><DIGIT>...<MINUS SIGN><DIGIT>...<MINUS SIGN><DIGIT>...<SPACE><DIGIT>...<COLON><DIGIT>...<COLON><DIGIT>...[<PERIOD>[<DIGIT>...]][{<PLUS SUGN>|<MINUS SIGN>}<DIGIT>...<COLON><DIGIT>...]<QUOTE>
```

### OTHER DELIMITER IDENTIFIER
#### BNF
```
<OTHER DELIMITER IDENTIFER> ::=
  <SQL special character>
  | <NOT EQUALS OPERATOR>
  | <GREATER THAN OR EQUALS OPERATOR>
  | <LESS THAN OR EQUALS OPERATOR>
  | <CONCATENATION OPERATOR>
  | <DOUBLE PERIOD>
  | <LEFT BRACKET>
  | <RIGHT BRACKET>
```
#### simplified BNF
```
<OTHER DELIMITER IDENTIFIER> ::= <SQL special character>...
```

## TOKEN conclusion
### simplified BNF
```
<REGULAR IDENTIFIER> ::= <SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]
<KEYWORD> ::= <SIMPLE LATIN LETTER>...
<UNSIGNED NUMERIC LITERAL> ::= {<DIGIT>... [<PERIOD> [<DIGIT>...]] | <PERIOD> <DIGIT>...} [E [<PLUS SIGN>|<MINUS SIGN>] <DIGIT>...]
<NATIONAL CHARACTER STRING LITERAL> ::= N <QUOTE> [<QUOTE ESCAPE CHARACTER>...]<QUOTE>[{<SEPARATOR>... <QUOTE>[<QUOTE ESCAPE CHARACTER>}...]<QUOTE>}...]
<SEPARATOR> ::= { <COMMENT> | <SPACE> | {\r | \n } }...
<COMMENT> ::= <MINUS SIGN><MINUS SIGN>[<MINUS SIGN>...][<SQL LANGUAGE CHARACTER>...]{\r | \n}
<BIT STRING LITERAL> ::= B <QUOTE>[{0|1}...]<QUOTE>[{<SEPARATOR>... <QUOTE>[{0|1}...]<QUOTE>}...]
<HEX STRING LITERAL> ::= X <QUOTE>[{<DIGIT>|A|B|C|D|E|F|a|b|c|d|e|f}...]<QUOTE>[{<SEPARATOR>... <QUOTE>[{<DIGIT>|A|B|C|D|E|F|a|b|c|d|e|f}...]<QUOTE>}...]
<CHARACTER STRING LITERAL> ::= [<UNDERSCORE>[[<IDENTIFIER><PERIOD>]<IDENTIFIER><PERIOD>]<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]]<QUOTE>[<QUOTE ESCAPE CHARACTER>...]<QUOTE>[{<SEPARATOR>... <QUOTE>[<QUOTE ESCAPE CHARACTER>...]<QUOTE>}...]
<IDENTIFIER> ::= [<UNDERSCORE>[[<IDENTIFIER><PERIOD>]<IDENTIFIER><PERIOD>]<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]]{{<SIMPLE LATIN LETTER>[{<UNDERSCORE>|<SIMPLE LATIN LETTER>|<DIGIT>}...]}|{<DOUBLE QUOTE><DOUBLEQUOTE ESCAPE CHARACTER>...<DOUBLE QUOTE>}}
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE><DOUBLEQUOTE ESCAPE CHARACTER>...<DOUBLE QUOTE>
<DATE STRING> ::= <QUOTE><DIGIT>...<MINUS SIGN><DIGIT>...<MINUS SIGN><DIGIT>...<QUOTE>
<TIME STRING> ::= <QUOTE><DIGIT>...<COLON><DIGIT>...<COLON><DIGIT>...[<PERIOD>[<DIGIT>...]][{<PLUS SUGN>|<MINUS SIGN>}<DIGIT>...<COLON><DIGIT>...]<QUOTE>
<TIMESTAMP STRING> ::= <QUOTE><DIGIT>...<MINUS SIGN><DIGIT>...<MINUS SIGN><DIGIT>...<SPACE><DIGIT>...<COLON><DIGIT>...<COLON><DIGIT>...[<PERIOD>[<DIGIT>...]][{<PLUS SUGN>|<MINUS SIGN>}<DIGIT>...<COLON><DIGIT>...]<QUOTE>
<OTHER DELIMITER IDENTIFIER> ::= <SQL special character>...
```

### First Charavter to Token Type
#### SIMPLE LATIN LETTER
##### B
- `<BIT STRING LETTER>`
- `<REGULAR IDENTIFIER>`
- `<KEYWORD>`
- `<IDENTIFIER>`
- `<CHARACTER STRING LITERAL>`

##### N
- `<NATIONAL CHARACTER STRING LETTER>`
- `<REGULAR IDENTIFIER>`
- `<KEYWORD>`
- `<IDENTIFIER>`
- `<CHARACTER STRING LITERAL>`

##### X
- `<HEX STRING LETTER>`
- `<REGULAR IDENTIFIER>`
- `<KEYWORD>`
- `<IDENTIFIER>`
- `<CHARACTER STRING LITERAL>`

##### Others
- `<REGULAR IDENTIFIER>`
- `<KEYWORD>`
- `<IDENTIFIER>`
- `<CHARACTER STRING LITERAL>`

#### DIGIT
- `<UNSIGNED NUMERIC LITERAL>`

#### SQL special character
##### PERIOD
- `<UNSIGNED NUMERIC LITERAL>`
- `<OTHER DELMITER IDENTIFIER>`

##### UNDERSCORE
- `<CHARACTER STRING LITERAL>`
- `<IDENTIFIER>`
- `<OTHER DELMITER IDENTIFIER>`

##### MINUS SIGN
- `<COMMENT>`
- `<OTHER DELMITER IDENTIFIER>`

##### DOUBLE QUOTE
- `<DELIMITED IDENTIFIER>`
- `<OTHER DELIMITER IDENTIFIER>`

##### QUOTE
- `<DATE STRING>`
- `<TIME STRING>`
- `<TIMESTAMP STRING>`
- `<OTHER DELIMITER IDENTIFIER>`

##### Others
- `<OTHER DELIMITER IDENTIFIER>`
