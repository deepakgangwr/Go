package variables

// Exported Variables
var Public = "DEEPAK" // 'Public' starts with a capital letter, so it is exported and accessible from other packages

// Unexported Variables
var private = "gangwar" // 'private' starts with a small letter, so it is not accessible from outside this package

// Declaring different types of variables

var (
	// String variables
	Name    = "Deepak"   // Explicit initialization
	Version = "2.2"      // Compiler automatically infers the type as string

	// Integer variables
	Money    = 5000       // Declared as int
	Currency = 6454       // Type inferred as int

	// Float variables
	Dimension = 248.23    // Explicit float64 declaration
	Flo       = 53.44     // Compiler infers as float64

	// Boolean variable
	Flag = false          // Default value false
)

// Constants
const PiConst = 3.14 // Constant value (cannot be changed later)

// Additional Variables
var Person = 123 // Correct way to declare at package level

// Important Note:
// We cannot write 'person := 123' here.
// The short variable declaration ':=' can ONLY be used inside a function like 'main'.
// At the package level (outside functions), we must use 'var' to declare variables.
