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

const yyNprod = 225
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 746

var yyAct = [...]int{

	124, 174, 166, 410, 79, 297, 380, 208, 121, 152,
	132, 338, 251, 288, 223, 122, 290, 109, 101, 218,
	81, 209, 3, 183, 184, 258, 111, 312, 313, 314,
	315, 316, 233, 317, 318, 68, 36, 37, 38, 39,
	94, 55, 57, 58, 86, 52, 83, 53, 115, 56,
	89, 419, 104, 91, 82, 304, 419, 70, 96, 47,
	419, 49, 176, 391, 155, 50, 106, 46, 344, 280,
	390, 389, 176, 176, 249, 116, 127, 249, 88, 107,
	90, 131, 54, 59, 137, 366, 289, 372, 336, 151,
	289, 114, 128, 129, 130, 362, 364, 84, 160, 282,
	119, 154, 162, 164, 135, 323, 167, 242, 169, 182,
	421, 173, 147, 142, 180, 420, 172, 69, 149, 418,
	270, 371, 210, 168, 118, 239, 211, 343, 133, 134,
	112, 333, 331, 281, 386, 138, 248, 339, 214, 300,
	363, 217, 83, 76, 237, 83, 221, 240, 228, 226,
	82, 165, 80, 82, 183, 184, 260, 159, 229, 170,
	231, 398, 399, 271, 136, 225, 173, 206, 207, 374,
	244, 63, 64, 65, 61, 62, 66, 183, 184, 116,
	257, 245, 259, 228, 93, 356, 255, 388, 263, 247,
	357, 268, 269, 387, 272, 273, 274, 275, 276, 277,
	278, 279, 264, 256, 191, 192, 193, 194, 195, 196,
	197, 198, 236, 238, 235, 116, 116, 360, 354, 140,
	83, 83, 143, 355, 293, 196, 197, 198, 82, 295,
	284, 286, 261, 262, 301, 144, 296, 359, 358, 144,
	292, 95, 161, 97, 302, 249, 339, 83, 175, 396,
	224, 307, 308, 224, 177, 82, 36, 37, 38, 39,
	325, 376, 306, 146, 255, 309, 322, 292, 85, 327,
	328, 406, 131, 299, 18, 137, 194, 195, 196, 197,
	198, 326, 84, 128, 129, 130, 246, 139, 116, 254,
	405, 163, 176, 310, 404, 135, 144, 220, 253, 346,
	163, 335, 345, 332, 254, 342, 219, 341, 215, 213,
	212, 227, 178, 253, 108, 321, 87, 220, 181, 133,
	134, 255, 255, 352, 353, 141, 138, 368, 69, 84,
	416, 320, 370, 367, 69, 365, 349, 348, 337, 373,
	305, 243, 241, 222, 98, 83, 171, 378, 417, 394,
	381, 77, 158, 377, 156, 136, 369, 153, 382, 191,
	192, 193, 194, 195, 196, 197, 198, 312, 313, 314,
	315, 316, 392, 317, 318, 329, 150, 393, 191, 192,
	193, 194, 195, 196, 197, 198, 148, 145, 173, 400,
	92, 67, 395, 402, 375, 75, 18, 330, 265, 381,
	266, 267, 409, 408, 102, 411, 411, 411, 83, 412,
	413, 285, 414, 127, 73, 100, 82, 103, 131, 291,
	424, 137, 99, 40, 425, 423, 426, 230, 114, 128,
	129, 130, 157, 401, 127, 403, 71, 119, 177, 131,
	385, 135, 137, 42, 43, 44, 45, 347, 298, 114,
	128, 129, 130, 384, 351, 60, 224, 105, 119, 78,
	422, 118, 135, 407, 18, 133, 134, 112, 41, 17,
	16, 15, 138, 14, 13, 12, 18, 232, 48, 303,
	234, 51, 118, 294, 415, 397, 133, 134, 112, 379,
	383, 127, 350, 138, 334, 216, 131, 287, 126, 137,
	123, 136, 125, 340, 120, 283, 84, 128, 129, 130,
	185, 117, 127, 361, 252, 119, 311, 131, 250, 135,
	137, 113, 136, 319, 179, 110, 72, 84, 128, 129,
	130, 35, 74, 11, 10, 9, 119, 8, 7, 118,
	135, 6, 5, 133, 134, 4, 2, 1, 0, 0,
	138, 0, 18, 0, 0, 0, 0, 0, 0, 0,
	118, 0, 0, 0, 133, 134, 0, 0, 0, 0,
	0, 138, 131, 0, 0, 137, 0, 0, 0, 136,
	0, 0, 84, 128, 129, 130, 131, 0, 0, 137,
	0, 163, 0, 0, 324, 135, 84, 128, 129, 130,
	136, 18, 19, 20, 21, 163, 0, 0, 0, 135,
	191, 192, 193, 194, 195, 196, 197, 198, 0, 133,
	134, 0, 0, 33, 0, 0, 138, 0, 0, 22,
	0, 0, 0, 133, 134, 0, 0, 0, 0, 0,
	138, 191, 192, 193, 194, 195, 196, 197, 198, 0,
	0, 0, 0, 0, 0, 136, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 136,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	27, 29, 30, 0, 28, 0, 0, 31, 34, 32,
	186, 190, 188, 189, 0, 23, 24, 26, 25, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 202,
	203, 204, 205, 0, 199, 200, 201, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 187, 191, 192,
	193, 194, 195, 196, 197, 198,
}
var yyPact = [...]int{

	596, -1000, -1000, 207, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -44, -60, -21, -62, -1000, -6, -1000,
	-1000, -1000, 77, 356, 293, 459, 419, -1000, -1000, -1000,
	396, -1000, 366, 316, 450, 62, -64, -64, -26, 293,
	-1000, -23, 293, -1000, 355, -68, -68, 293, -68, -1000,
	315, 399, 392, 394, 448, 293, -24, -1000, 270, -1000,
	-1000, -1000, 414, -1000, 248, 316, 292, 36, 316, 186,
	352, -1000, 218, -1000, 35, 351, 50, 341, 293, -1000,
	322, -1000, -42, 319, 412, 317, 93, 293, 316, 561,
	561, -1000, 492, 561, 394, 561, 448, 311, 561, 239,
	268, -1000, -1000, 299, 32, 111, 669, -1000, 492, 471,
	-1000, -1000, -1000, 561, 266, 265, -1000, 264, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 561, -1000,
	273, 294, 308, 446, 294, -1000, 247, 293, -1000, 407,
	-1000, -78, -1000, 112, -1000, 307, 30, -1000, -1000, 306,
	-1000, 253, 572, 547, 572, 111, 669, 572, -1000, 572,
	394, -1000, 24, 572, -1000, 254, 56, 561, -87, -1000,
	-1000, 293, 82, 492, 492, 561, 256, 377, 561, 561,
	95, 561, 561, 561, 561, 561, 561, 561, 561, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -43, 21, -13,
	669, -1000, 393, 56, -1000, 459, 10, 572, 391, 294,
	294, 243, -1000, 435, 492, -1000, 572, -1000, -1000, -1000,
	-1000, -1000, 75, 293, -1000, -51, -1000, -1000, -1000, -1000,
	-1000, -1000, 305, -1000, -1000, 391, 294, -1000, -1000, 561,
	240, 313, 296, 269, 28, -1000, -1000, 541, 423, -1000,
	-1000, -1000, -1000, 572, -1000, 256, 561, 561, 572, 309,
	-1000, 372, 204, 204, 204, 151, 151, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 20, 56, 19, 6, -1000, 492,
	73, 256, 207, 182, 15, -1000, 435, 423, 433, 111,
	302, -1000, -1000, 301, -1000, -1000, -1000, 186, 572, 443,
	254, 254, -1000, -1000, 164, 131, 184, 183, 163, 33,
	-1000, 300, -27, 298, 561, -1000, -1000, 572, 290, 561,
	-1000, -1000, 9, -1000, 4, -1000, 561, 88, -1000, 364,
	208, -1000, -1000, -1000, 294, 423, -1000, 561, -1000, -1000,
	441, 426, 313, 70, -1000, 139, -1000, 133, -1000, -1000,
	-1000, -1000, -33, -34, -41, -1000, -1000, -1000, 572, 561,
	572, -1000, -1000, 572, 561, 318, 256, -1000, -1000, 196,
	-1000, 135, -1000, 435, 492, 561, 492, -1000, -1000, 250,
	246, 227, 572, 572, 456, -1000, 561, -1000, -1000, -1000,
	423, 111, 192, 111, 293, 293, 293, 294, -1000, 314,
	7, -1000, 3, -2, 186, -1000, 453, 404, -1000, 293,
	-1000, -1000, -1000, 293, -1000, 293, -1000,
}
var yyPgo = [...]int{

	0, 547, 546, 21, 545, 542, 541, 538, 537, 535,
	534, 533, 423, 532, 531, 526, 17, 26, 524, 523,
	521, 518, 12, 516, 514, 143, 513, 3, 14, 48,
	511, 510, 16, 504, 2, 15, 7, 503, 502, 10,
	500, 8, 498, 497, 13, 495, 494, 492, 490, 5,
	489, 6, 485, 1, 484, 19, 483, 11, 4, 20,
	184, 268, 481, 480, 479, 478, 477, 0, 9, 475,
	474, 473, 471, 470, 469, 52, 18, 468,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 4, 4, 72, 72, 5, 6, 7, 7, 69,
	69, 70, 71, 74, 73, 73, 73, 73, 73, 73,
	73, 8, 8, 8, 8, 9, 9, 9, 10, 11,
	11, 11, 11, 11, 77, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	20, 20, 18, 18, 18, 21, 21, 22, 22, 22,
	22, 19, 19, 19, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 76, 76, 76, 75,
	75, 29, 29, 29, 29, 29, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 37, 37, 35, 35, 39,
	36, 36, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 38,
	38, 40, 40, 40, 42, 45, 45, 43, 43, 44,
	46, 46, 41, 41, 33, 33, 33, 33, 47, 47,
	48, 48, 49, 49, 50, 50, 51, 52, 52, 52,
	53, 53, 53, 54, 54, 54, 55, 55, 56, 56,
	57, 57, 58, 58, 59, 59, 60, 60, 61, 61,
	62, 62, 63, 63, 63, 63, 63, 64, 64, 65,
	65, 66, 66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 7, 12,
	3, 7, 7, 6, 6, 8, 7, 3, 4, 1,
	2, 1, 1, 5, 4, 4, 2, 3, 4, 5,
	4, 5, 5, 8, 4, 6, 7, 4, 5, 6,
	4, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 0, 2, 2, 0,
	2, 1, 3, 3, 2, 3, 3, 3, 4, 3,
	4, 5, 6, 3, 4, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 3, 3, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 1, 3, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 99, 100, 102, 101, 84, 88, 85,
	86, 91, 93, 27, 92, -14, 49, 50, 51, 52,
	-12, -77, -12, -12, -12, -12, 111, 103, -65, 105,
	109, -62, 105, 107, 103, 103, 111, 104, 105, 89,
	-12, 97, 98, 94, 95, 96, 99, 35, -67, 35,
	-3, 17, -15, 18, -13, 29, -25, 35, 9, -58,
	90, -59, -41, -67, 35, -61, 108, -61, 104, -67,
	103, -67, 35, -60, 108, -60, -67, -60, 29, 23,
	23, -76, 10, 23, -75, 9, -67, 103, 44, -16,
	111, -17, 74, -20, 35, -29, -34, -30, 68, 44,
	-33, -41, -35, -40, -67, -38, -42, 20, 36, 37,
	38, 25, -39, 72, 73, 48, 108, 28, 79, 39,
	-25, 33, 77, -25, 53, 35, 45, 77, 35, 68,
	35, -67, -68, 35, -68, 106, 35, 20, 35, 64,
	-67, -25, -34, 44, -34, -29, -34, -34, -76, -34,
	-75, 35, -36, -34, -53, 9, 53, 15, 44, -18,
	-67, 19, 77, 66, 67, -31, 21, 68, 23, 24,
	22, 69, 70, 71, 72, 73, 74, 75, 76, 45,
	46, 47, 40, 41, 42, 43, -29, -29, -36, -3,
	-34, -34, 44, 44, -39, 44, -45, -34, -55, 33,
	44, -58, 35, -28, 10, -59, -34, 64, -67, -68,
	20, -68, -66, 110, -63, 102, 100, 32, 101, 13,
	35, 35, 77, 35, -68, -55, 33, -76, 112, 53,
	-21, -22, -24, 44, 35, -39, -17, -34, 112, -67,
	74, -29, -29, -34, -35, 21, 23, 24, -34, -34,
	25, 68, -34, -34, -34, -34, -34, -34, -34, -34,
	112, 112, 112, 112, -16, 18, -16, -43, -44, 80,
	-32, 28, -3, -58, -56, -41, -28, -49, 13, -29,
	64, -67, -68, -64, 106, 35, -32, -58, -34, -28,
	53, -23, 54, 55, 56, 57, 58, 60, 61, -19,
	35, 19, -22, 77, 53, -53, -35, -34, -34, 66,
	25, 112, -16, 112, -46, -44, 82, -29, -57, 64,
	-37, -35, -57, 112, 53, -49, -53, 14, 35, 35,
	-47, 11, -22, -22, 54, 59, 54, 59, 54, 54,
	54, -26, 62, 107, 63, 35, 112, 35, -34, 66,
	-34, 112, 83, -34, 81, 30, 53, -41, -53, -50,
	-51, -34, -68, -48, 12, 14, 64, 54, 54, 104,
	104, 104, -34, -34, 31, -35, 53, -52, 26, 27,
	-49, -29, -36, -29, 44, 44, 44, 7, -51, -53,
	-27, -67, -27, -27, -58, -54, 16, 34, 112, 53,
	112, 112, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 54, 54,
	54, 54, 54, 219, 210, 0, 0, 29, 0, 31,
	32, 54, 0, 0, 0, 0, 58, 60, 61, 62,
	63, 56, 0, 0, 0, 0, 208, 208, 0, 0,
	220, 0, 0, 211, 0, 206, 206, 0, 206, 30,
	0, 0, 0, 106, 109, 0, 0, 36, 0, 223,
	20, 59, 0, 64, 55, 0, 0, 96, 0, 27,
	0, 202, 0, 172, 223, 0, 0, 0, 0, 224,
	0, 224, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 37, 0, 0, 106, 0, 109, 0, 0, 190,
	0, 65, 67, 72, 223, 70, 71, 111, 0, 0,
	142, 143, 144, 0, 172, 0, 158, 0, 174, 175,
	176, 177, 138, 161, 162, 163, 159, 160, 165, 57,
	196, 0, 0, 104, 0, 28, 0, 0, 224, 0,
	224, 221, 44, 0, 47, 0, 51, 207, 50, 0,
	224, 196, 34, 0, 35, 107, 0, 108, 38, 110,
	106, 40, 0, 140, 17, 0, 0, 0, 0, 68,
	73, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 126,
	127, 128, 129, 130, 131, 132, 114, 0, 0, 0,
	140, 153, 0, 0, 125, 0, 0, 166, 0, 0,
	0, 104, 97, 182, 0, 203, 204, 205, 173, 41,
	209, 42, 0, 0, 224, 217, 212, 213, 214, 215,
	216, 48, 0, 52, 53, 0, 0, 39, 33, 0,
	104, 75, 81, 0, 93, 95, 66, 191, 190, 74,
	69, 112, 113, 116, 117, 0, 0, 0, 119, 0,
	123, 0, 145, 146, 147, 148, 149, 150, 151, 152,
	115, 137, 139, 154, 0, 0, 0, 170, 167, 0,
	200, 0, 134, 200, 0, 198, 182, 190, 0, 105,
	0, 222, 45, 0, 218, 49, 23, 24, 141, 178,
	0, 0, 84, 85, 0, 0, 0, 0, 0, 98,
	82, 0, 0, 0, 0, 18, 118, 120, 0, 0,
	124, 155, 0, 157, 0, 168, 0, 0, 21, 0,
	133, 135, 22, 197, 0, 190, 26, 0, 224, 46,
	180, 0, 76, 79, 86, 0, 88, 0, 90, 91,
	92, 77, 0, 0, 0, 83, 78, 94, 192, 0,
	121, 156, 164, 171, 0, 0, 0, 199, 25, 183,
	184, 187, 43, 182, 0, 0, 0, 87, 89, 0,
	0, 0, 122, 169, 0, 136, 0, 186, 188, 189,
	190, 181, 179, 80, 0, 0, 0, 0, 185, 193,
	0, 102, 0, 0, 201, 19, 0, 0, 99, 0,
	100, 101, 194, 0, 103, 0, 195,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &Show{Section: "show create table"}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:327
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:331
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:335
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:346
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:350
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:355
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:361
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:367
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:371
		{
			yyVAL.statement = &DDL{Action: AST_DROP}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:375
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:379
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:384
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:389
		{
			SetAllowComments(yylex, true)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:393
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:399
		{
			yyVAL.bytes2 = nil
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:403
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:409
		{
			yyVAL.str = AST_UNION
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:413
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:417
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.str = AST_EXCEPT
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:425
		{
			yyVAL.str = AST_INTERSECT
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:430
		{
			yyVAL.str = ""
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.str = AST_DISTINCT
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:444
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:450
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:454
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:458
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:464
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:468
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:473
		{
			yyVAL.bytes = nil
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:477
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:481
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:497
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:501
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:505
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:509
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:514
		{
			yyVAL.bytes = nil
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:518
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:522
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:528
		{
			yyVAL.str = AST_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:532
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:536
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:540
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:548
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:552
		{
			yyVAL.str = AST_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:556
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:560
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:566
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:570
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:574
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:580
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:584
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:589
		{
			yyVAL.indexHints = nil
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:593
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:597
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:601
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:607
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:611
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:616
		{
			yyVAL.boolExpr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:620
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:625
		{
			yyVAL.expr = nil
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:629
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:633
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:638
		{
			yyVAL.valExpr = nil
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:642
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:649
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:667
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:671
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: NewConditionRight(yyDollar[3].tuple)}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:675
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: NewConditionRight(yyDollar[4].tuple)}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:679
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:695
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:699
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:703
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:709
		{
			yyVAL.str = AST_EQ
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:713
		{
			yyVAL.str = AST_LT
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:717
		{
			yyVAL.str = AST_GT
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:721
		{
			yyVAL.str = AST_LE
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:725
		{
			yyVAL.str = AST_GE
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:729
		{
			yyVAL.str = AST_NE
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:733
		{
			yyVAL.str = AST_NSE
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:739
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:743
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:749
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:753
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:759
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:769
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExprs = NewValExprs(yyDollar[1].valExpr)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:785
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:789
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:793
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:797
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:801
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:813
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:829
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
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:844
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:848
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 156:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:852
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:856
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:860
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:866
		{
			yyVAL.bytes = IF_BYTES
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:870
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:876
		{
			yyVAL.byt = AST_UPLUS
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:880
		{
			yyVAL.byt = AST_UMINUS
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:884
		{
			yyVAL.byt = AST_TILDA
		}
	case 164:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:890
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:895
		{
			yyVAL.valExpr = nil
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:905
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:909
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:915
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = nil
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:924
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:934
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:940
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:944
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:952
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:958
		{
			yyVAL.valExprs = nil
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:962
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:967
		{
			yyVAL.boolExpr = nil
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:971
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:976
		{
			yyVAL.orderBy = nil
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:980
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:986
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:990
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:996
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.str = AST_ASC
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.str = AST_ASC
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1009
		{
			yyVAL.str = AST_DESC
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.limit = nil
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1018
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1022
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.str = ""
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1035
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
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.columns = nil
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.updateExprs = nil
		}
	case 201:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.updateExprs = NewUpdateExprs(yyDollar[1].updateExpr)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1104
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1107
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
		//line sql.y:1113
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1145
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
