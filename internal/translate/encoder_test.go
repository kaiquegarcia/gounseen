package translate

import (
	"strconv"
	"testing"
)

func Test_ShouldDecodeLetters(t *testing.T) {
	InitializeLetters()

	t.Run("letter A", func(t *testing.T) {
		output := Decode("ㅏ")
		if output != "A" {
			t.Errorf("decoding ㅏ resulted %s instead of A", output)
			t.Fail()
		}
	})

	t.Run("letter B", func(t *testing.T) {
		output := Decode("Б")
		if output != "B" {
			t.Errorf("decoding Б resulted %s instead of B", output)
			t.Fail()
		}
	})

	t.Run("letter C", func(t *testing.T) {
		output := Decode("C")
		if output != "C" {
			t.Errorf("decoding C resulted %s instead of C", output)
			t.Fail()
		}
	})

	t.Run("letter C", func(t *testing.T) {
		output := Decode("𑀘")
		if output != "C" {
			t.Errorf("decoding 𑀘 resulted %s instead of C", output)
			t.Fail()
		}
	})

	t.Run("letter D", func(t *testing.T) {
		output := Decode("Δ")
		if output != "D" {
			t.Errorf("decoding Δ resulted %s instead of D", output)
			t.Fail()
		}
	})

	t.Run("letter E", func(t *testing.T) {
		output := Decode("Э")
		if output != "E" {
			t.Errorf("decoding Э resulted %s instead of E", output)
			t.Fail()
		}
	})

	t.Run("letter F", func(t *testing.T) {
		output := Decode("Ф")
		if output != "F" {
			t.Errorf("decoding Ф resulted %s instead of F", output)
			t.Fail()
		}
	})

	t.Run("letter G", func(t *testing.T) {
		output := Decode("Γ")
		if output != "G" {
			t.Errorf("decoding Γ resulted %s instead of G", output)
			t.Fail()
		}
	})

	t.Run("letter H", func(t *testing.T) {
		output := Decode("ㅎ")
		if output != "H" {
			t.Errorf("decoding ㅎ resulted %s instead of H", output)
			t.Fail()
		}
	})

	t.Run("letter I", func(t *testing.T) {
		output := Decode("И")
		if output != "I" {
			t.Errorf("decoding И resulted %s instead of I", output)
			t.Fail()
		}
	})

	t.Run("letter J", func(t *testing.T) {
		output := Decode("𑀛")
		if output != "J" {
			t.Errorf("decoding 𑀛 resulted %s instead of J", output)
			t.Fail()
		}
	})

	t.Run("letter J", func(t *testing.T) {
		output := Decode("ㅈ")
		if output != "J" {
			t.Errorf("decoding ㅈ resulted %s instead of J", output)
			t.Fail()
		}
	})

	t.Run("letter K", func(t *testing.T) {
		output := Decode("ㅋ")
		if output != "K" {
			t.Errorf("decoding ㅋ resulted %s instead of K", output)
			t.Fail()
		}
	})

	t.Run("letter L", func(t *testing.T) {
		output := Decode("Λ")
		if output != "L" {
			t.Errorf("decoding Λ resulted %s instead of L", output)
			t.Fail()
		}
	})

	t.Run("letter M", func(t *testing.T) {
		output := Decode("ㅁ")
		if output != "M" {
			t.Errorf("decoding ㅁ resulted %s instead of M", output)
			t.Fail()
		}
	})

	t.Run("letter N", func(t *testing.T) {
		output := Decode("𑀦")
		if output != "N" {
			t.Errorf("decoding 𑀦 resulted %s instead of N", output)
			t.Fail()
		}
	})

	t.Run("letter N", func(t *testing.T) {
		output := Decode("N")
		if output != "N" {
			t.Errorf("decoding N resulted %s instead of N", output)
			t.Fail()
		}
	})

	t.Run("letter O", func(t *testing.T) {
		output := Decode("Ω")
		if output != "O" {
			t.Errorf("decoding Ω resulted %s instead of O", output)
			t.Fail()
		}
	})

	t.Run("letter P", func(t *testing.T) {
		output := Decode("Π")
		if output != "P" {
			t.Errorf("decoding Π resulted %s instead of P", output)
			t.Fail()
		}
	})

	t.Run("letter Q", func(t *testing.T) {
		output := Decode("Q")
		if output != "Q" {
			t.Errorf("decoding Q resulted %s instead of Q", output)
			t.Fail()
		}
	})

	t.Run("letter R", func(t *testing.T) {
		output := Decode("ㄹ")
		if output != "R" {
			t.Errorf("decoding ㄹ resulted %s instead of R", output)
			t.Fail()
		}
	})

	t.Run("letter S", func(t *testing.T) {
		output := Decode("Σ")
		if output != "S" {
			t.Errorf("decoding Σ resulted %s instead of S", output)
			t.Fail()
		}
	})

	t.Run("letter T", func(t *testing.T) {
		output := Decode("T")
		if output != "T" {
			t.Errorf("decoding T resulted %s instead of T", output)
			t.Fail()
		}
	})

	t.Run("letter U", func(t *testing.T) {
		output := Decode("U")
		if output != "U" {
			t.Errorf("decoding U resulted %s instead of U", output)
			t.Fail()
		}
	})

	t.Run("letter V", func(t *testing.T) {
		output := Decode("V")
		if output != "V" {
			t.Errorf("decoding V resulted %s instead of V", output)
			t.Fail()
		}
	})

	t.Run("letter W", func(t *testing.T) {
		output := Decode("W")
		if output != "W" {
			t.Errorf("decoding W resulted %s instead of W", output)
			t.Fail()
		}
	})

	t.Run("letter X", func(t *testing.T) {
		output := Decode("X")
		if output != "X" {
			t.Errorf("decoding X resulted %s instead of X", output)
			t.Fail()
		}
	})

	t.Run("letter Y", func(t *testing.T) {
		output := Decode("Y")
		if output != "Y" {
			t.Errorf("decoding Y resulted %s instead of Y", output)
			t.Fail()
		}
	})

	t.Run("letter Z", func(t *testing.T) {
		output := Decode("З")
		if output != "Z" {
			t.Errorf("decoding З resulted %s instead of Z", output)
			t.Fail()
		}
	})
}

