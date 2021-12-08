package main

func main() {
	//多个 if 语句可以折叠成 switch

	// NOT BAD
	if foo() {
		// ...
	} else if bar == baz {
		// ...
	} else {
		// ...
	}

	// BETTER
	switch {
	case foo():
		// ...
	case bar == baz:
		// ...
	default:
		// ...
	}
}
