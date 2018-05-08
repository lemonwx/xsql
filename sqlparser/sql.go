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
const SHOW = 57420
const DATABASES = 57421
const TABLES = 57422
const PROXY = 57423
const VARIABLES = 57424
const STATUS = 57425
const CREATE = 57426
const ALTER = 57427
const DROP = 57428
const RENAME = 57429
const TABLE = 57430
const INDEX = 57431
const VIEW = 57432
const TO = 57433
const IGNORE = 57434
const IF = 57435
const UNIQUE = 57436
const USING = 57437
const DATABASE = 57438

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

const yyNprod = 224
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 688

var yyAct = [...]int{

	122, 171, 164, 407, 78, 294, 377, 205, 119, 150,
	130, 335, 248, 285, 220, 120, 287, 107, 100, 80,
	416, 206, 3, 255, 215, 416, 109, 309, 310, 311,
	312, 313, 230, 314, 315, 67, 180, 181, 93, 416,
	85, 36, 37, 38, 39, 301, 82, 173, 113, 341,
	88, 173, 173, 90, 81, 153, 246, 69, 95, 55,
	57, 58, 47, 52, 49, 53, 105, 56, 50, 103,
	46, 246, 388, 282, 114, 125, 359, 361, 387, 418,
	129, 386, 277, 135, 417, 363, 87, 89, 149, 59,
	112, 126, 127, 128, 54, 369, 286, 158, 415, 117,
	152, 160, 162, 133, 279, 165, 368, 167, 340, 170,
	330, 328, 177, 320, 169, 278, 236, 180, 181, 83,
	207, 360, 166, 116, 208, 239, 179, 131, 132, 110,
	245, 145, 371, 140, 136, 234, 211, 147, 237, 214,
	82, 75, 68, 82, 218, 383, 225, 223, 81, 385,
	163, 81, 286, 336, 333, 297, 226, 157, 228, 193,
	194, 195, 222, 134, 170, 203, 204, 280, 241, 63,
	64, 65, 61, 62, 79, 168, 114, 254, 384, 256,
	225, 257, 357, 252, 242, 260, 356, 244, 265, 266,
	142, 269, 270, 271, 272, 273, 274, 275, 276, 261,
	253, 336, 267, 233, 235, 232, 191, 192, 193, 194,
	195, 355, 114, 114, 180, 181, 138, 82, 82, 141,
	142, 290, 246, 393, 373, 81, 292, 281, 283, 258,
	259, 298, 92, 293, 221, 353, 84, 289, 351, 159,
	354, 299, 251, 352, 82, 268, 144, 221, 304, 305,
	403, 250, 81, 36, 37, 38, 39, 322, 402, 303,
	401, 252, 306, 319, 289, 172, 324, 325, 137, 129,
	296, 174, 135, 18, 161, 243, 216, 307, 323, 83,
	126, 127, 128, 212, 86, 114, 217, 217, 161, 94,
	142, 96, 133, 210, 209, 175, 343, 106, 332, 342,
	329, 68, 339, 251, 338, 83, 364, 362, 224, 173,
	318, 178, 250, 395, 396, 413, 131, 132, 252, 252,
	349, 350, 139, 136, 365, 346, 317, 68, 345, 367,
	302, 240, 238, 414, 327, 334, 370, 219, 76, 156,
	154, 97, 82, 151, 375, 372, 148, 378, 146, 143,
	374, 91, 134, 66, 391, 379, 188, 189, 190, 191,
	192, 193, 194, 195, 309, 310, 311, 312, 313, 389,
	314, 315, 366, 74, 390, 188, 189, 190, 191, 192,
	193, 194, 195, 99, 101, 170, 397, 98, 18, 392,
	399, 262, 420, 263, 264, 227, 378, 102, 155, 406,
	405, 70, 408, 408, 408, 82, 409, 410, 174, 411,
	125, 288, 72, 81, 382, 129, 344, 421, 135, 295,
	40, 422, 381, 423, 348, 112, 126, 127, 128, 221,
	398, 125, 400, 104, 117, 77, 129, 419, 133, 135,
	42, 43, 44, 45, 404, 18, 112, 126, 127, 128,
	41, 17, 60, 16, 15, 117, 14, 13, 116, 133,
	12, 229, 131, 132, 110, 48, 300, 231, 51, 136,
	291, 412, 18, 394, 376, 380, 347, 331, 213, 116,
	284, 124, 121, 131, 132, 110, 123, 125, 337, 118,
	136, 182, 129, 115, 358, 135, 249, 308, 134, 247,
	111, 108, 83, 126, 127, 128, 316, 176, 125, 71,
	35, 117, 73, 129, 11, 133, 135, 10, 129, 134,
	9, 135, 8, 83, 126, 127, 128, 7, 83, 126,
	127, 128, 117, 6, 5, 116, 133, 161, 4, 131,
	132, 133, 2, 1, 0, 0, 136, 0, 18, 0,
	0, 0, 0, 0, 0, 0, 116, 0, 0, 0,
	131, 132, 0, 0, 0, 131, 132, 136, 129, 0,
	0, 135, 136, 0, 0, 134, 0, 0, 83, 126,
	127, 128, 18, 19, 20, 21, 0, 161, 0, 0,
	0, 133, 0, 0, 0, 0, 134, 0, 0, 0,
	0, 134, 0, 0, 33, 183, 187, 185, 186, 0,
	22, 0, 0, 0, 0, 131, 132, 0, 0, 0,
	0, 0, 136, 0, 199, 200, 201, 202, 0, 196,
	197, 198, 326, 0, 0, 188, 189, 190, 191, 192,
	193, 194, 195, 188, 189, 190, 191, 192, 193, 194,
	195, 134, 184, 188, 189, 190, 191, 192, 193, 194,
	195, 27, 29, 30, 321, 28, 0, 0, 31, 34,
	32, 0, 0, 0, 0, 0, 23, 24, 26, 25,
	188, 189, 190, 191, 192, 193, 194, 195,
}
var yyPact = [...]int{

	577, -1000, -1000, 204, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -41, -42, -9, -44, -1000, 0, -1000,
	-1000, -1000, 75, 318, 266, 440, 384, -1000, -1000, -1000,
	394, -1000, 344, 303, 426, 84, -68, -68, -18, 266,
	-1000, -16, 266, -1000, 316, -70, -70, 266, -70, -1000,
	312, 364, 360, 374, 424, 266, -1000, 253, -1000, -1000,
	-1000, 390, -1000, 229, 303, 289, 56, 303, 167, 314,
	-1000, 201, -1000, 54, 313, 69, 311, 266, -1000, 308,
	-1000, -51, 305, 378, 304, 93, 266, 303, 493, 493,
	-1000, 488, 493, 374, 493, 424, 493, 256, 251, -1000,
	-1000, 292, 49, 148, 584, -1000, 488, 467, -1000, -1000,
	-1000, 493, 250, 249, -1000, 239, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 493, -1000, 243, 270,
	302, 419, 270, -1000, 244, 266, -1000, 375, -1000, -78,
	-1000, 103, -1000, 297, 48, -1000, -1000, 296, -1000, 242,
	574, 543, 574, 148, 584, 574, -1000, 574, 374, 18,
	574, -1000, 207, 411, 493, -89, -1000, -1000, 266, 107,
	488, 488, 493, 230, 370, 493, 493, 177, 493, 493,
	493, 493, 493, 493, 493, 493, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -30, 3, -8, 584, -1000, 55,
	411, -1000, 440, 16, 574, 383, 270, 270, 237, -1000,
	406, 488, -1000, 574, -1000, -1000, -1000, -1000, -1000, 91,
	266, -1000, -61, -1000, -1000, -1000, -1000, -1000, -1000, 295,
	-1000, -1000, 383, 270, -1000, -1000, 493, 224, 310, 291,
	268, 36, -1000, -1000, 611, 393, -1000, -1000, -1000, -1000,
	574, -1000, 230, 493, 493, 574, 566, -1000, 309, 134,
	134, 134, 85, 85, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1, 411, -2, 72, -1000, 488, 89, 230, 204,
	137, -4, -1000, 406, 393, 402, 148, 293, -1000, -1000,
	290, -1000, -1000, -1000, 167, 574, 413, 207, 207, -1000,
	-1000, 184, 181, 157, 132, 128, 14, -1000, 272, -27,
	271, 493, -1000, -1000, 574, 306, 493, -1000, -1000, -6,
	-1000, 12, -1000, 493, 51, -1000, 315, 171, -1000, -1000,
	-1000, 270, 393, -1000, 493, -1000, -1000, 410, 400, 310,
	81, -1000, 124, -1000, 95, -1000, -1000, -1000, -1000, -23,
	-26, -32, -1000, -1000, -1000, 574, 493, 574, -1000, -1000,
	574, 493, 323, 230, -1000, -1000, 170, -1000, 287, -1000,
	406, 488, 493, 488, -1000, -1000, 216, 214, 206, 574,
	574, 437, -1000, 493, -1000, -1000, -1000, 393, 148, 169,
	148, 266, 266, 266, 270, -1000, 299, -14, -1000, -28,
	-33, 167, -1000, 430, 371, -1000, 266, -1000, -1000, -1000,
	266, -1000, 266, -1000,
}
var yyPgo = [...]int{

	0, 543, 542, 21, 538, 534, 533, 527, 522, 520,
	517, 514, 420, 512, 510, 509, 17, 26, 507, 506,
	500, 499, 12, 497, 496, 141, 494, 3, 14, 48,
	493, 491, 16, 489, 2, 15, 7, 488, 486, 10,
	482, 8, 481, 480, 13, 478, 477, 476, 475, 5,
	474, 6, 473, 1, 471, 24, 470, 11, 4, 19,
	232, 236, 468, 467, 466, 465, 461, 0, 9, 460,
	457, 456, 454, 453, 451, 69, 18, 450,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 4, 4, 72, 72, 5, 6, 7, 7, 69,
	69, 70, 71, 74, 73, 73, 73, 73, 73, 73,
	8, 8, 8, 8, 9, 9, 9, 10, 11, 11,
	11, 11, 11, 77, 12, 13, 13, 14, 14, 14,
	14, 14, 15, 15, 16, 16, 17, 17, 17, 20,
	20, 18, 18, 18, 21, 21, 22, 22, 22, 22,
	19, 19, 19, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 24, 24, 24, 25, 25, 26, 26, 26,
	26, 27, 27, 28, 28, 76, 76, 76, 75, 75,
	29, 29, 29, 29, 29, 30, 30, 30, 30, 30,
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
	2, 1, 1, 5, 4, 4, 2, 3, 4, 5,
	5, 5, 8, 4, 6, 7, 4, 5, 6, 4,
	4, 5, 5, 0, 2, 0, 2, 1, 2, 1,
	1, 1, 0, 1, 1, 3, 1, 2, 3, 1,
	1, 0, 1, 2, 1, 3, 3, 3, 3, 5,
	0, 1, 2, 1, 1, 2, 3, 2, 3, 2,
	2, 2, 1, 3, 1, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 0, 2, 2, 0, 2,
	1, 3, 3, 2, 3, 3, 3, 4, 3, 4,
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
	7, 8, 33, 99, 100, 102, 101, 84, 88, 85,
	86, 91, 93, 27, 92, -14, 49, 50, 51, 52,
	-12, -77, -12, -12, -12, -12, 111, 103, -65, 105,
	109, -62, 105, 107, 103, 103, 111, 104, 105, 89,
	-12, 97, 98, 94, 95, 96, 35, -67, 35, -3,
	17, -15, 18, -13, 29, -25, 35, 9, -58, 90,
	-59, -41, -67, 35, -61, 108, -61, 104, -67, 103,
	-67, 35, -60, 108, -60, -67, -60, 29, 23, 23,
	-76, 10, 23, -75, 9, -67, 44, -16, 111, -17,
	74, -20, 35, -29, -34, -30, 68, 44, -33, -41,
	-35, -40, -67, -38, -42, 20, 36, 37, 38, 25,
	-39, 72, 73, 48, 108, 28, 79, 39, -25, 33,
	77, -25, 53, 35, 45, 77, 35, 68, 35, -67,
	-68, 35, -68, 106, 35, 20, 35, 64, -67, -25,
	-34, 44, -34, -29, -34, -34, -76, -34, -75, -36,
	-34, -53, 9, 53, 15, 44, -18, -67, 19, 77,
	66, 67, -31, 21, 68, 23, 24, 22, 69, 70,
	71, 72, 73, 74, 75, 76, 45, 46, 47, 40,
	41, 42, 43, -29, -29, -36, -3, -34, -34, 44,
	44, -39, 44, -45, -34, -55, 33, 44, -58, 35,
	-28, 10, -59, -34, 64, -67, -68, 20, -68, -66,
	110, -63, 102, 100, 32, 101, 13, 35, 35, 77,
	35, -68, -55, 33, -76, 112, 53, -21, -22, -24,
	44, 35, -39, -17, -34, 112, -67, 74, -29, -29,
	-34, -35, 21, 23, 24, -34, -34, 25, 68, -34,
	-34, -34, -34, -34, -34, -34, -34, 112, 112, 112,
	112, -16, 18, -16, -43, -44, 80, -32, 28, -3,
	-58, -56, -41, -28, -49, 13, -29, 64, -67, -68,
	-64, 106, 35, -32, -58, -34, -28, 53, -23, 54,
	55, 56, 57, 58, 60, 61, -19, 35, 19, -22,
	77, 53, -53, -35, -34, -34, 66, 25, 112, -16,
	112, -46, -44, 82, -29, -57, 64, -37, -35, -57,
	112, 53, -49, -53, 14, 35, 35, -47, 11, -22,
	-22, 54, 59, 54, 59, 54, 54, 54, -26, 62,
	107, 63, 35, 112, 35, -34, 66, -34, 112, 83,
	-34, 81, 30, 53, -41, -53, -50, -51, -34, -68,
	-48, 12, 14, 64, 54, 54, 104, 104, 104, -34,
	-34, 31, -35, 53, -52, 26, 27, -49, -29, -36,
	-29, 44, 44, 44, 7, -51, -53, -27, -67, -27,
	-27, -58, -54, 16, 34, 112, 53, 112, 112, 7,
	21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 53, 53,
	53, 53, 53, 218, 209, 0, 0, 29, 0, 31,
	32, 53, 0, 0, 0, 0, 57, 59, 60, 61,
	62, 55, 0, 0, 0, 0, 207, 207, 0, 0,
	219, 0, 0, 210, 0, 205, 205, 0, 205, 30,
	0, 0, 0, 105, 108, 0, 36, 0, 222, 20,
	58, 0, 63, 54, 0, 0, 95, 0, 27, 0,
	201, 0, 171, 222, 0, 0, 0, 0, 223, 0,
	223, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	37, 0, 0, 105, 0, 108, 0, 189, 0, 64,
	66, 71, 222, 69, 70, 110, 0, 0, 141, 142,
	143, 0, 171, 0, 157, 0, 173, 174, 175, 176,
	137, 160, 161, 162, 158, 159, 164, 56, 195, 0,
	0, 103, 0, 28, 0, 0, 223, 0, 223, 220,
	43, 0, 46, 0, 50, 206, 49, 0, 223, 195,
	34, 0, 35, 106, 0, 107, 38, 109, 105, 0,
	139, 17, 0, 0, 0, 0, 67, 72, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 125, 126, 127, 128,
	129, 130, 131, 113, 0, 0, 0, 139, 152, 0,
	0, 124, 0, 0, 165, 0, 0, 0, 103, 96,
	181, 0, 202, 203, 204, 172, 40, 208, 41, 0,
	0, 223, 216, 211, 212, 213, 214, 215, 47, 0,
	51, 52, 0, 0, 39, 33, 0, 103, 74, 80,
	0, 92, 94, 65, 190, 189, 73, 68, 111, 112,
	115, 116, 0, 0, 0, 118, 0, 122, 0, 144,
	145, 146, 147, 148, 149, 150, 151, 114, 136, 138,
	153, 0, 0, 0, 169, 166, 0, 199, 0, 133,
	199, 0, 197, 181, 189, 0, 104, 0, 221, 44,
	0, 217, 48, 23, 24, 140, 177, 0, 0, 83,
	84, 0, 0, 0, 0, 0, 97, 81, 0, 0,
	0, 0, 18, 117, 119, 0, 0, 123, 154, 0,
	156, 0, 167, 0, 0, 21, 0, 132, 134, 22,
	196, 0, 189, 26, 0, 223, 45, 179, 0, 75,
	78, 85, 0, 87, 0, 89, 90, 91, 76, 0,
	0, 0, 82, 77, 93, 191, 0, 120, 155, 163,
	170, 0, 0, 0, 198, 25, 182, 183, 186, 42,
	181, 0, 0, 0, 86, 88, 0, 0, 0, 121,
	168, 0, 135, 0, 185, 187, 188, 189, 180, 178,
	79, 0, 0, 0, 0, 184, 192, 0, 101, 0,
	0, 200, 19, 0, 0, 98, 0, 99, 100, 193,
	0, 102, 0, 194,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 76, 69, 3,
	44, 112, 74, 72, 53, 73, 77, 75, 3, 3,
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
	107, 108, 109, 110, 111,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:274
		{
			yyVAL.statement = &Begin{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &Commit{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &Rollback{}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:297
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:301
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:305
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:313
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:317
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:323
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:327
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:331
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:342
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:346
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:351
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:357
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:363
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:367
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:371
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:375
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:380
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:385
		{
			SetAllowComments(yylex, true)
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:389
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:395
		{
			yyVAL.bytes2 = nil
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:399
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:405
		{
			yyVAL.str = AST_UNION
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:409
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:413
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:417
		{
			yyVAL.str = AST_EXCEPT
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.str = AST_INTERSECT
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:426
		{
			yyVAL.str = ""
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.str = AST_DISTINCT
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:436
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:446
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:450
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:454
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:460
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:464
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:469
		{
			yyVAL.bytes = nil
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:473
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:477
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:493
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:497
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:501
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:505
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:510
		{
			yyVAL.bytes = nil
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:514
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:518
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:524
		{
			yyVAL.str = AST_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:528
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:532
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:536
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:544
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:548
		{
			yyVAL.str = AST_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:552
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:556
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:562
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:566
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:570
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:576
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:580
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:585
		{
			yyVAL.indexHints = nil
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:589
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:593
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:597
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:603
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:607
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:612
		{
			yyVAL.boolExpr = nil
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:616
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:621
		{
			yyVAL.expr = nil
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:625
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:629
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:634
		{
			yyVAL.valExpr = nil
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:638
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:645
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:649
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:653
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:663
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:667
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:671
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:675
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:679
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:695
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:699
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:705
		{
			yyVAL.str = AST_EQ
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:709
		{
			yyVAL.str = AST_LT
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:713
		{
			yyVAL.str = AST_GT
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:717
		{
			yyVAL.str = AST_LE
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:721
		{
			yyVAL.str = AST_GE
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:725
		{
			yyVAL.str = AST_NE
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:729
		{
			yyVAL.str = AST_NSE
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:735
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:739
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:745
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:749
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:755
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:759
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:765
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.valExprs = NewValExprs(yyDollar[1].valExpr)
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:781
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:785
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:789
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:793
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:797
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:801
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:813
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:825
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
		//line sql.y:840
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:844
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:848
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:852
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:856
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:862
		{
			yyVAL.bytes = IF_BYTES
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:866
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:872
		{
			yyVAL.byt = AST_UPLUS
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:876
		{
			yyVAL.byt = AST_UMINUS
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:880
		{
			yyVAL.byt = AST_TILDA
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:886
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:891
		{
			yyVAL.valExpr = nil
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:895
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:901
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:905
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:911
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:916
		{
			yyVAL.valExpr = nil
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:930
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:936
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:954
		{
			yyVAL.valExprs = nil
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:958
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:963
		{
			yyVAL.boolExpr = nil
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:967
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:972
		{
			yyVAL.orderBy = nil
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:976
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:986
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:992
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:997
		{
			yyVAL.str = AST_ASC
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.str = AST_ASC
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.str = AST_DESC
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.limit = nil
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1018
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1023
		{
			yyVAL.str = ""
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1031
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
		//line sql.y:1044
		{
			yyVAL.columns = nil
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.updateExprs = nil
		}
	case 200:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.updateExprs = NewUpdateExprs(yyDollar[1].updateExpr)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1095
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 223:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1141
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
