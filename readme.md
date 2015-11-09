# decimals
Decimals is a small library of functions for rounding and formatting base ten numbers in Go. These are functions that are either missing from the standard libraries or are more convenient for formatting numbers as a human would read them.

### Installation
Install with go get.

```sh
go get github.com/olihawkins/decimals
```

### Tests
Use `go test` to run the tests.

### Documentation
See the [godoc][gd] for the full documentation.

### Rounding
Round int64 and float64 numbers.
```golang
decimals.RoundInt(x int64, precision int) int64
decimals.RoundFloat(x float64, precision int) float64
```
Round an integer to the nearest power of ten using a negative value for precision.
```golang
i := decimals.RoundInt(555, 0)  // i = 555
i := decimals.RoundInt(555, -1) // i = 560
i := decimals.RoundInt(555, -2) // i = 600
i := decimals.RoundInt(555, -3) // i = 1000 
```
Round a float to the given number of decimal places using a positive value for precision, or to the nearest power of ten using a negative value for precision.
```golang
f := decimals.RoundFloat(55.555, 2)  // f = 55.56
f := decimals.RoundFloat(55.555, 1)  // f = 55.6
f := decimals.RoundFloat(55.555, 0)  // f = 56
f := decimals.RoundFloat(55.555, -1) // f = 60
f := decimals.RoundFloat(55.555, -2) // f = 100
```

### Formatting
Convert integers and floats to formatted strings with the given decimal precision, using a comma separator for thousands.
```golang
decimals.FormatInt(x int64, precision int) string
decimals.FormatFloat(x float64, precision int) string
```golang
Format an integer rounded to the nearest power of ten using a negative value for precision.
```golang
i := decimals.FormatInt(5555555, 0)  // i = "5,555,555"
i := decimals.FormatInt(5555555, -1) // i = "5,555,560"
i := decimals.FormatInt(5555555, -2) // i = "5,555,600"
i := decimals.FormatInt(5555555, -3) // i = "5,556,000" 
```
Format a float rounded to the given number of decimal places using a positive value for precision, or to the nearest power of ten using a negative value for precision.
```golang
f := decimals.FormatFloat(5555.555, 2)  // f = "5,555.56"
f := decimals.FormatFloat(5555.555, 1)  // f = "5,555.6"
f := decimals.FormatFloat(5555.555, 0)  // f = "5,556"
f := decimals.FormatFloat(5555.555, -1) // f = "5,560"
f := decimals.FormatFloat(5555.555, -2) // f = "5,600"
```
   [gd]: <https://godoc.org/github.com/olihawkins/decimals>



