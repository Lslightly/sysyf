// Code generated from SysYF.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SysYF
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SysYFParser struct {
	*antlr.BaseParser
}

var SysYFParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sysyfParserInit() {
	staticData := &SysYFParserStaticData
	staticData.LiteralNames = []string{
		"", "'const'", "','", "';'", "'='", "'{'", "'}'", "'('", "')'", "'if'",
		"'else'", "'while'", "'break'", "'continue'", "'return'", "'void'",
		"'int'", "'float'", "'+'", "'-'", "'*'", "'/'", "'%'", "'||'", "'&&'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'!'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "VOID_T",
		"INT_T", "FLOAT_T", "PLUS", "MINUS", "MUL", "DIV", "MOD", "LOR", "LAND",
		"EQ", "NEQ", "LT", "LTE", "GT", "GTE", "NOT", "LBRACKET", "RBRACKET",
		"NAME", "INT", "FLOAT", "C_COMMENT", "CPP_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"comp_unit", "global_def", "const_decl", "btype", "const_def", "const_init_val",
		"var_decl", "var_def", "func_def", "func_fparam", "block", "block_item",
		"stmt", "assign_stmt", "empty_stmt", "exp_stmt", "if_stmt", "while_stmt",
		"break_stmt", "continue_stmt", "return_stmt", "cond_exp", "exp", "otherexp",
		"func_call", "lval", "ident", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 291, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 4, 0, 58, 8, 0, 11, 0, 12, 0, 59, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 3, 1, 67, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 74,
		8, 2, 10, 2, 12, 2, 77, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		3, 4, 86, 8, 4, 1, 4, 3, 4, 89, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 101, 8, 5, 10, 5, 12, 5, 104, 9, 5, 1,
		5, 1, 5, 3, 5, 108, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 114, 8, 6, 10,
		6, 12, 6, 117, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 124, 8, 7, 1,
		7, 3, 7, 127, 8, 7, 1, 7, 1, 7, 3, 7, 131, 8, 7, 1, 8, 1, 8, 3, 8, 135,
		8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 142, 8, 8, 10, 8, 12, 8, 145,
		9, 8, 3, 8, 147, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		156, 8, 9, 1, 10, 1, 10, 5, 10, 160, 8, 10, 10, 10, 12, 10, 163, 9, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 170, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 181, 8, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 200, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 3, 20, 216, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 5, 21, 235, 8, 21, 10, 21, 12, 21, 238, 9, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 3, 22, 244, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5,
		22, 252, 8, 22, 10, 22, 12, 22, 255, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 264, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 271, 8, 24, 10, 24, 12, 24, 274, 9, 24, 3, 24, 276, 8, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 285, 8, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 0, 2, 42, 44, 28, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 0, 7, 1, 0, 16, 17, 1, 0, 27, 30, 1, 0, 25, 26, 2, 0, 18, 19, 31,
		31, 1, 0, 20, 22, 1, 0, 18, 19, 1, 0, 35, 36, 305, 0, 57, 1, 0, 0, 0, 2,
		66, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0, 6, 80, 1, 0, 0, 0, 8, 82, 1, 0, 0, 0,
		10, 107, 1, 0, 0, 0, 12, 109, 1, 0, 0, 0, 14, 120, 1, 0, 0, 0, 16, 134,
		1, 0, 0, 0, 18, 151, 1, 0, 0, 0, 20, 157, 1, 0, 0, 0, 22, 169, 1, 0, 0,
		0, 24, 180, 1, 0, 0, 0, 26, 182, 1, 0, 0, 0, 28, 187, 1, 0, 0, 0, 30, 189,
		1, 0, 0, 0, 32, 192, 1, 0, 0, 0, 34, 201, 1, 0, 0, 0, 36, 207, 1, 0, 0,
		0, 38, 210, 1, 0, 0, 0, 40, 213, 1, 0, 0, 0, 42, 219, 1, 0, 0, 0, 44, 243,
		1, 0, 0, 0, 46, 263, 1, 0, 0, 0, 48, 265, 1, 0, 0, 0, 50, 279, 1, 0, 0,
		0, 52, 286, 1, 0, 0, 0, 54, 288, 1, 0, 0, 0, 56, 58, 3, 2, 1, 0, 57, 56,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 61, 1, 0, 0, 0, 61, 62, 5, 0, 0, 1, 62, 1, 1, 0, 0, 0, 63, 67, 3, 4,
		2, 0, 64, 67, 3, 12, 6, 0, 65, 67, 3, 16, 8, 0, 66, 63, 1, 0, 0, 0, 66,
		64, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 3, 1, 0, 0, 0, 68, 69, 5, 1, 0,
		0, 69, 70, 3, 6, 3, 0, 70, 75, 3, 8, 4, 0, 71, 72, 5, 2, 0, 0, 72, 74,
		3, 8, 4, 0, 73, 71, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5,
		3, 0, 0, 79, 5, 1, 0, 0, 0, 80, 81, 7, 0, 0, 0, 81, 7, 1, 0, 0, 0, 82,
		88, 3, 52, 26, 0, 83, 85, 5, 32, 0, 0, 84, 86, 5, 35, 0, 0, 85, 84, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 5, 33, 0, 0, 88,
		83, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 5, 4, 0,
		0, 91, 92, 3, 10, 5, 0, 92, 9, 1, 0, 0, 0, 93, 108, 3, 44, 22, 0, 94, 95,
		5, 5, 0, 0, 95, 108, 5, 6, 0, 0, 96, 97, 5, 5, 0, 0, 97, 102, 3, 44, 22,
		0, 98, 99, 5, 2, 0, 0, 99, 101, 3, 44, 22, 0, 100, 98, 1, 0, 0, 0, 101,
		104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 6, 0, 0, 106, 108, 1, 0,
		0, 0, 107, 93, 1, 0, 0, 0, 107, 94, 1, 0, 0, 0, 107, 96, 1, 0, 0, 0, 108,
		11, 1, 0, 0, 0, 109, 110, 3, 6, 3, 0, 110, 115, 3, 14, 7, 0, 111, 112,
		5, 2, 0, 0, 112, 114, 3, 14, 7, 0, 113, 111, 1, 0, 0, 0, 114, 117, 1, 0,
		0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0,
		117, 115, 1, 0, 0, 0, 118, 119, 5, 3, 0, 0, 119, 13, 1, 0, 0, 0, 120, 126,
		3, 52, 26, 0, 121, 123, 5, 32, 0, 0, 122, 124, 5, 35, 0, 0, 123, 122, 1,
		0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 5, 33, 0,
		0, 126, 121, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128,
		129, 5, 4, 0, 0, 129, 131, 3, 10, 5, 0, 130, 128, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 15, 1, 0, 0, 0, 132, 135, 3, 6, 3, 0, 133, 135, 5, 15,
		0, 0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 137, 3, 52, 26, 0, 137, 146, 5, 7, 0, 0, 138, 143, 3, 18, 9, 0, 139,
		140, 5, 2, 0, 0, 140, 142, 3, 18, 9, 0, 141, 139, 1, 0, 0, 0, 142, 145,
		1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 147, 1, 0,
		0, 0, 145, 143, 1, 0, 0, 0, 146, 138, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 149, 5, 8, 0, 0, 149, 150, 3, 20, 10, 0, 150,
		17, 1, 0, 0, 0, 151, 152, 3, 6, 3, 0, 152, 155, 3, 52, 26, 0, 153, 154,
		5, 32, 0, 0, 154, 156, 5, 33, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 19, 1, 0, 0, 0, 157, 161, 5, 5, 0, 0, 158, 160, 3, 22, 11,
		0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165,
		5, 6, 0, 0, 165, 21, 1, 0, 0, 0, 166, 170, 3, 4, 2, 0, 167, 170, 3, 12,
		6, 0, 168, 170, 3, 24, 12, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0,
		0, 169, 168, 1, 0, 0, 0, 170, 23, 1, 0, 0, 0, 171, 181, 3, 26, 13, 0, 172,
		181, 3, 28, 14, 0, 173, 181, 3, 30, 15, 0, 174, 181, 3, 20, 10, 0, 175,
		181, 3, 32, 16, 0, 176, 181, 3, 34, 17, 0, 177, 181, 3, 36, 18, 0, 178,
		181, 3, 38, 19, 0, 179, 181, 3, 40, 20, 0, 180, 171, 1, 0, 0, 0, 180, 172,
		1, 0, 0, 0, 180, 173, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180, 175, 1, 0,
		0, 0, 180, 176, 1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0,
		180, 179, 1, 0, 0, 0, 181, 25, 1, 0, 0, 0, 182, 183, 3, 50, 25, 0, 183,
		184, 5, 4, 0, 0, 184, 185, 3, 44, 22, 0, 185, 186, 5, 3, 0, 0, 186, 27,
		1, 0, 0, 0, 187, 188, 5, 3, 0, 0, 188, 29, 1, 0, 0, 0, 189, 190, 3, 44,
		22, 0, 190, 191, 5, 3, 0, 0, 191, 31, 1, 0, 0, 0, 192, 193, 5, 9, 0, 0,
		193, 194, 5, 7, 0, 0, 194, 195, 3, 42, 21, 0, 195, 196, 5, 8, 0, 0, 196,
		199, 3, 24, 12, 0, 197, 198, 5, 10, 0, 0, 198, 200, 3, 24, 12, 0, 199,
		197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 33, 1, 0, 0, 0, 201, 202, 5,
		11, 0, 0, 202, 203, 5, 7, 0, 0, 203, 204, 3, 42, 21, 0, 204, 205, 5, 8,
		0, 0, 205, 206, 3, 24, 12, 0, 206, 35, 1, 0, 0, 0, 207, 208, 5, 12, 0,
		0, 208, 209, 5, 3, 0, 0, 209, 37, 1, 0, 0, 0, 210, 211, 5, 13, 0, 0, 211,
		212, 5, 3, 0, 0, 212, 39, 1, 0, 0, 0, 213, 215, 5, 14, 0, 0, 214, 216,
		3, 44, 22, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1,
		0, 0, 0, 217, 218, 5, 3, 0, 0, 218, 41, 1, 0, 0, 0, 219, 220, 6, 21, -1,
		0, 220, 221, 3, 44, 22, 0, 221, 236, 1, 0, 0, 0, 222, 223, 10, 4, 0, 0,
		223, 224, 7, 1, 0, 0, 224, 235, 3, 42, 21, 5, 225, 226, 10, 3, 0, 0, 226,
		227, 7, 2, 0, 0, 227, 235, 3, 42, 21, 4, 228, 229, 10, 2, 0, 0, 229, 230,
		5, 24, 0, 0, 230, 235, 3, 42, 21, 3, 231, 232, 10, 1, 0, 0, 232, 233, 5,
		23, 0, 0, 233, 235, 3, 42, 21, 2, 234, 222, 1, 0, 0, 0, 234, 225, 1, 0,
		0, 0, 234, 228, 1, 0, 0, 0, 234, 231, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0,
		236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 43, 1, 0, 0, 0, 238, 236,
		1, 0, 0, 0, 239, 240, 6, 22, -1, 0, 240, 244, 3, 46, 23, 0, 241, 242, 7,
		3, 0, 0, 242, 244, 3, 44, 22, 3, 243, 239, 1, 0, 0, 0, 243, 241, 1, 0,
		0, 0, 244, 253, 1, 0, 0, 0, 245, 246, 10, 2, 0, 0, 246, 247, 7, 4, 0, 0,
		247, 252, 3, 44, 22, 3, 248, 249, 10, 1, 0, 0, 249, 250, 7, 5, 0, 0, 250,
		252, 3, 44, 22, 2, 251, 245, 1, 0, 0, 0, 251, 248, 1, 0, 0, 0, 252, 255,
		1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 45, 1, 0,
		0, 0, 255, 253, 1, 0, 0, 0, 256, 264, 3, 50, 25, 0, 257, 264, 3, 54, 27,
		0, 258, 264, 3, 48, 24, 0, 259, 260, 5, 7, 0, 0, 260, 261, 3, 44, 22, 0,
		261, 262, 5, 8, 0, 0, 262, 264, 1, 0, 0, 0, 263, 256, 1, 0, 0, 0, 263,
		257, 1, 0, 0, 0, 263, 258, 1, 0, 0, 0, 263, 259, 1, 0, 0, 0, 264, 47, 1,
		0, 0, 0, 265, 266, 3, 52, 26, 0, 266, 275, 5, 7, 0, 0, 267, 272, 3, 44,
		22, 0, 268, 269, 5, 2, 0, 0, 269, 271, 3, 44, 22, 0, 270, 268, 1, 0, 0,
		0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273,
		276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 267, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 8, 0, 0, 278, 49, 1, 0,
		0, 0, 279, 284, 3, 52, 26, 0, 280, 281, 5, 32, 0, 0, 281, 282, 3, 44, 22,
		0, 282, 283, 5, 33, 0, 0, 283, 285, 1, 0, 0, 0, 284, 280, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 51, 1, 0, 0, 0, 286, 287, 5, 34, 0, 0, 287, 53, 1,
		0, 0, 0, 288, 289, 7, 6, 0, 0, 289, 55, 1, 0, 0, 0, 29, 59, 66, 75, 85,
		88, 102, 107, 115, 123, 126, 130, 134, 143, 146, 155, 161, 169, 180, 199,
		215, 234, 236, 243, 251, 253, 263, 272, 275, 284,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SysYFParserInit initializes any static state used to implement SysYFParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSysYFParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SysYFParserInit() {
	staticData := &SysYFParserStaticData
	staticData.once.Do(sysyfParserInit)
}

// NewSysYFParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSysYFParser(input antlr.TokenStream) *SysYFParser {
	SysYFParserInit()
	this := new(SysYFParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SysYFParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SysYF.g4"

	return this
}

// SysYFParser tokens.
const (
	SysYFParserEOF         = antlr.TokenEOF
	SysYFParserT__0        = 1
	SysYFParserT__1        = 2
	SysYFParserT__2        = 3
	SysYFParserT__3        = 4
	SysYFParserT__4        = 5
	SysYFParserT__5        = 6
	SysYFParserT__6        = 7
	SysYFParserT__7        = 8
	SysYFParserT__8        = 9
	SysYFParserT__9        = 10
	SysYFParserT__10       = 11
	SysYFParserT__11       = 12
	SysYFParserT__12       = 13
	SysYFParserT__13       = 14
	SysYFParserVOID_T      = 15
	SysYFParserINT_T       = 16
	SysYFParserFLOAT_T     = 17
	SysYFParserPLUS        = 18
	SysYFParserMINUS       = 19
	SysYFParserMUL         = 20
	SysYFParserDIV         = 21
	SysYFParserMOD         = 22
	SysYFParserLOR         = 23
	SysYFParserLAND        = 24
	SysYFParserEQ          = 25
	SysYFParserNEQ         = 26
	SysYFParserLT          = 27
	SysYFParserLTE         = 28
	SysYFParserGT          = 29
	SysYFParserGTE         = 30
	SysYFParserNOT         = 31
	SysYFParserLBRACKET    = 32
	SysYFParserRBRACKET    = 33
	SysYFParserNAME        = 34
	SysYFParserINT         = 35
	SysYFParserFLOAT       = 36
	SysYFParserC_COMMENT   = 37
	SysYFParserCPP_COMMENT = 38
	SysYFParserWS          = 39
)

// SysYFParser rules.
const (
	SysYFParserRULE_comp_unit      = 0
	SysYFParserRULE_global_def     = 1
	SysYFParserRULE_const_decl     = 2
	SysYFParserRULE_btype          = 3
	SysYFParserRULE_const_def      = 4
	SysYFParserRULE_const_init_val = 5
	SysYFParserRULE_var_decl       = 6
	SysYFParserRULE_var_def        = 7
	SysYFParserRULE_func_def       = 8
	SysYFParserRULE_func_fparam    = 9
	SysYFParserRULE_block          = 10
	SysYFParserRULE_block_item     = 11
	SysYFParserRULE_stmt           = 12
	SysYFParserRULE_assign_stmt    = 13
	SysYFParserRULE_empty_stmt     = 14
	SysYFParserRULE_exp_stmt       = 15
	SysYFParserRULE_if_stmt        = 16
	SysYFParserRULE_while_stmt     = 17
	SysYFParserRULE_break_stmt     = 18
	SysYFParserRULE_continue_stmt  = 19
	SysYFParserRULE_return_stmt    = 20
	SysYFParserRULE_cond_exp       = 21
	SysYFParserRULE_exp            = 22
	SysYFParserRULE_otherexp       = 23
	SysYFParserRULE_func_call      = 24
	SysYFParserRULE_lval           = 25
	SysYFParserRULE_ident          = 26
	SysYFParserRULE_number         = 27
)

// IComp_unitContext is an interface to support dynamic dispatch.
type IComp_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllGlobal_def() []IGlobal_defContext
	Global_def(i int) IGlobal_defContext

	// IsComp_unitContext differentiates from other interfaces.
	IsComp_unitContext()
}

type Comp_unitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_unitContext() *Comp_unitContext {
	var p = new(Comp_unitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_comp_unit
	return p
}

func InitEmptyComp_unitContext(p *Comp_unitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_comp_unit
}

func (*Comp_unitContext) IsComp_unitContext() {}

func NewComp_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_unitContext {
	var p = new(Comp_unitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_comp_unit

	return p
}

func (s *Comp_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Comp_unitContext) EOF() antlr.TerminalNode {
	return s.GetToken(SysYFParserEOF, 0)
}

func (s *Comp_unitContext) AllGlobal_def() []IGlobal_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobal_defContext); ok {
			len++
		}
	}

	tst := make([]IGlobal_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobal_defContext); ok {
			tst[i] = t.(IGlobal_defContext)
			i++
		}
	}

	return tst
}

