package framed

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

// Invalidate runs row validate using callback
func (t *Table) Invalidate(cb func(*Row) error) error {
	for _, row := range t.Rows {
		if err := cb(row); err != nil {
			return RowValidationFailedError(row.Index, err)
		}
	}

	return nil
}
