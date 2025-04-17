## Table
---

![Test Status](https://github.com/OpenRunic/framed/actions/workflows/lint-test.yml/badge.svg)

Data manipulation library inspired from Pandas (Python) for golang

#### Download
```
go get -u github.com/OpenRunic/framed
```

#### Options
---

```go
df := framed.New(

  // set max rows to import/load (default = -1)
  framed.WithMaxRows(100),

  // set column separater (default = ',')
  framed.WithSeparator(';'),

  // set column names for table (otherwise automatically generated from first row)
  framed.WithColumns([]string{"id", "col2", ... "colN"}),

  // set column definition
  framed.WithDefinitionType("id", framed.ToType(1)),

  // set column data type reader
  framed.WithTypeReader(func(idx int, value string) reflect.Type {
    return framed.ToType(bool)
  }),
)
```

#### Load data
---
Option callbacks can be used as above on table loaders below

```go
// load table from file
table, err := framed.File("path_to_file.csv")

// load table from url
table, err := framed.URL("http_https_url...csv")

// load table from io.Reader
table, err := framed.Reader(io.Reader)

// load table from slice of string
table, err := framed.Lines([]string{"1....", "2...."})

// load table from slice of raw data
table, err := framed.Raw([][]string{{"1", "..."}, {"2", "..."}})
```

#### Actions Pipeline
---
Run set of pipeline actions to modify data in the table

```go
// add column action
addAction := framed.AddColumn("name", "", func(s *framed.State, r *framed.Row) string {
  return fmt.Sprintf("%s %s", r.At(s.Index("last_name")), r.At(s.Index("first_name")))
})

// drop column action
dropAction := framed.DropColumn("last_name", "age")

// rename column action
renameAction := framed.RenameColumn("fname", "first_name")

// pick specific columns action
pickAction := framed.PickColumn("first_name", "last_name")

// omit specific columns action
omitAction := framed.PickColumn("id", "age")

// make specific changes to table or its options while in pipeline action
updateAction := framed.ChangeOptions(func(tbl *Table) (*Table, error) {
  return tbl, nil
})

// change column type action
changeColTypeAction := framed.ChangeColumnType("age", "", func(s *framed.State, r *framed.Row, a any) string {
  v := a.(int64)
  if v > 18 {
    return "adult"
  } else if v > 12 {
    return "teen"
  }

  return "kid"
})

// modify every table rows action
modifyAction := framed.ModifyRow(func(s *framed.State, r *framed.Row) *framed.Row {
  r.Set(0, framed.ColumnValue(r, 0, 0)+10) // add 10 to every row's column at index 0
  return r
})

// filter rows from table action
filterAction := framed.FilterRow(func(s *framed.State, r *framed.Row) bool {
  return r.Index > 9 // ignore first 10 rows
})

// execute actions to build new table from result
newTable, err := table.Execute(
  addAction,
  dropAction,
  renameAction,
  pickAction,
  omitAction,
  updateAction,
  changeColTypeAction,
  modifyAction,
  filterAction,
)
```

#### Table/Row utilities
---

```go
// manually set columns for table
table.UseColumns([]string{"col1", "col2", ...., "colN"})

// add row to table
table.AddRow(&framed.Row{...})

// get first rows as new table
df := table.Chunk(0, 100)

// add line to table
err := table.InsertLine("a..,b..,..z")

// add slice of strings as line
err := table.InsertSlice([]string{"a..", "b...", "...", "...z"})

// add lines from reader to table
err := table.Read(io.Reader)

// get column definition
def := table.State.Definition("col1")

// get column definition by index
def := table.State.DefinitionAt(0)

// loop through chunks of tables with 100 rows each
for _, cTable := range table.Chunks(100) {
  ...
}

// get first 100 rows as []*framed.Row
rows := table.Slice(0, 100)

// get first row
table.First()

// get last row
table.Last()

// get length of rows
table.Length()

// get row at x index
row := table.At(x)

// get column at x index of row
colValue := row.At(0)

// set column value at x index of row
row.Set(0, "updated_value")

// patch column value at x index of row with type-safety
err := row.Patch(0, "updated_value")

// pick only selected columns from row
columns, err := row.Pick("col2", "col3")

// clone row with selected columns
newRow, err := row.CloneP("col2", "col3")
```

#### Other Utilities
---

```go
// extract reflect.Type from any given value
framed.ToType(10)

// read typesafe column value from column as index 0
value := framed.ColumnValue[string](row, 0, "default_value")
```

### Support

You can file an [Issue](https://github.com/OpenRunic/framed/issues/new).

### Contribute

To contrib to this project, you can open a PR or an issue.
