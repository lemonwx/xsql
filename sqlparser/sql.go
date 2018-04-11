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
const CREATE = 57422
const ALTER = 57423
const DROP = 57424
const RENAME = 57425
const TABLE = 57426
const INDEX = 57427
const VIEW = 57428
const TO = 57429
const IGNORE = 57430
const IF = 57431
const UNIQUE = 57432
const USING = 57433
const DATABASE = 57434

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

const yyNprod = 218
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 649

var yyAct = [...]int{

	113, 159, 151, 390, 73, 192, 360, 121, 110, 279,
	141, 234, 270, 272, 318, 111, 207, 99, 75, 202,
	193, 3, 167, 168, 217, 100, 293, 294, 295, 296,
	297, 92, 298, 299, 62, 399, 88, 399, 104, 35,
	36, 37, 38, 80, 51, 77, 52, 399, 286, 83,
	161, 324, 85, 76, 161, 64, 89, 54, 55, 56,
	116, 161, 97, 232, 262, 120, 232, 144, 126, 105,
	371, 370, 95, 369, 82, 103, 117, 118, 119, 46,
	346, 48, 84, 140, 108, 49, 53, 45, 124, 352,
	401, 148, 400, 304, 271, 166, 143, 152, 264, 155,
	136, 158, 398, 164, 157, 351, 323, 223, 107, 313,
	131, 194, 122, 123, 101, 195, 311, 138, 263, 127,
	78, 231, 342, 344, 198, 70, 221, 154, 366, 224,
	201, 77, 150, 319, 77, 205, 282, 212, 210, 76,
	167, 168, 76, 271, 125, 316, 190, 191, 213, 368,
	215, 252, 209, 59, 60, 61, 158, 147, 367, 227,
	180, 181, 182, 343, 105, 240, 241, 212, 238, 228,
	156, 63, 245, 74, 340, 250, 251, 339, 254, 255,
	256, 257, 258, 259, 260, 261, 246, 239, 230, 87,
	220, 222, 219, 338, 253, 129, 167, 168, 132, 105,
	105, 133, 208, 232, 77, 77, 243, 244, 275, 133,
	242, 354, 76, 277, 266, 268, 336, 149, 283, 208,
	319, 337, 278, 274, 178, 179, 180, 181, 182, 284,
	77, 334, 376, 160, 288, 289, 335, 356, 76, 162,
	79, 135, 287, 229, 238, 291, 90, 281, 303, 274,
	290, 307, 308, 128, 204, 237, 18, 19, 20, 21,
	386, 385, 133, 306, 236, 293, 294, 295, 296, 297,
	105, 298, 299, 384, 153, 305, 199, 161, 31, 18,
	197, 326, 315, 196, 22, 312, 98, 81, 325, 321,
	322, 175, 176, 177, 178, 179, 180, 181, 182, 238,
	238, 302, 203, 332, 333, 63, 396, 78, 348, 237,
	317, 347, 350, 204, 345, 378, 379, 301, 236, 353,
	35, 36, 37, 38, 397, 77, 329, 358, 165, 328,
	361, 226, 225, 357, 206, 27, 28, 29, 71, 362,
	30, 33, 32, 145, 63, 142, 23, 24, 26, 25,
	139, 137, 372, 134, 86, 58, 130, 373, 175, 176,
	177, 178, 179, 180, 181, 182, 374, 355, 158, 91,
	69, 382, 375, 380, 18, 247, 310, 248, 249, 361,
	403, 93, 389, 388, 214, 391, 391, 391, 77, 392,
	393, 267, 394, 116, 94, 146, 76, 273, 120, 18,
	404, 126, 65, 381, 405, 383, 406, 67, 103, 117,
	118, 119, 162, 39, 116, 365, 327, 108, 280, 120,
	364, 124, 126, 331, 208, 96, 72, 402, 387, 78,
	117, 118, 119, 41, 42, 43, 44, 18, 108, 40,
	17, 107, 124, 16, 57, 122, 123, 101, 15, 14,
	349, 13, 127, 175, 176, 177, 178, 179, 180, 181,
	182, 12, 107, 216, 47, 285, 122, 123, 218, 50,
	18, 116, 276, 127, 395, 377, 120, 125, 359, 126,
	363, 265, 330, 314, 200, 269, 78, 117, 118, 119,
	120, 115, 112, 126, 114, 108, 320, 109, 125, 124,
	78, 117, 118, 119, 169, 106, 341, 235, 292, 153,
	233, 102, 300, 124, 163, 66, 34, 68, 11, 107,
	10, 9, 8, 122, 123, 7, 6, 5, 4, 120,
	127, 2, 126, 1, 0, 0, 0, 122, 123, 78,
	117, 118, 119, 0, 127, 0, 0, 120, 153, 0,
	126, 0, 124, 0, 0, 125, 0, 78, 117, 118,
	119, 0, 0, 0, 0, 0, 153, 0, 211, 125,
	124, 0, 0, 0, 0, 0, 122, 123, 0, 0,
	0, 0, 0, 127, 0, 170, 174, 172, 173, 0,
	0, 0, 0, 0, 122, 123, 0, 0, 0, 0,
	0, 127, 0, 0, 186, 187, 188, 189, 125, 183,
	184, 185, 309, 0, 0, 175, 176, 177, 178, 179,
	180, 181, 182, 0, 0, 0, 125, 0, 0, 0,
	0, 0, 171, 175, 176, 177, 178, 179, 180, 181,
	182, 175, 176, 177, 178, 179, 180, 181, 182,
}
var yyPact = [...]int{

	251, -1000, -1000, 271, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -20, -57, -13, -42, -1000, -1000, -1000,
	-1000, 320, 61, 270, 432, 385, -1000, -1000, -1000, 389,
	-1000, 341, 303, 417, 85, -61, -61, -26, 270, -1000,
	-17, 270, -1000, 319, -68, 270, -68, 340, -1000, 371,
	416, 270, 242, -1000, -1000, -1000, 40, -1000, 214, 303,
	323, 33, 303, 148, 318, -1000, 196, -1000, 23, 316,
	49, 315, 270, -1000, 310, -1000, -35, 308, 375, 93,
	270, 303, -1000, 451, 522, 371, 522, 416, 522, 224,
	-1000, -1000, 309, 18, 74, 564, -1000, 451, 394, -1000,
	-1000, -1000, 522, 239, 236, -1000, 232, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 522, -1000, 269,
	272, 299, 414, 272, -1000, 504, 270, -1000, 364, -1000,
	-82, -1000, 94, -1000, 297, -1000, -1000, 296, -1000, 210,
	74, 564, 572, 465, -1000, 572, 371, 13, 572, -1000,
	220, 40, 522, -1000, -1000, 270, 136, 451, 451, 522,
	230, 354, 522, 522, 126, 522, 522, 522, 522, 522,
	522, 522, 522, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -44, 10, -10, 564, -1000, 373, 40, -1000, 432,
	14, 572, 369, 272, 272, 209, -1000, 405, 451, -1000,
	572, -1000, -1000, -1000, -1000, -1000, 72, 270, -1000, -54,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 369, 272,
	-1000, -1000, 522, 192, 211, 282, 274, 16, -1000, -1000,
	222, -1000, -1000, -1000, -1000, 572, -1000, 230, 522, 522,
	572, 546, -1000, 351, 152, 152, 152, 86, 86, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 8, 40, 1, 63,
	-1000, 451, 69, 230, 271, 156, -2, -1000, 405, 397,
	402, 74, 294, -1000, -1000, 291, -1000, -1000, 148, 572,
	412, 220, 220, -1000, -1000, 177, 162, 139, 123, 120,
	60, -1000, 279, -28, 276, 522, -1000, 572, 384, 522,
	-1000, -1000, -3, -1000, 6, -1000, 522, 130, -1000, 337,
	184, -1000, -1000, -1000, 272, 397, -1000, 522, -1000, -1000,
	408, 401, 211, 64, -1000, 104, -1000, 95, -1000, -1000,
	-1000, -1000, -27, -29, -30, -1000, -1000, -1000, 572, 522,
	572, -1000, -1000, 572, 522, 335, 230, -1000, -1000, 179,
	-1000, 289, -1000, 405, 451, 522, 451, -1000, -1000, 229,
	217, 216, 572, 572, 421, -1000, 522, -1000, -1000, -1000,
	397, 74, 150, 74, 270, 270, 270, 272, -1000, 290,
	-6, -1000, -16, -18, 148, -1000, 420, 359, -1000, 270,
	-1000, -1000, -1000, 270, -1000, 270, -1000,
}
var yyPgo = [...]int{

	0, 533, 531, 20, 528, 527, 526, 525, 522, 521,
	520, 518, 413, 517, 516, 515, 17, 25, 514, 512,
	511, 510, 11, 508, 507, 125, 506, 3, 16, 38,
	505, 504, 13, 497, 2, 15, 5, 496, 494, 7,
	492, 8, 491, 485, 12, 484, 483, 482, 480, 9,
	478, 6, 475, 1, 474, 19, 472, 14, 4, 18,
	189, 240, 469, 468, 465, 464, 463, 0, 10, 461,
	451, 449, 448, 443, 440, 72, 31, 439,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 73, 73, 8, 8, 8, 8,
	9, 9, 9, 10, 11, 11, 11, 77, 12, 13,
	13, 14, 14, 14, 14, 14, 15, 15, 16, 16,
	17, 17, 17, 20, 20, 18, 18, 18, 21, 21,
	22, 22, 22, 22, 19, 19, 19, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 24, 24, 24, 25,
	25, 26, 26, 26, 26, 27, 27, 28, 28, 76,
	76, 76, 75, 75, 29, 29, 29, 29, 29, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 31,
	31, 31, 31, 31, 31, 31, 32, 32, 37, 37,
	35, 35, 39, 36, 36, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 38, 38, 40, 40, 40, 42, 45, 45,
	43, 43, 44, 46, 46, 41, 41, 33, 33, 33,
	33, 47, 47, 48, 48, 49, 49, 50, 50, 51,
	52, 52, 52, 53, 53, 53, 54, 54, 54, 55,
	55, 56, 56, 57, 57, 58, 58, 59, 59, 60,
	60, 61, 61, 62, 62, 63, 63, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 2, 3, 4, 5, 5, 5, 8, 4,
	6, 7, 4, 5, 4, 5, 5, 0, 2, 0,
	2, 1, 2, 1, 1, 1, 0, 1, 1, 3,
	1, 2, 3, 1, 1, 0, 1, 2, 1, 3,
	3, 3, 3, 5, 0, 1, 2, 1, 1, 2,
	3, 2, 3, 2, 2, 2, 1, 3, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 0,
	2, 2, 0, 2, 1, 3, 3, 2, 3, 3,
	3, 4, 3, 4, 5, 6, 3, 4, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 3,
	3, 1, 3, 1, 3, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 4, 5,
	4, 1, 1, 1, 1, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 1, 1, 1,
	1, 0, 3, 0, 2, 0, 3, 1, 3, 2,
	0, 1, 1, 0, 2, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 1, 3, 3, 3, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 95, 96, 98, 97, 84, 85, 86,
	89, 27, 91, 90, -14, 49, 50, 51, 52, -12,
	-77, -12, -12, -12, -12, 107, 99, -65, 101, 105,
	-62, 101, 103, 99, 99, 100, 101, -12, 35, 92,
	93, 94, -67, 35, -3, 17, -15, 18, -13, 29,
	-25, 35, 9, -58, 88, -59, -41, -67, 35, -61,
	104, -61, 100, -67, 99, -67, 35, -60, 104, -67,
	-60, 29, -76, 10, 23, -75, 9, -67, 44, -16,
	-17, 74, -20, 35, -29, -34, -30, 68, 44, -33,
	-41, -35, -40, -67, -38, -42, 20, 36, 37, 38,
	25, -39, 72, 73, 48, 104, 28, 79, 39, -25,
	33, 77, -25, 53, 35, 45, 77, 35, 68, 35,
	-67, -68, 35, -68, 102, 35, 20, 64, -67, -25,
	-29, -34, -34, 44, -76, -34, -75, -36, -34, -53,
	9, 53, 15, -18, -67, 19, 77, 66, 67, -31,
	21, 68, 23, 24, 22, 69, 70, 71, 72, 73,
	74, 75, 76, 45, 46, 47, 40, 41, 42, 43,
	-29, -29, -36, -3, -34, -34, 44, 44, -39, 44,
	-45, -34, -55, 33, 44, -58, 35, -28, 10, -59,
	-34, 64, -67, -68, 20, -68, -66, 106, -63, 98,
	96, 32, 97, 13, 35, 35, 35, -68, -55, 33,
	-76, 108, 53, -21, -22, -24, 44, 35, -39, -17,
	-34, -67, 74, -29, -29, -34, -35, 21, 23, 24,
	-34, -34, 25, 68, -34, -34, -34, -34, -34, -34,
	-34, -34, 108, 108, 108, 108, -16, 18, -16, -43,
	-44, 80, -32, 28, -3, -58, -56, -41, -28, -49,
	13, -29, 64, -67, -68, -64, 102, -32, -58, -34,
	-28, 53, -23, 54, 55, 56, 57, 58, 60, 61,
	-19, 35, 19, -22, 77, 53, -35, -34, -34, 66,
	25, 108, -16, 108, -46, -44, 82, -29, -57, 64,
	-37, -35, -57, 108, 53, -49, -53, 14, 35, 35,
	-47, 11, -22, -22, 54, 59, 54, 59, 54, 54,
	54, -26, 62, 103, 63, 35, 108, 35, -34, 66,
	-34, 108, 83, -34, 81, 30, 53, -41, -53, -50,
	-51, -34, -68, -48, 12, 14, 64, 54, 54, 100,
	100, 100, -34, -34, 31, -35, 53, -52, 26, 27,
	-49, -29, -36, -29, 44, 44, 44, 7, -51, -53,
	-27, -67, -27, -27, -58, -54, 16, 34, 108, 53,
	108, 108, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 47, 47,
	47, 47, 47, 212, 203, 0, 0, 28, 29, 30,
	47, 0, 0, 0, 0, 51, 53, 54, 55, 56,
	49, 0, 0, 0, 0, 201, 201, 0, 0, 213,
	0, 0, 204, 0, 199, 0, 199, 0, 32, 99,
	102, 0, 0, 216, 19, 52, 0, 57, 48, 0,
	0, 89, 0, 26, 0, 195, 0, 165, 216, 0,
	0, 0, 0, 217, 0, 217, 0, 0, 0, 0,
	0, 0, 33, 0, 0, 99, 0, 102, 0, 183,
	58, 60, 65, 216, 63, 64, 104, 0, 0, 135,
	136, 137, 0, 165, 0, 151, 0, 167, 168, 169,
	170, 131, 154, 155, 156, 152, 153, 158, 50, 189,
	0, 0, 97, 0, 27, 0, 0, 217, 0, 217,
	214, 39, 0, 42, 0, 44, 200, 0, 217, 189,
	100, 0, 101, 0, 34, 103, 99, 0, 133, 17,
	0, 0, 0, 61, 66, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 119, 120, 121, 122, 123, 124, 125,
	107, 0, 0, 0, 133, 146, 0, 0, 118, 0,
	0, 159, 0, 0, 0, 97, 90, 175, 0, 196,
	197, 198, 166, 36, 202, 37, 0, 0, 217, 210,
	205, 206, 207, 208, 209, 43, 45, 46, 0, 0,
	35, 31, 0, 97, 68, 74, 0, 86, 88, 59,
	184, 67, 62, 105, 106, 109, 110, 0, 0, 0,
	112, 0, 116, 0, 138, 139, 140, 141, 142, 143,
	144, 145, 108, 130, 132, 147, 0, 0, 0, 163,
	160, 0, 193, 0, 127, 193, 0, 191, 175, 183,
	0, 98, 0, 215, 40, 0, 211, 22, 23, 134,
	171, 0, 0, 77, 78, 0, 0, 0, 0, 0,
	91, 75, 0, 0, 0, 0, 111, 113, 0, 0,
	117, 148, 0, 150, 0, 161, 0, 0, 20, 0,
	126, 128, 21, 190, 0, 183, 25, 0, 217, 41,
	173, 0, 69, 72, 79, 0, 81, 0, 83, 84,
	85, 70, 0, 0, 0, 76, 71, 87, 185, 0,
	114, 149, 157, 164, 0, 0, 0, 192, 24, 176,
	177, 180, 38, 175, 0, 0, 0, 80, 82, 0,
	0, 0, 115, 162, 0, 129, 0, 179, 181, 182,
	183, 174, 172, 73, 0, 0, 0, 0, 178, 186,
	0, 95, 0, 0, 194, 18, 0, 0, 92, 0,
	93, 94, 187, 0, 96, 0, 188,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 76, 69, 3,
	44, 108, 74, 72, 53, 73, 77, 75, 3, 3,
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
	107,
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
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:200
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:204
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:211
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:215
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:227
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:231
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:244
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:250
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:257
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:261
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:267
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:279
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:285
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:309
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:313
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:317
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:322
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:328
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:332
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:337
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:343
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:349
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:353
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:358
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:363
		{
			SetAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:367
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:373
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:377
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:383
		{
			yyVAL.str = AST_UNION
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = AST_EXCEPT
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = AST_INTERSECT
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:408
		{
			yyVAL.str = AST_DISTINCT
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:414
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:418
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:424
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:428
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:432
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:438
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:442
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:447
		{
			yyVAL.bytes = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:451
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:455
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:461
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:465
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:471
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:475
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:488
		{
			yyVAL.bytes = nil
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:492
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:496
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:502
		{
			yyVAL.str = AST_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:506
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:510
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:514
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:518
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:540
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:544
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:548
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:554
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:558
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:563
		{
			yyVAL.indexHints = nil
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:567
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:571
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:575
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:581
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:585
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:590
		{
			yyVAL.boolExpr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:594
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:599
		{
			yyVAL.expr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:603
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:607
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:612
		{
			yyVAL.valExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:616
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:623
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:627
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:631
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:635
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:641
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:645
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:649
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:683
		{
			yyVAL.str = AST_EQ
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:687
		{
			yyVAL.str = AST_LT
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:691
		{
			yyVAL.str = AST_GT
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:695
		{
			yyVAL.str = AST_LE
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:699
		{
			yyVAL.str = AST_GE
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_NE
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_NSE
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:713
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:717
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:727
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:733
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:737
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:743
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:749
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:753
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:759
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:771
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:803
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
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:818
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 149:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:840
		{
			yyVAL.bytes = IF_BYTES
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:844
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:850
		{
			yyVAL.byt = AST_UPLUS
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:854
		{
			yyVAL.byt = AST_UMINUS
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.byt = AST_TILDA
		}
	case 157:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:864
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = nil
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:879
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:883
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:889
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:894
		{
			yyVAL.valExpr = nil
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:904
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:908
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:914
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:918
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExprs = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:940
		{
			yyVAL.boolExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:944
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:949
		{
			yyVAL.orderBy = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:953
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:959
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:963
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:969
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:974
		{
			yyVAL.str = AST_ASC
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:978
		{
			yyVAL.str = AST_ASC
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.str = AST_DESC
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:987
		{
			yyVAL.limit = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:991
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:995
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.str = ""
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1004
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1008
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
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.columns = nil
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1040
		{
			yyVAL.updateExprs = nil
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1044
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1118
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
