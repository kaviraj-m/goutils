package main

import (
	"fmt"
	"time"

	"github.com/kaviraj-m/goutils/errorutil"
	"github.com/kaviraj-m/goutils/strutil"
	"github.com/kaviraj-m/goutils/timeutil"
	"github.com/kaviraj-m/goutils/validation"
)

func main() {
	fmt.Println("GoUtils Library Examples")
	fmt.Println("------------------------")

	// Validation examples
	fmt.Println("\n=== Validation Examples ===")
	fmt.Printf("Is 'test@example.com' a valid email? %v\n", validation.IsEmail("test@example.com"))
	fmt.Printf("Is 'https://golang.org' a valid URL? %v\n", validation.IsURL("https://golang.org"))
	fmt.Printf("Is '12345' numeric? %v\n", validation.IsNumeric("12345"))
	fmt.Printf("Is '123abc' alphanumeric? %v\n", validation.IsAlphanumeric("123abc"))
	fmt.Printf("Is '+1234567890' a phone number? %v\n", validation.IsPhoneNumber("+1234567890"))

	// String utility examples
	fmt.Println("\n=== String Utility Examples ===")
	fmt.Printf("Reverse of 'hello': %s\n", strutil.Reverse("hello"))
	fmt.Printf("Truncating 'This is a long text' to 10 chars: %s\n", strutil.Truncate("This is a long text", 10))
	fmt.Printf("Truncating with ellipsis: %s\n", strutil.TruncateWithEllipsis("This is a long text", 10))
	fmt.Printf("Random string (10 chars): %s\n", strutil.RandomString(10))
	fmt.Printf("'HelloWorld' to snake_case: %s\n", strutil.ToSnakeCase("HelloWorld"))
	fmt.Printf("'hello_world' to camelCase: %s\n", strutil.ToCamelCase("hello_world"))
	fmt.Printf("'hello_world' to PascalCase: %s\n", strutil.ToPascalCase("hello_world"))

	// Error utility examples
	fmt.Println("\n=== Error Utility Examples ===")
	err := fmt.Errorf("original error")
	wrappedErr := errorutil.WrapError(err, "additional context")
	fmt.Printf("Wrapped error: %v\n", wrappedErr)

	errChain := errorutil.NewErrorChain()
	errChain.Add(fmt.Errorf("error 1"))
	errChain.Addf("error %d", 2)
	fmt.Printf("Error chain: \n%v\n", errChain.Error())
	fmt.Printf("Has errors: %v\n", errChain.HasErrors())

	// Time utility examples
	fmt.Println("\n=== Time Utility Examples ===")
	now := time.Now()
	fmt.Printf("Current date: %s\n", timeutil.FormatDate(now))
	fmt.Printf("Current date and time: %s\n", timeutil.FormatDateTime(now))

	duration := 36*time.Hour + 12*time.Minute + 5*time.Second
	fmt.Printf("Formatted duration: %s\n", timeutil.FormatDuration(duration))

	pastTime := now.Add(-24 * time.Hour)
	fmt.Printf("Time ago: %s\n", timeutil.TimeAgo(pastTime))

	fmt.Printf("Is today weekend? %v\n", timeutil.IsWeekend(now))
	fmt.Printf("Start of day: %s\n", timeutil.FormatDateTime(timeutil.StartOfDay(now)))
	fmt.Printf("End of day: %s\n", timeutil.FormatDateTime(timeutil.EndOfDay(now)))
	fmt.Printf("Start of month: %s\n", timeutil.FormatDateTime(timeutil.StartOfMonth(now)))
	fmt.Printf("End of month: %s\n", timeutil.FormatDateTime(timeutil.EndOfMonth(now)))
}
