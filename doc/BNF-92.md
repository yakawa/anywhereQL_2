# SQL-92 BNF
## Terminator / Characters
```
<SIMPLE LATIN UPPER CASE LETTER> ::= A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z
<SIMPLE LATIN LOWER CASE LETTER> ::= a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
<DIGIT> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
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
<NEWLINE> ::= \r | \n | \r\n
<NOT EQUALS OPERATOR> ::= <>
<GREATER THAN OR EQUALS OPERATOR> ::= >=
<LESS THAN OR EQUALS OPERATOR> ::= <=
<CONCATENATION OPERATOR> ::= ||
<DOUBLE PERIOD> ::= ..
```
## Terminator Token / Identifier
```
<TOKEN> ::= <nondelimiter token> | <delimiter token>
<nondelimiter token> ::=
  <regular identifier>
  | <key word>
  | <UNSIGED NUMERICAL LITERAL>
  | <NATINAL CHARACTER STRING LITERAL>
  | <BIT STRING LITERAL>
  | <HEX STRING LITERAL>
<REGULAR IDENTIFIER> ::= <identifier body>
<identifier body> ::= <identifier start> [ {<UNDERSCORE>|<identifier part>}...]
<identifier start> ::= !!!
<identifier part> ::= <identifier start> | <DIGIT>
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
<IDENTIFIER> ::= [ <introducer> <character set specification> ] <actual identifier>
<actual identifier> ::= <REGULAR IDENTIFIER> | <DELIMITED IDENTIFIER>
<DELIMITED IDENTIFIER> ::= <DOUBLE QUOTE> <delimited identifier body> <DOUBLE QUOTE>
<delimited identifier body> ::= <delimited identifier part> ...
<delimited identifier part> ::= <nondoublequote character> | <DOUBLEQUOTE SYMBOL>
<DOUBLEQUOTE SYMBOL> ::= <DOUBLE QUOTE> <DOUBLE QUOTE>
```
## Terminator / Literal
```
<unsigned literal> ::= <UNSIGNED NUMERIC LITERAL> | <general literal>
<UNSIGNED NUMERICAL LITERAL> ::= <exact numeric literal> | <approximate numeric literal>
<exact numeric literal> ::= <UNSIGNED INTEGER> [ <PERIOD> [ <UNSIGNED INTEGER> ] ] | <PERIOD> <UNSIGNED INTEGER>
<UNSIGNED INTEGER> ::= <DIGIT> ...
<approximate numeric literal> ::= <mantissa> E <exponent>
<mantissa> ::= <exact numeric literal>
<exponent> ::= <signed integer>
<signed integer> ::= [ <sign> ] <unsigned integer>
<sign> ::= <PLUS SIGN> | <MINUS SIGN>
<NATINAL CHARACTER STRING LITERAL> ::= N <QUOTE> [ <character representation> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <character representation> ... ] <QUOTE> }... ]
<character representation> ::= <nonquote character> | <quote symbol>
<nonquote character> ::= !!!
<quote symbol> ::= <QUOTE> <QUOTE>
<separator> ::= { <COMMENT> | <SPACE> | <NEWLINE> }...
<COMMENT> ::= <comment introducer> [ <comment character> ... ] <NEWLINE>
<comment introducer> ::= <MINUS SIGN> <MINUS SIGN> [<MINUS SIGN>...]
<comment character> ::= <nonquote character> | <QUOTE>
<BIT STRING LITERAL> ::= B <QUOTE> [ <BIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <BIT> ... ] <QUOTE> }... ]
<BIT> ::= 0 | 1
<HEX STRING LITERAL> ::= X <QUOTE> [ <HEXIT> ... ] <QUOTE> [ { <separator> ... <QUOTE> [ <HEXIT> ... ] <QUOTE> }... ]
<HEXIT> ::= <DIGIT> | A | B | C | D | E | F | a | b | c | d | e | f
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
<SQL language identifier> ::= <SQL language identifier start> [ { <UNDERSCORE> | <SQL language identifier part> }... ]
<SQL language identifier start> ::= <simple Latin letter>
<SQL language identifier part> ::= <simple Latin letter> | <DIGIT>
<unqualified schema name> ::= <IDENTIFIER>
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
<INTERVAL STRING> ::= <QUOTE> { <year-month literal> | <day-time literal> } <QUOTE>
<year-month literal> ::= <years value> | [ <years value> <MINUS SIGN> ] <months value>
<day-time literal> ::= <day-time interval> | <time interval>
<day-time interval> ::= <days value> [ <SPACE> <hours value> [ <COLON> <minutes value> [ <COLON> <seconds value> ] ] ]
<time-interval> ::=
  <hours value> [ <COLON> <minutes value> [ <COLON> <seconds value> ] ]
  | <minutes value> [ <COLON> <seconds value> ]
  | <seconds value>
<general literal> ::=
  <CHARACTER STRING LITERAL>
  | <NATINAL CHARACTER STRING LITERAL>
  | <BIT STRING LITERAL>
  | <HEX STRING LITERAL>
  | <DATETIME LITERAL>
  | <INTERVAL LITERAL>
<DATETIME LITERAL> ::= <DATE LITERAL> | <TIME LITERAL> | <TIMESTAMP LITERAL>
<DATE LITERAL> ::= DATE <DATE STRING>
<TIME LITERAL> ::= TIME <TIME STRING>
<TIMESTAMP LITERAL> ::= TIMESTAMP <TIMESTAMP STRING>
<INTERVAL LITERAL> ::= INTERVAL [<sign>] <INTERVAL STRING> <interval qualifier>
<interval qualifier> ::=
  <start field> TO <end field>
  | <single datetime field>
<start field> ::= <non-second datetime field> [<LEFT PAREN> <interval leading field precision> <RIGHT PAREN>]
<non-second datetime field> ::= YEAR | MONTH | DAY | HOUR | MINUTE
<interval leading field percision> ::= <UNSIGNED INTEGER>
<end field> ::=
  <non-second datetime field>
  | SECOND [<LEFT PAREN><interval fractional seconds precision> <RIGHT PAREN>]
<interval fractional seconds percision> ::= <UNSIGNED INTEGER>
<single datetime field> ::=
  <non-second datetime field> [<LEFT PAREN><interval leading field precision><RIGHT PAREN>]
  | SECOND [<LEFT PAREN><interval leading field precision>[<COMMA><LEFT PAREN><interval fractional seconds percision>]<RIGHT PAREN>]
```

