# BNF
## Queries by Original
```
<query expression> ::= <query expression body>
<query expression body> ::= <non-join query expression> | <joined table>
<non-join query expression> ::= <non-join query term>
<non-join query term> ::= <non-join query primary>
<non-join query primary> ::= <simple table>
<simple table> ::= <query specification>
<query specification> ::= SELECT [ <set quantifier> ] <select list> <table expression>
<select list> ::= * | <select sublist> [ { , <select sublist> }... ]
<select sublist> ::= <derived column> | <qualified asterisk>
<derived column> ::= <value expression> [ <as clause> ]
<as clause> ::= [ AS ] <column name>
<column name> ::= <IDENTIFIER>
<qualified asterisk> ::=  <asterisked identifier chain> <period> <asterisk> | <<all fields reference>
<asterisked identifier chain> ::= <asterisked identifier> [ { . <asterisked identifier> }... ]
<asterisked identifier> ::= <IDENTIFIER>
<all fields reference> ::= <value expression primary> . *
<value expression primary> ::= <parenthesized value expression> | <nonparenthesized value expression primary>
<parenthesized value expression> ::= ( <value expression> )
<nonparenthesized value expression primary> ::= <unsigned value specification> | <column reference> | <set function specification>
<unsigned value specification> ::= <unsigned literal> | <general value specification>
<unsigned literal> ::= <unsigned numeric literal> | <general literal>
<unsigned numeric literal> ::= <exact numeric literal> | <approximate numeric literal>
<approximate numeric literal> ::= <mantissa> E <exponent>
<mantissa> ::= <exact numeric literal>
<exponent> ::= <signed integer>
<signed integer> ::= [ <sign> ] <unsigned integer>
<exact numeric literal> ::=  <unsigned integer> [ . [ <unsigned integer> ] ] | . <unsigned integer>
<unsigned integer> ::= <digit> ...
<digit> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
<column reference> ::= <basic identifier chain>
<basic identifier chain> ::= <identifier chain>
<identifier chain> ::= <IDENTIFIER> [ { . <IDENTIFIER> }... ]
<set function specification> ::= COUNT ( * ) | <general set function>
<general set function> ::= <set function type> ( [ <set quantifier> ] <value expression> )
<set function type> ::= <computational operation>
<computational operation> ::= AVG | MAX | MIN | SUM | EVERY | ANY | SOME | COUNT
<set quantifier> ::= DISTINCT | ALL
<value expression> ::= <numeric value expression> | <string value expression> |  <boolean value expression> | <row value expression> | <reference value expression> | <collection value expression>
<numeric value expression> ::= <term> | <numeric value expression> + <term> | <numeric value expression> - <term>
<term> ::=  <factor> | <term> * <factor> | <term> / <factor>
<factor> ::= [ <sign> ] <numeric primary>
<numeric primary> ::= <value expression primary> | <numeric value function>
<sign> ::= + | -


<numeric value function>    ::=
```

## Query
```
<SQL> ::= <SELECT Statement><FROM Statement>[<WHERE Statement>][<GROUP BY Statement>][<HAVING Statement>][<ORDER BY Statement>][<LIMIT Statement>]
<SELECT Statement> ::= SELECT <Column List>
<Column List> ::= * | <Column Detail>+
```
