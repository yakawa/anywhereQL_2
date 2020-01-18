// Copyrights (C) 2019-2020 KAWASAKI Yasukazu

package token

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) Debug() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %s}\n", t.Type, t.Literal)
}

const (
	ILLIEGAL = "ILLIEGAL"
	EOF      = "EOF"

	// Data Type
	IDENTIFIER = "IDENTIFIER"
	COMMENT    = "COMMENT"

	// Characters
	SPACE_CHAR                      = " "
	DOUBLE_QUOTE                    = "\""
	PERCENT                         = "%"
	AMPERSAND                       = "&"
	QUOTE                           = "'"
	LEFT_PAREN                      = "("
	RIGHT_PAREN                     = ")"
	ASTERISK                        = "*"
	PLUS_SIGN                       = "+"
	COMMA                           = ","
	MINUS_SIGN                      = "-"
	PERIOD                          = "."
	SOLIDAS                         = "/"
	COLON                           = ":"
	SEMICOLON                       = ";"
	LESS_THAN_OPERATOR              = "<"
	GREATER_THAN_OPERATOR           = ">"
	EQUALS_OPERATOR                 = "="
	QUESTION_MARK                   = "?"
	UNDERSCORE                      = "_"
	VERTICAL_BAR                    = "|"
	LEFT_BRACKET                    = "["
	RIGHT_BRACKET                   = "]"
	NEWLINE                         = "NEWLINE"
	NOT_EQUALS_OPERATOR             = "<>"
	GREATER_THAN_OR_EQUALS_OPERATOR = ">="
	LESS_THAN_OR_EQUALS_OPERATOR    = "<="
	CONCATENATION_OPERATOR          = "||"
	DOUBLE_PERIOD                   = ".."

	// SQL Reserved Key Words
	ABSOLUTE          = "ABSOLUTE"
	ACTION            = "ACTION"
	ADD               = "ADD"
	ALL               = "ALL"
	ALLOCATE          = "ALLOCATE"
	ALTER             = "ALTER"
	AND               = "AND"
	ANY               = "ANY"
	ARE               = "ARE"
	AS                = "AS"
	ASC               = "ASC"
	ASSERTION         = "ASSERSION"
	AT                = "AT"
	AUTHORIZATION     = "AUTHORIZATION"
	AVG               = "AVG"
	BEGIN             = "BEGIN"
	BETWEEN           = "BETWEEN"
	BIT               = "BIT"
	BIT_LENGTH        = "BIT_LENGTH"
	BOTH              = "BOTH"
	BY                = "BY"
	CASCADE           = "CASCATE"
	CASCADED          = "CASCADED"
	CASE              = "CASE"
	CAST              = "CAST"
	CATALOG           = "CATALOG"
	CHAR              = "CHAR"
	CHARACTER         = "CHARACTER"
	CHARACTER_LENGTH  = "CHARACTER_LENGTH"
	CHAR_LENGTH       = "CHAR_LENGTH"
	CHECK             = "CHECK"
	CLOSE             = "CLOSE"
	COALESCE          = "COALESCE"
	COLLATE           = "COLLATE"
	COLLATION         = "COLLATION"
	COLUMN            = "COLUMN"
	COMMIT            = "COMMIT"
	CONNECT           = "CONNECT"
	CONNECTION        = "CONNECTION"
	CONSTRAINT        = "CONSTRAINT"
	CONSTRAINTS       = "CONSTRAINTS"
	CONTINUE          = "CONTINUE"
	CONVERT           = "CONVERT"
	CORRESPONDING     = "CORRESPONDING"
	CREATE            = "CREATE"
	CROSS             = "CROSS"
	CURRENT           = "CURRENT"
	CURRENT_DATE      = "CURRENT_DATE"
	CURRENT_TIME      = "CURRENT_TIME"
	CURRENT_TIMESTAMP = "CURRENT_TIMESTAMP"
	CURRENT_USER      = "CURRENT_USER"
	CURSOR            = "CURSOR"
	DATE              = "DATE"
	DAY               = "DAY"
	DEALLOCATE        = "DEALLOCATE"
	DEC               = "DEC"
	DECIMAL           = "DECIMAL"
	DECLARE           = "DECLARE"
	DEFAULT           = "DEFAULT"
	DEFERRABLE        = "DEFERRABLE"
	DEFERRED          = "DEFERRED"
	DELETE            = "DELETE"
	DESC              = "DESC"
	DESCRIBE          = "DESCRIBE"
	DESCRIPTOR        = "DESCRIPTOR"
	DIAGNOSTICS       = "DIAGNOSTICS"
	DISCONNECT        = "DISCONNECT"
	DISTINCT          = "DISTINCT"
	DOMAIN            = "DOMAIN"
	DOUBLE            = "DOUBLE"
	DROP              = "DROP"
	ELSE              = "ELSE"
	END               = "END"
	END_EXEC          = "END-EXEC"
	ESCAPE            = "ESCAPE"
	EXCEPT            = "EXCEPT"
	EXCEPTION         = "EXCEPTION"
	EXEC              = "EXEC"
	EXECUTE           = "EXECUTE"
	EXISTS            = "EXISTS"
	EXTERNAL          = "EXTERNAL"
	EXTRACT           = "EXTRACT"
	FALSE             = "FALSE"
	FETCH             = "FETCH"
	FIRST             = "FIRST"
	FLOAT             = "FLOAT"
	FOR               = "FOR"
	FOREIGN           = "FOREGIN"
	FOUND             = "FOUND"
	FROM              = "FROM"
	FULL              = "FULL"
	GET               = "GET"
	GLOBAL            = "GLOBAL"
	GO                = "GO"
	GOTO              = "GOTO"
	GRANT             = "GRANT"
	GROUP             = "GROUP"
	HAVING            = "HAVING"
	HOUR              = "HOUR"
	IDENTITY          = "IDENTITY"
	IMMEDIATE         = "IMMEDIATE"
	IN                = "IN"
	INDICATOR         = "INDICATOR"
	INITIALLY         = "INITIALLY"
	INNER             = "INNER"
	INPUT             = "INPUT"
	INSENSITIVE       = "INSENSITIVE"
	INSERT            = "INSERT"
	INT               = "INT"
	INTEGER           = "INTEGER"
	INTERSECT         = "INTERSECT"
	INTERVAL          = "INTERVAL"
	INTO              = "INTO"
	IS                = "IS"
	ISOLATION         = "ISOLATE"
	JOIN              = "JOIN"
	KEY               = "KEY"
	LANGUAGE          = "LANGUAGE"
	LAST              = "LATT"
	LEADING           = "LEADING"
	LEFT              = "LEFT"
	LEVEL             = "LEVEL"
	LIKE              = "LIKE"
	LOCAL             = "LOCAL"
	LOWER             = "LOWER"
	MATCH             = "MATCH"
	MAX               = "MAX"
	MIN               = "MIN"
	MINUTE            = "MINUTE"
	MODULE            = "MODULE"
	MONTH             = "MONTH"
	NAMES             = "NAMES"
	NATIONAL          = "NATINAL"
	NATURAL           = "NATURAL"
	NCHAR             = "NCAR"
	NEXT              = "NEXT"
	NO                = "NO"
	NOT               = "NOT"
	NULL              = "NULL"
	NULLIF            = "NULLIF"
	NUMERIC           = "NUMERIC"
	OCTET_LENGTH      = "OCTET_LENGTH"
	OF                = "OF"
	ON                = "ON"
	ONLY              = "ONLY"
	OPEN              = "OPEN"
	OPTION            = "OPTION"
	OR                = "OR"
	ORDER             = "ORDER"
	OUTER             = "OUTER"
	OUTPUT            = "OUTPUT"
	OVERLAPS          = "OVERLAPS"
	PAD               = "PAD"
	PARTIAL           = "PARTIAL"
	POSITION          = "POSITION"
	PRECISION         = "PRECISION"
	PREPARE           = "PREPARE"
	PRESERVE          = "PRESERVE"
	PRIMARY           = "PRIMARY"
	PRIOR             = "PRIOR"
	PRIVILEGES        = "PRIVILEGES"
	PROCEDURE         = "PROCEDURE"
	PUBLIC            = "PUBLIC"
	READ              = "READ"
	REAL              = "REAL"
	REFERENCES        = "REFERENCES"
	RELATIVE          = "RELATIVE"
	RESTRICT          = "PERSRICT"
	REVOKE            = "REVOKE"
	RIGHT             = "RIGHT"
	ROLLBACK          = "ROLLBACK"
	ROWS              = "ROWS"
	SCHEMA            = "SChEMA"
	SCROLL            = "SCROLL"
	SECOND            = "SECOND"
	SECTION           = "SECTION"
	SELECT            = "SELECT"
	SESSION           = "SESSION"
	SESSION_USER      = "SESSION_USER"
	SET               = "SET"
	SIZE              = "SIZE"
	SMALLINT          = "SMALLINT"
	SOME              = "SOME"
	SPACE             = "SPACE"
	SQL               = "SQL"
	SQLCODE           = "SQLCODE"
	SQLERROR          = "SQLERRROR"
	SQLSTATE          = "SQLSTATE"
	SUBSTRING         = "SUBSTRING"
	SUM               = "SUM"
	SYSTEM_USER       = "SYSTEM_USER"
	TABLE             = "TABLE"
	TEMPORARY         = "TEMPORARY"
	THEN              = "THEN"
	TIME              = "TIME"
	TIMESTAMP         = "TIMESTAMP"
	TIMEZONE_HOUR     = "TIMEZONE_HOUR"
	TIMEZONE_MINUTE   = "TIMEZONE_MINUTE"
	TO                = "TO"
	TRAILING          = "TRAILING"
	TRANSACTION       = "TRANSACTION"
	TRANSLATE         = "TRANSLATE"
	TRANSLATION       = "TRANSLATION"
	TRIM              = "TRIM"
	TRUE              = "TRUE"
	UNION             = "UNION"
	UNIQUE            = "UNIQUE"
	UNKNOWN           = "UNKNOWN"
	UPDATE            = "UPDATE"
	UPPER             = "UPPER"
	USAGE             = "USAGE"
	USER              = "USER"
	USING             = "USING"
	VALUE             = "VALUE"
	VALUES            = "VALUES"
	VARCHAR           = "VARCHAR"
	VARYING           = "VARYING"
	VIEW              = "VIEW"
	WHEN              = "WHEN"
	WHENEVER          = "WHENEVER"
	WHERE             = "WHERE"
	WITH              = "WITH"
	WORK              = "WORK"
	WRITE             = "WRITE"
	YEAR              = "YEAR"
	ZONE              = "ZONE"

	// SQL Reserved Key Words
	ADA                         = "ADA"
	C                           = "C"
	CATALOG_NAME                = "CATALOG_NAME"
	CHARACTER_SET_CATALOG       = "CHARACTER_SET_CATALOG"
	CHARACTER_SET_NAME          = "CHARACTER_SET_NAME"
	CHARACTER_SET_SCHEMA        = "CHARACTER_SET_SCHEMA"
	CLASS_ORIGIN                = "CLASS_ORIGIN"
	COBOL                       = "COBOL"
	COLLATION_CATALOG           = "COLLATION_CATALOG"
	COLLATION_NAME              = "COLLATION_NAME"
	COLLATION_SCHEMA            = "COLLATION_SCHEMA"
	COLUMN_NAME                 = "COLUMN_NAME"
	COMMAND_FUNCTION            = "COMMAND_FUNCTION"
	COMMITTED                   = "COMMITTED"
	CONDITION_NUMBER            = "CONDITION_NUMBER"
	CONNECTION_NAME             = "CONNECTION_NAME"
	CONSTRAINT_CATALOG          = "CONSTRAINT_CATALOG"
	CONSTRAINT_NAME             = "CONSTRAINT_NAME"
	CONSTRAINT_SCHEMA           = "CONSTRAINT_SCHEMA"
	CURSOR_NAME                 = "CURSOR_NAME"
	DATA                        = "DATA"
	DATETIME_INTERVAL_CODE      = "DATETIME_INTERVAL_CODE"
	DATETIME_INTERVAL_PRECISION = "DATETiME_INTERVAL_PRECISION"
	DYNAMIC_FUNCTION            = "DYNAMIC_FUNCTION"
	FORTRAN                     = "FORTRAN"
	LENGTH                      = "LENGTH"
	MESSAGE_LENGTH              = "MESSAGE_LENGTH"
	MESSAGE_OCTET_LENGTH        = "MESSAGE_OCTET_LENGTH"
	MESSAGE_TEXT                = "MESSAGE_TEXT"
	MORE                        = "MORE"
	MUMPS                       = "MUMPS"
	NAME                        = "NAME"
	NULLABLE                    = "NULLABLE"
	NUMBER                      = "NUMBER"
	PASCAL                      = "PASCAL"
	PLI                         = "PLI"
	REPEATABLE                  = "REPEATABLE"
	RETURNED_LENGTH             = "RETURNED_LENGTH"
	RETURNED_OCTET_LENGTH       = "RETURNED_OCTET_LENGTH"
	RETURNED_SQLSTATE           = "RETURNED_SQLSTATE"
	ROW_COUNT                   = "ROW_COUNT"
	SCALE                       = "SCALE"
	SCHEMA_NAME                 = "SCHEMA_NAME"
	SERIALIZABLE                = "SERIALIZABLE"
	SERVER_NAME                 = "SERVER_NAME"
	SUBCLASS_ORIGIN             = "SUBCLASS_ORIGIN"
	TABLE_NAME                  = "TABLE_NAME"
	TYPE                        = "TYPE"
	UNCOMMITTED                 = "UNCOMMITTED"
	UNNAMED                     = "UNNAMED"
)