## Terminator / Keyword
```
<key word> ::= <reserved word> | <non-reserved word>
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
## Terminator / Other
```
<SQL terminal character> ::= <SQL language character> | <SQL embedded language character>
<SQL language character> ::= <simple Latin letter> | <DIGIT> | <SQL special character>
<SQL embedded language character> ::= <LEFT BRACKET> | <RIGHT BRACKET>
<simple Latin letter> ::= <SIMPLE LATIN UPPER CASE LETTER> | <SIMPLE LATIN LOWER CASE LETTER>
```
## Data Type
```
<data type> ::=
  <character string type>[CHARACTER SET <character set specification>]
  | <natinal character string type>
  | <bit string type>
  | <numeric type>
  | <datetime type>
  | <interval type>
<character string type> ::=
  CHARACTER [<LEFT PAREN><length><RIGHT PAREN>]
  | CHAR [<LEFT PAREN><length><RIGHT PAREN>]
  | CHARACTER VARYING [<LEFT PAREN><length><RIGHT PAREN>]
  | CHAR VARYING [<LEFT PAREN><length><RIGHT PAREN>]
  | VARCHAR [<LEFT PAREN><length><RIGHT PAREN>]
<length> ::= <UNSIGNED INTEGER>
<natinal character string type> ::=
  NATINAL CHARACTER [<LEFT PAREN><length><RIGHT PAREN>]
  | NATINAL CHAR [<LEFT PAREN><length><RIGHT PAREN>]
  | NCHAR [<LEFT PAREN><length><RIGHT PAREN>]
  | NATINAL CHARACTER VARYING [<LEFT PAREN><length><RIGHT PAREN>]
  | NATINAL CHAR VARYING [<LEFT PAREN><length><RIGHT PAREN>]
  | NCHAR VARYING [<LEFT PAREN><length><RIGHT PAREN>]
<bit string type> ::=
  BIT [<LEFT PAREN><length><RIGHT PAREN>]
  | BIT VARYING [<LEFT PAREN><length><RIGHT PAREN>]
<numeric type> ::=
  <exact numeric type>
  | <approximate numeric type>