func (s *Comp_unitContext) Global_def(i int) IGlobal_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobal_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobal_defContext)
}

func (s *Comp_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comp_unitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterComp_unit(s)
	}
}

func (s *Comp_unitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitComp_unit(s)
	}
}

func (s *Comp_unitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitComp_unit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Comp_unit() (localctx IComp_unitContext) {
	localctx = NewComp_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SysYFParserRULE_comp_unit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229378) != 0) {
		{
			p.SetState(56)
			p.Global_def()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Match(SysYFParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobal_defContext is an interface to support dynamic dispatch.
type IGlobal_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_decl() IConst_declContext
	Var_decl() IVar_declContext
	Func_def() IFunc_defContext

	// IsGlobal_defContext differentiates from other interfaces.
	IsGlobal_defContext()
}

type Global_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobal_defContext() *Global_defContext {
	var p = new(Global_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_global_def
	return p
}

func InitEmptyGlobal_defContext(p *Global_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_global_def
}

func (*Global_defContext) IsGlobal_defContext() {}

func NewGlobal_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_defContext {
	var p = new(Global_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_global_def

	return p
}

func (s *Global_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_defContext) Const_decl() IConst_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_declContext)
}

func (s *Global_defContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Global_defContext) Func_def() IFunc_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_defContext)
}

