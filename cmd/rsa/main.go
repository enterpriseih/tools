package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
)

const Licences = "AnrThDTNcp6c3AN0zHROpkNkOSLYAA7CCNsKn+m3RkJz4ERe+l855J532r9zeQ0bh0ZwbEh4e5IvFGaOZvdmAAC6lcm5del2KBgGkfgujrn7pe3kwClw0v6j1CQIVxv9TJ8W6cuYtEx0S/I2zUnzche/anjesmXGDQnS1Wo5OEHHPFnxxbNZu+snpsQUeeYJc1m2Xk/H1Lb/YfgiH06ASs9yLZKXIuqiJofEUBXkBMhoEI4Blv4hKMdEm34IZGg1Qi3g/etYrIcBFyF5Ilj6MQX97gB0HhoyBFU0jU8iUg+KVuIvDHWsgZnE8zGiCNkfZbJlDVMWQEHEnUVNu6G0RArqKRAX3K7RinI5qHFeVe6MNViwYtb4ypRxeuntQ+nJcBb79Rv4vRuiHFm9zvu22uJ7lZhW5gQpIXL/hkKJy0z5ld+wLn5WHm6ilDPJ2XFeIMKDiDo+qlYuMJyD5rJ6fm8u/fZi6/wkbZSRwYvHwvk4TPreyUOoLD+VYCehn+rbJR7j82XCSalZulG3+JRoZ6CQq8igXe8bc8ficAwSFeYFswtt37Mnq5BszCdbnqs1TBNDky1wIVrwHrI5eS1k91LMHPPQB5cV/YjRWQSAaHAUnRXNX/aB9vH1Ge/OYFic5rnTebK/lzHqlzbwCr9LYldta3dK4tRdpOgr/cdl1VE2VcCkesaftl9BCh5HEgARcoRN9UYSOSPfJS/3iAHvDB9NYdZcxAOLi/6OfnOzRXHge3hAIRZvVVWHk3bZjsvxF4s8cRG/uFGX9sC4USwhydoO2ZQWCcXGBlDZmuZU8ZbsYandcvVuNIkI7rS4TCRBMPGTh5kI4iH425CxSS49IE4zo2lZDRXFBU2II8AEGz6TvXEcgZkhl43sjF7YZ8soYlGMTuEUrhSszfJ9WAfSAJhQzUyF5Fy5/wreEgYF/smsu9WLTSiWJWnOUCG12Dh5KUmZlzIEZVrC/VxTs6xVfRBdM1bNCfKaZfqz9AaubCpZRtrZZJkY3g4gZdWPzdEcBkd2JUgQXQCTBOHwYkaCpGMamVZjQCaIunLjyfhzQcHcEBylL2evIRvfTeBtjIuugXoss+44T/j53GWk5uxuLUIJBHWaaEz4E4NwWxRIutlJQXVCIjmRraHzD0NVTTZqaa4WerLtynbwD3wx7E5hRX1OgG9rDsSehTO83SYrzg8IqTOES3QHi0SxNag6J2ZWtdp6NS7VuY6iQ4w54CSI1w2uemFuOpH8t/L/GfpP0H65T54U7NET/S9u2prg6Brr6gTRGE7fjX8g9BAWp5a+N9b3LBfaYFezaq+vAAA/1hjHj/KaZq9g408I9A5EsxphvmQjiTPdxEb/0flnF8qwSg=="
const publicKey = `-----BEGIN PUBLIC KEY-----
MIIEIjANBgkqhkiG9w0BAQEFAAOCBA8AMIIECgKCBAEAupAnRDZwYHATqn/A8lkj
9+eyON6e3oHApqUEjbWUJErT28qD/iQfvkILn76kL/3BLzkk4tnI53kZ4k+L9ss9
fNKxozZGBz4hh6zzxcVkUj/aI51sH2AfXZeScZ7YS7k+9PtR87RPddUORSsKQycj
KgMJNpcZu+yumHGrgJbOlYwCN97L4gygaY7e2lqJevrq8mtYu/F7+waei5PHJPHy
6j1ebTUV6zA/hMK63/cHSACA6f4zaSz96oZnLiyPbN/KOR1lgz9JfhH9rSmxZX8R
irQJKWJdYgzcdLdUxyqOdR9aNVyG/xgeCTSNIgX2+1+oyRm98lzbocXt8syjG1aR
L1NOgubt9HmQCzsAsm1RcWjLmd9n48MlUwep6FYMhiTUqydBUabD2AmDo+LP7rmE
ytDkZT2OkZGL8Jg5D3GpRp7jaG0rKe1zcGLec2+HiZz6KxLTH5dxdH3CsUaKGIJK
yqQYfF8Aic7Nr82nCnUY6pYsGvfdGbEAJJ11WdXKQr/mqU9OrV9DaK59FNHvEnvo
X9oOP3IyJd3to53twGvRoaRupxnSyRVIHVvB+mMZ1puAUa44DNTDyXlLalxZEO46
UaADlSEdpum8E1qXi8+kdhSTyFX1YE54qqe+3fJ3SLFwCNKrlWmqjKfeDm8il/IO
iGqcwJxwbaWlEAZoKLFJ5pjJCAXtTgZfupFuZzXvjiOdLA3uHO5kIXERSvuzMun7
SxkSOtFeLaD5azvuHh7xqvpV+5e+a1739kK1qtK95D9AsPh07XTP3gzga/rs7QOA
7fzg8ayMCUTdX0pahUxM0mGx7OP1+8hGUwpgGdeTLUUG1Wi8mhAmktxosoVg+tb7
FvJPJGAFwO0PS81bxvW7zk3onaFBTsNZJR3MMwV9HseR3Fdn1r2tE935TmPQG4nC
heqBlX/SNqls70/q4JNHgZxd+Kh8bluVoNEaDLCZPZSrivSHGIaD7ZjLfOCkd8DB
IrrNyPmtGh8Dt3D9bWQyNW2wDz4var8fSpH5BSqGsnXz00nlCkZqxLajz8Lz98ea
koizXF5YHuNABcwetLHZG51daSvtlwNBpkcD01BH4YYI2MZrcOc10UBbLgr8Z5v+
1Dgm058haGLceFyQYL/1/uj0MrBXz5q7yzg+dXwsmAirctYCibno2rLgqvtIRTKZ
37Pev0WvVwH3mGRJi2K0Diu/hzyPnpB/2TaeUVnrbNnZ2Mu9ZNi/UJqhAtszyZ/y
nlhFdIbF0rnmRdTfKWlI1kiuXbFT72aUguhXNJ1hpMxGIWzXGhWAde+58R5d07BQ
sMFtnRGsp1tDp9Y9/Ux3F/EX8G3Rf9Cw9AVIPahdy22fNI/zqROjIE199kK4Z1ke
ZwIDAQAB
-----END PUBLIC KEY-----
`

func main() {
	licences, _ := base64.StdEncoding.DecodeString(Licences)

	//fmt.Println(string(licences))

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	out := RSA_public_decrypt(pub.(*rsa.PublicKey), licences)
	fmt.Println(string(out))
}
func RSA_public_decrypt(pubKey *rsa.PublicKey, data []byte) []byte {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}
	return out[skip:]
}
