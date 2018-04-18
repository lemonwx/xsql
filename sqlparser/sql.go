//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const OFF = 57401
const AND = 57402
const OR = 57403
const NOT = 57404
const UNARY = 57405
const CASE = 57406
const WHEN = 57407
const THEN = 57408
const ELSE = 57409
const END = 57410
const BEGIN = 57411
const COMMIT = 57412
const ROLLBACK = 57413
const AUTOCOMMIT = 57414
const NAMES = 57415
const REPLACE = 57416
const ADMIN = 57417
const SHOW = 57418
const DATABASES = 57419
const TABLES = 57420
const PROXY = 57421
const VARIABLES = 57422
const STATUS = 57423
const CREATE = 57424
const ALTER = 57425
const DROP = 57426
const RENAME = 57427
const TABLE = 57428
const INDEX = 57429
const VIEW = 57430
const TO = 57431
const IGNORE = 57432
const IF = 57433
const UNIQUE = 57434
const USING = 57435
const DATABASE = 57436

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OFF",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"AUTOCOMMIT",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"VARIABLES",
	"STATUS",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"DATABASE",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 222
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 710

var yyAct = [...]int{

	120, 169, 162, 403, 76, 291, 373, 203, 117, 128,
	331, 245, 148, 282, 98, 118, 218, 105, 284, 78,
	204, 3, 123, 46, 213, 48, 107, 127, 412, 49,
	133, 45, 412, 252, 65, 178, 179, 110, 124, 125,
	126, 35, 36, 37, 38, 80, 115, 412, 228, 86,
	131, 111, 88, 79, 91, 67, 171, 93, 83, 305,
	306, 307, 308, 309, 103, 310, 311, 337, 171, 171,
	114, 243, 112, 243, 129, 130, 108, 298, 151, 274,
	101, 134, 55, 56, 57, 414, 147, 384, 383, 413,
	54, 355, 357, 365, 51, 156, 52, 382, 85, 158,
	160, 150, 276, 163, 411, 165, 283, 168, 132, 87,
	175, 106, 167, 364, 53, 359, 164, 234, 205, 81,
	178, 179, 206, 316, 336, 326, 324, 283, 275, 329,
	242, 66, 73, 209, 356, 367, 232, 212, 80, 235,
	177, 80, 216, 143, 223, 221, 79, 138, 145, 79,
	140, 161, 61, 62, 63, 59, 60, 224, 264, 226,
	220, 332, 168, 191, 192, 193, 201, 202, 381, 238,
	254, 379, 77, 332, 112, 251, 294, 253, 223, 155,
	249, 241, 239, 257, 166, 380, 262, 263, 353, 266,
	267, 268, 269, 270, 271, 272, 273, 258, 250, 352,
	351, 265, 231, 233, 230, 136, 178, 179, 139, 349,
	112, 112, 347, 140, 350, 80, 80, 348, 90, 287,
	243, 389, 369, 79, 289, 278, 280, 142, 157, 295,
	255, 256, 82, 290, 286, 189, 190, 191, 192, 193,
	399, 80, 296, 398, 219, 300, 301, 219, 397, 79,
	35, 36, 37, 38, 318, 159, 170, 249, 299, 315,
	286, 302, 172, 320, 321, 210, 305, 306, 307, 308,
	309, 293, 310, 311, 92, 319, 94, 240, 18, 84,
	248, 214, 112, 18, 19, 20, 21, 303, 215, 247,
	140, 135, 215, 339, 208, 328, 338, 325, 335, 207,
	171, 334, 173, 104, 66, 32, 314, 176, 248, 409,
	81, 22, 360, 249, 249, 345, 346, 247, 358, 342,
	361, 341, 313, 66, 237, 363, 236, 410, 39, 217,
	74, 154, 366, 152, 149, 330, 146, 144, 80, 141,
	371, 89, 64, 374, 137, 387, 370, 368, 41, 42,
	43, 44, 95, 72, 375, 259, 18, 260, 261, 58,
	323, 97, 27, 28, 29, 385, 96, 30, 33, 31,
	386, 99, 68, 416, 225, 23, 24, 26, 25, 285,
	153, 168, 393, 70, 100, 388, 395, 172, 378, 340,
	292, 377, 374, 344, 219, 402, 401, 102, 404, 404,
	404, 80, 405, 406, 279, 407, 123, 75, 415, 79,
	400, 127, 18, 417, 133, 40, 17, 418, 16, 419,
	123, 110, 124, 125, 126, 127, 15, 14, 133, 394,
	115, 396, 13, 12, 131, 110, 124, 125, 126, 227,
	47, 297, 229, 50, 115, 288, 408, 390, 131, 372,
	376, 343, 327, 211, 114, 18, 281, 122, 129, 130,
	108, 119, 121, 333, 116, 134, 180, 113, 114, 354,
	123, 246, 129, 130, 108, 127, 304, 244, 133, 134,
	109, 312, 174, 69, 34, 81, 124, 125, 126, 71,
	11, 10, 132, 9, 115, 8, 277, 7, 131, 6,
	5, 4, 2, 1, 0, 0, 132, 0, 18, 123,
	0, 0, 0, 0, 127, 0, 0, 133, 114, 0,
	0, 0, 129, 130, 81, 124, 125, 126, 127, 134,
	0, 133, 0, 115, 0, 0, 0, 131, 81, 124,
	125, 126, 0, 0, 0, 0, 0, 159, 0, 0,
	0, 131, 0, 0, 0, 0, 132, 114, 0, 0,
	127, 129, 130, 133, 0, 0, 0, 0, 134, 0,
	81, 124, 125, 126, 127, 129, 130, 133, 0, 159,
	0, 0, 134, 131, 81, 124, 125, 126, 0, 0,
	0, 0, 0, 159, 0, 132, 0, 131, 0, 222,
	391, 392, 0, 0, 0, 0, 0, 129, 130, 132,
	181, 185, 183, 184, 134, 0, 0, 0, 0, 0,
	0, 129, 130, 0, 0, 0, 0, 0, 134, 197,
	198, 199, 200, 0, 194, 195, 196, 0, 0, 0,
	0, 132, 0, 186, 187, 188, 189, 190, 191, 192,
	193, 0, 0, 0, 0, 132, 0, 182, 186, 187,
	188, 189, 190, 191, 192, 193, 362, 0, 0, 186,
	187, 188, 189, 190, 191, 192, 193, 322, 317, 0,
	186, 187, 188, 189, 190, 191, 192, 193, 0, 0,
	0, 0, 0, 0, 186, 187, 188, 189, 190, 191,
	192, 193, 186, 187, 188, 189, 190, 191, 192, 193,
}
var yyPact = [...]int{

	278, -1000, -1000, 201, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -78, -9, 13, -19, -1000, -1000, -1000,
	-1000, 60, 307, 269, 407, 355, -1000, -1000, -1000, 365,
	-1000, 324, 295, 398, 84, -48, -48, -4, 269, -1000,
	8, 269, -1000, 306, -52, -52, 269, -52, 323, 343,
	338, 361, 388, 269, -1000, 259, -1000, -1000, -1000, 2,
	-1000, 252, 295, 311, 70, 295, 160, 304, -1000, 182,
	-1000, 66, 302, 80, 301, 269, -1000, 299, -1000, -26,
	298, 360, 296, 115, 269, 295, 549, 549, -1000, 489,
	549, 361, 549, 388, 549, 247, 258, -1000, -1000, 288,
	63, 140, 589, -1000, 489, 450, -1000, -1000, -1000, 549,
	255, 250, -1000, 221, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 549, -1000, 248, 275, 294, 384,
	275, -1000, 535, 269, -1000, 354, -1000, -60, -1000, 104,
	-1000, 291, -1000, -1000, -1000, 289, -1000, 244, 633, 503,
	633, 140, 589, 633, -1000, 633, 361, 20, 633, -1000,
	245, 400, 549, -77, -1000, -1000, 269, 96, 489, 489,
	549, 211, 334, 549, 549, 133, 549, 549, 549, 549,
	549, 549, 549, 549, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -31, 18, -8, 589, -1000, 386, 400, -1000,
	407, 26, 633, 351, 275, 275, 237, -1000, 377, 489,
	-1000, 633, -1000, -1000, -1000, -1000, -1000, 112, 269, -1000,
	-27, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 351,
	275, -1000, -1000, 549, 234, 212, 287, 273, 46, -1000,
	-1000, 625, 372, -1000, -1000, -1000, -1000, 633, -1000, 211,
	549, 549, 633, 611, -1000, 335, 163, 163, 163, 89,
	89, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 16, 400,
	15, 47, -1000, 489, 109, 211, 201, 97, 14, -1000,
	377, 372, 375, 140, 286, -1000, -1000, 284, -1000, -1000,
	160, 633, 382, 245, 245, -1000, -1000, 158, 155, 146,
	145, 134, 29, -1000, 283, 5, 277, 549, -1000, -1000,
	633, 600, 549, -1000, -1000, 3, -1000, 10, -1000, 549,
	54, -1000, 317, 169, -1000, -1000, -1000, 275, 372, -1000,
	549, -1000, -1000, 379, 374, 212, 107, -1000, 131, -1000,
	114, -1000, -1000, -1000, -1000, -5, -14, -15, -1000, -1000,
	-1000, 633, 549, 633, -1000, -1000, 633, 549, 314, 211,
	-1000, -1000, 168, -1000, 574, -1000, 377, 489, 549, 489,
	-1000, -1000, 204, 199, 196, 633, 633, 403, -1000, 549,
	-1000, -1000, -1000, 372, 140, 167, 140, 269, 269, 269,
	275, -1000, 293, -6, -1000, -21, -25, 160, -1000, 401,
	352, -1000, 269, -1000, -1000, -1000, 269, -1000, 269, -1000,
}
var yyPgo = [...]int{

	0, 503, 502, 20, 501, 500, 499, 497, 495, 493,
	491, 490, 328, 489, 484, 483, 17, 26, 482, 481,
	480, 477, 11, 476, 471, 132, 469, 3, 16, 51,
	467, 466, 18, 464, 2, 15, 7, 463, 462, 9,
	461, 8, 457, 456, 13, 453, 452, 451, 450, 5,
	449, 6, 447, 1, 446, 24, 445, 10, 4, 19,
	218, 232, 443, 442, 441, 440, 439, 0, 12, 433,
	432, 427, 426, 418, 416, 80, 14, 415,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 4, 4, 72, 72, 5, 6, 7, 7, 69,
	70, 71, 74, 73, 73, 73, 73, 73, 73, 8,
	8, 8, 8, 9, 9, 9, 10, 11, 11, 11,
	11, 77, 12, 13, 13, 14, 14, 14, 14, 14,
	15, 15, 16, 16, 17, 17, 17, 20, 20, 18,
	18, 18, 21, 21, 22, 22, 22, 22, 19, 19,
	19, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	24, 24, 24, 25, 25, 26, 26, 26, 26, 27,
	27, 28, 28, 76, 76, 76, 75, 75, 29, 29,
	29, 29, 29, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 31, 31, 31, 31, 31, 31, 31,
	32, 32, 37, 37, 35, 35, 39, 36, 36, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 38, 38, 40, 40,
	40, 42, 45, 45, 43, 43, 44, 46, 46, 41,
	41, 33, 33, 33, 33, 47, 47, 48, 48, 49,
	49, 50, 50, 51, 52, 52, 52, 53, 53, 53,
	54, 54, 54, 55, 55, 56, 56, 57, 57, 58,
	58, 59, 59, 60, 60, 61, 61, 62, 62, 63,
	63, 63, 63, 63, 64, 64, 65, 65, 66, 66,
	67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 7, 12,
	3, 7, 7, 6, 6, 8, 7, 3, 4, 1,
	1, 1, 5, 4, 4, 2, 3, 4, 5, 5,
	5, 8, 4, 6, 7, 4, 5, 4, 4, 5,
	5, 0, 2, 0, 2, 1, 2, 1, 1, 1,
	0, 1, 1, 3, 1, 2, 3, 1, 1, 0,
	1, 2, 1, 3, 3, 3, 3, 5, 0, 1,
	2, 1, 1, 2, 3, 2, 3, 2, 2, 2,
	1, 3, 1, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 0, 2, 2, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 3, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	0, 2, 4, 0, 3, 1, 3, 0, 5, 1,
	3, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 97, 98, 100, 99, 84, 85, 86,
	89, 91, 27, 90, -14, 49, 50, 51, 52, -12,
	-77, -12, -12, -12, -12, 109, 101, -65, 103, 107,
	-62, 103, 105, 101, 109, 101, 102, 103, -12, 95,
	96, 92, 93, 94, 35, -67, 35, -3, 17, -15,
	18, -13, 29, -25, 35, 9, -58, 88, -59, -41,
	-67, 35, -61, 106, -61, 102, -67, 101, -67, 35,
	-60, 106, -60, -67, -60, 29, 23, 23, -76, 10,
	23, -75, 9, -67, 44, -16, 109, -17, 74, -20,
	35, -29, -34, -30, 68, 44, -33, -41, -35, -40,
	-67, -38, -42, 20, 36, 37, 38, 25, -39, 72,
	73, 48, 106, 28, 79, 39, -25, 33, 77, -25,
	53, 35, 45, 77, 35, 68, 35, -67, -68, 35,
	-68, 104, 35, 20, 35, 64, -67, -25, -34, 44,
	-34, -29, -34, -34, -76, -34, -75, -36, -34, -53,
	9, 53, 15, 44, -18, -67, 19, 77, 66, 67,
	-31, 21, 68, 23, 24, 22, 69, 70, 71, 72,
	73, 74, 75, 76, 45, 46, 47, 40, 41, 42,
	43, -29, -29, -36, -3, -34, -34, 44, 44, -39,
	44, -45, -34, -55, 33, 44, -58, 35, -28, 10,
	-59, -34, 64, -67, -68, 20, -68, -66, 108, -63,
	100, 98, 32, 99, 13, 35, 35, 35, -68, -55,
	33, -76, 110, 53, -21, -22, -24, 44, 35, -39,
	-17, -34, 110, -67, 74, -29, -29, -34, -35, 21,
	23, 24, -34, -34, 25, 68, -34, -34, -34, -34,
	-34, -34, -34, -34, 110, 110, 110, 110, -16, 18,
	-16, -43, -44, 80, -32, 28, -3, -58, -56, -41,
	-28, -49, 13, -29, 64, -67, -68, -64, 104, -32,
	-58, -34, -28, 53, -23, 54, 55, 56, 57, 58,
	60, 61, -19, 35, 19, -22, 77, 53, -53, -35,
	-34, -34, 66, 25, 110, -16, 110, -46, -44, 82,
	-29, -57, 64, -37, -35, -57, 110, 53, -49, -53,
	14, 35, 35, -47, 11, -22, -22, 54, 59, 54,
	59, 54, 54, 54, -26, 62, 105, 63, 35, 110,
	35, -34, 66, -34, 110, 83, -34, 81, 30, 53,
	-41, -53, -50, -51, -34, -68, -48, 12, 14, 64,
	54, 54, 102, 102, 102, -34, -34, 31, -35, 53,
	-52, 26, 27, -49, -29, -36, -29, 44, 44, 44,
	7, -51, -53, -27, -67, -27, -27, -58, -54, 16,
	34, 110, 53, 110, 110, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 51, 51,
	51, 51, 51, 216, 207, 0, 0, 29, 30, 31,
	51, 0, 0, 0, 0, 55, 57, 58, 59, 60,
	53, 0, 0, 0, 0, 205, 205, 0, 0, 217,
	0, 0, 208, 0, 203, 203, 0, 203, 0, 0,
	0, 103, 106, 0, 35, 0, 220, 20, 56, 0,
	61, 52, 0, 0, 93, 0, 27, 0, 199, 0,
	169, 220, 0, 0, 0, 0, 221, 0, 221, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	0, 103, 0, 106, 0, 187, 0, 62, 64, 69,
	220, 67, 68, 108, 0, 0, 139, 140, 141, 0,
	169, 0, 155, 0, 171, 172, 173, 174, 135, 158,
	159, 160, 156, 157, 162, 54, 193, 0, 0, 101,
	0, 28, 0, 0, 221, 0, 221, 218, 42, 0,
	45, 0, 47, 204, 48, 0, 221, 193, 33, 0,
	34, 104, 0, 105, 37, 107, 103, 0, 137, 17,
	0, 0, 0, 0, 65, 70, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 123, 124, 125, 126, 127, 128,
	129, 111, 0, 0, 0, 137, 150, 0, 0, 122,
	0, 0, 163, 0, 0, 0, 101, 94, 179, 0,
	200, 201, 202, 170, 39, 206, 40, 0, 0, 221,
	214, 209, 210, 211, 212, 213, 46, 49, 50, 0,
	0, 38, 32, 0, 101, 72, 78, 0, 90, 92,
	63, 188, 187, 71, 66, 109, 110, 113, 114, 0,
	0, 0, 116, 0, 120, 0, 142, 143, 144, 145,
	146, 147, 148, 149, 112, 134, 136, 151, 0, 0,
	0, 167, 164, 0, 197, 0, 131, 197, 0, 195,
	179, 187, 0, 102, 0, 219, 43, 0, 215, 23,
	24, 138, 175, 0, 0, 81, 82, 0, 0, 0,
	0, 0, 95, 79, 0, 0, 0, 0, 18, 115,
	117, 0, 0, 121, 152, 0, 154, 0, 165, 0,
	0, 21, 0, 130, 132, 22, 194, 0, 187, 26,
	0, 221, 44, 177, 0, 73, 76, 83, 0, 85,
	0, 87, 88, 89, 74, 0, 0, 0, 80, 75,
	91, 189, 0, 118, 153, 161, 168, 0, 0, 0,
	196, 25, 180, 181, 184, 41, 179, 0, 0, 0,
	84, 86, 0, 0, 0, 119, 166, 0, 133, 0,
	183, 185, 186, 187, 178, 176, 77, 0, 0, 0,
	0, 182, 190, 0, 99, 0, 0, 198, 19, 0,
	0, 96, 0, 97, 98, 191, 0, 100, 0, 192,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 76, 69, 3,
	44, 110, 74, 72, 53, 73, 77, 75, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 71, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:170
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:176
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:196
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:200
		{
			yyVAL.selStmt = &SimpleSelect{}
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:204
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:208
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:215
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:219
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:231
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:235
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:248
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:254
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:261
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &Begin{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.statement = &Commit{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:283
		{
			yyVAL.statement = &Rollback{}
		}
	case 32:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:289
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:307
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:311
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:315
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:325
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:329
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:334
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:344
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:349
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:355
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:361
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:365
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:369
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:374
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:379
		{
			SetAllowComments(yylex, true)
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:383
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:389
		{
			yyVAL.bytes2 = nil
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:393
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = AST_UNION
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.str = AST_EXCEPT
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:415
		{
			yyVAL.str = AST_INTERSECT
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:420
		{
			yyVAL.str = ""
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:424
		{
			yyVAL.str = AST_DISTINCT
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:434
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:444
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:454
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:463
		{
			yyVAL.bytes = nil
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:467
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:471
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:477
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:481
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:504
		{
			yyVAL.bytes = nil
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:508
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:512
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:518
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:542
		{
			yyVAL.str = AST_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:546
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:550
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:556
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:560
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:564
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:570
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:574
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:579
		{
			yyVAL.indexHints = nil
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:583
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:587
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:591
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:597
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:601
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:606
		{
			yyVAL.boolExpr = nil
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:615
		{
			yyVAL.expr = nil
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:619
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:623
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:628
		{
			yyVAL.valExpr = nil
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:632
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:639
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:689
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:693
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:699
		{
			yyVAL.str = AST_EQ
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_LT
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_GT
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_LE
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:715
		{
			yyVAL.str = AST_GE
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:719
		{
			yyVAL.str = AST_NE
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.str = AST_NSE
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:729
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:733
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:739
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:743
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:749
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:753
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:759
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:765
		{
			yyVAL.valExprs = ValExprs{NumVal{}, yyDollar[1].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:769
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:815
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:819
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:856
		{
			yyVAL.bytes = IF_BYTES
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:860
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:866
		{
			yyVAL.byt = AST_UPLUS
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.byt = AST_UMINUS
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:874
		{
			yyVAL.byt = AST_TILDA
		}
	case 161:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:880
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = nil
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:895
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:899
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:905
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:910
		{
			yyVAL.valExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:914
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:920
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:924
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:938
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:942
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExprs = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:956
		{
			yyVAL.boolExpr = nil
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:960
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:965
		{
			yyVAL.orderBy = nil
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:969
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:975
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:979
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:985
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:990
		{
			yyVAL.str = AST_ASC
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:994
		{
			yyVAL.str = AST_ASC
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.str = AST_DESC
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.limit = nil
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.str = ""
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1024
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.columns = nil
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.updateExprs = nil
		}
	case 198:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.updateExprs = UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("version")}, Expr: NumVal([]byte{48})}, yyDollar[1].updateExpr}
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1134
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