func (s *Global_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterGlobal_def(s)
	}
}

func (s *Global_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitGlobal_def(s)
	}
}

func (s *Global_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitGlobal_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Global_def() (localctx IGlobal_defContext) {
	localctx = NewGlobal_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SysYFParserRULE_global_def)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Const_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Var_decl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.Func_def()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_declContext is an interface to support dynamic dispatch.
type IConst_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Btype() IBtypeContext
	AllConst_def() []IConst_defContext
	Const_def(i int) IConst_defContext

	// IsConst_declContext differentiates from other interfaces.
	IsConst_declContext()
}

type Const_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_declContext() *Const_declContext {
	var p = new(Const_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_decl
	return p
}

func InitEmptyConst_declContext(p *Const_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_decl
}

func (*Const_declContext) IsConst_declContext() {}

func NewConst_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_declContext {
	var p = new(Const_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_const_decl

	return p
}

func (s *Const_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_declContext) Btype() IBtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBtypeContext)
}

func (s *Const_declContext) AllConst_def() []IConst_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_defContext); ok {
			len++
		}
	}

	tst := make([]IConst_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_defContext); ok {
			tst[i] = t.(IConst_defContext)
			i++
		}
	}

	return tst
}

func (s *Const_declContext) Const_def(i int) IConst_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_defContext)
}

