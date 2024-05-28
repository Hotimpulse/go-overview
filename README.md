### Assignment Expression & Implicit vs Explicit

```
        var name string
	name = "Lev"
	number := 9

	fmt.Println("Hello World!", name, number)
	fmt.Printf("%T", number)
```

- all of the types have a default value that will be assigned to a variable (i.e. uint64 is 0)

### Printing to Console & fmt

General

- %v (value default format)
- %T (type)
- %% (literal "%")

Boolean 

- %t (true or false)

Integer

- %b (base2) 
- %o (base8)
- %d (base10)
- %h (base16)

binary / octal / decimal / hexadecimal

Floating Points

- %e (scientific notation)
- %f / %F => same thing (decimal / no exponent)
- %g (for large exponents)

Strings

- %s (default)
- %q (double quoted string)

Width & Precision

- %f (default width, default precision)
- %9f (width 9, default precision)
- %.2f (default width, precision 2)
- %9.2f (width 9, precision 2)
- %9.f (width 9, precision 0)
- 
- 

