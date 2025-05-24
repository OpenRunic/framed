package framed

// PipelineAction defines an action interface to modify
// table rows and options as needed.
type PipelineAction interface {

	// ActionName returns the name of pipeline action
	ActionName() string

	// Execute executes the pipeline action
	Execute(*Table) (*Table, error)
}

// IsEmptyChecker defines interface for checking data is empty
type IsEmptyChecker interface {
	IsEmptyCheck() bool
}

// NumericReader defines interface for reading numeric value
type NumericReader interface {
	NumericRead() float64
}
