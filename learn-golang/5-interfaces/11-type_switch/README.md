# Type Switches

A [type switch](https://go.dev/tour/methods/16) makes it easy to do several type assertions in a series.

A type switch is similar to a regular switch statement, but the cases specify _types_ instead of _values_.

```go
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
```

`fmt.Printf("%T\n", v)` prints the _type_ of a variable.

## Assignment

After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the `getExpenseReport` function _using a type switch_.

1. [ ] If the `expense` is an `email`, return the email's `toAddress` and the `cost` of the email.
2. [ ] If the `expense` is an `sms`, return the sms's `toPhoneNumber` and its `cost`.
3. [ ] If the `expense` has any other underlying type, return an empty string and `0.0` for the cost.
