// 	int128_amd.s

#include "textflag.h"

TEXT ·Add(SB),NOSPLIT,$0
	MOVQ 	a+0(FP), AX
	MOVQ 	a+8(FP), BX
	ADDQ 	b+16(FP), AX
	ADDQ 	b+24(FP), BX
	MOVQ 	AX, z+32(FP)
	MOVQ 	BX, z+40(FP)
	RET