func (s *Const_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterConst_decl(s)
	}
}

func (s *Const_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitConst_decl(s)
	}
}

func (s *Const_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitConst_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Const_decl() (localctx IConst_declContext) {
	localctx = NewConst_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SysYFParserRULE_const_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(SysYFParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Btype()
	}
	{
		p.SetState(70)
		p.Const_def()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SysYFParserT__1 {
		{
			p.SetState(71)
			p.Match(SysYFParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Const_def()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBtypeContext is an interface to support dynamic dispatch.
type IBtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_T() antlr.TerminalNode
	FLOAT_T() antlr.TerminalNode

	// IsBtypeContext differentiates from other interfaces.
	IsBtypeContext()
}

type BtypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBtypeContext() *BtypeContext {
	var p = new(BtypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_btype
	return p
}

func InitEmptyBtypeContext(p *BtypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_btype
}

func (*BtypeContext) IsBtypeContext() {}

func NewBtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BtypeContext {
	var p = new(BtypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_btype

	return p
}

func (s *BtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BtypeContext) INT_T() antlr.TerminalNode {
	return s.GetToken(SysYFParserINT_T, 0)
}

func (s *BtypeContext) FLOAT_T() antlr.TerminalNode {
	return s.GetToken(SysYFParserFLOAT_T, 0)
}

func (s *BtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBtype(s)
	}
}

func (s *BtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBtype(s)
	}
}

func (s *BtypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBtype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Btype() (localctx IBtypeContext) {
	localctx = NewBtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SysYFParserRULE_btype)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SysYFParserINT_T || _la == SysYFParserFLOAT_T) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_defContext is an interface to support dynamic dispatch.
type IConst_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	Const_init_val() IConst_init_valContext
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsConst_defContext differentiates from other interfaces.
	IsConst_defContext()
}

type Const_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_defContext() *Const_defContext {
	var p = new(Const_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_def
	return p
}

func InitEmptyConst_defContext(p *Const_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_def
}

func (*Const_defContext) IsConst_defContext() {}

func NewConst_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_defContext {
	var p = new(Const_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_const_def

	return p
}

func (s *Const_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_defContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Const_defContext) Const_init_val() IConst_init_valContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_init_valContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_init_valContext)
}

func (s *Const_defContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserLBRACKET, 0)
}

func (s *Const_defContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserRBRACKET, 0)
}

func (s *Const_defContext) INT() antlr.TerminalNode {
	return s.GetToken(SysYFParserINT, 0)
}

func (s *Const_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterConst_def(s)
	}
}

func (s *Const_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitConst_def(s)
	}
}

