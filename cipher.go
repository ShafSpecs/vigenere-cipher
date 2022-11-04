package main

func main() {
	const cipherText = "CSOITEUIWUIZNSROCNKFD"
	const keyword = "GOLANG"

	var actualText = ""

	// a b c d e f g h i j k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25

	var (
		gValue = 6
		oValue = 14
		lValue = 11
		aValue = 0
		nValue = 13
	)

	for i := 0; i < len(cipherText); i++ {
		switch i % 6 {
		case 0:
			{
				//actualText = cipherText[i] - gValue
			}
		case 1:
			{
			}
		case 2:
			{
			}
		case 3:
			{
			}
		case 4:
			{
			}
		case 5:
			{
			}
		}
	}
}