var reservedKeyWord = map[string]TokenType{
	"ABSOLUTE":          ABSOLUTE,
	"ACTION":            ACTION,
	"ADD":               ADD,
	"ALL":               ALL,
	"ALLOCATE":          ALLOCATE,
	"ALTER":             ALTER,
	"AND":               AND,
	"ANY":               ANY,
	"ARE":               ARE,
	"AS":                AS,
	"ASC":               ASC,
	"ASSERTION":         ASSERTION,
	"AT":                AT,
	"AUTHORIZATION":     AUTHORIZATION,
	"AVG":               AVG,
	"BEGIN":             BEGIN,
	"BETWEEN":           BETWEEN,
	"BIT":               BIT,
	"BIT_LENGTH":        BIT_LENGTH,
	"BOTH":              BOTH,
	"BY":                BY,
	"CASCADE":           CASCADE,
	"CASCADED":          CASCADED,
	"CASE":              CASE,
	"CAST":              CAST,
	"CATALOG":           CATALOG,
	"CHAR":              CHAR,
	"CHARACTER":         CHARACTER,
	"CHARACTER_LENGTH":  CHARACTER_LENGTH,
	"CHAR_LENGTH":       CHAR_LENGTH,
	"CHECK":             CHECK,
	"CLOSE":             CLOSE,
	"COALESCE":          COALESCE,
	"COLLATE":           COLLATE,
	"COLLATION":         COLLATION,
	"COLUMN":            COLUMN,
	"COMMIT":            COMMIT,
	"CONNECT":           CONNECT,
	"CONNECTION":        CONNECTION,
	"CONSTRAINT":        CONSTRAINT,
	"CONSTRAINTS":       CONSTRAINTS,
	"CONTINUE":          CONTINUE,
	"CONVERT":           CONVERT,
	"CORRESPONDING":     CORRESPONDING,
	"CREATE":            CREATE,
	"CROSS":             CROSS,
	"CURRENT":           CURRENT,
	"CURRENT_DATE":      CURRENT_DATE,
	"CURRENT_TIME":      CURRENT_TIME,
	"CURRENT_TIMESTAMP": CURRENT_TIMESTAMP,
	"CURRENT_USER":      CURRENT_USER,
	"CURSOR":            CURSOR,
	"DATE":              DATE,
	"DAY":               DAY,
	"DEALLOCATE":        DEALLOCATE,
	"DEC":               DEC,
	"DECIMAL":           DECIMAL,
	"DECLARE":           DECLARE,
	"DEFAULT":           DEFAULT,
	"DEFERRABLE":        DEFERRABLE,
	"DEFERRED":          DEFERRED,
	"DELETE":            DELETE,
	"DESC":              DESC,
	"DESCRIBE":          DESCRIBE,
	"DESCRIPTOR":        DESCRIPTOR,
	"DIAGNOSTICS":       DIAGNOSTICS,
	"DISCONNECT":        DISCONNECT,
	"DISTINCT":          DISTINCT,
	"DOMAIN":            DOMAIN,
	"DOUBLE":            DOUBLE,
	"DROP":              DROP,
	"ELSE":              ELSE,
	"END":               END,
	"END_EXEC":          END_EXEC,
	"ESCAPE":            ESCAPE,
	"EXCEPT":            EXCEPT,
	"EXCEPTION":         EXCEPTION,
	"EXEC":              EXEC,
	"EXECUTE":           EXECUTE,
	"EXISTS":            EXISTS,
	"EXTERNAL":          EXTERNAL,
	"EXTRACT":           EXTRACT,
	"FALSE":             FALSE,
	"FETCH":             FETCH,
	"FIRST":             FIRST,
	"FLOAT":             FLOAT,
	"FOR":               FOR,
	"FOREIGN":           FOREIGN,
	"FOUND":             FOUND,
	"FROM":              FROM,
	"FULL":              FULL,
	"GET":               GET,
	"GLOBAL":            GLOBAL,
	"GO":                GO,
	"GOTO":              GOTO,
	"GRANT":             GRANT,
	"GROUP":             GROUP,
	"HAVING":            HAVING,
	"HOUR":              HOUR,
	"IDENTITY":          IDENTITY,
	"IMMEDIATE":         IMMEDIATE,
	"IN":                IN,
	"INDICATOR":         INDICATOR,
	"INITIALLY":         INITIALLY,
	"INNER":             INNER,
	"INPUT":             INPUT,
	"INSENSITIVE":       INSENSITIVE,
	"INSERT":            INSERT,
	"INT":               INT,
	"INTEGER":           INTEGER,
	"INTERSECT":         INTERSECT,
	"INTERVAL":          INTERVAL,
	"INTO":              INTO,
	"IS":                IS,
	"ISOLATION":         ISOLATION,
	"JOIN":              JOIN,
	"KEY":               KEY,
	"LANGUAGE":          LANGUAGE,
	"LAST":              LAST,
	"LEADING":           LEADING,
	"LEFT":              LEFT,
	"LEVEL":             LEVEL,
	"LIKE":              LIKE,
	"LOCAL":             LOCAL,
	"LOWER":             LOWER,
	"MATCH":             MATCH,
	"MAX":               MAX,
	"MIN":               MIN,
	"MINUTE":            MINUTE,
	"MODULE":            MODULE,
	"MONTH":             MONTH,
	"NAMES":             NAMES,
	"NATIONAL":          NATIONAL,
	"NATURAL":           NATURAL,
	"NCHAR":             NCHAR,
	"NEXT":              NEXT,
	"NO":                NO,
	"NOT":               NOT,
	"NULL":              NULL,
	"NULLIF":            NULLIF,
	"NUMERIC":           NUMERIC,
	"OCTET_LENGTH":      OCTET_LENGTH,
	"OF":                OF,
	"ON":                ON,
	"ONLY":              ONLY,
	"OPEN":              OPEN,
	"OPTION":            OPTION,
	"OR":                OR,
	"ORDER":             ORDER,
	"OUTER":             OUTER,
	"OUTPUT":            OUTPUT,
	"OVERLAPS":          OVERLAPS,
	"PAD":               PAD,
	"PARTIAL":           PARTIAL,
	"POSITION":          POSITION,
	"PRECISION":         PRECISION,
	"PREPARE":           PREPARE,
	"PRESERVE":          PRESERVE,
	"PRIMARY":           PRIMARY,
	"PRIOR":             PRIOR,
	"PRIVILEGES":        PRIVILEGES,
	"PROCEDURE":         PROCEDURE,
	"PUBLIC":            PUBLIC,
	"READ":              READ,
	"REAL":              REAL,
	"REFERENCES":        REFERENCES,
	"RELATIVE":          RELATIVE,
	"RESTRICT":          RESTRICT,
	"REVOKE":            REVOKE,
	"RIGHT":             RIGHT,
	"ROLLBACK":          ROLLBACK,
	"ROWS":              ROWS,
	"SCHEMA":            SCHEMA,
	"SCROLL":            SCROLL,
	"SECOND":            SECOND,
	"SECTION":           SECTION,
	"SELECT":            SELECT,
	"SESSION":           SESSION,
	"SESSION_USER":      SESSION_USER,
	"SET":               SET,
	"SIZE":              SIZE,
	"SMALLINT":          SMALLINT,
	"SOME":              SOME,
	"SPACE":             SPACE,
	"SQL":               SQL,
	"SQLCODE":           SQLCODE,
	"SQLERROR":          SQLERROR,
	"SQLSTATE":          SQLSTATE,
	"SUBSTRING":         SUBSTRING,
	"SUM":               SUM,
	"SYSTEM_USER":       SYSTEM_USER,
	"TABLE":             TABLE,
	"TEMPORARY":         TEMPORARY,
	"THEN":              THEN,
	"TIME":              TIME,
	"TIMESTAMP":         TIMESTAMP,
	"TIMEZONE_HOUR":     TIMEZONE_HOUR,
	"TIMEZONE_MINUTE":   TIMEZONE_MINUTE,
	"TO":                TO,
	"TRAILING":          TRAILING,
	"TRANSACTION":       TRANSACTION,
	"TRANSLATE":         TRANSLATE,
	"TRANSLATION":       TRANSLATION,
	"TRIM":              TRIM,
	"TRUE":              TRUE,
	"UNION":             UNION,
	"UNIQUE":            UNIQUE,
	"UNKNOWN":           UNKNOWN,
	"UPDATE":            UPDATE,
	"UPPER":             UPPER,
	"USAGE":             USAGE,
	"USER":              USER,
	"USING":             USING,
	"VALUE":             VALUE,
	"VALUES":            VALUES,
	"VARCHAR":           VARCHAR,
	"VARYING":           VARYING,
	"VIEW":              VIEW,
	"WHEN":              WHEN,
	"WHENEVER":          WHENEVER,
	"WHERE":             WHERE,
	"WITH":              WITH,
	"WORK":              WORK,
	"WRITE":             WRITE,
	"YEAR":              YEAR,
	"ZONE":              ZONE,
}