<exact numeric type> ::=
  NUMERIC [<LEFT PAREN><percision>[<COMMA><scale>]<RIGHT PAREN>]
  | DECIMAL [<LEFT PAREN><percision>[<COMMA><scale>]<RIGHT PAREN>]
  | DEC [<LEFT PAREN><percision>[<COMMA><scale>]<RIGHT PAREN>]
  | INTEGER | INT | SMALLINT
<percision> ::= <UNSIGNED INTEGER>
<scale> ::= <UNSIGNED INTEGER>
<approximate numeric type> ::=
  FLOAT [<LEFT PAREN><percision><RIGHT PAREN>]
  | REAL | DOUBLE PERCISION
<datetime type> ::=
  DATE
  | TIME [<LEFT PAREN><time precision><RIGHT PAREN>][WITH TIME ZONE]
  | TIMESTAMP [<LEFT PAREN><timestamp percision><RIGHT PAREN>][WITH TIME ZONE]
<interval type> ::= INTERVAL <interval qualifier>
```
## Queries
```
<scalar subquery> ::= <subquery>
<subquery> ::= <LEFT PAREN> <query expression> <RIGHT PAREN>
<query expression> ::= <non-join query expression> | <joined table>
<non-join query expression> ::=
  <non-join query term>
  |  <query expression> UNION [ALL] [corresponding spec>] <query term>
  |  <query expression> EXCEPT [ALL] [corresponding spec>] <query term>
<non-join query term> ::=
  <non-join query primary>
  | <query term> INTERSECT [ALL] [<coressponding spec>] <query primary>
<non-join query primary> ::= <simple table> | <LEFT PAREN> <non-join query expression> <RIGHT PAREN>
<simple table> ::= <query specification> | <table value constructor> | <explicit table>
<query specification> ::= SELECT [<set quantifier>] <select list> <table expression>
<select list> ::= <ASTERISK> | <select sublist>[ {<COMMA><select sublist>}...]
<select sublist> ::= <derived column> | <quanlifier> <PERIOD> <ASTERISK>
<derived column> ::= <value expression> [<as clause>]
<quanlifier> ::= <table name> | <correlation name>
<correlation name> ::= <IDENTIFIER>
<table name> ::= <qualified name> | <qualified local table name>
<qualified name> ::= [<schema name><PERIOD>]<qualified identifier>
<qualified identifier> ::= <IDENTIFIER>
<qualified local table name> ::= MODULE <PERIOD> <local table name>
<local table name> ::= <qualified identifier>
<set quantifier> ::= DISTINCT | ALL
<explicit table> ::= TABLE <table name>
<query term> ::= <non-join query term> | <joined table>
<corresponding spec> ::= CORRESPONDING [ BY <LEFT PAREN> <corresponding column list> <RIGHT PAREN> ]
<corresponding column list> ::= <column name list>
<column name list> ::= <column name>[ {<COMMA><column name>}...]
<query primary> ::= <non-join query primary> | <joined table>
<joined table> ::=
  <cross join>
  | <qualified join>
  | <LEFT PAREN><joined table><RIGHT PAREN>
<cross join> ::= <table reference> CROSS JOIN <table reference>
<qualified join> ::= <table reference> [NATURAL] [<join type>] JOIN <table reference> [<join specification>]
<join type> ::= INNER | <outer join type> [OUTER] | UNION
<outer join type> ::= LEFT | RIGHT | FULL
<join specification> ::= <join condition> | <named columns join>
<join condition> ::= ON <search condition>
<named columns join> ::= USING <LEFT PAREN> <join column list> <RIGHT PAREN>
<join column list> ::= <column name list>
<table reference> ::=
  <table name> [<correlation specification>]
  | <derived table> <correlation specification>
  | <joined table>
<correlation specification> ::= [AS] <correlation name> [<LEFT PAREN><derived column list><RIGHT PAREN>]
<derived column list> ::= <column name list>
<derived table> ::= <table subquery>
<table subquery> ::= <subquery>
<search condition> ::=
  <boolean term>
  | <search condition> OR <boolean term>
<boolean term> ::=
  <boolean factor>
  | <boolean term> AND <boolean factor>
