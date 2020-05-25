# Matrix Library

## Table of Contents

- [Matrix Library](#matrix-library)
  - [Table of Contents](#table-of-contents)
  - [Matrix](#matrix)
    - [CRUD](#crud)
      - [Create](#create)
        - [One](#one)
        - [Many](#many)
      - [Read](#read)
        - [One](#one-1)
        - [Many](#many-1)
        - [All](#all)
        - [Row](#row)
        - [Rows](#rows)
        - [Column](#column)
        - [Columns](#columns)
      - [Update](#update)
        - [One](#one-2)
        - [Many](#many-2)
        - [All](#all-1)
        - [Row](#row-1)
        - [Rows](#rows-1)
        - [Column](#column-1)
        - [Columns](#columns-1)
      - [Delete](#delete)
        - [One](#one-3)
        - [Many](#many-3)
        - [All](#all-2)
        - [Row](#row-2)
        - [Rows](#rows-2)
        - [Column](#column-2)
        - [Columns](#columns-2)
    - [Arithmetic](#arithmetic)
      - [Multiply](#multiply)
        - [One](#one-4)
        - [Many](#many-4)
      - [Divide](#divide)
        - [One](#one-5)
        - [Many](#many-5)
      - [Add](#add)
        - [One](#one-6)
        - [Many](#many-6)
      - [Subtract](#subtract)
        - [One](#one-7)
        - [Many](#many-7)
      - [Exponent](#exponent)
        - [One](#one-8)
        - [Many](#many-8)
    - [Other](#other)
      - [Scale](#scale)
      - [AddRow](#addrow)
      - [RemoveRow](#removerow)
      - [AddColumn](#addcolumn)
      - [RemoveColumn](#removecolumn)
    - [Metadata](#metadata)
      - [GetRows](#getrows)
      - [GetColumns](#getcolumns)
      - [GetScale](#getscale)
      - [Print (Human Readable)](#print-human-readable)
    - [String Serializer](#string-serializer)

---

## Matrix

### CRUD
#### Create
##### One

**Description:** Creates one matrix.<br>

- *Leaving `rows` empty instantiates an empty matrix.*
  
**Interface:** ```Create(capacityX int, capacityY int, rows ...[]int) ([][]int, error)```

##### Many

**Description:** Creates ***n*** matrices.

#### Read
##### One

**Description:** Returns the value of 1 element.<br>
**Interface:** ```Read(matrix [][]int, coordinateX int, coordinateY int) (int, error)``` 

##### Many

**Description:** Returns values based on an list of coordinates.<br>
**Interface:** ```ReadMany(matrix [][]int, coordinateX int, coordinateY int) (int, error)```

{coordinateX int, coordinateY int}

##### All

**Description:** <br>
**Interface:** ```ReadAll(matrix [][]int) ([][]int, error)```

##### Row
##### Rows
##### Column
##### Columns

#### Update
##### One

**Description:** <br>
**Interface:** ``` ```

##### Many
##### All
##### Row
##### Rows
##### Column
##### Columns

#### Delete
*Sets values to 0. Does not remove the element.*
##### One
##### Many
##### All
##### Row
##### Rows
##### Column
##### Columns

### Arithmetic
#### Multiply
##### One
##### Many
#### Divide
##### One
##### Many
#### Add
##### One
##### Many
#### Subtract
##### One
##### Many
#### Exponent
##### One
##### Many

### Other
#### Scale
> **Description:** Scales to a specific size.
> `Scale(xCapacity, yCapacity) ([][]int, error)`
#### AddRow
#### RemoveRow
#### AddColumn
#### RemoveColumn

### Metadata
#### GetRows
#### GetColumns
#### GetScale
> **Description**: Get scale (rows * columns) <br>
> `GetScale(matrix) (map[string]int{columns int, rows int}, error)`
#### Print (Human Readable)

### String Serializer
> **Description:** Parses a with P.E.M.D.A.S. rules.
> **Example:** `(matrix1 + matrix2 - matrix3) / matrix4 * matrix5 ^ 3`


