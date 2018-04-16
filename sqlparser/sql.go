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

const yyNprod = 220
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 640

var yyAct = [...]int{

	117, 165, 158, 396, 75, 198, 366, 125, 114, 285,
	324, 240, 145, 276, 278, 115, 213, 103, 223, 77,
	199, 3, 208, 405, 405, 104, 299, 300, 301, 302,
	303, 96, 304, 305, 64, 173, 174, 405, 108, 167,
	46, 90, 48, 82, 18, 79, 49, 330, 45, 85,
	292, 148, 87, 78, 377, 66, 91, 376, 375, 35,
	36, 37, 38, 101, 124, 167, 167, 130, 51, 84,
	52, 109, 238, 238, 80, 121, 122, 123, 99, 268,
	407, 406, 352, 155, 86, 144, 53, 128, 358, 348,
	350, 80, 142, 152, 404, 277, 357, 154, 156, 310,
	147, 159, 172, 161, 329, 164, 140, 170, 163, 135,
	137, 126, 127, 65, 124, 200, 258, 130, 131, 201,
	270, 325, 319, 317, 80, 121, 122, 123, 204, 269,
	237, 160, 349, 155, 207, 79, 157, 128, 79, 211,
	372, 218, 216, 78, 76, 129, 78, 229, 173, 174,
	196, 197, 248, 217, 219, 374, 221, 215, 164, 259,
	325, 126, 127, 360, 288, 233, 227, 72, 131, 230,
	109, 246, 247, 218, 244, 151, 234, 277, 251, 322,
	162, 256, 257, 373, 260, 261, 262, 263, 264, 265,
	266, 267, 252, 245, 236, 129, 54, 55, 56, 60,
	61, 62, 58, 59, 346, 109, 109, 186, 187, 188,
	79, 79, 249, 250, 281, 173, 174, 342, 78, 283,
	272, 274, 343, 345, 289, 340, 344, 214, 284, 280,
	341, 89, 226, 228, 225, 137, 79, 290, 238, 133,
	294, 295, 136, 214, 78, 35, 36, 37, 38, 293,
	244, 382, 362, 287, 309, 280, 296, 313, 314, 139,
	392, 153, 184, 185, 186, 187, 188, 166, 81, 312,
	297, 132, 391, 168, 355, 390, 109, 181, 182, 183,
	184, 185, 186, 187, 188, 155, 137, 332, 92, 321,
	205, 318, 328, 315, 331, 327, 181, 182, 183, 184,
	185, 186, 187, 188, 18, 244, 244, 384, 385, 338,
	339, 167, 203, 202, 354, 83, 323, 102, 356, 243,
	299, 300, 301, 302, 303, 359, 304, 305, 242, 235,
	209, 79, 380, 364, 243, 65, 367, 80, 353, 363,
	210, 210, 351, 242, 308, 171, 335, 368, 311, 334,
	181, 182, 183, 184, 185, 186, 187, 188, 378, 232,
	307, 65, 402, 379, 181, 182, 183, 184, 185, 186,
	187, 188, 231, 212, 164, 73, 149, 388, 381, 386,
	403, 316, 146, 143, 141, 367, 138, 88, 395, 394,
	63, 397, 397, 397, 79, 398, 399, 273, 400, 120,
	134, 361, 78, 18, 124, 93, 410, 130, 95, 387,
	411, 389, 412, 120, 107, 121, 122, 123, 124, 71,
	97, 130, 253, 112, 254, 255, 279, 128, 107, 121,
	122, 123, 94, 98, 409, 220, 150, 112, 69, 67,
	168, 128, 371, 333, 286, 370, 337, 111, 18, 214,
	100, 126, 127, 105, 74, 408, 393, 18, 131, 40,
	17, 111, 16, 120, 15, 126, 127, 105, 124, 14,
	13, 130, 131, 39, 12, 222, 47, 291, 80, 121,
	122, 123, 224, 50, 282, 129, 401, 112, 383, 271,
	365, 128, 369, 41, 42, 43, 44, 336, 320, 129,
	206, 275, 120, 119, 57, 116, 118, 124, 326, 113,
	130, 111, 175, 110, 347, 126, 127, 80, 121, 122,
	123, 124, 131, 241, 130, 298, 112, 239, 106, 306,
	128, 80, 121, 122, 123, 18, 19, 20, 21, 169,
	155, 68, 34, 70, 128, 11, 10, 9, 8, 129,
	111, 7, 6, 5, 126, 127, 4, 32, 2, 1,
	0, 131, 0, 22, 0, 0, 0, 0, 126, 127,
	0, 0, 0, 0, 0, 131, 181, 182, 183, 184,
	185, 186, 187, 188, 176, 180, 178, 179, 129, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 129, 192, 193, 194, 195, 0, 189, 190,
	191, 0, 0, 0, 27, 28, 29, 0, 0, 30,
	33, 31, 0, 0, 0, 0, 0, 23, 24, 26,
	25, 177, 181, 182, 183, 184, 185, 186, 187, 188,
}
var yyPact = [...]int{

	530, -1000, -1000, 196, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -61, -35, -15, 95, -1000, -1000, -1000,
	-1000, 107, 355, 300, 452, 422, -1000, -1000, -1000, 420,
	-1000, 390, 340, 445, 56, -63, -63, -33, 300, -1000,
	-17, 300, -1000, 352, -65, 300, -65, 376, 409, 385,
	410, 441, 300, -1000, 273, -1000, -1000, -1000, 393, -1000,
	232, 340, 367, 32, 340, 182, 351, -1000, 214, -1000,
	29, 349, 24, 348, 300, -1000, 347, -1000, -53, 341,
	416, 111, 300, 340, 496, 496, -1000, 482, 496, 410,
	496, 441, 496, 258, -1000, -1000, 326, 25, 149, 563,
	-1000, 482, 443, -1000, -1000, -1000, 496, 269, 268, -1000,
	246, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 496, -1000, 297, 302, 338, 439, 302, -1000, 89,
	300, -1000, 415, -1000, -90, -1000, 134, -1000, 337, -1000,
	-1000, 324, -1000, 296, 507, 39, 507, 149, 563, 507,
	-1000, 507, 410, 20, 507, -1000, 284, 393, 496, -1000,
	-1000, 300, 78, 482, 482, 496, 241, 401, 496, 496,
	91, 496, 496, 496, 496, 496, 496, 496, 496, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -31, 19, 10,
	563, -1000, 379, 393, -1000, 452, 15, 507, 398, 302,
	302, 233, -1000, 431, 482, -1000, 507, -1000, -1000, -1000,
	-1000, -1000, 100, 300, -1000, -54, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 398, 302, -1000, -1000, 496, 217,
	266, 325, 299, 22, -1000, -1000, 295, -1000, -1000, -1000,
	-1000, 507, -1000, 241, 496, 496, 507, 227, -1000, 356,
	190, 190, 190, 133, 133, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 13, 393, 12, 97, -1000, 482, 96, 241,
	196, 57, -6, -1000, 431, 425, 429, 149, 314, -1000,
	-1000, 311, -1000, -1000, 182, 507, 435, 284, 284, -1000,
	-1000, 171, 163, 172, 169, 150, 27, -1000, 307, -28,
	303, 496, -1000, 507, 208, 496, -1000, -1000, -14, -1000,
	5, -1000, 496, 82, -1000, 371, 199, -1000, -1000, -1000,
	302, 425, -1000, 496, -1000, -1000, 433, 428, 266, 76,
	-1000, 129, -1000, 101, -1000, -1000, -1000, -1000, -44, -45,
	-48, -1000, -1000, -1000, 507, 496, 507, -1000, -1000, 507,
	496, 301, 241, -1000, -1000, 198, -1000, 281, -1000, 431,
	482, 496, 482, -1000, -1000, 231, 228, 216, 507, 507,
	449, -1000, 496, -1000, -1000, -1000, 425, 149, 185, 149,
	300, 300, 300, 302, -1000, 346, -16, -1000, -29, -30,
	182, -1000, 448, 413, -1000, 300, -1000, -1000, -1000, 300,
	-1000, 300, -1000,
}
var yyPgo = [...]int{

	0, 559, 558, 20, 556, 553, 552, 551, 548, 547,
	546, 545, 473, 543, 542, 541, 17, 25, 539, 529,
	528, 527, 11, 525, 523, 167, 514, 3, 16, 38,
	513, 512, 14, 509, 2, 15, 5, 508, 506, 7,
	505, 8, 503, 501, 13, 500, 498, 497, 492, 9,
	490, 6, 488, 1, 486, 22, 484, 10, 4, 19,
	231, 268, 483, 482, 477, 476, 475, 0, 12, 474,
	470, 469, 464, 462, 460, 78, 31, 459,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 73, 73, 73, 73, 8, 8,
	8, 8, 9, 9, 9, 10, 11, 11, 11, 77,
	12, 13, 13, 14, 14, 14, 14, 14, 15, 15,
	16, 16, 17, 17, 17, 20, 20, 18, 18, 18,
	21, 21, 22, 22, 22, 22, 19, 19, 19, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 24, 24,
	24, 25, 25, 26, 26, 26, 26, 27, 27, 28,
	28, 76, 76, 76, 75, 75, 29, 29, 29, 29,
	29, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 31, 31, 31, 31, 31, 31, 31, 32, 32,
	37, 37, 35, 35, 39, 36, 36, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 38, 38, 40, 40, 40, 42,
	45, 45, 43, 43, 44, 46, 46, 41, 41, 33,
	33, 33, 33, 47, 47, 48, 48, 49, 49, 50,
	50, 51, 52, 52, 52, 53, 53, 53, 54, 54,
	54, 55, 55, 56, 56, 57, 57, 58, 58, 59,
	59, 60, 60, 61, 61, 62, 62, 63, 63, 63,
	63, 63, 64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 5, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 4, 4, 2, 3, 4, 5, 5, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 0,
	2, 0, 2, 1, 2, 1, 1, 1, 0, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 1, 2,
	1, 3, 3, 3, 3, 5, 0, 1, 2, 1,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 3,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 0, 2, 2, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 3, 3, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 1, 3, 3,
	3, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 97, 98, 100, 99, 84, 85, 86,
	89, 91, 27, 90, -14, 49, 50, 51, 52, -12,
	-77, -12, -12, -12, -12, 109, 101, -65, 103, 107,
	-62, 103, 105, 101, 101, 102, 103, -12, 95, 96,
	92, 93, 94, 35, -67, 35, -3, 17, -15, 18,
	-13, 29, -25, 35, 9, -58, 88, -59, -41, -67,
	35, -61, 106, -61, 102, -67, 101, -67, 35, -60,
	106, -67, -60, 29, 23, 23, -76, 10, 23, -75,
	9, -67, 44, -16, -17, 74, -20, 35, -29, -34,
	-30, 68, 44, -33, -41, -35, -40, -67, -38, -42,
	20, 36, 37, 38, 25, -39, 72, 73, 48, 106,
	28, 79, 39, -25, 33, 77, -25, 53, 35, 45,
	77, 35, 68, 35, -67, -68, 35, -68, 104, 35,
	20, 64, -67, -25, -34, 44, -34, -29, -34, -34,
	-76, -34, -75, -36, -34, -53, 9, 53, 15, -18,
	-67, 19, 77, 66, 67, -31, 21, 68, 23, 24,
	22, 69, 70, 71, 72, 73, 74, 75, 76, 45,
	46, 47, 40, 41, 42, 43, -29, -29, -36, -3,
	-34, -34, 44, 44, -39, 44, -45, -34, -55, 33,
	44, -58, 35, -28, 10, -59, -34, 64, -67, -68,
	20, -68, -66, 108, -63, 100, 98, 32, 99, 13,
	35, 35, 35, -68, -55, 33, -76, 110, 53, -21,
	-22, -24, 44, 35, -39, -17, -34, -67, 74, -29,
	-29, -34, -35, 21, 23, 24, -34, -34, 25, 68,
	-34, -34, -34, -34, -34, -34, -34, -34, 110, 110,
	110, 110, -16, 18, -16, -43, -44, 80, -32, 28,
	-3, -58, -56, -41, -28, -49, 13, -29, 64, -67,
	-68, -64, 104, -32, -58, -34, -28, 53, -23, 54,
	55, 56, 57, 58, 60, 61, -19, 35, 19, -22,
	77, 53, -35, -34, -34, 66, 25, 110, -16, 110,
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
	9, 10, 11, 12, 13, 14, 15, 16, 49, 49,
	49, 49, 49, 214, 205, 0, 0, 28, 29, 30,
	49, 0, 0, 0, 0, 53, 55, 56, 57, 58,
	51, 0, 0, 0, 0, 203, 203, 0, 0, 215,
	0, 0, 206, 0, 201, 0, 201, 0, 0, 0,
	101, 104, 0, 34, 0, 218, 19, 54, 0, 59,
	50, 0, 0, 91, 0, 26, 0, 197, 0, 167,
	218, 0, 0, 0, 0, 219, 0, 219, 0, 0,
	0, 0, 0, 0, 0, 0, 35, 0, 0, 101,
	0, 104, 0, 185, 60, 62, 67, 218, 65, 66,
	106, 0, 0, 137, 138, 139, 0, 167, 0, 153,
	0, 169, 170, 171, 172, 133, 156, 157, 158, 154,
	155, 160, 52, 191, 0, 0, 99, 0, 27, 0,
	0, 219, 0, 219, 216, 41, 0, 44, 0, 46,
	202, 0, 219, 191, 32, 0, 33, 102, 0, 103,
	36, 105, 101, 0, 135, 17, 0, 0, 0, 63,
	68, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 121,
	122, 123, 124, 125, 126, 127, 109, 0, 0, 0,
	135, 148, 0, 0, 120, 0, 0, 161, 0, 0,
	0, 99, 92, 177, 0, 198, 199, 200, 168, 38,
	204, 39, 0, 0, 219, 212, 207, 208, 209, 210,
	211, 45, 47, 48, 0, 0, 37, 31, 0, 99,
	70, 76, 0, 88, 90, 61, 186, 69, 64, 107,
	108, 111, 112, 0, 0, 0, 114, 0, 118, 0,
	140, 141, 142, 143, 144, 145, 146, 147, 110, 132,
	134, 149, 0, 0, 0, 165, 162, 0, 195, 0,
	129, 195, 0, 193, 177, 185, 0, 100, 0, 217,
	42, 0, 213, 22, 23, 136, 173, 0, 0, 79,
	80, 0, 0, 0, 0, 0, 93, 77, 0, 0,
	0, 0, 113, 115, 0, 0, 119, 150, 0, 152,
	0, 163, 0, 0, 20, 0, 128, 130, 21, 192,
	0, 185, 25, 0, 219, 43, 175, 0, 71, 74,
	81, 0, 83, 0, 85, 86, 87, 72, 0, 0,
	0, 78, 73, 89, 187, 0, 116, 151, 159, 166,
	0, 0, 0, 194, 24, 178, 179, 182, 40, 177,
	0, 0, 0, 82, 84, 0, 0, 0, 117, 164,
	0, 131, 0, 181, 183, 184, 185, 176, 174, 75,
	0, 0, 0, 0, 180, 188, 0, 97, 0, 0,
	196, 18, 0, 0, 94, 0, 95, 96, 189, 0,
	98, 0, 190,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:291
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &Show{Section: "variables"}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &Show{Section: "desc "}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:303
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:307
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:311
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:317
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:321
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:325
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:340
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:351
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:357
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:361
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:366
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:371
		{
			SetAllowComments(yylex, true)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:375
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:381
		{
			yyVAL.bytes2 = nil
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:385
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = AST_UNION
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.str = AST_EXCEPT
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.str = AST_INTERSECT
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:412
		{
			yyVAL.str = ""
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:416
		{
			yyVAL.str = AST_DISTINCT
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:422
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:426
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:432
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:436
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:440
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:446
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:450
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:455
		{
			yyVAL.bytes = nil
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:459
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:463
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:469
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:473
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:496
		{
			yyVAL.bytes = nil
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:500
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:504
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:510
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:514
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:518
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:522
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:526
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:534
		{
			yyVAL.str = AST_JOIN
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:542
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:548
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:552
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:556
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:562
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:566
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:571
		{
			yyVAL.indexHints = nil
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:575
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 95:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:579
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:583
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:589
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:593
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:598
		{
			yyVAL.boolExpr = nil
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:602
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:607
		{
			yyVAL.expr = nil
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:611
		{
			yyVAL.expr = yyDollar[2].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:615
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:620
		{
			yyVAL.valExpr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:624
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:631
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:635
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:639
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:649
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:653
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:657
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:665
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:669
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:673
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:677
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:681
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:685
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:691
		{
			yyVAL.str = AST_EQ
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:695
		{
			yyVAL.str = AST_LT
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:699
		{
			yyVAL.str = AST_GT
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:703
		{
			yyVAL.str = AST_LE
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:707
		{
			yyVAL.str = AST_GE
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:711
		{
			yyVAL.str = AST_NE
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:715
		{
			yyVAL.str = AST_NSE
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:721
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:725
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:731
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:735
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:741
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:745
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:751
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:757
		{
			yyVAL.valExprs = ValExprs{NumVal{}, yyDollar[1].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:761
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:775
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:779
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:811
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
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:834
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:838
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:842
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:848
		{
			yyVAL.bytes = IF_BYTES
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:852
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:858
		{
			yyVAL.byt = AST_UPLUS
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:862
		{
			yyVAL.byt = AST_UMINUS
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:866
		{
			yyVAL.byt = AST_TILDA
		}
	case 159:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:872
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = nil
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:887
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:891
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:897
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = nil
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:906
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:912
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:916
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:922
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExprs = nil
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:948
		{
			yyVAL.boolExpr = nil
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:952
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:957
		{
			yyVAL.orderBy = nil
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:961
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:967
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:971
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:977
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:982
		{
			yyVAL.str = AST_ASC
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:986
		{
			yyVAL.str = AST_ASC
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.str = AST_DESC
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:995
		{
			yyVAL.limit = nil
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:999
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1008
		{
			yyVAL.str = ""
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 190:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1016
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
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.columns = nil
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: &ColName{Name: []byte("version")}}, &NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1048
		{
			yyVAL.updateExprs = nil
		}
	case 196:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.updateExprs = UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("version")}}, yyDollar[1].updateExpr}
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: StrVal([]byte("ON"))}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1126
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