<boolean factor> ::= [NOT] <boolean test>
<boolean test> ::= <boolean primary> [IS [NOT] <truth value>]
<boolean primary> ::= <predicate> | <LEFT PAREN><search condotion><RIGHT PAREN>
<predicate> ::=
  <comparison predicate>
  | <between predicate>
  | <in predicate>
  | <like predicate>
  | <null predicate>
  | <quantified comparsion predicate>
  | <exists predicate>
  | <match predicate>
  | <overlaps predicate>
<comparison predicate> ::= <row value constructro> <comp op> <row value constructor>
<between predicate> ::= <row value constructor> [NOT] BETWEEN <row value constructor> AND <row value constructor>
<in predicate> ::= <row value constructor> [NOT] IN <in predicate value>
<in predicate value> ::= <table subquery> | <LEFT PAREN> <in value list> <RIGHT PAREN>
<in value list> ::= <value expression>{<COMMA><value expression>}...
<like predicate> ::= <match value> [NOT] LIKE <pattern> [ESCAPE <escape character>]
<match value> ::= <character value expression>
<pattern> ::= <character value expression>
<escape character> ::= <character value expression>
<null predicate> ::= <row value constructor> IS [NOT] NULL
<quantified comparison predicate> ::= <row value constructor> <comp op> <quantifier> <table subquery>
<quantifier> ::= <all>|<some>
<all> ::= ALL
<some> ::= SOME | ANY
<exists predicate> ::= EXISTS <table subquery>
<match predicate> ::= <row value constructor> MATCH [UNIQUE][PARTICAL|FULL] <table subquery>
<overlaps predicate> ::= <row value constructor 1>OVERLAPS <row value constructor 2>
<row value constructor 1> ::= <row value constructor>
<row value constructor 2> ::= <row value constructor>
<truth value> ::= TRUE | FALSE | UNKNOWN
<comp op> ::=
  <EQUALS OPERATOR>
  | <NOT EQUALS OPERATOR>
  | <LESS THAN OPERATOR>
  | <GREATER THAN OPERATOR>
  | <LESS THAN OR EQUALS OPERATOR>
  | <GREATER THAN OR EQUALS OPERATOR>
```
## Expressions
```
<value expression> ::=
  <numeric value expression>
  | <string value expression>
  | <datetime value expression>
  | <interval value expression>
<numeric value expression> ::=
  <term>
  | <numeric value expression> <PLUS SIGN> <term>
  | <numeric value expression> <MINUS SIGN> <term>
<term> ::=
  <factor>
  | <term> <ASTERISK> <factor>
  | <term> <SOLIDAS> <factor>
<factor> ::= [<sign>] <numeric primary>
<numeric primary> ::= <value expression primary> | <numeric value function>
<value expression primary> ::=
  <unsigned value specification>
  | <column reference>
  | <set function specification>
  | <scalar subquery>
  | <case expression>
  | <LEFT PAREN> <value expression> <RIGHT PAREN>
  | <cast specification>
<unsigned value specification> ::= <unsigned literal> | <general value specification>
<genera value specification> ::=
  <parameter specification>
  | <dynamic parameter specification>
  | <variable specification>
  | USER
  | CURRENT USER
  | SESSION USER
  | SYSTEM USER
  | VALUE
<parameter specification> ::= <parameter name>[<indicator parameter>]
<parameter name> ::= <COLON><IDENTIFIER>
<indicator parameter> ::= [INDICATOR]<parameter name>
<dynamic parameter specification> ::= <QUESTION MARK>
<variable specification> ::= <embedded variable name>[ <indicator variable>]
<embedded variable name> ::= <COLON><host identifier>
<indicator variable> ::= [INDICATOR] <embedded variable name>
<column reference> ::= [<qualifier><PERIOD>]<column name>
<qualifier> ::= <table name> | <correlation name>
<set function specification> ::=
  COUNT <LEFT PAREN> <ASTERISK> <RIGHT PAREN>
  | <general set function>
<general set function> ::=
  <set function type> <LEFT PAREN>[<set quantifier>]<value expression><RIGHT PAREN>
<set function type> ::= AVG | MAX | MIN | SUM | COUNT
<case expression> ::= <case abbreviation> | <case specification>
<case abbreviation> ::=
  NULLIF <LEFT PAREN><value expression><COMMA><value expression><RIGHT PAREN>
  | COALESCE <LEFT PAREN><value expression>{<COMMA><value expression>}... <RIGHT PAREN>
