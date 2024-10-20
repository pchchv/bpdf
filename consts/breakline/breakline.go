// The breakline package contains all break line strategies.
package breakline

// Strategy represents a break line strategy.
type Strategy string

const (
	// EmptySpaceStrategy is a break line strategy that counts the length of words to create a new line.
	// Only works in languages that use spaces to separate words.
	EmptySpaceStrategy Strategy = "empty_space_strategy"
	// DashStrategy is a break line strategy that calculates the length for a set of characters unrelated to the meaning of words.
	// This strategy is useful for languages that do not use space between words.
	// A dash at the end of a line is used to separate strings.
	DashStrategy Strategy = "dash_strategy"
)
