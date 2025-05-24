package framed

import (
	"reflect"
)

type ActionSelect struct {
	configs []*ActionSelectConfig
}

func (a ActionSelect) ActionName() string {
	return "select"
}

func (a ActionSelect) Execute(src *Table) (*Table, error) {
	maxRows := 0
	colLength := 0
	colData := make([][]any, 0)
	columns := make([]string, 0)
	colsFound := make(map[string]bool)
	colDefs := make(map[string]*Definition)

	for _, config := range a.configs {
		cols, err := config.ResolveColumns(src.State.Columns)
		if err != nil {
			return nil, err
		}

		cColumns := make([]string, 0)
		for _, col := range cols {
			if !colsFound[col] {
				colsFound[col] = true
				colLength++
				cColumns = append(cColumns, col)
			}
		}

		for _, col := range cColumns {
			idx := len(colData)
			alias := config.ResolveName(col)
			def := src.State.Definition(col)
			columns = append(columns, alias)

			var colType reflect.Type = nil
			if def != nil {
				colType = def.Type
			}

			if config.modifier != nil {
				mData := make([]any, src.Length())
				for rIndex, row := range src.Rows {
					cv, err := config.modifier(src.State, row)
					if err != nil {
						return nil, err
					}

					mData[rIndex] = cv
				}

				colData = append(colData, mData)
			} else if config.builder != nil {
				// raw assign column values from reader
				values, err := config.builder(src)
				if err != nil {
					return nil, err
				}

				colData = append(colData, values)
			} else if config.values != nil {
				// raw assign column values
				colData = append(colData, config.values)
			} else if def != nil {
				// load from existing table
				colData = append(colData, src.Col(col))
			} else {
				return nil, FailedToResolveColumnDataError(alias)
			}

			// resolve column definition
			if len(config.alias) > 0 && src.State.HasDefinition(alias) {
				colDefs[alias] = src.State.Definition(alias)
			} else {
				if len(colData[idx]) > 0 {
					colType = ToType(colData[idx][0])
				}

				colDefs[alias] = NewDefinition(colType).Copy(def)
			}

			// store max rows count
			if len(colData[idx]) > maxRows {
				maxRows = len(colData[idx])
			}
		}
	}

	// create new table
	df := New().SetOptions(src.Options)
	df.SetDefinitions(colDefs)
	df.UseColumns(columns)
	df.MarkResolved()

	// loop through max rows count and set when value is available
	for i := range maxRows {
		cLen := len(columns)
		rowData := make([]any, cLen)
		for j := range cLen {
			if len(colData[j]) > i {
				rowData[j] = colData[j][i]
			}
		}

		df.AddRow(&Row{
			Index:   i,
			Columns: rowData,
		})
	}

	return df, nil
}

// Select provides access to specific column settings
// and generates the new table.
//
//	newTable, err := table.Execute(
//		...
//		framed.Select(
//			framed.Col("col1"),
//			framed.Col(regexp.MustCompile("^T_")),
//		),
//		...
//	)
func Select(configs ...*ActionSelectConfig) *ActionSelect {
	return &ActionSelect{
		configs: configs,
	}
}
