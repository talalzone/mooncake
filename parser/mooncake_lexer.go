// Code generated from /Users/talal/go/src/mooncake/grammar/Mooncake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 307, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 108, 10, 12, 3, 13, 3, 
	13, 3, 13, 3, 13, 5, 13, 114, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 
	5, 14, 121, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 127, 10, 15, 3, 
	16, 3, 16, 3, 16, 5, 16, 132, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 137, 
	10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 144, 10, 18, 3, 19, 3, 
	19, 3, 19, 3, 19, 3, 19, 5, 19, 151, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 
	3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 
	25, 3, 26, 5, 26, 211, 10, 26, 3, 26, 3, 26, 3, 26, 6, 26, 216, 10, 26, 
	13, 26, 14, 26, 217, 3, 27, 3, 27, 7, 27, 222, 10, 27, 12, 27, 14, 27, 
	225, 11, 27, 3, 28, 3, 28, 5, 28, 229, 10, 28, 3, 29, 3, 29, 3, 29, 3, 
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 
	3, 31, 3, 32, 3, 32, 7, 32, 248, 10, 32, 12, 32, 14, 32, 251, 11, 32, 3, 
	32, 3, 32, 3, 33, 6, 33, 256, 10, 33, 13, 33, 14, 33, 257, 3, 33, 3, 33, 
	3, 34, 6, 34, 263, 10, 34, 13, 34, 14, 34, 264, 3, 34, 3, 34, 3, 35, 3, 
	35, 3, 35, 3, 35, 3, 35, 6, 35, 274, 10, 35, 13, 35, 14, 35, 275, 3, 35, 
	3, 35, 5, 35, 280, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 
	37, 3, 37, 6, 37, 290, 10, 37, 13, 37, 14, 37, 291, 3, 38, 6, 38, 295, 
	10, 38, 13, 38, 14, 38, 296, 3, 39, 3, 39, 7, 39, 301, 10, 39, 12, 39, 
	14, 39, 304, 11, 39, 3, 39, 3, 39, 3, 302, 2, 40, 3, 3, 5, 4, 7, 5, 9, 
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 
	65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 3, 2, 7, 3, 2, 
	50, 59, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 4, 2, 67, 92, 99, 124, 
	5, 2, 50, 59, 67, 92, 99, 124, 2, 327, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 5, 81, 3, 2, 2, 2, 7, 
	83, 3, 2, 2, 2, 9, 86, 3, 2, 2, 2, 11, 88, 3, 2, 2, 2, 13, 90, 3, 2, 2, 
	2, 15, 92, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19, 99, 3, 2, 2, 2, 21, 101, 
	3, 2, 2, 2, 23, 107, 3, 2, 2, 2, 25, 113, 3, 2, 2, 2, 27, 120, 3, 2, 2, 
	2, 29, 126, 3, 2, 2, 2, 31, 131, 3, 2, 2, 2, 33, 136, 3, 2, 2, 2, 35, 143, 
	3, 2, 2, 2, 37, 150, 3, 2, 2, 2, 39, 152, 3, 2, 2, 2, 41, 157, 3, 2, 2, 
	2, 43, 164, 3, 2, 2, 2, 45, 178, 3, 2, 2, 2, 47, 196, 3, 2, 2, 2, 49, 203, 
	3, 2, 2, 2, 51, 210, 3, 2, 2, 2, 53, 219, 3, 2, 2, 2, 55, 228, 3, 2, 2, 
	2, 57, 230, 3, 2, 2, 2, 59, 235, 3, 2, 2, 2, 61, 241, 3, 2, 2, 2, 63, 245, 
	3, 2, 2, 2, 65, 255, 3, 2, 2, 2, 67, 262, 3, 2, 2, 2, 69, 273, 3, 2, 2, 
	2, 71, 281, 3, 2, 2, 2, 73, 287, 3, 2, 2, 2, 75, 294, 3, 2, 2, 2, 77, 298, 
	3, 2, 2, 2, 79, 80, 7, 125, 2, 2, 80, 4, 3, 2, 2, 2, 81, 82, 7, 127, 2, 
	2, 82, 6, 3, 2, 2, 2, 83, 84, 7, 63, 2, 2, 84, 85, 7, 64, 2, 2, 85, 8, 
	3, 2, 2, 2, 86, 87, 7, 93, 2, 2, 87, 10, 3, 2, 2, 2, 88, 89, 7, 46, 2, 
	2, 89, 12, 3, 2, 2, 2, 90, 91, 7, 95, 2, 2, 91, 14, 3, 2, 2, 2, 92, 93, 
	7, 35, 2, 2, 93, 94, 7, 35, 2, 2, 94, 95, 7, 35, 2, 2, 95, 16, 3, 2, 2, 
	2, 96, 97, 7, 35, 2, 2, 97, 98, 7, 35, 2, 2, 98, 18, 3, 2, 2, 2, 99, 100, 
	7, 35, 2, 2, 100, 20, 3, 2, 2, 2, 101, 102, 7, 128, 2, 2, 102, 22, 3, 2, 
	2, 2, 103, 104, 7, 103, 2, 2, 104, 108, 7, 115, 2, 2, 105, 106, 7, 63, 
	2, 2, 106, 108, 7, 63, 2, 2, 107, 103, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 
	108, 24, 3, 2, 2, 2, 109, 110, 7, 112, 2, 2, 110, 114, 7, 103, 2, 2, 111, 
	112, 7, 35, 2, 2, 112, 114, 7, 63, 2, 2, 113, 109, 3, 2, 2, 2, 113, 111, 
	3, 2, 2, 2, 114, 26, 3, 2, 2, 2, 115, 116, 7, 99, 2, 2, 116, 117, 7, 112, 
	2, 2, 117, 121, 7, 102, 2, 2, 118, 119, 7, 40, 2, 2, 119, 121, 7, 40, 2, 
	2, 120, 115, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 28, 3, 2, 2, 2, 122, 
	123, 7, 113, 2, 2, 123, 127, 7, 116, 2, 2, 124, 125, 7, 126, 2, 2, 125, 
	127, 7, 126, 2, 2, 126, 122, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 30, 
	3, 2, 2, 2, 128, 129, 7, 105, 2, 2, 129, 132, 7, 118, 2, 2, 130, 132, 7, 
	64, 2, 2, 131, 128, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 32, 3, 2, 2, 
	2, 133, 134, 7, 110, 2, 2, 134, 137, 7, 118, 2, 2, 135, 137, 7, 62, 2, 
	2, 136, 133, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 34, 3, 2, 2, 2, 138, 
	139, 7, 105, 2, 2, 139, 140, 7, 118, 2, 2, 140, 144, 7, 103, 2, 2, 141, 
	142, 7, 64, 2, 2, 142, 144, 7, 63, 2, 2, 143, 138, 3, 2, 2, 2, 143, 141, 
	3, 2, 2, 2, 144, 36, 3, 2, 2, 2, 145, 146, 7, 110, 2, 2, 146, 147, 7, 118, 
	2, 2, 147, 151, 7, 103, 2, 2, 148, 149, 7, 62, 2, 2, 149, 151, 7, 63, 2, 
	2, 150, 145, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 38, 3, 2, 2, 2, 152, 
	153, 7, 66, 2, 2, 153, 154, 7, 110, 2, 2, 154, 155, 7, 103, 2, 2, 155, 
	156, 7, 112, 2, 2, 156, 40, 3, 2, 2, 2, 157, 158, 7, 66, 2, 2, 158, 159, 
	7, 104, 2, 2, 159, 160, 7, 110, 2, 2, 160, 161, 7, 113, 2, 2, 161, 162, 
	7, 99, 2, 2, 162, 163, 7, 118, 2, 2, 163, 42, 3, 2, 2, 2, 164, 165, 7, 
	66, 2, 2, 165, 166, 7, 102, 2, 2, 166, 167, 7, 99, 2, 2, 167, 168, 7, 118, 
	2, 2, 168, 169, 7, 103, 2, 2, 169, 170, 7, 86, 2, 2, 170, 171, 7, 107, 
	2, 2, 171, 172, 7, 111, 2, 2, 172, 173, 7, 103, 2, 2, 173, 174, 7, 78, 
	2, 2, 174, 175, 7, 113, 2, 2, 175, 176, 7, 112, 2, 2, 176, 177, 7, 105, 
	2, 2, 177, 44, 3, 2, 2, 2, 178, 179, 7, 66, 2, 2, 179, 180, 7, 99, 2, 2, 
	180, 181, 7, 104, 2, 2, 181, 182, 7, 118, 2, 2, 182, 183, 7, 103, 2, 2, 
	183, 184, 7, 116, 2, 2, 184, 185, 7, 69, 2, 2, 185, 186, 7, 119, 2, 2, 
	186, 187, 7, 116, 2, 2, 187, 188, 7, 116, 2, 2, 188, 189, 7, 103, 2, 2, 
	189, 190, 7, 112, 2, 2, 190, 191, 7, 118, 2, 2, 191, 192, 7, 86, 2, 2, 
	192, 193, 7, 107, 2, 2, 193, 194, 7, 111, 2, 2, 194, 195, 7, 103, 2, 2, 
	195, 46, 3, 2, 2, 2, 196, 197, 7, 103, 2, 2, 197, 198, 7, 122, 2, 2, 198, 
	199, 7, 107, 2, 2, 199, 200, 7, 117, 2, 2, 200, 201, 7, 118, 2, 2, 201, 
	202, 7, 117, 2, 2, 202, 48, 3, 2, 2, 2, 203, 204, 7, 103, 2, 2, 204, 205, 
	7, 111, 2, 2, 205, 206, 7, 114, 2, 2, 206, 207, 7, 118, 2, 2, 207, 208, 
	7, 123, 2, 2, 208, 50, 3, 2, 2, 2, 209, 211, 7, 47, 2, 2, 210, 209, 3, 
	2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 5, 53, 27, 
	2, 213, 215, 7, 48, 2, 2, 214, 216, 9, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 
	217, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 52, 3, 
	2, 2, 2, 219, 223, 9, 2, 2, 2, 220, 222, 9, 2, 2, 2, 221, 220, 3, 2, 2, 
	2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 
	54, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 229, 5, 57, 29, 2, 227, 229, 
	5, 59, 30, 2, 228, 226, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229, 56, 3, 2, 
	2, 2, 230, 231, 7, 118, 2, 2, 231, 232, 7, 116, 2, 2, 232, 233, 7, 119, 
	2, 2, 233, 234, 7, 103, 2, 2, 234, 58, 3, 2, 2, 2, 235, 236, 7, 104, 2, 
	2, 236, 237, 7, 99, 2, 2, 237, 238, 7, 110, 2, 2, 238, 239, 7, 117, 2, 
	2, 239, 240, 7, 103, 2, 2, 240, 60, 3, 2, 2, 2, 241, 242, 7, 112, 2, 2, 
	242, 243, 7, 107, 2, 2, 243, 244, 7, 110, 2, 2, 244, 62, 3, 2, 2, 2, 245, 
	249, 7, 37, 2, 2, 246, 248, 10, 3, 2, 2, 247, 246, 3, 2, 2, 2, 248, 251, 
	3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 
	2, 2, 251, 249, 3, 2, 2, 2, 252, 253, 8, 32, 2, 2, 253, 64, 3, 2, 2, 2, 
	254, 256, 9, 4, 2, 2, 255, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 
	255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 
	8, 33, 3, 2, 260, 66, 3, 2, 2, 2, 261, 263, 9, 3, 2, 2, 262, 261, 3, 2, 
	2, 2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 
	265, 266, 3, 2, 2, 2, 266, 267, 8, 34, 3, 2, 267, 68, 3, 2, 2, 2, 268, 
	274, 9, 5, 2, 2, 269, 270, 7, 93, 2, 2, 270, 271, 5, 53, 27, 2, 271, 272, 
	7, 95, 2, 2, 272, 274, 3, 2, 2, 2, 273, 268, 3, 2, 2, 2, 273, 269, 3, 2, 
	2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 
	276, 279, 3, 2, 2, 2, 277, 278, 7, 48, 2, 2, 278, 280, 5, 69, 35, 2, 279, 
	277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 70, 3, 2, 2, 2, 281, 282, 7, 
	38, 2, 2, 282, 283, 7, 125, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 5, 69, 
	35, 2, 285, 286, 7, 127, 2, 2, 286, 72, 3, 2, 2, 2, 287, 289, 7, 97, 2, 
	2, 288, 290, 9, 5, 2, 2, 289, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 
	289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 74, 3, 2, 2, 2, 293, 295, 9, 
	6, 2, 2, 294, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 294, 3, 2, 2, 
	2, 296, 297, 3, 2, 2, 2, 297, 76, 3, 2, 2, 2, 298, 302, 7, 41, 2, 2, 299, 
	301, 11, 2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 303, 
	3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 303, 305, 3, 2, 2, 2, 304, 302, 3, 2, 
	2, 2, 305, 306, 7, 41, 2, 2, 306, 78, 3, 2, 2, 2, 25, 2, 107, 113, 120, 
	126, 131, 136, 143, 150, 210, 217, 223, 228, 249, 257, 264, 273, 275, 279, 
	291, 294, 296, 302, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'", "'=>'", "'['", "','", "']'", "'!!!'", "'!!'", "'!'", 
	"'~'", "", "", "", "", "", "", "", "", "'@len'", "'@float'", "'@dateTimeLong'", 
	"'@afterCurrentTime'", "'exists'", "'empty'", "", "", "", "'true'", "'false'", 
	"'nil'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "FATAL", "SEVERE", "WARNING", "LINKED", "EQ", 
	"NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", "FLOAT_FUNC", 
	"DATETIME_LONG", "AFTER_CURR_TIME", "EXISTS", "EMPTY", "FLOAT", "INT", 
	"BOOL", "TRUE", "FALSE", "NULL", "COMMENT", "WS", "TERMINATOR", "ATTR_ID", 
	"CTX_ID", "DECL_ID", "ERROR_CODE", "ERROR_INFO",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "FATAL", "SEVERE", "WARNING", 
	"LINKED", "EQ", "NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", 
	"FLOAT_FUNC", "DATETIME_LONG", "AFTER_CURR_TIME", "EXISTS", "EMPTY", "FLOAT", 
	"INT", "BOOL", "TRUE", "FALSE", "NULL", "COMMENT", "WS", "TERMINATOR", 
	"ATTR_ID", "CTX_ID", "DECL_ID", "ERROR_CODE", "ERROR_INFO",
}

type MooncakeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMooncakeLexer(input antlr.CharStream) *MooncakeLexer {

	l := new(MooncakeLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Mooncake.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MooncakeLexer tokens.
const (
	MooncakeLexerT__0 = 1
	MooncakeLexerT__1 = 2
	MooncakeLexerT__2 = 3
	MooncakeLexerT__3 = 4
	MooncakeLexerT__4 = 5
	MooncakeLexerT__5 = 6
	MooncakeLexerFATAL = 7
	MooncakeLexerSEVERE = 8
	MooncakeLexerWARNING = 9
	MooncakeLexerLINKED = 10
	MooncakeLexerEQ = 11
	MooncakeLexerNE = 12
	MooncakeLexerAND = 13
	MooncakeLexerOR = 14
	MooncakeLexerGT = 15
	MooncakeLexerLT = 16
	MooncakeLexerGTE = 17
	MooncakeLexerLTE = 18
	MooncakeLexerLEN_FUNC = 19
	MooncakeLexerFLOAT_FUNC = 20
	MooncakeLexerDATETIME_LONG = 21
	MooncakeLexerAFTER_CURR_TIME = 22
	MooncakeLexerEXISTS = 23
	MooncakeLexerEMPTY = 24
	MooncakeLexerFLOAT = 25
	MooncakeLexerINT = 26
	MooncakeLexerBOOL = 27
	MooncakeLexerTRUE = 28
	MooncakeLexerFALSE = 29
	MooncakeLexerNULL = 30
	MooncakeLexerCOMMENT = 31
	MooncakeLexerWS = 32
	MooncakeLexerTERMINATOR = 33
	MooncakeLexerATTR_ID = 34
	MooncakeLexerCTX_ID = 35
	MooncakeLexerDECL_ID = 36
	MooncakeLexerERROR_CODE = 37
	MooncakeLexerERROR_INFO = 38
)

