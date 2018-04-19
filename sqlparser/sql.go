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

const yyNprod = 223
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 634

var yyAct = [...]int{

	120, 169, 162, 405, 76, 292, 375, 203, 117, 148,
	128, 333, 246, 283, 218, 118, 285, 105, 98, 78,
	253, 204, 3, 123, 213, 228, 107, 414, 127, 91,
	414, 133, 83, 299, 65, 178, 179, 414, 110, 124,
	125, 126, 54, 56, 57, 80, 151, 115, 111, 86,
	55, 131, 88, 79, 101, 386, 67, 93, 18, 307,
	308, 309, 310, 311, 103, 312, 313, 35, 36, 37,
	38, 114, 112, 123, 171, 129, 130, 108, 127, 275,
	385, 133, 134, 339, 416, 384, 147, 415, 81, 124,
	125, 126, 85, 171, 413, 156, 171, 115, 150, 158,
	160, 131, 87, 163, 244, 165, 367, 168, 244, 132,
	175, 46, 167, 48, 51, 361, 52, 49, 205, 45,
	164, 114, 206, 53, 284, 129, 130, 81, 277, 357,
	359, 366, 134, 318, 209, 237, 73, 212, 80, 177,
	338, 80, 216, 143, 223, 221, 79, 138, 161, 79,
	328, 234, 145, 326, 224, 284, 226, 331, 166, 132,
	220, 276, 168, 201, 202, 243, 239, 191, 192, 193,
	232, 265, 358, 235, 112, 252, 381, 254, 223, 66,
	77, 250, 240, 258, 334, 242, 263, 264, 295, 267,
	268, 269, 270, 271, 272, 273, 274, 259, 251, 61,
	62, 63, 59, 60, 155, 178, 179, 178, 179, 136,
	112, 112, 139, 383, 266, 80, 80, 382, 255, 288,
	369, 90, 140, 79, 290, 279, 281, 256, 257, 296,
	355, 291, 157, 334, 354, 287, 231, 233, 230, 297,
	351, 353, 80, 349, 140, 352, 302, 303, 350, 219,
	79, 35, 36, 37, 38, 320, 244, 301, 391, 250,
	304, 317, 287, 371, 322, 323, 142, 82, 294, 170,
	401, 18, 19, 20, 21, 172, 321, 92, 219, 94,
	400, 364, 399, 112, 186, 187, 188, 189, 190, 191,
	192, 193, 305, 32, 341, 159, 330, 340, 327, 22,
	337, 324, 336, 210, 186, 187, 188, 189, 190, 191,
	192, 193, 241, 171, 84, 249, 250, 250, 347, 348,
	208, 140, 363, 215, 248, 207, 173, 365, 104, 18,
	393, 394, 214, 332, 368, 189, 190, 191, 192, 193,
	80, 135, 373, 215, 316, 376, 176, 66, 372, 81,
	27, 28, 29, 377, 362, 30, 33, 31, 360, 249,
	315, 344, 66, 23, 24, 26, 25, 387, 248, 343,
	137, 300, 388, 186, 187, 188, 189, 190, 191, 192,
	193, 411, 238, 168, 395, 236, 217, 390, 397, 307,
	308, 309, 310, 311, 376, 312, 313, 404, 403, 412,
	406, 406, 406, 80, 407, 408, 280, 409, 123, 74,
	154, 79, 152, 127, 149, 419, 133, 146, 144, 420,
	141, 421, 89, 110, 124, 125, 126, 64, 396, 123,
	398, 389, 115, 370, 127, 18, 131, 133, 325, 95,
	72, 99, 97, 96, 110, 124, 125, 126, 260, 418,
	261, 262, 319, 115, 100, 225, 114, 131, 286, 153,
	129, 130, 108, 70, 68, 172, 380, 134, 186, 187,
	188, 189, 190, 191, 192, 193, 342, 114, 293, 379,
	346, 129, 130, 108, 123, 219, 102, 75, 134, 127,
	18, 417, 133, 402, 132, 18, 40, 17, 278, 81,
	124, 125, 126, 16, 39, 15, 14, 13, 115, 12,
	127, 227, 131, 133, 47, 132, 298, 229, 106, 50,
	81, 124, 125, 126, 41, 42, 43, 44, 127, 159,
	289, 133, 114, 131, 410, 58, 129, 130, 81, 124,
	125, 126, 127, 134, 392, 133, 374, 159, 378, 345,
	329, 131, 81, 124, 125, 126, 211, 129, 130, 282,
	122, 159, 119, 121, 134, 131, 335, 222, 116, 180,
	132, 113, 356, 247, 306, 129, 130, 245, 181, 185,
	183, 184, 134, 109, 314, 174, 69, 34, 71, 129,
	130, 132, 11, 10, 9, 8, 134, 197, 198, 199,
	200, 7, 194, 195, 196, 6, 5, 4, 2, 132,
	186, 187, 188, 189, 190, 191, 192, 193, 1, 0,
	0, 0, 0, 132, 0, 182, 186, 187, 188, 189,
	190, 191, 192, 193,
}
var yyPact = [...]int{

	266, -1000, -1000, 202, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 10, 11, 22, -59, -1000, -1000, -1000,
	-1000, 107, 392, 312, 490, 447, -1000, -1000, -1000, 445,
	-1000, 411, 374, 478, 92, -74, -74, -10, 312, -1000,
	1, 312, -1000, 387, -77, -77, 312, -77, 410, 420,
	419, 431, 477, 312, -1000, 284, -1000, -1000, -1000, 409,
	-1000, 302, 374, 337, 70, 374, 191, 385, -1000, 221,
	-1000, 66, 383, 84, 382, 312, -1000, 379, -1000, -58,
	377, 439, 375, 140, 312, 374, 517, 517, -1000, 464,
	517, 431, 517, 477, 517, 260, 282, -1000, -1000, 327,
	62, 141, 557, -1000, 464, 53, -1000, -1000, -1000, 517,
	281, 276, -1000, 259, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 517, -1000, 299, 314, 351, 475,
	314, -1000, 503, 312, -1000, 435, -1000, -83, -1000, 138,
	-1000, 350, 58, -1000, -1000, 347, -1000, 279, 541, 485,
	541, 141, 557, 541, -1000, 541, 431, 55, 541, -1000,
	280, 3, 517, -90, -1000, -1000, 312, 144, 464, 464,
	517, 251, 427, 517, 517, 146, 517, 517, 517, 517,
	517, 517, 517, 517, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -31, 51, 18, 557, -1000, 388, 3, -1000,
	490, 44, 541, 430, 314, 314, 268, -1000, 465, 464,
	-1000, 541, -1000, -1000, -1000, -1000, -1000, 124, 312, -1000,
	-71, -1000, -1000, -1000, -1000, -1000, -1000, 336, -1000, -1000,
	430, 314, -1000, -1000, 517, 239, 335, 325, 324, 56,
	-1000, -1000, 399, 450, -1000, -1000, -1000, -1000, 541, -1000,
	251, 517, 517, 541, 235, -1000, 413, 263, 263, 263,
	93, 93, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 43,
	3, 40, 75, -1000, 464, 120, 251, 202, 169, 30,
	-1000, 465, 450, 462, 141, 334, -1000, -1000, 326, -1000,
	-1000, -1000, 191, 541, 469, 280, 280, -1000, -1000, 189,
	186, 187, 180, 176, 67, -1000, 323, 5, 319, 517,
	-1000, -1000, 541, 215, 517, -1000, -1000, 21, -1000, 23,
	-1000, 517, 139, -1000, 403, 210, -1000, -1000, -1000, 314,
	450, -1000, 517, -1000, -1000, 467, 452, 335, 112, -1000,
	163, -1000, 159, -1000, -1000, -1000, -1000, -17, -22, -47,
	-1000, -1000, -1000, 541, 517, 541, -1000, -1000, 541, 517,
	400, 251, -1000, -1000, 205, -1000, 304, -1000, 465, 464,
	517, 464, -1000, -1000, 238, 236, 226, 541, 541, 486,
	-1000, 517, -1000, -1000, -1000, 450, 141, 203, 141, 312,
	312, 312, 314, -1000, 365, -16, -1000, -23, -26, 191,
	-1000, 484, 428, -1000, 312, -1000, -1000, -1000, 312, -1000,
	312, -1000,
}
var yyPgo = [...]int{

	0, 618, 608, 21, 607, 606, 605, 601, 595, 594,
	593, 592, 504, 588, 587, 586, 17, 26, 585, 584,
	583, 577, 12, 574, 573, 136, 572, 3, 14, 48,
	571, 569, 16, 568, 2, 15, 7, 566, 563, 10,
	562, 8, 560, 559, 13, 556, 550, 549, 548, 5,
	546, 6, 544, 1, 534, 24, 530, 11, 4, 19,
	221, 267, 519, 517, 516, 514, 511, 0, 9, 509,
	507, 506, 505, 503, 497, 54, 18, 496,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 4, 4, 72, 72, 5, 6, 7, 7, 69,
	70, 71, 74, 73, 73, 73, 73, 73, 73, 8,
	8, 8, 8, 9, 9, 9, 10, 11, 11, 11,
	11, 11, 77, 12, 13, 13, 14, 14, 14, 14,
	14, 15, 15, 16, 16, 17, 17, 17, 20, 20,
	18, 18, 18, 21, 21, 22, 22, 22, 22, 19,
	19, 19, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 24, 24, 24, 25, 25, 26, 26, 26, 26,
	27, 27, 28, 28, 76, 76, 76, 75, 75, 29,
	29, 29, 29, 29, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 31, 31, 31, 31, 31, 31,
	31, 32, 32, 37, 37, 35, 35, 39, 36, 36,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 38, 38, 40,
	40, 40, 42, 45, 45, 43, 43, 44, 46, 46,
	41, 41, 33, 33, 33, 33, 47, 47, 48, 48,
	49, 49, 50, 50, 51, 52, 52, 52, 53, 53,
	53, 54, 54, 54, 55, 55, 56, 56, 57, 57,
	58, 58, 59, 59, 60, 60, 61, 61, 62, 62,
	63, 63, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 7, 12,
	3, 7, 7, 6, 6, 8, 7, 3, 4, 1,
	1, 1, 5, 4, 4, 2, 3, 4, 5, 5,
	5, 8, 4, 6, 7, 4, 5, 6, 4, 4,
	5, 5, 0, 2, 0, 2, 1, 2, 1, 1,
	1, 0, 1, 1, 3, 1, 2, 3, 1, 1,
	0, 1, 2, 1, 3, 3, 3, 3, 5, 0,
	1, 2, 1, 1, 2, 3, 2, 3, 2, 2,
	2, 1, 3, 1, 1, 3, 0, 5, 5, 5,
	1, 3, 0, 2, 0, 2, 2, 0, 2, 1,
	3, 3, 2, 3, 3, 3, 4, 3, 4, 5,
	6, 3, 4, 2, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 3, 3, 1, 3, 1, 3,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 4, 5, 4, 1, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 1, 1, 1, 1, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 0, 2, 4, 0, 3, 1, 3, 0, 5,
	1, 3, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 97, 98, 100, 99, 84, 85, 86,
	89, 91, 27, 90, -14, 49, 50, 51, 52, -12,
	-77, -12, -12, -12, -12, 109, 101, -65, 103, 107,
	-62, 103, 105, 101, 101, 109, 102, 103, -12, 95,
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
	100, 98, 32, 99, 13, 35, 35, 77, 35, -68,
	-55, 33, -76, 110, 53, -21, -22, -24, 44, 35,
	-39, -17, -34, 110, -67, 74, -29, -29, -34, -35,
	21, 23, 24, -34, -34, 25, 68, -34, -34, -34,
	-34, -34, -34, -34, -34, 110, 110, 110, 110, -16,
	18, -16, -43, -44, 80, -32, 28, -3, -58, -56,
	-41, -28, -49, 13, -29, 64, -67, -68, -64, 104,
	35, -32, -58, -34, -28, 53, -23, 54, 55, 56,
	57, 58, 60, 61, -19, 35, 19, -22, 77, 53,
	-53, -35, -34, -34, 66, 25, 110, -16, 110, -46,
	-44, 82, -29, -57, 64, -37, -35, -57, 110, 53,
	-49, -53, 14, 35, 35, -47, 11, -22, -22, 54,
	59, 54, 59, 54, 54, 54, -26, 62, 105, 63,
	35, 110, 35, -34, 66, -34, 110, 83, -34, 81,
	30, 53, -41, -53, -50, -51, -34, -68, -48, 12,
	14, 64, 54, 54, 102, 102, 102, -34, -34, 31,
	-35, 53, -52, 26, 27, -49, -29, -36, -29, 44,
	44, 44, 7, -51, -53, -27, -67, -27, -27, -58,
	-54, 16, 34, 110, 53, 110, 110, 7, 21, -67,
	-67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 52, 52,
	52, 52, 52, 217, 208, 0, 0, 29, 30, 31,
	52, 0, 0, 0, 0, 56, 58, 59, 60, 61,
	54, 0, 0, 0, 0, 206, 206, 0, 0, 218,
	0, 0, 209, 0, 204, 204, 0, 204, 0, 0,
	0, 104, 107, 0, 35, 0, 221, 20, 57, 0,
	62, 53, 0, 0, 94, 0, 27, 0, 200, 0,
	170, 221, 0, 0, 0, 0, 222, 0, 222, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	0, 104, 0, 107, 0, 188, 0, 63, 65, 70,
	221, 68, 69, 109, 0, 0, 140, 141, 142, 0,
	170, 0, 156, 0, 172, 173, 174, 175, 136, 159,
	160, 161, 157, 158, 163, 55, 194, 0, 0, 102,
	0, 28, 0, 0, 222, 0, 222, 219, 42, 0,
	45, 0, 49, 205, 48, 0, 222, 194, 33, 0,
	34, 105, 0, 106, 37, 108, 104, 0, 138, 17,
	0, 0, 0, 0, 66, 71, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 124, 125, 126, 127, 128, 129,
	130, 112, 0, 0, 0, 138, 151, 0, 0, 123,
	0, 0, 164, 0, 0, 0, 102, 95, 180, 0,
	201, 202, 203, 171, 39, 207, 40, 0, 0, 222,
	215, 210, 211, 212, 213, 214, 46, 0, 50, 51,
	0, 0, 38, 32, 0, 102, 73, 79, 0, 91,
	93, 64, 189, 188, 72, 67, 110, 111, 114, 115,
	0, 0, 0, 117, 0, 121, 0, 143, 144, 145,
	146, 147, 148, 149, 150, 113, 135, 137, 152, 0,
	0, 0, 168, 165, 0, 198, 0, 132, 198, 0,
	196, 180, 188, 0, 103, 0, 220, 43, 0, 216,
	47, 23, 24, 139, 176, 0, 0, 82, 83, 0,
	0, 0, 0, 0, 96, 80, 0, 0, 0, 0,
	18, 116, 118, 0, 0, 122, 153, 0, 155, 0,
	166, 0, 0, 21, 0, 131, 133, 22, 195, 0,
	188, 26, 0, 222, 44, 178, 0, 74, 77, 84,
	0, 86, 0, 88, 89, 90, 75, 0, 0, 0,
	81, 76, 92, 190, 0, 119, 154, 162, 169, 0,
	0, 0, 197, 25, 181, 182, 185, 41, 180, 0,
	0, 0, 85, 87, 0, 0, 0, 120, 167, 0,
	134, 0, 184, 186, 187, 188, 179, 177, 78, 0,
	0, 0, 0, 183, 191, 0, 100, 0, 0, 199,
	19, 0, 0, 97, 0, 98, 99, 192, 0, 101,
	0, 193,
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
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:361
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:365
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:369
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:373
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:378
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:383
		{
			SetAllowComments(yylex, true)
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:387
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:393
		{
			yyVAL.bytes2 = nil
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:397
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = AST_UNION
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:415
		{
			yyVAL.str = AST_EXCEPT
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:419
		{
			yyVAL.str = AST_INTERSECT
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:424
		{
			yyVAL.str = ""
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:428
		{
			yyVAL.str = AST_DISTINCT
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:438
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:444
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:452
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:462
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:467
		{
			yyVAL.bytes = nil
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:471
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:475
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:481
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:485
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:503
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:508
		{
			yyVAL.bytes = nil
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:512
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:516
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:542
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:546
		{
			yyVAL.str = AST_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:550
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:554
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:560
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:564
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:568
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:574
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:578
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:583
		{
			yyVAL.indexHints = nil
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:587
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:591
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:595
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:601
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:605
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:614
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:619
		{
			yyVAL.expr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:623
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:627
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:632
		{
			yyVAL.valExpr = nil
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:636
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:651
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:655
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:689
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:693
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:697
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_EQ
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_LT
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_GT
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:715
		{
			yyVAL.str = AST_LE
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:719
		{
			yyVAL.str = AST_GE
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.str = AST_NE
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:727
		{
			yyVAL.str = AST_NSE
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:733
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:737
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:743
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:747
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:753
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:757
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:763
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:769
		{
			yyVAL.valExprs = ValExprs{NumVal{}, yyDollar[1].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:773
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:815
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:823
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
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 154:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:860
		{
			yyVAL.bytes = IF_BYTES
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:864
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.byt = AST_UPLUS
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:874
		{
			yyVAL.byt = AST_UMINUS
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:878
		{
			yyVAL.byt = AST_TILDA
		}
	case 162:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:884
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = nil
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:903
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:909
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:914
		{
			yyVAL.valExpr = nil
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:924
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:928
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:938
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:942
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:946
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExprs = nil
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:960
		{
			yyVAL.boolExpr = nil
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:964
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:969
		{
			yyVAL.orderBy = nil
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:973
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:979
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:983
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:989
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:994
		{
			yyVAL.str = AST_ASC
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.str = AST_ASC
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.str = AST_DESC
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.limit = nil
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 190:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.str = ""
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1028
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
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.columns = nil
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.updateExprs = nil
		}
	case 199:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.updateExprs = UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("version")}, Expr: NumVal([]byte{48})}, yyDollar[1].updateExpr}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1095
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1100
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
		//line sql.y:1106
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1138
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
