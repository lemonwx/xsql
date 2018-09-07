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
const START = 57415
const TRANSACTION = 57416
const NAMES = 57417
const REPLACE = 57418
const ADMIN = 57419
const KILL = 57420
const SHOW = 57421
const DATABASES = 57422
const TABLES = 57423
const PROXY = 57424
const VARIABLES = 57425
const STATUS = 57426
const CREATE = 57427
const ALTER = 57428
const DROP = 57429
const RENAME = 57430
const TABLE = 57431
const INDEX = 57432
const VIEW = 57433
const TO = 57434
const IGNORE = 57435
const IF = 57436
const UNIQUE = 57437
const USING = 57438
const DATABASE = 57439

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
	"START",
	"TRANSACTION",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"KILL",
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

const yyNprod = 227
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 668

var yyAct = [...]int{

	127, 177, 169, 414, 82, 300, 384, 211, 124, 135,
	342, 254, 291, 155, 221, 125, 226, 112, 293, 104,
	212, 3, 84, 57, 59, 60, 114, 288, 261, 130,
	423, 58, 186, 187, 134, 236, 70, 140, 38, 39,
	40, 41, 97, 89, 117, 131, 132, 133, 86, 308,
	423, 118, 92, 122, 107, 94, 85, 138, 73, 49,
	99, 51, 395, 423, 19, 52, 179, 48, 109, 316,
	317, 318, 319, 320, 348, 321, 322, 121, 119, 283,
	158, 136, 137, 115, 134, 394, 179, 140, 141, 54,
	425, 55, 154, 393, 87, 131, 132, 133, 179, 252,
	252, 163, 285, 166, 91, 165, 167, 138, 157, 170,
	424, 172, 110, 61, 176, 366, 368, 183, 139, 175,
	242, 93, 286, 422, 56, 213, 375, 171, 370, 214,
	152, 136, 137, 376, 347, 87, 292, 327, 141, 240,
	217, 292, 243, 340, 220, 86, 337, 245, 86, 224,
	185, 231, 229, 85, 150, 145, 85, 168, 335, 284,
	251, 367, 186, 187, 173, 232, 390, 343, 139, 176,
	228, 360, 304, 209, 210, 162, 361, 247, 71, 248,
	79, 147, 119, 260, 358, 262, 231, 273, 258, 359,
	83, 266, 343, 250, 271, 272, 392, 275, 276, 277,
	278, 279, 280, 281, 282, 267, 259, 391, 239, 241,
	238, 65, 66, 67, 63, 64, 68, 263, 119, 119,
	186, 187, 96, 86, 86, 149, 364, 296, 363, 362,
	274, 85, 298, 287, 289, 378, 147, 305, 264, 265,
	252, 299, 295, 197, 198, 199, 200, 201, 303, 400,
	86, 306, 380, 249, 311, 312, 227, 227, 85, 143,
	410, 409, 146, 329, 223, 408, 258, 310, 326, 295,
	313, 153, 331, 332, 199, 200, 201, 166, 222, 302,
	130, 98, 164, 100, 330, 134, 178, 88, 140, 223,
	19, 119, 180, 234, 218, 117, 131, 132, 133, 314,
	147, 216, 350, 339, 122, 349, 336, 346, 138, 257,
	345, 38, 39, 40, 41, 402, 403, 215, 256, 181,
	257, 142, 111, 72, 258, 258, 356, 357, 121, 256,
	179, 372, 136, 137, 115, 71, 374, 90, 87, 141,
	325, 371, 369, 377, 341, 353, 352, 309, 398, 86,
	246, 382, 184, 420, 385, 244, 324, 381, 194, 195,
	196, 197, 198, 199, 200, 201, 386, 225, 71, 139,
	174, 421, 113, 80, 161, 159, 396, 156, 151, 373,
	148, 397, 194, 195, 196, 197, 198, 199, 200, 201,
	95, 69, 176, 404, 144, 379, 399, 406, 316, 317,
	318, 319, 320, 385, 321, 322, 413, 412, 19, 415,
	415, 415, 86, 416, 417, 101, 418, 130, 78, 334,
	85, 103, 134, 19, 428, 140, 42, 102, 429, 427,
	430, 294, 117, 131, 132, 133, 233, 160, 130, 105,
	405, 122, 407, 134, 76, 138, 140, 44, 45, 46,
	47, 74, 106, 87, 131, 132, 133, 180, 268, 62,
	269, 270, 122, 389, 351, 121, 138, 301, 388, 136,
	137, 115, 355, 227, 108, 81, 141, 130, 426, 411,
	19, 43, 134, 18, 17, 140, 121, 16, 15, 14,
	136, 137, 87, 131, 132, 133, 13, 141, 12, 134,
	235, 122, 140, 50, 307, 138, 139, 237, 53, 87,
	131, 132, 133, 19, 20, 21, 22, 297, 166, 419,
	401, 134, 138, 383, 140, 121, 387, 139, 354, 136,
	137, 87, 131, 132, 133, 34, 141, 338, 230, 219,
	166, 23, 290, 129, 138, 126, 136, 137, 128, 344,
	123, 333, 188, 141, 194, 195, 196, 197, 198, 199,
	200, 201, 120, 365, 255, 315, 139, 253, 136, 137,
	116, 323, 182, 75, 37, 141, 77, 11, 10, 9,
	8, 7, 6, 139, 194, 195, 196, 197, 198, 199,
	200, 201, 28, 30, 31, 5, 29, 4, 2, 32,
	35, 36, 33, 328, 1, 139, 0, 0, 24, 25,
	27, 26, 189, 193, 191, 192, 0, 0, 0, 194,
	195, 196, 197, 198, 199, 200, 201, 0, 0, 0,
	0, 205, 206, 207, 208, 0, 202, 203, 204, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 190,
	194, 195, 196, 197, 198, 199, 200, 201,
}
var yyPact = [...]int{

	508, -1000, -1000, 262, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -45, -17, 20, -81, -1000, 24,
	-1000, -1000, -1000, 116, 356, 300, 286, 475, 434, -1000,
	-1000, -1000, 426, -1000, 389, 338, 466, 100, -66, -66,
	-1, 300, -1000, 17, 300, -1000, 355, -67, -67, 300,
	-67, -1000, 386, 404, 398, 429, 465, 300, 8, -1000,
	278, -1000, -1000, -1000, -1000, 260, -1000, 282, 338, 361,
	78, 338, 183, 345, -1000, 180, -1000, 77, 343, 62,
	338, 300, -1000, 342, -1000, -27, 340, 417, 339, 111,
	300, 338, 496, 496, -1000, 457, 496, 429, 496, 465,
	335, 496, 277, 275, -1000, -1000, 333, 73, 96, 591,
	-1000, 457, 418, -1000, -1000, -1000, 496, 273, 257, -1000,
	250, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 496, -1000, 245, 303, 332, 463, 303, -1000, 474,
	300, -1000, 416, 249, -76, -1000, 107, -1000, 320, 70,
	-1000, -1000, 315, -1000, 220, 515, 59, 515, 96, 591,
	515, -1000, 515, 429, -1000, 47, 515, -1000, 274, 397,
	496, -85, -1000, -1000, 300, 143, 457, 457, 496, 233,
	437, 496, 496, 162, 496, 496, 496, 496, 496, 496,
	496, 496, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-34, 46, -11, 591, -1000, 9, 397, -1000, 475, 56,
	515, 403, 303, 303, 247, -1000, 454, 457, -1000, 515,
	-1000, -1000, -1000, -1000, -1000, 108, 300, -1000, -58, -1000,
	-1000, -1000, -1000, -1000, -1000, 312, -1000, -1000, 403, 303,
	-1000, -1000, 496, 246, 344, 321, 285, 60, -1000, -1000,
	550, 442, -1000, -1000, -1000, -1000, 515, -1000, 233, 496,
	496, 515, 485, -1000, 394, 171, 171, 171, 200, 200,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 45, 397, 33,
	61, -1000, 457, 103, 233, 262, 128, 21, -1000, 454,
	442, 450, 96, -1000, 311, -1000, -1000, 310, -1000, -1000,
	-1000, 183, 515, 461, 274, 274, -1000, -1000, 130, 117,
	175, 174, 172, 53, -1000, 307, 15, 306, 496, -1000,
	-1000, 515, 313, 496, -1000, -1000, 13, -1000, 50, -1000,
	496, 154, -1000, 365, 199, -1000, -1000, -1000, 303, 442,
	-1000, 496, -1000, -1000, 456, 449, 344, 102, -1000, 153,
	-1000, 142, -1000, -1000, -1000, -1000, -12, -20, -43, -1000,
	-1000, -1000, 515, 496, 515, -1000, -1000, 515, 496, 317,
	233, -1000, -1000, 196, -1000, 289, -1000, 454, 457, 496,
	457, -1000, -1000, 221, 217, 216, 515, 515, 472, -1000,
	496, -1000, -1000, -1000, 442, 96, 187, 96, 300, 300,
	300, 303, -1000, 337, 10, -1000, -3, -23, 183, -1000,
	471, 408, -1000, 300, -1000, -1000, -1000, 300, -1000, 300,
	-1000,
}
var yyPgo = [...]int{

	0, 604, 598, 20, 597, 595, 582, 581, 580, 579,
	578, 577, 426, 576, 574, 573, 17, 26, 572, 571,
	570, 567, 11, 565, 564, 180, 563, 3, 16, 51,
	562, 552, 18, 550, 2, 15, 7, 549, 548, 9,
	545, 8, 543, 542, 12, 539, 537, 528, 526, 5,
	523, 6, 520, 1, 519, 14, 517, 10, 4, 22,
	222, 287, 508, 507, 504, 503, 500, 0, 13, 498,
	496, 489, 488, 487, 484, 483, 54, 19, 481,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 4, 4, 72, 72, 5, 6, 7, 7,
	69, 69, 70, 71, 74, 75, 73, 73, 73, 73,
	73, 73, 73, 8, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 11, 11, 78, 12, 13, 13,
	14, 14, 14, 14, 14, 15, 15, 16, 16, 17,
	17, 17, 20, 20, 18, 18, 18, 21, 21, 22,
	22, 22, 22, 19, 19, 19, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 24, 24, 24, 25, 25,
	26, 26, 26, 26, 27, 27, 28, 28, 77, 77,
	77, 76, 76, 29, 29, 29, 29, 29, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 31, 31,
	31, 31, 31, 31, 31, 32, 32, 37, 37, 35,
	35, 39, 36, 36, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 38, 38, 40, 40, 40, 42, 45, 45, 43,
	43, 44, 46, 46, 41, 41, 33, 33, 33, 33,
	47, 47, 48, 48, 49, 49, 50, 50, 51, 52,
	52, 52, 53, 53, 53, 54, 54, 54, 55, 55,
	56, 56, 57, 57, 58, 58, 59, 59, 60, 60,
	61, 61, 62, 62, 63, 63, 63, 63, 63, 64,
	64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 5, 7,
	12, 3, 7, 7, 6, 6, 8, 7, 3, 4,
	1, 2, 1, 1, 5, 2, 4, 4, 2, 3,
	4, 5, 4, 5, 6, 8, 4, 6, 7, 4,
	5, 6, 4, 4, 5, 5, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 0, 2,
	2, 0, 2, 1, 3, 3, 2, 3, 3, 3,
	4, 3, 4, 5, 6, 3, 4, 2, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 3, 3,
	1, 3, 1, 3, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 4, 5, 4,
	1, 1, 1, 1, 1, 1, 5, 0, 1, 1,
	2, 4, 0, 2, 1, 3, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 1, 3, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, -75, 5,
	6, 7, 8, 33, 100, 101, 103, 102, 84, 88,
	85, 86, 91, 94, 27, 92, 93, -14, 49, 50,
	51, 52, -12, -78, -12, -12, -12, -12, 112, 104,
	-65, 106, 110, -62, 106, 108, 104, 104, 112, 105,
	106, 89, -12, 98, 99, 95, 96, 97, 100, 35,
	-67, 35, 37, -3, 17, -15, 18, -13, 29, -25,
	35, 9, -58, 90, -59, -41, -67, 35, -61, 109,
	-61, 105, -67, 104, -67, 35, -60, 109, -60, -67,
	-60, 29, 23, 23, -77, 10, 23, -76, 9, -67,
	104, 44, -16, 112, -17, 74, -20, 35, -29, -34,
	-30, 68, 44, -33, -41, -35, -40, -67, -38, -42,
	20, 36, 37, 38, 25, -39, 72, 73, 48, 109,
	28, 79, 39, -25, 33, 77, -25, 53, 35, 45,
	77, 35, 68, -25, -67, -68, 35, -68, 107, 35,
	20, 35, 64, -67, -25, -34, 44, -34, -29, -34,
	-34, -77, -34, -76, 35, -36, -34, -53, 9, 53,
	15, 44, -18, -67, 19, 77, 66, 67, -31, 21,
	68, 23, 24, 22, 69, 70, 71, 72, 73, 74,
	75, 76, 45, 46, 47, 40, 41, 42, 43, -29,
	-29, -36, -3, -34, -34, 44, 44, -39, 44, -45,
	-34, -55, 33, 44, -58, 35, -28, 10, -59, -34,
	64, -67, -68, 20, 44, -66, 111, -63, 103, 101,
	32, 102, 13, 35, 35, 77, 35, -68, -55, 33,
	-77, 113, 53, -21, -22, -24, 44, 35, -39, -17,
	-34, 113, -67, 74, -29, -29, -34, -35, 21, 23,
	24, -34, -34, 25, 68, -34, -34, -34, -34, -34,
	-34, -34, -34, 113, 113, 113, 113, -16, 18, -16,
	-43, -44, 80, -32, 28, -3, -58, -56, -41, -28,
	-49, 13, -29, -68, 64, -67, -68, -64, 107, 35,
	-32, -58, -34, -28, 53, -23, 54, 55, 56, 57,
	58, 60, 61, -19, 35, 19, -22, 77, 53, -53,
	-35, -34, -34, 66, 25, 113, -16, 113, -46, -44,
	82, -29, -57, 64, -37, -35, -57, 113, 53, -49,
	-53, 14, 35, 35, -47, 11, -22, -22, 54, 59,
	54, 59, 54, 54, 54, -26, 62, 108, 63, 35,
	113, 35, -34, 66, -34, 113, 83, -34, 81, 30,
	53, -41, -53, -50, -51, -34, -68, -48, 12, 14,
	64, 54, 54, 105, 105, 105, -34, -34, 31, -35,
	53, -52, 26, 27, -49, -29, -36, -29, 44, 44,
	44, 7, -51, -53, -27, -67, -27, -27, -58, -54,
	16, 34, 113, 53, 113, 113, 7, 21, -67, -67,
	-67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 56,
	56, 56, 56, 56, 221, 212, 0, 0, 30, 0,
	32, 33, 56, 0, 0, 0, 0, 0, 60, 62,
	63, 64, 65, 58, 0, 0, 0, 0, 210, 210,
	0, 0, 222, 0, 0, 213, 0, 208, 208, 0,
	208, 31, 0, 0, 0, 108, 111, 0, 0, 38,
	0, 225, 35, 21, 61, 0, 66, 57, 0, 0,
	98, 0, 28, 0, 204, 0, 174, 225, 0, 0,
	0, 0, 226, 0, 226, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 39, 0, 0, 108, 0, 111,
	0, 0, 192, 0, 67, 69, 74, 225, 72, 73,
	113, 0, 0, 144, 145, 146, 0, 174, 0, 160,
	0, 176, 177, 178, 179, 140, 163, 164, 165, 161,
	162, 167, 59, 198, 0, 0, 106, 0, 29, 0,
	0, 226, 0, 0, 223, 46, 0, 49, 0, 53,
	209, 52, 0, 226, 198, 36, 0, 37, 109, 0,
	110, 40, 112, 108, 42, 0, 142, 18, 0, 0,
	0, 0, 70, 75, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 128, 129, 130, 131, 132, 133, 134, 116,
	0, 0, 0, 142, 155, 0, 0, 127, 0, 0,
	168, 0, 0, 0, 106, 99, 184, 0, 205, 206,
	207, 175, 43, 211, 226, 0, 0, 226, 219, 214,
	215, 216, 217, 218, 50, 0, 54, 55, 0, 0,
	41, 34, 0, 106, 77, 83, 0, 95, 97, 68,
	193, 192, 76, 71, 114, 115, 118, 119, 0, 0,
	0, 121, 0, 125, 0, 147, 148, 149, 150, 151,
	152, 153, 154, 117, 139, 141, 156, 0, 0, 0,
	172, 169, 0, 202, 0, 136, 202, 0, 200, 184,
	192, 0, 107, 44, 0, 224, 47, 0, 220, 51,
	24, 25, 143, 180, 0, 0, 86, 87, 0, 0,
	0, 0, 0, 100, 84, 0, 0, 0, 0, 19,
	120, 122, 0, 0, 126, 157, 0, 159, 0, 170,
	0, 0, 22, 0, 135, 137, 23, 199, 0, 192,
	27, 0, 226, 48, 182, 0, 78, 81, 88, 0,
	90, 0, 92, 93, 94, 79, 0, 0, 0, 85,
	80, 96, 194, 0, 123, 158, 166, 173, 0, 0,
	0, 201, 26, 185, 186, 189, 45, 184, 0, 0,
	0, 89, 91, 0, 0, 0, 124, 171, 0, 138,
	0, 188, 190, 191, 192, 183, 181, 82, 0, 0,
	0, 0, 187, 195, 0, 104, 0, 0, 203, 20,
	0, 0, 101, 0, 102, 103, 196, 0, 105, 0,
	197,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 76, 69, 3,
	44, 113, 74, 72, 53, 73, 77, 75, 3, 3,
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
	107, 108, 109, 110, 111, 112,
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
		//line sql.y:171
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:177
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:198
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:202
		{
			yyVAL.selStmt = &SimpleSelect{}
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:206
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, ExtraCols: NewSelectExprs(yyDollar[4].selectExprs, yyDollar[6].tableExprs), From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: NewGroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:210
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 22:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:217
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: NewIstRows(yyDollar[6].insRows), OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:221
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:233
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:237
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:250
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:256
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:267
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:276
		{
			yyVAL.statement = &Begin{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.statement = &Commit{}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:287
		{
			yyVAL.statement = &Rollback{}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:293
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Kill{Id: yyDollar[2].bytes}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:305
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:313
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:317
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:325
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:329
		{
			yyVAL.statement = &Show{Section: "show create table"}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:335
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes, Type: DATABASE}
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:339
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, TableName: yyDollar[4].tableName, Type: TABLE}
		}
	case 45:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:343
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:348
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:354
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:358
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:363
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:369
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:375
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:379
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:383
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:387
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:392
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:397
		{
			SetAllowComments(yylex, true)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:401
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:407
		{
			yyVAL.bytes2 = nil
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:411
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:417
		{
			yyVAL.str = AST_UNION
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:421
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:425
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:429
		{
			yyVAL.str = AST_EXCEPT
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:433
		{
			yyVAL.str = AST_INTERSECT
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:438
		{
			yyVAL.str = ""
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:442
		{
			yyVAL.str = AST_DISTINCT
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:452
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:458
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:462
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:466
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:472
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:476
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:481
		{
			yyVAL.bytes = nil
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:485
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:489
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:495
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:499
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:505
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:509
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:513
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:517
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:522
		{
			yyVAL.bytes = nil
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:526
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:530
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:536
		{
			yyVAL.str = AST_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:540
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:548
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:552
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:556
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:560
		{
			yyVAL.str = AST_JOIN
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:564
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:568
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:574
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:578
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:582
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:588
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:592
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:597
		{
			yyVAL.indexHints = nil
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:601
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:605
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 103:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:609
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:615
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:619
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = nil
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:633
		{
			yyVAL.expr = nil
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:637
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:641
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:646
		{
			yyVAL.valExpr = nil
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:650
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:675
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:679
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: NewConditionRight(yyDollar[3].tuple)}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: NewConditionRight(yyDollar[4].tuple)}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:695
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:699
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:703
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:707
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:711
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:717
		{
			yyVAL.str = AST_EQ
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:721
		{
			yyVAL.str = AST_LT
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:725
		{
			yyVAL.str = AST_GT
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:729
		{
			yyVAL.str = AST_LE
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:733
		{
			yyVAL.str = AST_GE
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:737
		{
			yyVAL.str = AST_NE
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:741
		{
			yyVAL.str = AST_NSE
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:747
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:751
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:757
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:761
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:767
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:777
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExprs = NewValExprs(yyDollar[1].valExpr)
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExprs = AppendValExpr(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:793
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:797
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:801
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:813
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:829
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:833
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:837
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
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:852
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:856
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 158:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:860
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:864
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:868
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:874
		{
			yyVAL.bytes = IF_BYTES
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:878
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:884
		{
			yyVAL.byt = AST_UPLUS
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:888
		{
			yyVAL.byt = AST_UMINUS
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:892
		{
			yyVAL.byt = AST_TILDA
		}
	case 166:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:898
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:903
		{
			yyVAL.valExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:907
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:913
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:917
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:923
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:928
		{
			yyVAL.valExpr = nil
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:932
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:938
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:942
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:952
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:956
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:960
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:966
		{
			yyVAL.valExprs = nil
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:975
		{
			yyVAL.boolExpr = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:979
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:984
		{
			yyVAL.orderBy = nil
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:988
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:994
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:998
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1009
		{
			yyVAL.str = AST_ASC
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1013
		{
			yyVAL.str = AST_ASC
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1017
		{
			yyVAL.str = AST_DESC
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1022
		{
			yyVAL.limit = nil
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1026
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.str = ""
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1043
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
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.columns = nil
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.updateExprs = nil
		}
	case 203:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.updateExprs = NewUpdateExprs(yyDollar[1].updateExpr)
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1095
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1139
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1153
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
