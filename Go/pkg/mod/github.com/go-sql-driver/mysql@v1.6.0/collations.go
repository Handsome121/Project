// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2014 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. control.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const defaultCollation = "utf8mb4_general_ci"
const binaryCollation = "binary"

// A list of available collations mapped to the internal ID.
// To update this map use the following MySQL query:
//     SELECT COLLATION_NAME, ID FROM information_schema.COLLATIONS WHERE ID<256 ORDER BY ID
//
// Handshake packet have only 1 byte for collation_id.  So we can't use collations with ID > 255.
//
// ucs2, utf16, and utf32 can't be used for connection charset.
// https://dev.mysql.com/doc/refman/5.7/en/charset-connection.html#charset-connection-impermissible-client-charset
// They are commented out to reduce this map.
var collations = map[string]byte{
	"big5_chinese_ci":      1,
	"latin2_czech_cs":      2,
	"dec8_swedish_ci":      3,
	"cp850_general_ci":     4,
	"latin1_german1_ci":    5,
	"hp8_english_ci":       6,
	"koi8r_general_ci":     7,
	"latin1_swedish_ci":    8,
	"latin2_general_ci":    9,
	"swe7_swedish_ci":      10,
	"ascii_general_ci":     11,
	"ujis_japanese_ci":     12,
	"sjis_japanese_ci":     13,
	"cp1251_bulgarian_ci":  14,
	"latin1_danish_ci":     15,
	"hebrew_general_ci":    16,
	"tis620_thai_ci":       18,
	"euckr_korean_ci":      19,
	"latin7_estonian_cs":   20,
	"latin2_hungarian_ci":  21,
	"koi8u_general_ci":     22,
	"cp1251_ukrainian_ci":  23,
	"gb2312_chinese_ci":    24,
	"greek_general_ci":     25,
	"cp1250_general_ci":    26,
	"latin2_croatian_ci":   27,
	"gbk_chinese_ci":       28,
	"cp1257_lithuanian_ci": 29,
	"latin5_turkish_ci":    30,
	"latin1_german2_ci":    31,
	"armscii8_general_ci":  32,
	"utf8_general_ci":      33,
	"cp1250_czech_cs":      34,
	//"ucs2_general_ci":          35,
	"cp866_general_ci":    36,
	"keybcs2_general_ci":  37,
	"macce_general_ci":    38,
	"macroman_general_ci": 39,
	"cp852_general_ci":    40,
	"latin7_general_ci":   41,
	"latin7_general_cs":   42,
	"macce_bin":           43,
	"cp1250_croatian_ci":  44,
	"utf8mb4_general_ci":  45,
	"utf8mb4_bin":         46,
	"latin1_bin":          47,
	"latin1_general_ci":   48,
	"latin1_general_cs":   49,
	"cp1251_bin":          50,
	"cp1251_general_ci":   51,
	"cp1251_general_cs":   52,
	"macroman_bin":        53,
	//"utf16_general_ci":         54,
	//"utf16_bin":                55,
	//"utf16le_general_ci":       56,
	"cp1256_general_ci": 57,
	"cp1257_bin":        58,
	"cp1257_general_ci": 59,
	//"utf32_general_ci":         60,
	//"utf32_bin":                61,
	//"utf16le_bin":              62,
	"binary":          63,
	"armscii8_bin":    64,
	"ascii_bin":       65,
	"cp1250_bin":      66,
	"cp1256_bin":      67,
	"cp866_bin":       68,
	"dec8_bin":        69,
	"greek_bin":       70,
	"hebrew_bin":      71,
	"hp8_bin":         72,
	"keybcs2_bin":     73,
	"koi8r_bin":       74,
	"koi8u_bin":       75,
	"utf8_tolower_ci": 76,
	"latin2_bin":      77,
	"latin5_bin":      78,
	"latin7_bin":      79,
	"cp850_bin":       80,
	"cp852_bin":       81,
	"swe7_bin":        82,
	"utf8_bin":        83,
	"big5_bin":        84,
	"euckr_bin":       85,
	"gb2312_bin":      86,
	"gbk_bin":         87,
	"sjis_bin":        88,
	"tis620_bin":      89,
	//"ucs2_bin":                 90,
	"ujis_bin":            91,
	"geostd8_general_ci":  92,
	"geostd8_bin":         93,
	"latin1_spanish_ci":   94,
	"cp932_japanese_ci":   95,
	"cp932_bin":           96,
	"eucjpms_japanese_ci": 97,
	"eucjpms_bin":         98,
	"cp1250_polish_ci":    99,
	//"utf16_unicode_ci":         101,
	//"utf16_icelandic_ci":       102,
	//"utf16_latvian_ci":         103,
	//"utf16_romanian_ci":        104,
	//"utf16_slovenian_ci":       105,
	//"utf16_polish_ci":          106,
	//"utf16_estonian_ci":        107,
	//"utf16_spanish_ci":         108,
	//"utf16_swedish_ci":         109,
	//"utf16_turkish_ci":         110,
	//"utf16_czech_ci":           111,
	//"utf16_danish_ci":          112,
	//"utf16_lithuanian_ci":      113,
	//"utf16_slovak_ci":          114,
	//"utf16_spanish2_ci":        115,
	//"utf16_roman_ci":           116,
	//"utf16_persian_ci":         117,
	//"utf16_esperanto_ci":       118,
	//"utf16_hungarian_ci":       119,
	//"utf16_sinhala_ci":         120,
	//"utf16_german2_ci":         121,
	//"utf16_croatian_ci":        122,
	//"utf16_unicode_520_ci":     123,
	//"utf16_vietnamese_ci":      124,
	//"ucs2_unicode_ci":          128,
	//"ucs2_icelandic_ci":        129,
	//"ucs2_latvian_ci":          130,
	//"ucs2_romanian_ci":         131,
	//"ucs2_slovenian_ci":        132,
	//"ucs2_polish_ci":           133,
	//"ucs2_estonian_ci":         134,
	//"ucs2_spanish_ci":          135,
	//"ucs2_swedish_ci":          136,
	//"ucs2_turkish_ci":          137,
	//"ucs2_czech_ci":            138,
	//"ucs2_danish_ci":           139,
	//"ucs2_lithuanian_ci":       140,
	//"ucs2_slovak_ci":           141,
	//"ucs2_spanish2_ci":         142,
	//"ucs2_roman_ci":            143,
	//"ucs2_persian_ci":          144,
	//"ucs2_esperanto_ci":        145,
	//"ucs2_hungarian_ci":        146,
	//"ucs2_sinhala_ci":          147,
	//"ucs2_german2_ci":          148,
	//"ucs2_croatian_ci":         149,
	//"ucs2_unicode_520_ci":      150,
	//"ucs2_vietnamese_ci":       151,
	//"ucs2_general_mysql500_ci": 159,
	//"utf32_unicode_ci":         160,
	//"utf32_icelandic_ci":       161,
	//"utf32_latvian_ci":         162,
	//"utf32_romanian_ci":        163,
	//"utf32_slovenian_ci":       164,
	//"utf32_polish_ci":          165,
	//"utf32_estonian_ci":        166,
	//"utf32_spanish_ci":         167,
	//"utf32_swedish_ci":         168,
	//"utf32_turkish_ci":         169,
	//"utf32_czech_ci":           170,
	//"utf32_danish_ci":          171,
	//"utf32_lithuanian_ci":      172,
	//"utf32_slovak_ci":          173,
	//"utf32_spanish2_ci":        174,
	//"utf32_roman_ci":           175,
	//"utf32_persian_ci":         176,
	//"utf32_esperanto_ci":       177,
	//"utf32_hungarian_ci":       178,
	//"utf32_sinhala_ci":         179,
	//"utf32_german2_ci":         180,
	//"utf32_croatian_ci":        181,
	//"utf32_unicode_520_ci":     182,
	//"utf32_vietnamese_ci":      183,
	"utf8_unicode_ci":          192,
	"utf8_icelandic_ci":        193,
	"utf8_latvian_ci":          194,
	"utf8_romanian_ci":         195,
	"utf8_slovenian_ci":        196,
	"utf8_polish_ci":           197,
	"utf8_estonian_ci":         198,
	"utf8_spanish_ci":          199,
	"utf8_swedish_ci":          200,
	"utf8_turkish_ci":          201,
	"utf8_czech_ci":            202,
	"utf8_danish_ci":           203,
	"utf8_lithuanian_ci":       204,
	"utf8_slovak_ci":           205,
	"utf8_spanish2_ci":         206,
	"utf8_roman_ci":            207,
	"utf8_persian_ci":          208,
	"utf8_esperanto_ci":        209,
	"utf8_hungarian_ci":        210,
	"utf8_sinhala_ci":          211,
	"utf8_german2_ci":          212,
	"utf8_croatian_ci":         213,
	"utf8_unicode_520_ci":      214,
	"utf8_vietnamese_ci":       215,
	"utf8_general_mysql500_ci": 223,
	"utf8mb4_unicode_ci":       224,
	"utf8mb4_icelandic_ci":     225,
	"utf8mb4_latvian_ci":       226,
	"utf8mb4_romanian_ci":      227,
	"utf8mb4_slovenian_ci":     228,
	"utf8mb4_polish_ci":        229,
	"utf8mb4_estonian_ci":      230,
	"utf8mb4_spanish_ci":       231,
	"utf8mb4_swedish_ci":       232,
	"utf8mb4_turkish_ci":       233,
	"utf8mb4_czech_ci":         234,
	"utf8mb4_danish_ci":        235,
	"utf8mb4_lithuanian_ci":    236,
	"utf8mb4_slovak_ci":        237,
	"utf8mb4_spanish2_ci":      238,
	"utf8mb4_roman_ci":         239,
	"utf8mb4_persian_ci":       240,
	"utf8mb4_esperanto_ci":     241,
	"utf8mb4_hungarian_ci":     242,
	"utf8mb4_sinhala_ci":       243,
	"utf8mb4_german2_ci":       244,
	"utf8mb4_croatian_ci":      245,
	"utf8mb4_unicode_520_ci":   246,
	"utf8mb4_vietnamese_ci":    247,
	"gb18030_chinese_ci":       248,
	"gb18030_bin":              249,
	"gb18030_unicode_520_ci":   250,
	"utf8mb4_0900_ai_ci":       255,
}

// A denylist of collations which is unsafe to interpolate parameters.
// These multibyte encodings may contains 0x5c (`\`) in their trailing bytes.
var unsafeCollations = map[string]bool{
	"big5_chinese_ci":        true,
	"sjis_japanese_ci":       true,
	"gbk_chinese_ci":         true,
	"big5_bin":               true,
	"gb2312_bin":             true,
	"gbk_bin":                true,
	"sjis_bin":               true,
	"cp932_japanese_ci":      true,
	"cp932_bin":              true,
	"gb18030_chinese_ci":     true,
	"gb18030_bin":            true,
	"gb18030_unicode_520_ci": true,
}
