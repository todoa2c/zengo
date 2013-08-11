zengo
=====

A grep like tool that shows around hit position

INSTALL
-----------
    go get github.com/todoa2c/zengo

Usage
-----------
example:

    zengo -pattern=pos < main.go

then, outputs lines of two column: number of line and shows around hit position

      14	 // foundPos stores position of occurence a
      43	eachRunePos returns position of sep in s fu
      49	rsep := []rune(sep) pos := runeIndex(rs[ind
      50	[index:], rsep) for pos != -1 { result = ap
      51	ppend(result, index+pos) index += pos + 1 p
      52	index+pos) index += pos + 1 pos = runeIndex
      53	s) index += pos + 1 pos = runeIndex(rs[inde
      78	runes, r) }  for _, pos := range eachRunePo
      79	 foundPos{lastLen + pos, lineNo}) }  lastLe
      92	ange foundPosList { pos := foundPos.runePos
      93	ineNo, string(runes[pos-*chars:pos+len(rune
      93	ng(runes[pos-*chars:pos+len(runePattern)+*c

chars specifies number of characters around pattern (default 20 characters)

    zengo -pattern=pos -chars=10 < main.go

      14	os stores position of o
      43	s returns position of s
      49	rune(sep) pos := runeIn
      50	rsep) for pos != -1 { r
      51	lt, index+pos) index +=
      52	 index += pos + 1 pos =
      53	= pos + 1 pos = runeInd
      78	}  for _, pos := range 
      79	lastLen + pos, lineNo})
      92	PosList { pos := foundP
      93	ing(runes[pos-*chars:po
      93	os-*chars:pos+len(runeP


