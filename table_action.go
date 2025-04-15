package framed

// PipelineAction defines an action interface to modify
// table rows and options as needed.
type PipelineAction interface {

	// ActionName returns the name of pipeline action
	ActionName() string

	// Execute executes the pipeline action
	Execute(*Table) (*Table, error)
}

// Execute is called by [Table] when pipeline actions are executed
func (t *Table) Execute(actions ...PipelineAction) (*Table, error) {
	df := t

	var err error
	for _, action := range actions {
		df, err = action.Execute(df)
		if err != nil {
			return nil, err
		}
	}

	return df, nil
}

// ExecuteS runs pipeline actions and dereferences the resulting [Table] back to self
func (t *Table) ExecuteS(actions ...PipelineAction) error {
	df, err := t.Execute(actions...)

	if err != nil {
		return err
	}

	*t = *df

	return nil
}
