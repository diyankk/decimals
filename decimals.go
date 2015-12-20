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
// If the rounded number falls outside the minimum and maximum for int64
// the minimum or maximum will be returned instead.
func RoundInt(x int64, precision int) int64 {

	var (
		xstr string = strconv.FormatInt(x, 10)
		xslice = []byte(xstr)
		zeroFrom int = -1
		roundFrom int
	)

	// Map for converting decimal bytes to int64
	decimalInts := map[byte]int64{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}	

	// Array for converting decimal ints to bytes
	decimalBytes := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',}	

	// If precision is not negative return x
	if precision > -1 {

		return x
	}

	// If x is negative remove the sign
	if x < 0 {
	
		xslice = xslice[1:]
	}
	
	// Set the index of the digit to round from
	roundFrom = len(xslice) + precision

	// If rounding to more than one order of magnitude larger than x return 0 
	if roundFrom < 0 {
	
		return 0
	}

	// If rounding to one order of magnitude larger than x round from first digit
	if roundFrom == 0 {
	
		firstDigit := decimalInts[xslice[0]]
		
		if firstDigit < 5 {
				
			return 0
		
		} else {
				
			xslice = append([]byte{'1'}, xslice...)
			zeroFrom = 1
		}
	
	// Otherwise round through the slice from right to left	
	} else {
	
		// Start rounding from the round digit
		roundDigit := decimalInts[xslice[roundFrom]]
	
		// If less than five round from there
		if roundDigit < 5 {
	
			zeroFrom = roundFrom
	
		// Otherwise keep moving left to find the rounding point
		} else {
	
			for i := roundFrom; i > 0; i-- {
			
				j := i - 1
				nextDigit := decimalInts[xslice[j]]
		
				if nextDigit < 9 {
			
					xslice[j] = decimalBytes[nextDigit + 1]
					zeroFrom = i
					break
				}
			}
		
			// If not found add a leading one and round from there
			if zeroFrom == -1 {
		
				xslice = append([]byte{'1'}, xslice...)
				zeroFrom = 1
			}
		}
	}
	
	// Zero all digits after the rounding point
	for i := zeroFrom; i < len(xslice); i++ {
		
		xslice[i] = '0'
	} 
	
	// If x is negative add the sign back
	if x < 0 {
		
		xslice = append([]byte("-"), xslice...)
	}
	
	// Convert the slice back to an int64
	rstr := string(xslice)
	r, _ := strconv.ParseInt(rstr, 10, 64)
	
	return r
}

// RoundFloat rounds a base ten float64 to the given decimal precision.
// Precision may be positive, representing the number of decimal places,
// or negative, representing the nearest power of ten to which the float 
// should be rounded.
func RoundFloat(x float64, precision int) float64 {
	
	// Handle negative precision with integer rounding
	if precision < 0 {
		
		i, _ := math.Modf(x)
		return float64(RoundInt(int64(i), precision)) 
	}

	// Handle positive precision with strconv.FormatFloat()
	rstr := strconv.FormatFloat(x, 'f', precision, 64)
	r, _ := strconv.ParseFloat(rstr, 64)

	return r
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
