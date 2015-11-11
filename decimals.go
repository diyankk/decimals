/*
Package decimals is a small library of functions for rounding and
formatting base ten numbers. These are functions that are either
missing from the standard libraries or are more convenient for
presenting numbers in a human-readable format.
*/
package decimals

import (
	"strconv"
	"math"
)

// RoundInt rounds a base ten int64 to the given precision. Precision is a
// negative number that represents the nearest power of ten to which the 
// integer should be rounded. It is expressed as a negative number to be 
// consistent with the decimal precision arguments used in rounding floats.
func RoundInt(x int64, power int) int64 {

	var r float64

	// Ensure power is not positive then invert
	if power > 0 {
		
		power = 0
	}

	power = -power
	
	// Get the absolute value of x
	y := float64(x)
	a := math.Abs(y)

	// Scale down the integer to a float for rounding
	pow := float64(power)
	p := math.Pow(10, pow)
	s := float64(a) / p

	// Get the fractional part to be rounded
	_, f := math.Modf(s)
	
	// Round up or down appropriately
	if f >= 0.5 {
		
		r = math.Ceil(s)
	
	} else {
		
		r = math.Floor(s)
	}
	
	// Multiply by the scaling term to return to the original magnitude
	ar := r * p

	// Add the sign back and return as int
	return int64(math.Copysign(ar, y))
}

// RoundFloat rounds a base ten float64 to the given decimal precision.
// Precision may be positive, representing the number of decimal places,
// or negative, representing the nearest power of ten to which the float 
// should be rounded.
func RoundFloat(x float64, precision int) float64 {
	
	var r float64

	// Handle negative precision with integer rounding
	if precision < 0 {
		
		i, _ := math.Modf(x)
		return float64(RoundInt(int64(i), precision)) 
	}

	// Get the absolute value of x
	a := math.Abs(x)
	
	// Scale up the float for rounding
	p := math.Pow(10, float64(precision))
	s := a * p

	// Get the fractional part to be rounded
	_, f := math.Modf(s)
	
	// Round up or down appropriately
	if f >= 0.5 {
		
		r = math.Ceil(s)
	
	} else {
		
		r = math.Floor(s)
	}

	// Divide by the scaling term to return to the original magnitude
	ar := r / p

	// Add the sign back
	return math.Copysign(ar, x)
}

// FormatThousands converts an int64 into a string formatted using a comma 
// separator for thousands.
func FormatThousands(x int64) string {

	var (
		xstr string
		xslice []byte
		fslice []byte
		lenx int
		lenf int
		commas int
	)

	// Get the number as a byte slice
	xstr = strconv.FormatInt(x, 10)
	xslice = []byte(xstr)
	lenx = len(xslice)

	// Determine the number of commas depending on the sign of x
	if x < 0 {

		commas = (lenx -2) / 3
		lenf = lenx + commas

	} else {

		commas = (lenx -1) / 3
		lenf = lenx + commas
		
	}

	// Create an empty byte slice for the formatted number
	fslice = make([]byte, lenf)

	// Copy the digits from right to left, adding commas
	i := lenx - 1 
	j := lenf - 1

	// Copy the digits in batches of three
	for k := 0; k < commas; k++ {

		for l := 0; l < 3; l++ {

			fslice[j] = xslice[i]
			i--
			j--
		}

		// Add the comma
		fslice[j] = []byte(",")[0]
		j--
	}

	// Copy the remaining digits
	for ; i >= 0; i, j = i - 1, j - 1 {

		fslice[j] = xslice[i]
	}

	return string(fslice)
}

// FormatInt converts an int64 to a formatted string. The int is rounded
// to the given precision and formatted using a comma separator for thousands.
func FormatInt(x int64, precision int) string {

	return FormatThousands(RoundInt(x, precision))
}

// FormatFloat converts a float64 to a formatted string. The float is rounded
// to the given precision and formatted using a comma separator for thousands.
func FormatFloat(x float64, precision int) string {

	// Round the float and get the decimal and fractional parts
	r := RoundFloat(x, precision)
	i, f := math.Modf(r)
	is := FormatThousands(int64(i))

	// If precision is less than one return the formatted integer part
	if precision <= 0 {

		return is
	}

	// Otherwise convert the fractional part to a string 
	fs := strconv.FormatFloat(f, 'f', precision, 64)

	// And get the digits after the decimal point
	if x < 0 {

		fs = fs[3:]

	} else {

		fs = fs[2:]
	}

	// Concatenate the decimal and fractional parts and return
	return is + "." + fs
}
