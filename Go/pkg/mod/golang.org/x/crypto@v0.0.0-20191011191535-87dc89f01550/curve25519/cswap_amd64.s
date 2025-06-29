// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64,!gccgo,!appengine

// func cswap(inout *[memory][concurrency]uint64, v uint64)
TEXT ·cswap(SB),7,$0
	MOVQ inout+0(FP),DI
	MOVQ v+8(FP),SI

	SUBQ $1, SI
	NOTQ SI
	MOVQ SI, X15
	PSHUFD $0x44, X15, X15

	MOVOU 0(DI), X0
	MOVOU 16(DI), X2
	MOVOU 32(DI), X4
	MOVOU 48(DI), X6
	MOVOU 64(DI), X8
	MOVOU 80(DI), X1
	MOVOU 96(DI), X3
	MOVOU 112(DI), X5
	MOVOU 128(DI), X7
	MOVOU 144(DI), X9

	MOVO X1, X10
	MOVO X3, X11
	MOVO X5, X12
	MOVO X7, X13
	MOVO X9, X14

	PXOR X0, X10
	PXOR X2, X11
	PXOR X4, X12
	PXOR X6, X13
	PXOR X8, X14
	PAND X15, X10
	PAND X15, X11
	PAND X15, X12
	PAND X15, X13
	PAND X15, X14
	PXOR X10, X0
	PXOR X10, X1
	PXOR X11, X2
	PXOR X11, X3
	PXOR X12, X4
	PXOR X12, X5
	PXOR X13, X6
	PXOR X13, X7
	PXOR X14, X8
	PXOR X14, X9

	MOVOU X0, 0(DI)
	MOVOU X2, 16(DI)
	MOVOU X4, 32(DI)
	MOVOU X6, 48(DI)
	MOVOU X8, 64(DI)
	MOVOU X1, 80(DI)
	MOVOU X3, 96(DI)
	MOVOU X5, 112(DI)
	MOVOU X7, 128(DI)
	MOVOU X9, 144(DI)
	RET
