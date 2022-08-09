package mont_arith

import (
    "fmt"
    "testing"
    "math/big"
)

func benchmarkMulMont(b *testing.B, preset ArithPreset, limbCount uint) {
	mod := GenTestModulus(limbCount)
	montCtx := NewField(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

    x := big.NewInt(1)
    y := big.NewInt(1)
    /*
	x := LimbsToInt(mod)
    x = x.Sub(x, big.NewInt(1))

	y := new(big.Int).SetBytes(LimbsToInt(mod).Bytes())
    y = y.Sub(y, big.NewInt(1))
    */

	outLimbs := make([]uint64, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, limbCount)
	yLimbs := IntToLimbs(y, limbCount)

    outBytes := LimbsToLEBytes(outLimbs)
    xBytes := LimbsToLEBytes(xLimbs)
    yBytes := LimbsToLEBytes(yLimbs)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulMont(montCtx, outBytes, xBytes, yBytes)
	}
}

func BenchmarkMulMontGo(b *testing.B) {
    preset := DefaultPreset()
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkMulMont(b, preset, uint(i))
			})
		}
	}

	bench(b, 1, 12)
}


func BenchmarkMulMontAsm(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkMulMont(b, Asm384Preset(), uint(i))
			})
		}
	}

	bench(b, 6, 6)
}

func benchmarkAddMod(b *testing.B, preset ArithPreset, limbCount uint) {
    modLimbs := MaxModulus(limbCount)
    mod := LimbsToInt(modLimbs)
	montCtx := NewField(preset)

    // worst-case performance: unecessary final subtraction
	err := montCtx.SetMod(modLimbs)
	if err != nil {
		panic("error")
	}
	x := new(big.Int).SetBytes(mod.Bytes())
	x = x.Sub(x, big.NewInt(2))
    y := big.NewInt(1)
    outBytes := make([]byte, limbCount * 8)
    xBytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
    yBytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        montCtx.AddMod(montCtx, outBytes, xBytes, yBytes)
    }
}

func BenchmarkAddModGo(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkAddMod(b, DefaultPreset(), uint(i))
			})
		}
	}

	bench(b, 1, 12)
}

func BenchmarkAddModAsm(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkAddMod(b, Asm384Preset(), uint(i))
			})
		}
	}

	bench(b, 6, 6)
}

func benchmarkSubMod(b *testing.B, preset ArithPreset, limbCount uint) {
    modLimbs := MaxModulus(limbCount)
	montCtx := NewField(preset)

    // worst-case performance: unecessary final subtraction
	err := montCtx.SetMod(modLimbs)
	if err != nil {
		panic("error")
	}
	x := big.NewInt(1)
    y := big.NewInt(0)
    outBytes := make([]byte, limbCount * 8)
    xBytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
    yBytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        montCtx.SubMod(montCtx, outBytes, xBytes, yBytes)
    }
}

func BenchmarkSubModGo(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkSubMod(b, DefaultPreset(), uint(i))
			})
		}
	}

	bench(b, 1, 12)
}

func BenchmarkSubModAsm(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkSubMod(b, Asm384Preset(), uint(i))
			})
		}
	}

	bench(b, 6, 6)
}

func benchmarkSetMod(b *testing.B, limbCount uint) {
    modLimbs := MidModulus(limbCount)
    montCtx := NewField(DefaultPreset())

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        montCtx.SetMod(modLimbs)
    }
}

func BenchmarkSetMod(b *testing.B) {
	bench := func(b *testing.B, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkSetMod(b, uint(i))
			})
		}
	}

	bench(b, 1, 12)

}