func Test_ShouldNotDecodeUnmapedLetters(t *testing.T) {
	InitializeLetters()

	t.Run("whitespace", func(t *testing.T) {
		output := Decode(" ")
		if output != " " {
			t.Errorf("decoding whitespace returned '%s' instead of whitespace", output)
			t.Fail()
		}
	})

	for number := 0; number < 10; number++ {
		num := strconv.Itoa(number)
		t.Run("number "+num, func(t *testing.T) {
			output := Decode(num)
			if output != num {
				t.Errorf("decoding number %s returned '%s' instead of %s", num, output, num)
				t.Fail()
			}
		})
	}
}

func Test_ShouldEncodeLetters(t *testing.T) {
	InitializeLetters()

	t.Run("letter ㅏ", func(t *testing.T) {
		output := Encode("A")
		if output != "ㅏ" {
			t.Errorf("encoding A resulted %s instead of ㅏ", output)
			t.Fail()
		}
	})

	t.Run("letter Б", func(t *testing.T) {
		output := Encode("B")
		if output != "Б" {
			t.Errorf("encoding B resulted %s instead of Б", output)
			t.Fail()
		}
	})

	t.Run("letter C", func(t *testing.T) {
		output := Encode("C")
		if output != "C" && output != "𑀘" {
			t.Errorf("encoding C resulted %s instead of C or 𑀘", output)
			t.Fail()
		}
	})

	t.Run("letter Δ", func(t *testing.T) {
		output := Encode("D")
		if output != "Δ" {
			t.Errorf("encoding D resulted %s instead of Δ", output)
			t.Fail()
		}
	})

	t.Run("letter Э", func(t *testing.T) {
		output := Encode("E")
		if output != "Э" {
			t.Errorf("encoding E resulted %s instead of Э", output)
			t.Fail()
		}
	})

	t.Run("letter Ф", func(t *testing.T) {
		output := Encode("F")
		if output != "Ф" {
			t.Errorf("encoding F resulted %s instead of Ф", output)
			t.Fail()
		}
	})

	t.Run("letter Γ", func(t *testing.T) {
		output := Encode("G")
		if output != "Γ" {
			t.Errorf("encoding G resulted %s instead of Γ", output)
			t.Fail()
		}
	})

	t.Run("letter ㅎ", func(t *testing.T) {
		output := Encode("H")
		if output != "ㅎ" {
			t.Errorf("encoding H resulted %s instead of ㅎ", output)
			t.Fail()
		}
	})

	t.Run("letter И", func(t *testing.T) {
		output := Encode("I")
		if output != "И" {
			t.Errorf("encoding I resulted %s instead of И", output)
			t.Fail()
		}
	})

	t.Run("letter 𑀛", func(t *testing.T) {
		output := Encode("J")
		if output != "𑀛" && output != "ㅈ" {
			t.Errorf("encoding J resulted %s instead of 𑀛 or ㅈ", output)
			t.Fail()
		}
	})

	t.Run("letter ㅋ", func(t *testing.T) {
		output := Encode("K")
		if output != "ㅋ" {
			t.Errorf("encoding K resulted %s instead of ㅋ", output)
			t.Fail()
		}
	})

	t.Run("letter Λ", func(t *testing.T) {
		output := Encode("L")
		if output != "Λ" {
			t.Errorf("encoding L resulted %s instead of Λ", output)
			t.Fail()
		}
	})

	t.Run("letter ㅁ", func(t *testing.T) {
		output := Encode("M")
		if output != "ㅁ" {
			t.Errorf("encoding M resulted %s instead of ㅁ", output)
			t.Fail()
		}
	})

	t.Run("letter 𑀦", func(t *testing.T) {
		output := Encode("N")
		if output != "𑀦" && output != "N" {
			t.Errorf("encoding N resulted %s instead of 𑀦 or N", output)
			t.Fail()
		}
	})

	t.Run("letter Ω", func(t *testing.T) {
		output := Encode("O")
		if output != "Ω" {
			t.Errorf("encoding O resulted %s instead of Ω", output)
			t.Fail()
		}
	})

	t.Run("letter Π", func(t *testing.T) {
		output := Encode("P")
		if output != "Π" {
			t.Errorf("encoding P resulted %s instead of Π", output)
			t.Fail()
		}
	})

	t.Run("letter Q", func(t *testing.T) {
		output := Encode("Q")
		if output != "Q" {
			t.Errorf("encoding Q resulted %s instead of Q", output)
			t.Fail()
		}
	})

	t.Run("letter ㄹ", func(t *testing.T) {
		output := Encode("R")
		if output != "ㄹ" {
			t.Errorf("encoding R resulted %s instead of ㄹ", output)
			t.Fail()
		}
	})

	t.Run("letter Σ", func(t *testing.T) {
		output := Encode("S")
		if output != "Σ" {
			t.Errorf("encoding S resulted %s instead of Σ", output)
			t.Fail()
		}
	})

	t.Run("letter T", func(t *testing.T) {
		output := Encode("T")
		if output != "T" {
			t.Errorf("encoding T resulted %s instead of T", output)
			t.Fail()
		}
	})

	t.Run("letter U", func(t *testing.T) {
		output := Encode("U")
		if output != "U" {
			t.Errorf("encoding U resulted %s instead of U", output)
			t.Fail()
		}
	})

	t.Run("letter V", func(t *testing.T) {
		output := Encode("V")
		if output != "V" {
			t.Errorf("encoding V resulted %s instead of V", output)
			t.Fail()
		}
	})

	t.Run("letter W", func(t *testing.T) {
		output := Encode("W")
		if output != "W" {
			t.Errorf("encoding W resulted %s instead of W", output)
			t.Fail()
		}
	})

	t.Run("letter X", func(t *testing.T) {
		output := Encode("X")
		if output != "X" {
			t.Errorf("encoding X resulted %s instead of X", output)
			t.Fail()
		}
	})

	t.Run("letter Y", func(t *testing.T) {
		output := Encode("Y")
		if output != "Y" {
			t.Errorf("encoding Y resulted %s instead of Y", output)
			t.Fail()
		}
	})

	t.Run("letter З", func(t *testing.T) {
		output := Encode("Z")
		if output != "З" {
			t.Errorf("encoding Z resulted %s instead of З", output)
			t.Fail()
		}
	})
}

func Test_ShouldNotEncodeUnmapedLetters(t *testing.T) {
	InitializeLetters()

	t.Run("whitespace", func(t *testing.T) {
		output := Encode(" ")
		if output != " " {
			t.Errorf("encoding whitespace returned '%s' instead of whitespace", output)
			t.Fail()
		}
	})

	for number := 0; number < 10; number++ {
		num := strconv.Itoa(number)
		t.Run("number "+num, func(t *testing.T) {
			output := Encode(num)
			if output != num {
				t.Errorf("encoding number %s returned '%s' instead of %s", num, output, num)
				t.Fail()
			}
		})
	}
}
