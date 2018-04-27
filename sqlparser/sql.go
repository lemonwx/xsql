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
	"'?'",
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

const yyNprod = 224
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 704

var yyAct = [...]int{

	120, 169, 162, 406, 76, 293, 376, 203, 117, 148,
	128, 334, 246, 284, 218, 118, 286, 105, 98, 78,
	127, 204, 3, 133, 213, 415, 107, 178, 179, 253,
	81, 124, 125, 126, 65, 35, 36, 37, 38, 159,
	415, 228, 91, 131, 83, 80, 300, 415, 111, 86,
	151, 171, 88, 79, 340, 101, 67, 93, 387, 308,
	309, 310, 311, 312, 103, 313, 314, 129, 130, 171,
	171, 276, 112, 244, 134, 54, 56, 57, 244, 46,
	87, 48, 417, 55, 386, 49, 147, 45, 358, 360,
	53, 51, 385, 52, 85, 156, 278, 416, 150, 158,
	160, 132, 81, 163, 414, 165, 258, 168, 367, 285,
	175, 339, 167, 368, 234, 362, 319, 285, 205, 332,
	164, 237, 206, 177, 178, 179, 329, 327, 143, 138,
	277, 359, 145, 232, 209, 243, 235, 212, 80, 370,
	66, 80, 216, 382, 223, 221, 79, 140, 161, 79,
	191, 192, 193, 335, 224, 77, 226, 73, 335, 166,
	220, 296, 168, 201, 202, 384, 239, 61, 62, 63,
	59, 60, 178, 179, 112, 252, 155, 254, 223, 255,
	90, 250, 240, 259, 383, 242, 264, 265, 356, 268,
	269, 270, 271, 272, 273, 274, 275, 260, 251, 231,
	233, 230, 186, 187, 188, 189, 190, 191, 192, 193,
	112, 112, 352, 266, 170, 80, 80, 353, 355, 289,
	172, 354, 140, 79, 291, 280, 282, 256, 257, 297,
	136, 292, 350, 139, 244, 288, 92, 351, 94, 298,
	219, 241, 80, 394, 395, 392, 303, 304, 372, 142,
	79, 402, 215, 157, 249, 321, 267, 302, 171, 250,
	305, 318, 288, 248, 401, 323, 324, 400, 295, 159,
	210, 208, 18, 19, 20, 21, 207, 322, 189, 190,
	191, 192, 193, 306, 112, 82, 186, 187, 188, 189,
	190, 191, 192, 193, 32, 342, 219, 331, 341, 328,
	22, 338, 365, 337, 173, 186, 187, 188, 189, 190,
	191, 192, 193, 35, 36, 37, 38, 250, 250, 348,
	349, 104, 214, 364, 135, 66, 317, 81, 366, 176,
	18, 363, 84, 215, 333, 369, 361, 412, 345, 140,
	95, 80, 316, 374, 344, 66, 377, 301, 238, 373,
	236, 27, 28, 29, 378, 413, 30, 33, 31, 217,
	249, 74, 154, 152, 23, 24, 26, 25, 388, 248,
	149, 325, 146, 389, 186, 187, 188, 189, 190, 191,
	192, 193, 144, 141, 168, 396, 89, 64, 391, 398,
	308, 309, 310, 311, 312, 377, 313, 314, 405, 404,
	137, 407, 407, 407, 80, 408, 409, 281, 410, 123,
	390, 371, 79, 18, 127, 72, 420, 133, 326, 39,
	421, 99, 422, 419, 110, 124, 125, 126, 97, 397,
	123, 399, 96, 115, 100, 127, 287, 131, 133, 41,
	42, 43, 44, 225, 153, 110, 124, 125, 126, 261,
	58, 262, 263, 320, 115, 70, 68, 114, 131, 172,
	381, 129, 130, 108, 343, 294, 380, 347, 134, 186,
	187, 188, 189, 190, 191, 192, 193, 219, 114, 102,
	75, 418, 129, 130, 108, 123, 403, 18, 40, 134,
	127, 18, 17, 133, 16, 132, 15, 14, 13, 279,
	110, 124, 125, 126, 12, 227, 123, 47, 299, 115,
	229, 127, 50, 131, 133, 290, 132, 411, 393, 106,
	375, 81, 124, 125, 126, 379, 346, 330, 211, 283,
	115, 122, 119, 114, 131, 121, 336, 129, 130, 108,
	116, 180, 113, 357, 134, 123, 247, 18, 307, 245,
	127, 109, 315, 133, 114, 174, 69, 34, 129, 130,
	81, 124, 125, 126, 71, 134, 11, 127, 10, 115,
	133, 132, 127, 131, 9, 133, 8, 81, 124, 125,
	126, 7, 81, 124, 125, 126, 159, 6, 5, 4,
	131, 159, 132, 114, 2, 131, 1, 129, 130, 0,
	0, 0, 0, 0, 134, 0, 0, 0, 0, 0,
	0, 222, 0, 0, 129, 130, 0, 0, 0, 129,
	130, 134, 127, 0, 0, 133, 134, 0, 0, 0,
	0, 132, 81, 124, 125, 126, 181, 185, 183, 184,
	0, 159, 0, 0, 0, 131, 0, 0, 132, 0,
	0, 0, 0, 132, 0, 197, 198, 199, 200, 0,
	194, 195, 196, 0, 0, 0, 0, 0, 0, 129,
	130, 0, 0, 0, 0, 0, 134, 0, 0, 0,
	0, 0, 0, 182, 186, 187, 188, 189, 190, 191,
	192, 193, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 132,
}
var yyPact = [...]int{

	267, -1000, -1000, 264, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -22, -12, -11, -26, -1000, -1000, -1000,
	-1000, 75, 352, 290, 482, 439, -1000, -1000, -1000, 437,
	-1000, 386, 326, 471, 67, -62, -62, -8, 290, -1000,
	-21, 290, -1000, 351, -64, -64, 290, -64, 311, 409,
	405, 411, 470, 290, -1000, 277, -1000, -1000, -1000, 410,
	-1000, 285, 326, 367, 52, 326, 169, 348, -1000, 204,
	-1000, 51, 347, 64, 337, 290, -1000, 335, -1000, -54,
	328, 424, 327, 112, 290, 326, 597, 597, -1000, 525,
	597, 411, 597, 470, 597, 205, 260, -1000, -1000, 310,
	46, 106, 615, -1000, 525, 486, -1000, -1000, -1000, 597,
	232, 227, -1000, 226, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 597, -1000, 289, 292, 324, 467,
	292, -1000, 547, 290, -1000, 423, -1000, -67, -1000, 101,
	-1000, 315, 44, -1000, -1000, 313, -1000, 208, 133, 542,
	133, 106, 615, 133, -1000, 133, 411, 25, 133, -1000,
	219, 465, 597, -81, -1000, -1000, 290, 105, 525, 525,
	-5, 225, 428, 597, 597, 188, 597, 597, 597, 597,
	597, 597, 597, 597, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -39, 20, -14, 615, -1000, 389, 465, -1000,
	482, 29, 133, 408, 292, 292, 286, -1000, 452, 525,
	-1000, 133, -1000, -1000, -1000, -1000, -1000, 97, 290, -1000,
	-58, -1000, -1000, -1000, -1000, -1000, -1000, 312, -1000, -1000,
	408, 292, -1000, -1000, 597, 230, 336, 307, 325, 39,
	-1000, -1000, 400, 444, -1000, -1000, -1000, -1000, -1000, 133,
	-1000, 225, 597, 597, 133, 305, -1000, 393, 206, 206,
	206, 76, 76, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	17, 465, 16, 37, -1000, 525, 89, 225, 264, 94,
	1, -1000, 452, 444, 450, 106, 309, -1000, -1000, 303,
	-1000, -1000, -1000, 169, 133, 456, 219, 219, -1000, -1000,
	178, 158, 167, 164, 134, 26, -1000, 301, 5, 296,
	597, -1000, -1000, 133, 236, 597, -1000, -1000, -2, -1000,
	30, -1000, 597, 58, -1000, 381, 195, -1000, -1000, -1000,
	292, 444, -1000, 597, -1000, -1000, 454, 446, 336, 79,
	-1000, 130, -1000, 111, -1000, -1000, -1000, -1000, -10, -18,
	-44, -1000, -1000, -1000, 133, 597, 133, -1000, -1000, 133,
	597, 379, 225, -1000, -1000, 192, -1000, 217, -1000, 452,
	525, 597, 525, -1000, -1000, 223, 220, 207, 133, 133,
	479, -1000, 597, -1000, -1000, -1000, 444, 106, 181, 106,
	290, 290, 290, 292, -1000, 321, -6, -1000, -13, -28,
	169, -1000, 474, 402, -1000, 290, -1000, -1000, -1000, 290,
	-1000, 290, -1000,
}
var yyPgo = [...]int{

	0, 596, 594, 21, 589, 588, 587, 581, 576, 574,
	568, 566, 419, 564, 557, 556, 17, 26, 555, 552,
	551, 549, 12, 548, 546, 157, 543, 3, 14, 48,
	542, 541, 16, 540, 2, 15, 7, 536, 535, 10,
	532, 8, 531, 529, 13, 528, 527, 526, 525, 5,
	520, 6, 518, 1, 517, 24, 515, 11, 4, 19,
	180, 285, 512, 510, 508, 507, 505, 0, 9, 504,
	498, 497, 496, 494, 492, 55, 18, 488,
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
	30, 30, 30, 30, 30, 31, 31, 31, 31, 31,
	31, 31, 32, 32, 37, 37, 35, 35, 39, 36,
	36, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 38, 38,
	40, 40, 40, 42, 45, 45, 43, 43, 44, 46,
	46, 41, 41, 33, 33, 33, 33, 47, 47, 48,
	48, 49, 49, 50, 50, 51, 52, 52, 52, 53,
	53, 53, 54, 54, 54, 55, 55, 56, 56, 57,
	57, 58, 58, 59, 59, 60, 60, 61, 61, 62,
	62, 63, 63, 63, 63, 63, 64, 64, 65, 65,
	66, 66, 67, 68,
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
	3, 3, 2, 3, 3, 3, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 3, 3, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 5, 4, 1, 1, 1,
	1, 1, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 1, 1, 1, 1, 0, 3, 0,
	2, 0, 3, 1, 3, 2, 0, 1, 1, 0,
	2, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
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
	-39, -17, -34, 110, -67, 74, -29, -29, 111, -34,
	-35, 21, 23, 24, -34, -34, 25, 68, -34, -34,
	-34, -34, -34, -34, -34, -34, 110, 110, 110, 110,
	-16, 18, -16, -43, -44, 80, -32, 28, -3, -58,
	-56, -41, -28, -49, 13, -29, 64, -67, -68, -64,
	104, 35, -32, -58, -34, -28, 53, -23, 54, 55,
	56, 57, 58, 60, 61, -19, 35, 19, -22, 77,
	53, -53, -35, -34, -34, 66, 25, 110, -16, 110,
	-46, -44, 82, -29, -57, 64, -37, -35, -57, 110,
	53, -49, -53, 14, 35, 35, -47, 11, -22, -22,
	54, 59, 54, 59, 54, 54, 54, -26, 62, 105,
	63, 35, 110, 35, -34, 66, -34, 110, 83, -34,
	81, 30, 53, -41, -53, -50, -51, -34, -68, -48,
	12, 14, 64, 54, 54, 102, 102, 102, -34, -34,
	31, -35, 53, -52, 26, 27, -49, -29, -36, -29,
	44, 44, 44, 7, -51, -53, -27, -67, -27, -27,
	-58, -54, 16, 34, 110, 53, 110, 110, 7, 21,
	-67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 52, 52,
	52, 52, 52, 218, 209, 0, 0, 29, 30, 31,
	52, 0, 0, 0, 0, 56, 58, 59, 60, 61,
	54, 0, 0, 0, 0, 207, 207, 0, 0, 219,
	0, 0, 210, 0, 205, 205, 0, 205, 0, 0,
	0, 104, 107, 0, 35, 0, 222, 20, 57, 0,
	62, 53, 0, 0, 94, 0, 27, 0, 201, 0,
	171, 222, 0, 0, 0, 0, 223, 0, 223, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	0, 104, 0, 107, 0, 189, 0, 63, 65, 70,
	222, 68, 69, 109, 0, 0, 141, 142, 143, 0,
	171, 0, 157, 0, 173, 174, 175, 176, 137, 160,
	161, 162, 158, 159, 164, 55, 195, 0, 0, 102,
	0, 28, 0, 0, 223, 0, 223, 220, 42, 0,
	45, 0, 49, 206, 48, 0, 223, 195, 33, 0,
	34, 105, 0, 106, 37, 108, 104, 0, 139, 17,
	0, 0, 0, 0, 66, 71, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 125, 126, 127, 128, 129, 130,
	131, 112, 0, 0, 0, 139, 152, 0, 0, 124,
	0, 0, 165, 0, 0, 0, 102, 95, 181, 0,
	202, 203, 204, 172, 39, 208, 40, 0, 0, 223,
	216, 211, 212, 213, 214, 215, 46, 0, 50, 51,
	0, 0, 38, 32, 0, 102, 73, 79, 0, 91,
	93, 64, 190, 189, 72, 67, 110, 111, 114, 115,
	116, 0, 0, 0, 118, 0, 122, 0, 144, 145,
	146, 147, 148, 149, 150, 151, 113, 136, 138, 153,
	0, 0, 0, 169, 166, 0, 199, 0, 133, 199,
	0, 197, 181, 189, 0, 103, 0, 221, 43, 0,
	217, 47, 23, 24, 140, 177, 0, 0, 82, 83,
	0, 0, 0, 0, 0, 96, 80, 0, 0, 0,
	0, 18, 117, 119, 0, 0, 123, 154, 0, 156,
	0, 167, 0, 0, 21, 0, 132, 134, 22, 196,
	0, 189, 26, 0, 223, 44, 179, 0, 74, 77,
	84, 0, 86, 0, 88, 89, 90, 75, 0, 0,
	0, 81, 76, 92, 191, 0, 120, 155, 163, 170,
	0, 0, 0, 198, 25, 182, 183, 186, 41, 181,
	0, 0, 0, 85, 87, 0, 0, 0, 121, 168,
	0, 135, 0, 185, 187, 188, 189, 180, 178, 78,
	0, 0, 0, 0, 184, 192, 0, 100, 0, 0,
	200, 19, 0, 0, 97, 0, 98, 99, 193, 0,
	101, 0, 194,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 76, 69, 3,
	44, 110, 74, 72, 53, 73, 77, 75, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 111, 3, 3, 3, 3, 3, 3,
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
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: NewSelectExprs(yyDollar[4].selectExprs), From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: NewGroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
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
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: ValArg("?")}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:689
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:693
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:697
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:701
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_EQ
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_LT
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:715
		{
			yyVAL.str = AST_GT
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:719
		{
			yyVAL.str = AST_LE
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.str = AST_GE
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:727
		{
			yyVAL.str = AST_NE
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:731
		{
			yyVAL.str = AST_NSE
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:737
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:741
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:747
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:751
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:757
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:761
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:767
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:773
		{
			yyVAL.valExprs = ValExprs{NumVal{}, yyDollar[1].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:777
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:815
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:823
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:827
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
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:846
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:850
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:854
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:864
		{
			yyVAL.bytes = IF_BYTES
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:868
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:874
		{
			yyVAL.byt = AST_UPLUS
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:878
		{
			yyVAL.byt = AST_UMINUS
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:882
		{
			yyVAL.byt = AST_TILDA
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:888
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = nil
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:903
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:907
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:913
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = nil
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:928
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:932
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:938
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:942
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:946
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:950
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExprs = nil
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:959
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:964
		{
			yyVAL.boolExpr = nil
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:968
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:973
		{
			yyVAL.orderBy = nil
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:977
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:983
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:987
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:993
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:998
		{
			yyVAL.str = AST_ASC
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.str = AST_ASC
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.str = AST_DESC
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.limit = nil
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.str = ""
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1032
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
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.columns = nil
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1059
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.updateExprs = nil
		}
	case 200:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.updateExprs = UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("version")}, Expr: NumVal([]byte{48})}, yyDollar[1].updateExpr}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1104
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
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1142
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