<case specification> ::= <simple case> | <searched case>
<simple case> ::= CASE <case operand> <simple when clause>... [<else clause>] END
<case operand> ::= <value expression>
<searched case> ::= CASE <searched when clause>... [<else clause>] END
<cast specification> ::= CAST <LEFT PAREN><cast operand> AS <cast target> <RIGHT PAREN>
<cast operand> ::= <value expression> | NULL
<cast target> ::= <domain name>|<data type>
<domain name> ::= <qualified name>
<string value expression> ::= <character value expression> | <bit value expression>
<character value expression> ::= <concatenation> | <character factor>
<concatenation> ::= <character value expression><CONCATENATION OPERATOR><character factor>
<character factor> ::= <character primary>[ <collate clause>]
<character primary> ::= <value expression primary> | <string value function>
<string value function> ::= <character value function> | <bit value function>
<character value function> ::=
  <character sunstring function>
  | <fold>
  | <form-of-use conversion>
  | <character translation>
  | <trim function>
<character substring function> ::=
  SUBSTRING <LEFT PAREN> <character value expression> FROM <start position> [FOR <string length>] <RIGHT PAREN>
<start position> ::= <numeric value expression>
<string length> ::= <numeric value expression>
<fold> ::= {UPPER|LOWER} <LEFT PAREN><character value expression><RIGHT PAREN>
<form-of-use conversion> ::=
  CONVERT <LEFT PAREN><character value expression>USING <form-of-use conversion name><RIGHT PAREN>
<form-of-use conversion name> ::= <qualified name>
<charavter translation> ::=
  TRANSLATE <LEFT PAREN><character value expression> USING <translation name> <RIGHT PAREN>
<translation name> ::= <qualified name>
<trim function> ::= TRIM <LEFT PAREN> <trim operands> <RIGHT PAREN>
<trim operands> ::= [[<trim specification>][<trim character>] FROM ] <trim source>
<trim specification> ::= LEADING | TRAILING | BOTH
<trim character> ::= <character value expression>
<trim source> ::= <character value expression>
<bit value function> ::= <bit substring function>
<bit substring function> ::=
  SUBSTRING <LEFT PAREN><bit value expression> FROM <start position>  [ FOR <string length> ] <RIGHT PAREN>
<bit value expression> ::= <bit concatenation> | <bit factor>
<bit concatenation> ::= <bit value expression><CONCATENATION OPERATOR><bit factor>
<bit factor> ::= <bit primary>
<bit primary> ::= <value expression primary>|<string value function>
<datetime value expression> ::=
  <datetime term>
  | <interval value expression><PLUS SIGN><datetime term>
  | <datetime value expression><PLUS SIGN><interval term>
  | <datetime value expression><MINUS SIGN><interval term>
<datetime term> ::= <datetime factor>
<datetime factor> ::= <datetime primary>[<time zone>]
<datetime primary> ::= <value expression primary> | <datetime value function>
<datetime value function> ::=
  <current date value function>
  | <current time value function>
  | <current timestamp value function>
<current date value function> ::= CURRENT DATE
<current time value function> ::= CURRENT TIME [<LEFT PAREN><time percision><RIGHT PAREN>]
<current timestamp value function> ::= CURRENT TIMESTAMP [<LEFT PAREN><timestamp percision><RIGHT PAREN>]
<time percision> ::= <time fractional seconds percision>
<time fractional seconds percision> ::= <UNSIGNED INTEGER>
<timestamp percision> ::= <time fractional seconds percision>
<time zone> ::= AT <time zone specifier>
<time zone specifier> ::= LOCAL | TIME ZONE <interval value expression>
<interval value expression> ::=
  <interval term>
  | <interval value expression 1><PLUS SIGN><interval term 1>
  | <interval value expression 1><MINUS SIGN><interval term 1>
  | <LEFT PAREN><datetime value expression><MINUS SIGN><datetime term><RIGHT PAREN><interval qualifier>
<interval term> ::=
  <interval factor>
  | <interval term 2><ASTERISK><factor>
  | <interval term 2><SOLIDUS><factor>
  | <term><ASTERISK><interval factor>
