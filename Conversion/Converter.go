package Conversion

import (
	"strconv"
	"strings"
)

func ConvertIEEE754(number float64, indicator int) (binary string)  {
	 binary = Sign(number) + ExponentAndMantissa(number, indicator)
	 return

}

func Sign(number float64) string {
	if(number > 0){
		return "0"
	}else {
		return "1"
	}

}

func ExponentAndMantissa(number float64, indicator int) (xman string){
	numberString := strconv.FormatFloat(number,'f',10,64)
	result := strings.Split(numberString, ".")
	Exponent := Exponent(result[0])
	Mantissa := Mantissa(result)
	xman = Exponent + Mantissa
	return

}

func Exponent(str string) string {
	bi, _ := strconv.Atoi(str)
	binary := strconv.FormatInt(int64(bi), 2)
	return strconv.FormatInt(int64((len(binary)-1) + 127), 2)
}

func Mantissa(args []string) (mantissa string){

	i := len(args[1])-1
	for i >= 0 {
		if (string(args[1][i]) != "0"){
			break
		}
		i--
	}

	f, _ := strconv.ParseFloat("."+args[1][:i+1], 64)


	for i := 0; i < 50; i++{
		f = f*2
		if(f > 1){
			f = f-1
			mantissa = mantissa + "1"
		}else {
			mantissa = mantissa + "0"
		}
	}

	bi, _ := strconv.Atoi(args[0])
	binary := strconv.FormatInt(int64(bi), 2)
	mantissa = binary[1:] + mantissa[0:23]
	return

}

