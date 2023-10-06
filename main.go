package main

import (
	"fmt"
	"strings"
)

func main() {
	key := "SUCHONG"
	// msg := "CATHODETUBE"
	suchong := "never underestimate the fallibility of an egomaniac"
	fmt.Println((98 - 32) % 'A')
	codes := encodeSuchong(suchong, key)
	fmt.Println("Phrase", codes)
	decode := atomicNumberDecode(codes)
	fmt.Println("decode", decode)
	vdecoded := vigenereDecipherFromSlice(decode, key)
	fmt.Println("decoded", vdecoded)
}

var (
	sample        = []string{"C", "Mn", "Cr", "Mg", "C", "O", "Ca", "Ti", "Mn", "Ca", "Mg", "N", "N", "P", "B", "Sc", "Ti", "Mg", "O", "Sc", "Na", "Cr", "Sc", "Si", "K", "V", "P", "P", "Be", "Li", "Ti", "C", "Li", "K", "N", "C", "Mn", "F", "Ti", "H", "Si", "Ca", "H", "Sc", "B"}
	atomicNumbers = []string{"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne", "Na", "Mg", "Al", "Si", "P", "S", "Cl", "Ar", "K", "Ca", "Sc", "Ti", "V", "Cr", "Mn", "Fe"}
)

func encodeSuchong(msg, key string) []string {
	return atomicNumberEncode(vigenereEncipherToRune(msg, key))
}

func caseStandardize(in string) string {
	out := []rune{}
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}
	fmt.Println("standardize", string(out))
	return string(out)
}

func vEncode(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func vDecode(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func vigenereEncipherToString(msg, key string) string {
	smsg, skey := caseStandardize(msg), caseStandardize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {

		out = append(out, vEncode(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func vigenereEncipherToRune(msg, key string) []rune {
	smsg, skey := caseStandardize(msg), caseStandardize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, vEncode(v, rune(skey[i%len(skey)])))
	}
	fmt.Println(out)
	return out
}

func vigenereDecipher(msg, key string) string {
	smsg, skey := caseStandardize(msg), caseStandardize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, vDecode(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func vigenereDecipherFromSlice(msg, key string) string {
	out := make([]rune, 0, len(msg))
	for i, v := range msg {
		out = append(out, vDecode(v, rune(key[i%len(key)])))
	}
	return string(out)
}

func atomicNumberEncode(runes []rune) []string {
	out := make([]string, 0, len(runes))
	for _, v := range runes {
		if 65 <= v && v <= 90 {
			out = append(out, atomicNumbers[v%'A'])
		} else if 97 <= v && v <= 122 {
			out = append(out, atomicNumbers[v%'a'])
		}
	}
	return out
}

func atomicNumberDecode(msg []string) string {
	out := make([]string, 0, len(msg))
	for _, v := range msg {
		for i, e := range atomicNumbers {
			if v == e {
				out = append(out, string(rune(i+'A')))
			}
		}
	}
	return strings.Join(out, "")
}

/*
Encryption:
raw text -> Vigenere encipher -> alphabetically position mapping(rune mapping) -> Periodic table atomic number mapping
Decryption:
Chemical element text -> atomic number mapping -> rune mapping -> Vigenere decipher
*/