func (s *Const_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitConst_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Const_def() (localctx IConst_defContext) {
	localctx = NewConst_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SysYFParserRULE_const_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Ident()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SysYFParserLBRACKET {
		{
			p.SetState(83)
			p.Match(SysYFParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SysYFParserINT {
			{
				p.SetState(84)
				p.Match(SysYFParserINT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(87)
			p.Match(SysYFParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(90)
		p.Match(SysYFParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Const_init_val()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConst_init_valContext is an interface to support dynamic dispatch.
type IConst_init_valContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConst_init_valContext differentiates from other interfaces.
	IsConst_init_valContext()
}

type Const_init_valContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_init_valContext() *Const_init_valContext {
	var p = new(Const_init_valContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_init_val
	return p
}

func InitEmptyConst_init_valContext(p *Const_init_valContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_const_init_val
}

func (*Const_init_valContext) IsConst_init_valContext() {}

func NewConst_init_valContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_init_valContext {
	var p = new(Const_init_valContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_const_init_val

	return p
}

func (s *Const_init_valContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_init_valContext) CopyAll(ctx *Const_init_valContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Const_init_valContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_init_valContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleContext struct {
	Const_init_valContext
}

func NewSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleContext {
	var p = new(SingleContext)

	InitEmptyConst_init_valContext(&p.Const_init_valContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_init_valContext))

	return p
}

func (s *SingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterSingle(s)
	}
}

func (s *SingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitSingle(s)
	}
}

func (s *SingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayMulContext struct {
	Const_init_valContext
}

func NewArrayMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayMulContext {
	var p = new(ArrayMulContext)

	InitEmptyConst_init_valContext(&p.Const_init_valContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_init_valContext))

	return p
}

func (s *ArrayMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayMulContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ArrayMulContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ArrayMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterArrayMul(s)
	}
}

func (s *ArrayMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitArrayMul(s)
	}
}

func (s *ArrayMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitArrayMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayEmptyContext struct {
	Const_init_valContext
}

func NewArrayEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayEmptyContext {
	var p = new(ArrayEmptyContext)

	InitEmptyConst_init_valContext(&p.Const_init_valContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_init_valContext))

	return p
}

func (s *ArrayEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterArrayEmpty(s)
	}
}

func (s *ArrayEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitArrayEmpty(s)
	}
}

func (s *ArrayEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitArrayEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Const_init_val() (localctx IConst_init_valContext) {
	localctx = NewConst_init_valContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SysYFParserRULE_const_init_val)
	var _la int

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.exp(0)
		}

	case 2:
		localctx = NewArrayEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(SysYFParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(SysYFParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewArrayMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(SysYFParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.exp(0)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SysYFParserT__1 {
			{
				p.SetState(98)
				p.Match(SysYFParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(99)
				p.exp(0)
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(SysYFParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Btype() IBtypeContext
	AllVar_def() []IVar_defContext
	Var_def(i int) IVar_defContext

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) Btype() IBtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBtypeContext)
}

func (s *Var_declContext) AllVar_def() []IVar_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_defContext); ok {
			len++
		}
	}

	tst := make([]IVar_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_defContext); ok {
			tst[i] = t.(IVar_defContext)
			i++
		}
	}

	return tst
}

func (s *Var_declContext) Var_def(i int) IVar_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_defContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (s *Var_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitVar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SysYFParserRULE_var_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Btype()
	}
	{
		p.SetState(110)
		p.Var_def()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SysYFParserT__1 {
		{
			p.SetState(111)
			p.Match(SysYFParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Var_def()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_defContext is an interface to support dynamic dispatch.
type IVar_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	Const_init_val() IConst_init_valContext
	INT() antlr.TerminalNode

	// IsVar_defContext differentiates from other interfaces.
	IsVar_defContext()
}

type Var_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_defContext() *Var_defContext {
	var p = new(Var_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_var_def
	return p
}

func InitEmptyVar_defContext(p *Var_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_var_def
}

func (*Var_defContext) IsVar_defContext() {}

func NewVar_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_defContext {
	var p = new(Var_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_var_def

	return p
}

func (s *Var_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_defContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Var_defContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserLBRACKET, 0)
}

func (s *Var_defContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserRBRACKET, 0)
}

func (s *Var_defContext) Const_init_val() IConst_init_valContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_init_valContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_init_valContext)
}

func (s *Var_defContext) INT() antlr.TerminalNode {
	return s.GetToken(SysYFParserINT, 0)
}

func (s *Var_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterVar_def(s)
	}
}

func (s *Var_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitVar_def(s)
	}
}

func (s *Var_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitVar_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Var_def() (localctx IVar_defContext) {
	localctx = NewVar_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SysYFParserRULE_var_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Ident()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SysYFParserLBRACKET {
		{
			p.SetState(121)
			p.Match(SysYFParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SysYFParserINT {
			{
				p.SetState(122)
				p.Match(SysYFParserINT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(125)
			p.Match(SysYFParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SysYFParserT__3 {
		{
			p.SetState(128)
			p.Match(SysYFParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Const_init_val()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_defContext is an interface to support dynamic dispatch.
type IFunc_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	Block() IBlockContext
	Btype() IBtypeContext
	VOID_T() antlr.TerminalNode
	AllFunc_fparam() []IFunc_fparamContext
	Func_fparam(i int) IFunc_fparamContext

	// IsFunc_defContext differentiates from other interfaces.
	IsFunc_defContext()
}

type Func_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_defContext() *Func_defContext {
	var p = new(Func_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_def
	return p
}

func InitEmptyFunc_defContext(p *Func_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_def
}

func (*Func_defContext) IsFunc_defContext() {}

func NewFunc_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_defContext {
	var p = new(Func_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_func_def

	return p
}

func (s *Func_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_defContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Func_defContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_defContext) Btype() IBtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBtypeContext)
}

func (s *Func_defContext) VOID_T() antlr.TerminalNode {
	return s.GetToken(SysYFParserVOID_T, 0)
}

func (s *Func_defContext) AllFunc_fparam() []IFunc_fparamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_fparamContext); ok {
			len++
		}
	}

	tst := make([]IFunc_fparamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_fparamContext); ok {
			tst[i] = t.(IFunc_fparamContext)
			i++
		}
	}

	return tst
}

func (s *Func_defContext) Func_fparam(i int) IFunc_fparamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_fparamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_fparamContext)
}

func (s *Func_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterFunc_def(s)
	}
}

func (s *Func_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitFunc_def(s)
	}
}

func (s *Func_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitFunc_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Func_def() (localctx IFunc_defContext) {
	localctx = NewFunc_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SysYFParserRULE_func_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SysYFParserINT_T, SysYFParserFLOAT_T:
		{
			p.SetState(132)
			p.Btype()
		}

	case SysYFParserVOID_T:
		{
			p.SetState(133)
			p.Match(SysYFParserVOID_T)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(136)
		p.Ident()
	}
	{
		p.SetState(137)
		p.Match(SysYFParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SysYFParserINT_T || _la == SysYFParserFLOAT_T {
		{
			p.SetState(138)
			p.Func_fparam()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SysYFParserT__1 {
			{
				p.SetState(139)
				p.Match(SysYFParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(140)
				p.Func_fparam()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(148)
		p.Match(SysYFParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_fparamContext is an interface to support dynamic dispatch.
type IFunc_fparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Btype() IBtypeContext
	Ident() IIdentContext
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode

	// IsFunc_fparamContext differentiates from other interfaces.
	IsFunc_fparamContext()
}

type Func_fparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_fparamContext() *Func_fparamContext {
	var p = new(Func_fparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_fparam
	return p
}

func InitEmptyFunc_fparamContext(p *Func_fparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_fparam
}

func (*Func_fparamContext) IsFunc_fparamContext() {}

func NewFunc_fparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_fparamContext {
	var p = new(Func_fparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_func_fparam

	return p
}

func (s *Func_fparamContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_fparamContext) Btype() IBtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBtypeContext)
}

func (s *Func_fparamContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Func_fparamContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserLBRACKET, 0)
}

func (s *Func_fparamContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserRBRACKET, 0)
}

func (s *Func_fparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_fparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_fparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterFunc_fparam(s)
	}
}

func (s *Func_fparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitFunc_fparam(s)
	}
}

func (s *Func_fparamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitFunc_fparam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Func_fparam() (localctx IFunc_fparamContext) {
	localctx = NewFunc_fparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SysYFParserRULE_func_fparam)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Btype()
	}
	{
		p.SetState(152)
		p.Ident()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SysYFParserLBRACKET {
		{
			p.SetState(153)
			p.Match(SysYFParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(SysYFParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBlock_item() []IBlock_itemContext
	Block_item(i int) IBlock_itemContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllBlock_item() []IBlock_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_itemContext); ok {
			len++
		}
	}

	tst := make([]IBlock_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_itemContext); ok {
			tst[i] = t.(IBlock_itemContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Block_item(i int) IBlock_itemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_itemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_itemContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SysYFParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(SysYFParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122407582378) != 0 {
		{
			p.SetState(158)
			p.Block_item()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(SysYFParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_itemContext is an interface to support dynamic dispatch.
type IBlock_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_decl() IConst_declContext
	Var_decl() IVar_declContext
	Stmt() IStmtContext

	// IsBlock_itemContext differentiates from other interfaces.
	IsBlock_itemContext()
}

type Block_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_itemContext() *Block_itemContext {
	var p = new(Block_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_block_item
	return p
}

func InitEmptyBlock_itemContext(p *Block_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_block_item
}

func (*Block_itemContext) IsBlock_itemContext() {}

func NewBlock_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_itemContext {
	var p = new(Block_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_block_item

	return p
}

func (s *Block_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_itemContext) Const_decl() IConst_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_declContext)
}

func (s *Block_itemContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Block_itemContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Block_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBlock_item(s)
	}
}

func (s *Block_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBlock_item(s)
	}
}

func (s *Block_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBlock_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Block_item() (localctx IBlock_itemContext) {
	localctx = NewBlock_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SysYFParserRULE_block_item)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SysYFParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Const_decl()
		}

	case SysYFParserINT_T, SysYFParserFLOAT_T:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Var_decl()
		}

	case SysYFParserT__2, SysYFParserT__4, SysYFParserT__6, SysYFParserT__8, SysYFParserT__10, SysYFParserT__11, SysYFParserT__12, SysYFParserT__13, SysYFParserPLUS, SysYFParserMINUS, SysYFParserNOT, SysYFParserNAME, SysYFParserINT, SysYFParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign_stmt() IAssign_stmtContext
	Empty_stmt() IEmpty_stmtContext
	Exp_stmt() IExp_stmtContext
	Block() IBlockContext
	If_stmt() IIf_stmtContext
	While_stmt() IWhile_stmtContext
	Break_stmt() IBreak_stmtContext
	Continue_stmt() IContinue_stmtContext
	Return_stmt() IReturn_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Empty_stmt() IEmpty_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmpty_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmpty_stmtContext)
}

func (s *StmtContext) Exp_stmt() IExp_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp_stmtContext)
}

func (s *StmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) Break_stmt() IBreak_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *StmtContext) Continue_stmt() IContinue_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
}

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SysYFParserRULE_stmt)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Empty_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Exp_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.Block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.If_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(176)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(177)
			p.Break_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(178)
			p.Continue_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(179)
			p.Return_stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lval() ILvalContext
	Exp() IExpContext

	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) Lval() ILvalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalContext)
}

func (s *Assign_stmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitAssign_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SysYFParserRULE_assign_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Lval()
	}
	{
		p.SetState(183)
		p.Match(SysYFParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.exp(0)
	}
	{
		p.SetState(185)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmpty_stmtContext is an interface to support dynamic dispatch.
type IEmpty_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEmpty_stmtContext differentiates from other interfaces.
	IsEmpty_stmtContext()
}

type Empty_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmpty_stmtContext() *Empty_stmtContext {
	var p = new(Empty_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_empty_stmt
	return p
}

func InitEmptyEmpty_stmtContext(p *Empty_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_empty_stmt
}

func (*Empty_stmtContext) IsEmpty_stmtContext() {}

func NewEmpty_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Empty_stmtContext {
	var p = new(Empty_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_empty_stmt

	return p
}

func (s *Empty_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Empty_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Empty_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterEmpty_stmt(s)
	}
}

func (s *Empty_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitEmpty_stmt(s)
	}
}

func (s *Empty_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitEmpty_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Empty_stmt() (localctx IEmpty_stmtContext) {
	localctx = NewEmpty_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SysYFParserRULE_empty_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExp_stmtContext is an interface to support dynamic dispatch.
type IExp_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsExp_stmtContext differentiates from other interfaces.
	IsExp_stmtContext()
}

type Exp_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_stmtContext() *Exp_stmtContext {
	var p = new(Exp_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_exp_stmt
	return p
}

func InitEmptyExp_stmtContext(p *Exp_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_exp_stmt
}

func (*Exp_stmtContext) IsExp_stmtContext() {}

func NewExp_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_stmtContext {
	var p = new(Exp_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_exp_stmt

	return p
}

func (s *Exp_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_stmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Exp_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterExp_stmt(s)
	}
}

func (s *Exp_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitExp_stmt(s)
	}
}

func (s *Exp_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitExp_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Exp_stmt() (localctx IExp_stmtContext) {
	localctx = NewExp_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SysYFParserRULE_exp_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.exp(0)
	}
	{
		p.SetState(190)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cond_exp() ICond_expContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) Cond_exp() ICond_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICond_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICond_expContext)
}

func (s *If_stmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SysYFParserRULE_if_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(SysYFParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(SysYFParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.cond_exp(0)
	}
	{
		p.SetState(195)
		p.Match(SysYFParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Stmt()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(197)
			p.Match(SysYFParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Stmt()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cond_exp() ICond_expContext
	Stmt() IStmtContext

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) Cond_exp() ICond_expContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICond_expContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICond_expContext)
}

func (s *While_stmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (s *While_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitWhile_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SysYFParserRULE_while_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(SysYFParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(SysYFParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.cond_exp(0)
	}
	{
		p.SetState(204)
		p.Match(SysYFParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Stmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_break_stmt
	return p
}

func InitEmptyBreak_stmtContext(p *Break_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_break_stmt
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBreak_stmt(s)
	}
}

func (s *Break_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBreak_stmt(s)
	}
}

func (s *Break_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBreak_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Break_stmt() (localctx IBreak_stmtContext) {
	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SysYFParserRULE_break_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(SysYFParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_continue_stmt
	return p
}

func InitEmptyContinue_stmtContext(p *Continue_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_continue_stmt
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterContinue_stmt(s)
	}
}

func (s *Continue_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitContinue_stmt(s)
	}
}

func (s *Continue_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitContinue_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Continue_stmt() (localctx IContinue_stmtContext) {
	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SysYFParserRULE_continue_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(SysYFParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_return_stmt
	return p
}

func InitEmptyReturn_stmtContext(p *Return_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_return_stmt
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterReturn_stmt(s)
	}
}

func (s *Return_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitReturn_stmt(s)
	}
}

func (s *Return_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitReturn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SysYFParserRULE_return_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(SysYFParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122407354496) != 0 {
		{
			p.SetState(214)
			p.exp(0)
		}

	}
	{
		p.SetState(217)
		p.Match(SysYFParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICond_expContext is an interface to support dynamic dispatch.
type ICond_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	AllCond_exp() []ICond_expContext
	Cond_exp(i int) ICond_expContext
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	EQ() antlr.TerminalNode
	LAND() antlr.TerminalNode
	LOR() antlr.TerminalNode

	// IsCond_expContext differentiates from other interfaces.
	IsCond_expContext()
}

type Cond_expContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCond_expContext() *Cond_expContext {
	var p = new(Cond_expContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_cond_exp
	return p
}

func InitEmptyCond_expContext(p *Cond_expContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_cond_exp
}

func (*Cond_expContext) IsCond_expContext() {}

func NewCond_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cond_expContext {
	var p = new(Cond_expContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_cond_exp

	return p
}

func (s *Cond_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Cond_expContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Cond_expContext) AllCond_exp() []ICond_expContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICond_expContext); ok {
			len++
		}
	}

	tst := make([]ICond_expContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICond_expContext); ok {
			tst[i] = t.(ICond_expContext)
			i++
		}
	}

	return tst
}

func (s *Cond_expContext) Cond_exp(i int) ICond_expContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICond_expContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICond_expContext)
}

func (s *Cond_expContext) LT() antlr.TerminalNode {
	return s.GetToken(SysYFParserLT, 0)
}

func (s *Cond_expContext) GT() antlr.TerminalNode {
	return s.GetToken(SysYFParserGT, 0)
}

func (s *Cond_expContext) LTE() antlr.TerminalNode {
	return s.GetToken(SysYFParserLTE, 0)
}

func (s *Cond_expContext) GTE() antlr.TerminalNode {
	return s.GetToken(SysYFParserGTE, 0)
}

func (s *Cond_expContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SysYFParserNEQ, 0)
}

func (s *Cond_expContext) EQ() antlr.TerminalNode {
	return s.GetToken(SysYFParserEQ, 0)
}

func (s *Cond_expContext) LAND() antlr.TerminalNode {
	return s.GetToken(SysYFParserLAND, 0)
}

func (s *Cond_expContext) LOR() antlr.TerminalNode {
	return s.GetToken(SysYFParserLOR, 0)
}

func (s *Cond_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cond_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterCond_exp(s)
	}
}

func (s *Cond_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitCond_exp(s)
	}
}

func (s *Cond_expContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitCond_exp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Cond_exp() (localctx ICond_expContext) {
	return p.cond_exp(0)
}

func (p *SysYFParser) cond_exp(_p int) (localctx ICond_expContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCond_expContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICond_expContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SysYFParserRULE_cond_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.exp(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCond_expContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_cond_exp)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(223)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2013265920) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(224)
					p.cond_exp(5)
				}

			case 2:
				localctx = NewCond_expContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_cond_exp)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(226)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SysYFParserEQ || _la == SysYFParserNEQ) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(227)
					p.cond_exp(4)
				}

			case 3:
				localctx = NewCond_expContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_cond_exp)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(229)
					p.Match(SysYFParserLAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(230)
					p.cond_exp(3)
				}

			case 4:
				localctx = NewCond_expContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_cond_exp)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(232)
					p.Match(SysYFParserLOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(233)
					p.cond_exp(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyAll(ctx *ExpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryHighContext struct {
	ExpContext
}

func NewBinaryHighContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryHighContext {
	var p = new(BinaryHighContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *BinaryHighContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryHighContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *BinaryHighContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *BinaryHighContext) MUL() antlr.TerminalNode {
	return s.GetToken(SysYFParserMUL, 0)
}

func (s *BinaryHighContext) DIV() antlr.TerminalNode {
	return s.GetToken(SysYFParserDIV, 0)
}

func (s *BinaryHighContext) MOD() antlr.TerminalNode {
	return s.GetToken(SysYFParserMOD, 0)
}

func (s *BinaryHighContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBinaryHigh(s)
	}
}

func (s *BinaryHighContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBinaryHigh(s)
	}
}

func (s *BinaryHighContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBinaryHigh(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryLowContext struct {
	ExpContext
}

func NewBinaryLowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLowContext {
	var p = new(BinaryLowContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *BinaryLowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLowContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *BinaryLowContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *BinaryLowContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SysYFParserPLUS, 0)
}

func (s *BinaryLowContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SysYFParserMINUS, 0)
}

func (s *BinaryLowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterBinaryLow(s)
	}
}

func (s *BinaryLowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitBinaryLow(s)
	}
}

func (s *BinaryLowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitBinaryLow(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryContext struct {
	ExpContext
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *UnaryContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SysYFParserPLUS, 0)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SysYFParserMINUS, 0)
}

func (s *UnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(SysYFParserNOT, 0)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type OtherexpAltContext struct {
	ExpContext
}

func NewOtherexpAltContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OtherexpAltContext {
	var p = new(OtherexpAltContext)

	InitEmptyExpContext(&p.ExpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpContext))

	return p
}

func (s *OtherexpAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherexpAltContext) Otherexp() IOtherexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOtherexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOtherexpContext)
}

