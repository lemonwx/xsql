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
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const PROXY = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const TABLE = 57424
const INDEX = 57425
const VIEW = 57426
const TO = 57427
const IGNORE = 57428
const IF = 57429
const UNIQUE = 57430
const USING = 57431
const DATABASE = 57432

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

const yyNprod = 217
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 599

var yyAct = [...]int{

	113, 159, 151, 389, 73, 192, 359, 121, 110, 278,
	141, 317, 233, 269, 271, 111, 207, 99, 75, 202,
	92, 398, 167, 168, 193, 3, 100, 216, 88, 292,
	293, 294, 295, 296, 62, 297, 298, 80, 104, 398,
	398, 35, 36, 37, 38, 77, 51, 18, 52, 83,
	341, 343, 85, 76, 161, 323, 89, 161, 46, 64,
	48, 370, 97, 261, 49, 161, 45, 120, 231, 105,
	126, 285, 231, 144, 400, 369, 368, 78, 117, 118,
	119, 345, 82, 140, 270, 78, 153, 84, 53, 342,
	124, 148, 399, 397, 138, 95, 143, 152, 263, 155,
	351, 158, 303, 164, 157, 365, 166, 350, 322, 136,
	312, 194, 131, 122, 123, 195, 154, 70, 310, 63,
	127, 262, 167, 168, 198, 230, 54, 55, 56, 251,
	201, 77, 150, 318, 77, 205, 74, 211, 210, 76,
	281, 270, 76, 315, 125, 147, 190, 191, 212, 367,
	214, 366, 209, 59, 60, 61, 158, 241, 133, 226,
	180, 181, 182, 133, 105, 239, 240, 211, 237, 227,
	339, 252, 244, 222, 318, 249, 250, 229, 253, 254,
	255, 256, 257, 258, 259, 260, 245, 129, 238, 335,
	132, 338, 220, 156, 336, 223, 167, 168, 333, 105,
	105, 337, 231, 334, 77, 77, 242, 243, 274, 149,
	208, 353, 76, 276, 265, 267, 375, 282, 35, 36,
	37, 38, 277, 355, 87, 228, 236, 273, 283, 77,
	377, 378, 135, 287, 288, 235, 204, 76, 79, 160,
	385, 384, 286, 237, 128, 162, 208, 280, 302, 289,
	306, 307, 273, 290, 219, 221, 218, 178, 179, 180,
	181, 182, 305, 383, 292, 293, 294, 295, 296, 105,
	297, 298, 175, 176, 177, 178, 179, 180, 181, 182,
	325, 90, 314, 161, 311, 81, 321, 324, 320, 133,
	203, 153, 18, 199, 197, 266, 196, 116, 237, 237,
	98, 204, 120, 331, 332, 126, 63, 347, 78, 316,
	395, 349, 103, 117, 118, 119, 346, 301, 352, 165,
	344, 108, 236, 328, 77, 124, 357, 327, 396, 360,
	225, 235, 356, 300, 224, 63, 206, 71, 361, 145,
	142, 139, 137, 134, 107, 86, 58, 130, 122, 123,
	101, 371, 373, 354, 348, 127, 372, 175, 176, 177,
	178, 179, 180, 181, 182, 91, 69, 158, 309, 402,
	381, 374, 379, 18, 246, 213, 247, 248, 360, 125,
	93, 388, 387, 264, 390, 390, 390, 77, 391, 392,
	65, 393, 116, 94, 146, 76, 272, 120, 18, 403,
	126, 67, 380, 404, 382, 405, 162, 103, 117, 118,
	119, 364, 326, 116, 279, 363, 108, 330, 120, 208,
	124, 126, 96, 72, 401, 386, 39, 18, 78, 117,
	118, 119, 40, 17, 16, 15, 14, 108, 13, 107,
	12, 124, 215, 122, 123, 101, 41, 42, 43, 44,
	127, 116, 47, 284, 217, 50, 120, 57, 275, 126,
	107, 394, 376, 358, 122, 123, 78, 117, 118, 119,
	120, 127, 362, 126, 125, 108, 329, 313, 200, 124,
	78, 117, 118, 119, 18, 19, 20, 21, 268, 153,
	115, 112, 114, 124, 319, 125, 109, 169, 107, 106,
	340, 234, 122, 123, 291, 232, 31, 102, 299, 127,
	163, 66, 22, 34, 68, 11, 122, 123, 10, 9,
	8, 308, 7, 127, 175, 176, 177, 178, 179, 180,
	181, 182, 6, 125, 170, 174, 172, 173, 175, 176,
	177, 178, 179, 180, 181, 182, 5, 125, 4, 2,
	1, 0, 0, 186, 187, 188, 189, 0, 183, 184,
	185, 0, 27, 28, 29, 0, 30, 33, 32, 0,
	0, 0, 23, 24, 26, 25, 304, 0, 0, 0,
	171, 175, 176, 177, 178, 179, 180, 181, 182, 0,
	0, 175, 176, 177, 178, 179, 180, 181, 182,
}
var yyPact = [...]int{

	479, -1000, -1000, 169, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -39, -53, -9, 29, -1000, -1000, -1000,
	-1000, 311, 63, 271, 422, 373, -1000, -1000, -1000, 383,
	-1000, 337, 302, 414, 50, -65, -65, -16, 271, -1000,
	-10, 271, -1000, 310, -74, 271, -74, 336, -1000, 370,
	413, 271, 256, -1000, -1000, -1000, 372, -1000, 205, 302,
	314, 36, 302, 105, 308, -1000, 187, -1000, 33, 307,
	27, 306, 271, -1000, 305, -1000, -27, 304, 374, 81,
	271, 302, -1000, 431, 445, 370, 445, 413, 445, 230,
	-1000, -1000, 300, 30, 57, 513, -1000, 431, 393, -1000,
	-1000, -1000, 445, 252, 250, -1000, 249, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 445, -1000, 257,
	273, 301, 409, 273, -1000, 445, 271, -1000, 355, -1000,
	-77, -1000, 160, -1000, 299, -1000, -1000, 295, -1000, 192,
	57, 513, 470, 42, -1000, 470, 370, 19, 470, -1000,
	191, 372, 445, -1000, -1000, 271, 84, 431, 431, 445,
	247, 353, 445, 445, 104, 445, 445, 445, 445, 445,
	445, 445, 445, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -43, 15, -8, 513, -1000, 277, 372, -1000, 422,
	5, 470, 368, 273, 273, 236, -1000, 401, 431, -1000,
	470, -1000, -1000, -1000, -1000, 76, 271, -1000, -29, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 368, 273, -1000,
	-1000, 445, 200, 210, 298, 287, 26, -1000, -1000, 523,
	-1000, -1000, -1000, -1000, 470, -1000, 247, 445, 445, 470,
	456, -1000, 343, 186, 186, 186, 87, 87, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 12, 372, 4, 62, -1000,
	431, 69, 247, 169, 110, 2, -1000, 401, 391, 398,
	57, 292, -1000, -1000, 288, -1000, -1000, 105, 470, 406,
	191, 191, -1000, -1000, 144, 135, 147, 137, 116, -12,
	-1000, 285, -25, 281, 445, -1000, 470, 289, 445, -1000,
	-1000, 1, -1000, 18, -1000, 445, 131, -1000, 323, 170,
	-1000, -1000, -1000, 273, 391, -1000, 445, -1000, -1000, 403,
	397, 210, 41, -1000, 97, -1000, 95, -1000, -1000, -1000,
	-1000, -22, -23, -37, -1000, -1000, -1000, 470, 445, 470,
	-1000, -1000, 470, 445, 321, 247, -1000, -1000, 163, -1000,
	204, -1000, 401, 431, 445, 431, -1000, -1000, 219, 197,
	196, 470, 470, 418, -1000, 445, -1000, -1000, -1000, 391,
	57, 149, 57, 271, 271, 271, 273, -1000, 294, -13,
	-1000, -14, -32, 105, -1000, 417, 348, -1000, 271, -1000,
	-1000, -1000, 271, -1000, 271, -1000,
}
var yyPgo = [...]int{

	0, 550, 549, 24, 548, 546, 532, 522, 520, 519,
	518, 515, 426, 514, 513, 511, 17, 26, 510, 508,
	507, 505, 12, 504, 501, 117, 500, 3, 16, 38,
	499, 497, 14, 496, 2, 15, 5, 494, 492, 7,
	491, 8, 490, 488, 13, 478, 477, 476, 472, 9,
	463, 6, 462, 1, 461, 19, 458, 11, 4, 18,
	224, 238, 455, 454, 453, 452, 442, 0, 10, 440,
	438, 436, 435, 434, 433, 95, 20, 432,
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
	55, 56, 56, 57, 57, 58, 58, 59, 60, 60,
	61, 61, 62, 62, 63, 63, 63, 63, 63, 64,
	64, 65, 65, 66, 66, 67, 68,
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
	3, 1, 3, 0, 5, 1, 3, 3, 0, 2,
	0, 3, 0, 1, 1, 1, 1, 1, 1, 0,
	1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 93, 94, 96, 95, 83, 84, 85,
	87, 27, 89, 88, -14, 49, 50, 51, 52, -12,
	-77, -12, -12, -12, -12, 105, 97, -65, 99, 103,
	-62, 99, 101, 97, 97, 98, 99, -12, 35, 90,
	91, 92, -67, 35, -3, 17, -15, 18, -13, 29,
	-25, 35, 9, -58, 86, -59, -41, -67, 35, -61,
	102, -61, 98, -67, 97, -67, 35, -60, 102, -67,
	-60, 29, -76, 10, 23, -75, 9, -67, 44, -16,
	-17, 73, -20, 35, -29, -34, -30, 67, 44, -33,
	-41, -35, -40, -67, -38, -42, 20, 36, 37, 38,
	25, -39, 71, 72, 48, 102, 28, 78, 39, -25,
	33, 76, -25, 53, 35, 45, 76, 35, 67, 35,
	-67, -68, 35, -68, 100, 35, 20, 64, -67, -25,
	-29, -34, -34, 44, -76, -34, -75, -36, -34, -53,
	9, 53, 15, -18, -67, 19, 76, 65, 66, -31,
	21, 67, 23, 24, 22, 68, 69, 70, 71, 72,
	73, 74, 75, 45, 46, 47, 40, 41, 42, 43,
	-29, -29, -36, -3, -34, -34, 44, 44, -39, 44,
	-45, -34, -55, 33, 44, -58, 35, -28, 10, -59,
	-34, -67, -68, 20, -68, -66, 104, -63, 96, 94,
	32, 95, 13, 35, 35, 35, -68, -55, 33, -76,
	106, 53, -21, -22, -24, 44, 35, -39, -17, -34,
	-67, 73, -29, -29, -34, -35, 21, 23, 24, -34,
	-34, 25, 67, -34, -34, -34, -34, -34, -34, -34,
	-34, 106, 106, 106, 106, -16, 18, -16, -43, -44,
	79, -32, 28, -3, -58, -56, -41, -28, -49, 13,
	-29, 64, -67, -68, -64, 100, -32, -58, -34, -28,
	53, -23, 54, 55, 56, 57, 58, 60, 61, -19,
	35, 19, -22, 76, 53, -35, -34, -34, 65, 25,
	106, -16, 106, -46, -44, 81, -29, -57, 64, -37,
	-35, -57, 106, 53, -49, -53, 14, 35, 35, -47,
	11, -22, -22, 54, 59, 54, 59, 54, 54, 54,
	-26, 62, 101, 63, 35, 106, 35, -34, 65, -34,
	106, 82, -34, 80, 30, 53, -41, -53, -50, -51,
	-34, -68, -48, 12, 14, 64, 54, 54, 98, 98,
	98, -34, -34, 31, -35, 53, -52, 26, 27, -49,
	-29, -36, -29, 44, 44, 44, 7, -51, -53, -27,
	-67, -27, -27, -58, -54, 16, 34, 106, 53, 106,
	106, 7, 21, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 47, 47,
	47, 47, 47, 211, 202, 0, 0, 28, 29, 30,
	47, 0, 0, 0, 0, 51, 53, 54, 55, 56,
	49, 0, 0, 0, 0, 200, 200, 0, 0, 212,
	0, 0, 203, 0, 198, 0, 198, 0, 32, 99,
	102, 0, 0, 215, 19, 52, 0, 57, 48, 0,
	0, 89, 0, 26, 0, 195, 0, 165, 215, 0,
	0, 0, 0, 216, 0, 216, 0, 0, 0, 0,
	0, 0, 33, 0, 0, 99, 0, 102, 0, 183,
	58, 60, 65, 215, 63, 64, 104, 0, 0, 135,
	136, 137, 0, 165, 0, 151, 0, 167, 168, 169,
	170, 131, 154, 155, 156, 152, 153, 158, 50, 189,
	0, 0, 97, 0, 27, 0, 0, 216, 0, 216,
	213, 39, 0, 42, 0, 44, 199, 0, 216, 189,
	100, 0, 101, 0, 34, 103, 99, 0, 133, 17,
	0, 0, 0, 61, 66, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 119, 120, 121, 122, 123, 124, 125,
	107, 0, 0, 0, 133, 146, 0, 0, 118, 0,
	0, 159, 0, 0, 0, 97, 90, 175, 0, 196,
	197, 166, 36, 201, 37, 0, 0, 216, 209, 204,
	205, 206, 207, 208, 43, 45, 46, 0, 0, 35,
	31, 0, 97, 68, 74, 0, 86, 88, 59, 184,
	67, 62, 105, 106, 109, 110, 0, 0, 0, 112,
	0, 116, 0, 138, 139, 140, 141, 142, 143, 144,
	145, 108, 130, 132, 147, 0, 0, 0, 163, 160,
	0, 193, 0, 127, 193, 0, 191, 175, 183, 0,
	98, 0, 214, 40, 0, 210, 22, 23, 134, 171,
	0, 0, 77, 78, 0, 0, 0, 0, 0, 91,
	75, 0, 0, 0, 0, 111, 113, 0, 0, 117,
	148, 0, 150, 0, 161, 0, 0, 20, 0, 126,
	128, 21, 190, 0, 183, 25, 0, 216, 41, 173,
	0, 69, 72, 79, 0, 81, 0, 83, 84, 85,
	70, 0, 0, 0, 76, 71, 87, 185, 0, 114,
	149, 157, 164, 0, 0, 0, 192, 24, 176, 177,
	180, 38, 175, 0, 0, 0, 80, 82, 0, 0,
	0, 115, 162, 0, 129, 0, 179, 181, 182, 183,
	174, 172, 73, 0, 0, 0, 0, 178, 186, 0,
	95, 0, 0, 194, 18, 0, 0, 92, 0, 93,
	94, 187, 0, 96, 0, 188,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 106, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105,
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
		//line sql.y:256
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:260
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:266
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:272
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:278
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:284
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:290
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:294
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:298
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:302
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:312
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 38:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:316
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:327
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:331
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:342
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:348
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:352
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:357
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:362
		{
			SetAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:366
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:372
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:376
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = AST_UNION
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:390
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:394
		{
			yyVAL.str = AST_EXCEPT
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:398
		{
			yyVAL.str = AST_INTERSECT
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = AST_DISTINCT
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:413
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:417
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:423
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:427
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:431
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:437
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:441
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:446
		{
			yyVAL.bytes = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:450
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:454
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:460
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:464
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:470
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:474
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:478
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:482
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:487
		{
			yyVAL.bytes = nil
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:491
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:495
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:501
		{
			yyVAL.str = AST_JOIN
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:505
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:509
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:513
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:517
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:521
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:525
		{
			yyVAL.str = AST_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:529
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:539
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:543
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:547
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:553
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:557
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:562
		{
			yyVAL.indexHints = nil
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:566
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:570
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:574
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:580
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:584
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:589
		{
			yyVAL.boolExpr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:593
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:598
		{
			yyVAL.expr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:602
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:606
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:611
		{
			yyVAL.valExpr = nil
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:615
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:630
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:634
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:672
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:682
		{
			yyVAL.str = AST_EQ
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:686
		{
			yyVAL.str = AST_LT
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:690
		{
			yyVAL.str = AST_GT
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:694
		{
			yyVAL.str = AST_LE
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.str = AST_GE
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = AST_NE
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = AST_NSE
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:712
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:716
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:722
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:726
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:732
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:742
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:752
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:766
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:798
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:802
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
		//line sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 149:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:829
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:833
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:839
		{
			yyVAL.bytes = IF_BYTES
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:843
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:849
		{
			yyVAL.byt = AST_UPLUS
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:853
		{
			yyVAL.byt = AST_UMINUS
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:857
		{
			yyVAL.byt = AST_TILDA
		}
	case 157:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:863
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:868
		{
			yyVAL.valExpr = nil
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:872
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:878
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:882
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:888
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 163:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = nil
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:903
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:907
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:913
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:921
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:925
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExprs = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:939
		{
			yyVAL.boolExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:943
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:948
		{
			yyVAL.orderBy = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:952
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:958
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:962
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:968
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:973
		{
			yyVAL.str = AST_ASC
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:977
		{
			yyVAL.str = AST_ASC
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.str = AST_DESC
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:986
		{
			yyVAL.limit = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:990
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:994
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:999
		{
			yyVAL.str = ""
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1007
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
		//line sql.y:1020
		{
			yyVAL.columns = nil
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1030
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.updateExprs = nil
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1049
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1059
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1112
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