var nonreservedKeyWord = map[string]TokenType{
	"ADA":                         ADA,
	"C":                           C,
	"CATALOG_NAME":                CATALOG_NAME,
	"CHARACTER_SET_CATALOG":       CHARACTER_SET_CATALOG,
	"CHARACTER_SET_NAME":          CHARACTER_SET_NAME,
	"CHARACTER_SET_SCHEMA":        CHARACTER_SET_SCHEMA,
	"CLASS_ORIGIN":                CLASS_ORIGIN,
	"COBOL":                       COBOL,
	"COLLATION_CATALOG":           COLLATION_CATALOG,
	"COLLATION_NAME":              COLLATION_NAME,
	"COLLATION_SCHEMA":            COLLATION_SCHEMA,
	"COLUMN_NAME":                 COLUMN_NAME,
	"COMMAND_FUNCTION":            COMMAND_FUNCTION,
	"COMMITTED":                   COMMITTED,
	"CONDITION_NUMBER":            CONDITION_NUMBER,
	"CONNECTION_NAME":             CONNECTION_NAME,
	"CONSTRAINT_CATALOG":          CONSTRAINT_CATALOG,
	"CONSTRAINT_NAME":             CONSTRAINT_NAME,
	"CONSTRAINT_SCHEMA":           CONSTRAINT_SCHEMA,
	"CURSOR_NAME":                 CURSOR_NAME,
	"DATA":                        DATA,
	"DATETIME_INTERVAL_CODE":      DATETIME_INTERVAL_CODE,
	"DATETIME_INTERVAL_PRECISION": DATETIME_INTERVAL_PRECISION,
	"DYNAMIC_FUNCTION":            DYNAMIC_FUNCTION,
	"FORTRAN":                     FORTRAN,
	"LENGTH":                      LENGTH,
	"MESSAGE_LENGTH":              MESSAGE_LENGTH,
	"MESSAGE_OCTET_LENGTH":        MESSAGE_OCTET_LENGTH,
	"MESSAGE_TEXT":                MESSAGE_TEXT,
	"MORE":                        MORE,
	"MUMPS":                       MUMPS,
	"NAME":                        NAME,
	"NULLABLE":                    NULLABLE,
	"NUMBER":                      NUMBER,
	"PASCAL":                      PASCAL,
	"PLI":                         PLI,
	"REPEATABLE":                  REPEATABLE,
	"RETURNED_LENGTH":             RETURNED_LENGTH,
	"RETURNED_OCTET_LENGTH":       RETURNED_OCTET_LENGTH,
	"RETURNED_SQLSTATE":           RETURNED_SQLSTATE,
	"ROW_COUNT":                   ROW_COUNT,
	"SCALE":                       SCALE,
	"SCHEMA_NAME":                 SCHEMA_NAME,
	"SERIALIZABLE":                SERIALIZABLE,
	"SERVER_NAME":                 SERVER_NAME,
	"SUBCLASS_ORIGIN":             SUBCLASS_ORIGIN,
	"TABLE_NAME":                  TABLE_NAME,
	"TYPE":                        TYPE,
	"UNCOMMITTED":                 UNCOMMITTED,
	"UNNAMED":                     UNNAMED,
}

func LookupKeyWord(w string) TokenType {
	if t, ok := reservedKeyWord[w]; ok {
		return t
	}
	if t, ok := nonreservedKeyWord[w]; ok {
		return t
	}
	return IDENTIFIER
}