func (s *OtherexpAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterOtherexpAlt(s)
	}
}

func (s *OtherexpAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitOtherexpAlt(s)
	}
}

func (s *OtherexpAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitOtherexpAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *SysYFParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, SysYFParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SysYFParserT__6, SysYFParserNAME, SysYFParserINT, SysYFParserFLOAT:
		localctx = NewOtherexpAltContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(240)
			p.Otherexp()
		}

	case SysYFParserPLUS, SysYFParserMINUS, SysYFParserNOT:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2148270080) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(242)
			p.exp(3)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryHighContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_exp)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(246)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(247)
					p.exp(3)
				}

			case 2:
				localctx = NewBinaryLowContext(p, NewExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SysYFParserRULE_exp)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(249)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SysYFParserPLUS || _la == SysYFParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(250)
					p.exp(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOtherexpContext is an interface to support dynamic dispatch.
type IOtherexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lval() ILvalContext
	Number() INumberContext
	Func_call() IFunc_callContext
	Exp() IExpContext

	// IsOtherexpContext differentiates from other interfaces.
	IsOtherexpContext()
}

type OtherexpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOtherexpContext() *OtherexpContext {
	var p = new(OtherexpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_otherexp
	return p
}

func InitEmptyOtherexpContext(p *OtherexpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_otherexp
}

func (*OtherexpContext) IsOtherexpContext() {}

func NewOtherexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OtherexpContext {
	var p = new(OtherexpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_otherexp

	return p
}

func (s *OtherexpContext) GetParser() antlr.Parser { return s.parser }

func (s *OtherexpContext) Lval() ILvalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalContext)
}

