package framed

// PipelineAction defines an interface to modify table data and options
type PipelineAction interface {
	ExecName() string
	Execute(*Table) (*Table, error)
}

// Execute is called by Table when pipeline actions are executed
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

// Runs pipeline actions and points the resulting table back to self
func (t *Table) ExecuteS(actions ...PipelineAction) error {
	df, err := t.Execute(actions...)

	if err != nil {
		return err
	}

	*t = *df

	return nil
}
