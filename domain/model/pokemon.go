package model

type Pokemon struct {
  Id int
  Name string
  Type1 string
  Type2 string `csv:"Type2,omitempty"`
  Total int
  HP int
  Attack int
  Defense int
  SpAtk int
  SpDef int
  Generation int
  Legendary bool
}