func (s *OtherexpContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *OtherexpContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *OtherexpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *OtherexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OtherexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterOtherexp(s)
	}
}

func (s *OtherexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitOtherexp(s)
	}
}

func (s *OtherexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitOtherexp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Otherexp() (localctx IOtherexpContext) {
	localctx = NewOtherexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SysYFParserRULE_otherexp)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Lval()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(258)
			p.Func_call()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(259)
			p.Match(SysYFParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.exp(0)
		}
		{
			p.SetState(261)
			p.Match(SysYFParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	AllExp() []IExpContext
	Exp(i int) IExpContext

	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Func_callContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *Func_callContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterFunc_call(s)
	}
}

func (s *Func_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitFunc_call(s)
	}
}

func (s *Func_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitFunc_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SysYFParserRULE_func_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Ident()
	}
	{
		p.SetState(266)
		p.Match(SysYFParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122407354496) != 0 {
		{
			p.SetState(267)
			p.exp(0)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SysYFParserT__1 {
			{
				p.SetState(268)
				p.Match(SysYFParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(269)
				p.exp(0)
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(277)
		p.Match(SysYFParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvalContext is an interface to support dynamic dispatch.
type ILvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	LBRACKET() antlr.TerminalNode
	Exp() IExpContext
	RBRACKET() antlr.TerminalNode

	// IsLvalContext differentiates from other interfaces.
	IsLvalContext()
}

type LvalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalContext() *LvalContext {
	var p = new(LvalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_lval
	return p
}

func InitEmptyLvalContext(p *LvalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_lval
}

func (*LvalContext) IsLvalContext() {}

func NewLvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalContext {
	var p = new(LvalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_lval

	return p
}

func (s *LvalContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *LvalContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserLBRACKET, 0)
}

func (s *LvalContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *LvalContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SysYFParserRBRACKET, 0)
}

func (s *LvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterLval(s)
	}
}

func (s *LvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitLval(s)
	}
}

func (s *LvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitLval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Lval() (localctx ILvalContext) {
	localctx = NewLvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SysYFParserRULE_lval)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Ident()
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(280)
			p.Match(SysYFParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.exp(0)
		}
		{
			p.SetState(282)
			p.Match(SysYFParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) NAME() antlr.TerminalNode {
	return s.GetToken(SysYFParserNAME, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SysYFParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(SysYFParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SysYFParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SysYFParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(SysYFParserINT, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SysYFParserFLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SysYFListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SysYFVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SysYFParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SysYFParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SysYFParserINT || _la == SysYFParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SysYFParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *Cond_expContext = nil
		if localctx != nil {
			t = localctx.(*Cond_expContext)
		}
		return p.Cond_exp_Sempred(t, predIndex)

	case 22:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SysYFParser) Cond_exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SysYFParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