<interval factor> ::= [<sign>]<interval primary>
<interval primary> ::= <value expression primary> [<interval qualifier>]
<interval term 2> ::= <interval term>
<interval value expression 1> ::= <interval value expression>
<interval term 1> ::= <interval term>
<table expression> ::= <from clause> [<where clause>][<group by clause>][<having clause>]
<numeric value function> ::= <position expression> | <extract expression> | <length expression?
<position expression> ::= POSITION <LEFT PAREN><character value expression> IN <character value expression> <RIGHT PAREN>
<extract expression> ::= EXTRACT <LEFT PAREN><extract field> FROM <extract source> <RIGHT PAREN>
<extract field> ::= <datetime field> | <time zone field>
<datetime field> ::= <non-second datetime field> | SECOND
<time zone field> ::= TIMEZONE_HOUR | TIMEZONE_MINUTE
<extract source> ::= <datetime value expression>|<interval value expression>
<length expression> ::= <char length expression> | <octet length expression> | <bit length expression>
<char length expression> ::= {CHAR_LENGTH|CHARACTER_LENGTH} <LEFT PAREN> <string value expression> <RIGHT PAREN>
<octet length expression> ::= OCTET_LENGTH <LEFT PAREN> <string value expression> <RIGHT PAREN>
<bit length expression> ::= BIT_LENGTH <LEFT PAREN> <string value expression> <RIGHT PAREN>
```
## Clauses
```
<as clause> ::= [AS] <column name>
<column name> ::= <IDENTIFIER>
<simple when clause> ::= WHEN <when operand> THEN <result>
<when operand> ::= <value expression>
<result> ::= <value expression> | NULL
<result expression> ::= <value expression>
<else clause> ::= ELSE <result>
<searched when clause> ::= WHEN <search condition> THEN <result>
<collate clause> ::= COLLATE <collation name>
<collation name> ::= <qualifed name>
<from clause> ::= FROM <table reference>[ {<COMMA><table reference>}...]
<where clause> ::= WHERE <search condition>
<group by clause> ::= GROUP BY <grouping column reference list>
<grouping column reference list> ::=
  <grouping column reference>[{<COMMA><grouping column reference>}...]
<grouping column reference> ::= <column reference>[<collate clause>]
<having clause> ::= HAVING <search condition>
<order by clause> ::= ORDER BY <sort specification list>
<sort specification list> ::= <sort specification>[{<COMMA><sort specification>}...]
<sort specification> ::= <sort key>[<collate clause>][<ordering specification>]
<sort key> ::= <column name> | <UNSIGNED INTEGER>
<ordering specification> ::= ASC | DESC
<updatability clause> ::= FOR {READ ONLY|UPDATE [OF <column name list>]}
```
## Constructors
```
<table value constructor> ::= VALUES <table value constructor list>
<table value constructor list> ::= <row value constructor> [{<COMMA><row value constructor>}...]
<row value constructor> ::=
  <row value constructor element>
  | <LEFT PAREN> <row value constructor list> <RIGHT PAREN>
  | <row subquery>
<row value constructor element> ::=
  <value expression>
  | <null specification>
  | <default specification>
<null specification> ::= NULL
<default specification> ::= DEFAULT
<row value constructor list> ::= <row value constructor element>[ {<COMMA><row value constructor element>}...]
<row subquery> ::= <subquery>
```
## Host Identifier
```
<host identifier> ::=
  <Ada host identifier>
  | <C host identifier>
  | <Cobol host identifier>
  | <Fortran host identifier>
  | <MUMPS host identifier>
  | <Pascal host identifier>
  | <PL/I host identifier>
<Ada host identifier> ::= !!!
<C host identifier> ::= !!!
<Cobol host identifier> ::= !!!
<Fortran host identifier> ::= !!!
<MUMPS host identifier> ::= !!!
<Pascal host identifier> ::= !!!
<PL/I host identifier> ::= !!!

```
## SQL
```
<preparable statement> ::= <preparable SQL data statement>
<prepareable SQL data statement> ::=
  <dynamic single row select statement>
  | <dynamic select statement>
<dynamic sigle row select statement> ::= <query specification>
<dynamic select statement> ::= <cursor specification>
<cursor specification> ::= <query expression>[<order by clause>][<updatability clause>]

<direct SQL statement> ::= <direct SQL data statement>
<dirct SQL data statement> ::= <direct select statement: multiple rows>
<direct select statement: multiple rows> ::= <query expression>[<oder by clause>]
```
