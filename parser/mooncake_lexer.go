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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 37, 327, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 
	12, 122, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 128, 10, 13, 3, 14, 
	3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 135, 10, 14, 3, 15, 3, 15, 3, 15, 3, 
	15, 5, 15, 141, 10, 15, 3, 16, 3, 16, 3, 16, 5, 16, 146, 10, 16, 3, 17, 
	3, 17, 3, 17, 5, 17, 151, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 
	18, 158, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 165, 10, 19, 
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 6, 24, 214, 10, 24, 13, 24, 14, 24, 
	215, 3, 25, 5, 25, 219, 10, 25, 3, 25, 6, 25, 222, 10, 25, 13, 25, 14, 
	25, 223, 3, 26, 3, 26, 5, 26, 228, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 
	3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 
	30, 3, 30, 3, 30, 3, 31, 3, 31, 7, 31, 249, 10, 31, 12, 31, 14, 31, 252, 
	11, 31, 3, 31, 3, 31, 3, 32, 6, 32, 257, 10, 32, 13, 32, 14, 32, 258, 3, 
	32, 3, 32, 3, 33, 6, 33, 264, 10, 33, 13, 33, 14, 33, 265, 3, 33, 3, 33, 
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 6, 35, 279, 
	10, 35, 13, 35, 14, 35, 280, 3, 35, 3, 35, 5, 35, 285, 10, 35, 3, 36, 3, 
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 
	5, 39, 299, 10, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 
	43, 3, 43, 7, 43, 310, 10, 43, 12, 43, 14, 43, 313, 11, 43, 3, 43, 3, 43, 
	3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 321, 10, 44, 3, 45, 3, 45, 3, 45, 3, 
	46, 3, 46, 2, 2, 47, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 
	55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 
	73, 2, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 3, 
	2, 8, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 4, 2, 67, 92, 99, 124, 
	3, 2, 50, 59, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 10, 2, 36, 36, 41, 
	41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 2, 340, 2, 
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 
	57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 
	2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 
	2, 3, 93, 3, 2, 2, 2, 5, 95, 3, 2, 2, 2, 7, 97, 3, 2, 2, 2, 9, 100, 3, 
	2, 2, 2, 11, 102, 3, 2, 2, 2, 13, 104, 3, 2, 2, 2, 15, 106, 3, 2, 2, 2, 
	17, 110, 3, 2, 2, 2, 19, 113, 3, 2, 2, 2, 21, 115, 3, 2, 2, 2, 23, 121, 
	3, 2, 2, 2, 25, 127, 3, 2, 2, 2, 27, 134, 3, 2, 2, 2, 29, 140, 3, 2, 2, 
	2, 31, 145, 3, 2, 2, 2, 33, 150, 3, 2, 2, 2, 35, 157, 3, 2, 2, 2, 37, 164, 
	3, 2, 2, 2, 39, 166, 3, 2, 2, 2, 41, 171, 3, 2, 2, 2, 43, 178, 3, 2, 2, 
	2, 45, 192, 3, 2, 2, 2, 47, 210, 3, 2, 2, 2, 49, 218, 3, 2, 2, 2, 51, 227, 
	3, 2, 2, 2, 53, 229, 3, 2, 2, 2, 55, 233, 3, 2, 2, 2, 57, 235, 3, 2, 2, 
	2, 59, 240, 3, 2, 2, 2, 61, 246, 3, 2, 2, 2, 63, 256, 3, 2, 2, 2, 65, 263, 
	3, 2, 2, 2, 67, 269, 3, 2, 2, 2, 69, 278, 3, 2, 2, 2, 71, 286, 3, 2, 2, 
	2, 73, 292, 3, 2, 2, 2, 75, 294, 3, 2, 2, 2, 77, 298, 3, 2, 2, 2, 79, 300, 
	3, 2, 2, 2, 81, 302, 3, 2, 2, 2, 83, 304, 3, 2, 2, 2, 85, 306, 3, 2, 2, 
	2, 87, 316, 3, 2, 2, 2, 89, 322, 3, 2, 2, 2, 91, 325, 3, 2, 2, 2, 93, 94, 
	7, 125, 2, 2, 94, 4, 3, 2, 2, 2, 95, 96, 7, 127, 2, 2, 96, 6, 3, 2, 2, 
	2, 97, 98, 7, 63, 2, 2, 98, 99, 7, 64, 2, 2, 99, 8, 3, 2, 2, 2, 100, 101, 
	7, 93, 2, 2, 101, 10, 3, 2, 2, 2, 102, 103, 7, 46, 2, 2, 103, 12, 3, 2, 
	2, 2, 104, 105, 7, 95, 2, 2, 105, 14, 3, 2, 2, 2, 106, 107, 7, 35, 2, 2, 
	107, 108, 7, 35, 2, 2, 108, 109, 7, 35, 2, 2, 109, 16, 3, 2, 2, 2, 110, 
	111, 7, 35, 2, 2, 111, 112, 7, 35, 2, 2, 112, 18, 3, 2, 2, 2, 113, 114, 
	7, 35, 2, 2, 114, 20, 3, 2, 2, 2, 115, 116, 7, 128, 2, 2, 116, 22, 3, 2, 
	2, 2, 117, 118, 7, 103, 2, 2, 118, 122, 7, 115, 2, 2, 119, 120, 7, 63, 
	2, 2, 120, 122, 7, 63, 2, 2, 121, 117, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 
	122, 24, 3, 2, 2, 2, 123, 124, 7, 112, 2, 2, 124, 128, 7, 103, 2, 2, 125, 
	126, 7, 35, 2, 2, 126, 128, 7, 63, 2, 2, 127, 123, 3, 2, 2, 2, 127, 125, 
	3, 2, 2, 2, 128, 26, 3, 2, 2, 2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 112, 
	2, 2, 131, 135, 7, 102, 2, 2, 132, 133, 7, 40, 2, 2, 133, 135, 7, 40, 2, 
	2, 134, 129, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 28, 3, 2, 2, 2, 136, 
	137, 7, 113, 2, 2, 137, 141, 7, 116, 2, 2, 138, 139, 7, 126, 2, 2, 139, 
	141, 7, 126, 2, 2, 140, 136, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 30, 
	3, 2, 2, 2, 142, 143, 7, 105, 2, 2, 143, 146, 7, 118, 2, 2, 144, 146, 7, 
	64, 2, 2, 145, 142, 3, 2, 2, 2, 145, 144, 3, 2, 2, 2, 146, 32, 3, 2, 2, 
	2, 147, 148, 7, 110, 2, 2, 148, 151, 7, 118, 2, 2, 149, 151, 7, 62, 2, 
	2, 150, 147, 3, 2, 2, 2, 150, 149, 3, 2, 2, 2, 151, 34, 3, 2, 2, 2, 152, 
	153, 7, 105, 2, 2, 153, 154, 7, 118, 2, 2, 154, 158, 7, 103, 2, 2, 155, 
	156, 7, 64, 2, 2, 156, 158, 7, 63, 2, 2, 157, 152, 3, 2, 2, 2, 157, 155, 
	3, 2, 2, 2, 158, 36, 3, 2, 2, 2, 159, 160, 7, 110, 2, 2, 160, 161, 7, 118, 
	2, 2, 161, 165, 7, 103, 2, 2, 162, 163, 7, 62, 2, 2, 163, 165, 7, 63, 2, 
	2, 164, 159, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 38, 3, 2, 2, 2, 166, 
	167, 7, 66, 2, 2, 167, 168, 7, 110, 2, 2, 168, 169, 7, 103, 2, 2, 169, 
	170, 7, 112, 2, 2, 170, 40, 3, 2, 2, 2, 171, 172, 7, 66, 2, 2, 172, 173, 
	7, 104, 2, 2, 173, 174, 7, 110, 2, 2, 174, 175, 7, 113, 2, 2, 175, 176, 
	7, 99, 2, 2, 176, 177, 7, 118, 2, 2, 177, 42, 3, 2, 2, 2, 178, 179, 7, 
	66, 2, 2, 179, 180, 7, 102, 2, 2, 180, 181, 7, 99, 2, 2, 181, 182, 7, 118, 
	2, 2, 182, 183, 7, 103, 2, 2, 183, 184, 7, 86, 2, 2, 184, 185, 7, 107, 
	2, 2, 185, 186, 7, 111, 2, 2, 186, 187, 7, 103, 2, 2, 187, 188, 7, 78, 
	2, 2, 188, 189, 7, 113, 2, 2, 189, 190, 7, 112, 2, 2, 190, 191, 7, 105, 
	2, 2, 191, 44, 3, 2, 2, 2, 192, 193, 7, 66, 2, 2, 193, 194, 7, 99, 2, 2, 
	194, 195, 7, 104, 2, 2, 195, 196, 7, 118, 2, 2, 196, 197, 7, 103, 2, 2, 
	197, 198, 7, 116, 2, 2, 198, 199, 7, 69, 2, 2, 199, 200, 7, 119, 2, 2, 
	200, 201, 7, 116, 2, 2, 201, 202, 7, 116, 2, 2, 202, 203, 7, 103, 2, 2, 
	203, 204, 7, 112, 2, 2, 204, 205, 7, 118, 2, 2, 205, 206, 7, 86, 2, 2, 
	206, 207, 7, 107, 2, 2, 207, 208, 7, 111, 2, 2, 208, 209, 7, 103, 2, 2, 
	209, 46, 3, 2, 2, 2, 210, 211, 5, 49, 25, 2, 211, 213, 7, 48, 2, 2, 212, 
	214, 5, 75, 38, 2, 213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 213, 
	3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 48, 3, 2, 2, 2, 217, 219, 7, 47, 
	2, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 221, 3, 2, 2, 2, 
	220, 222, 5, 75, 38, 2, 221, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 
	221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 50, 3, 2, 2, 2, 225, 228, 5, 
	57, 29, 2, 226, 228, 5, 59, 30, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 
	2, 2, 228, 52, 3, 2, 2, 2, 229, 230, 7, 112, 2, 2, 230, 231, 7, 107, 2, 
	2, 231, 232, 7, 110, 2, 2, 232, 54, 3, 2, 2, 2, 233, 234, 5, 85, 43, 2, 
	234, 56, 3, 2, 2, 2, 235, 236, 7, 118, 2, 2, 236, 237, 7, 116, 2, 2, 237, 
	238, 7, 119, 2, 2, 238, 239, 7, 103, 2, 2, 239, 58, 3, 2, 2, 2, 240, 241, 
	7, 104, 2, 2, 241, 242, 7, 99, 2, 2, 242, 243, 7, 110, 2, 2, 243, 244, 
	7, 117, 2, 2, 244, 245, 7, 103, 2, 2, 245, 60, 3, 2, 2, 2, 246, 250, 7, 
	37, 2, 2, 247, 249, 10, 2, 2, 2, 248, 247, 3, 2, 2, 2, 249, 252, 3, 2, 
	2, 2, 250, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 253, 3, 2, 2, 2, 
	252, 250, 3, 2, 2, 2, 253, 254, 8, 31, 2, 2, 254, 62, 3, 2, 2, 2, 255, 
	257, 9, 3, 2, 2, 256, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 256, 
	3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 8, 32, 
	3, 2, 261, 64, 3, 2, 2, 2, 262, 264, 9, 2, 2, 2, 263, 262, 3, 2, 2, 2, 
	264, 265, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 
	267, 3, 2, 2, 2, 267, 268, 8, 33, 3, 2, 268, 66, 3, 2, 2, 2, 269, 270, 
	5, 79, 40, 2, 270, 271, 5, 77, 39, 2, 271, 68, 3, 2, 2, 2, 272, 279, 5, 
	77, 39, 2, 273, 274, 7, 93, 2, 2, 274, 275, 5, 75, 38, 2, 275, 276, 7, 
	95, 2, 2, 276, 279, 3, 2, 2, 2, 277, 279, 5, 79, 40, 2, 278, 272, 3, 2, 
	2, 2, 278, 273, 3, 2, 2, 2, 278, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 
	280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 284, 3, 2, 2, 2, 282, 
	283, 7, 48, 2, 2, 283, 285, 5, 69, 35, 2, 284, 282, 3, 2, 2, 2, 284, 285, 
	3, 2, 2, 2, 285, 70, 3, 2, 2, 2, 286, 287, 7, 38, 2, 2, 287, 288, 7, 125, 
	2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 5, 69, 35, 2, 290, 291, 7, 127, 2, 
	2, 291, 72, 3, 2, 2, 2, 292, 293, 9, 4, 2, 2, 293, 74, 3, 2, 2, 2, 294, 
	295, 9, 5, 2, 2, 295, 76, 3, 2, 2, 2, 296, 299, 5, 73, 37, 2, 297, 299, 
	5, 75, 38, 2, 298, 296, 3, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 78, 3, 2, 
	2, 2, 300, 301, 7, 97, 2, 2, 301, 80, 3, 2, 2, 2, 302, 303, 7, 41, 2, 2, 
	303, 82, 3, 2, 2, 2, 304, 305, 7, 36, 2, 2, 305, 84, 3, 2, 2, 2, 306, 311, 
	5, 81, 41, 2, 307, 310, 5, 87, 44, 2, 308, 310, 10, 6, 2, 2, 309, 307, 
	3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 313, 3, 2, 2, 2, 311, 309, 3, 2, 
	2, 2, 311, 312, 3, 2, 2, 2, 312, 314, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 
	314, 315, 5, 81, 41, 2, 315, 86, 3, 2, 2, 2, 316, 320, 5, 91, 46, 2, 317, 
	321, 9, 7, 2, 2, 318, 321, 11, 2, 2, 2, 319, 321, 7, 2, 2, 3, 320, 317, 
	3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 321, 88, 3, 2, 
	2, 2, 322, 323, 5, 91, 46, 2, 323, 324, 11, 2, 2, 2, 324, 90, 3, 2, 2, 
	2, 325, 326, 7, 94, 2, 2, 326, 92, 3, 2, 2, 2, 25, 2, 121, 127, 134, 140, 
	145, 150, 157, 164, 215, 218, 223, 227, 250, 258, 265, 278, 280, 284, 298, 
	309, 311, 320, 4, 8, 2, 2, 2, 3, 2,
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
	"'@afterCurrentTime'", "", "", "", "'nil'", "", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "FATAL", "SEVERE", "WARNING", "LINKED", "EQ", 
	"NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", "FLOAT_FUNC", 
	"DATETIME_LONG", "AFTER_CURR_TIME", "FLOAT", "INT", "BOOL", "NULL", "STRING", 
	"TRUE", "FALSE", "COMMENT", "WS", "TERMINATOR", "DECL_ID", "ATTR_ID", "CTX_ID",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "FATAL", "SEVERE", "WARNING", 
	"LINKED", "EQ", "NE", "AND", "OR", "GT", "LT", "GTE", "LTE", "LEN_FUNC", 
	"FLOAT_FUNC", "DATETIME_LONG", "AFTER_CURR_TIME", "FLOAT", "INT", "BOOL", 
	"NULL", "STRING", "TRUE", "FALSE", "COMMENT", "WS", "TERMINATOR", "DECL_ID", 
	"ATTR_ID", "CTX_ID", "Character", "Numeric", "AlphaNumeric", "Underscore", 
	"SingleQuote", "DoubleQuote", "SingleQuoteString", "EscSeq", "EscAny", 
	"Esc",
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
	MooncakeLexerFLOAT = 23
	MooncakeLexerINT = 24
	MooncakeLexerBOOL = 25
	MooncakeLexerNULL = 26
	MooncakeLexerSTRING = 27
	MooncakeLexerTRUE = 28
	MooncakeLexerFALSE = 29
	MooncakeLexerCOMMENT = 30
	MooncakeLexerWS = 31
	MooncakeLexerTERMINATOR = 32
	MooncakeLexerDECL_ID = 33
	MooncakeLexerATTR_ID = 34
	MooncakeLexerCTX_ID = 35
)

