package model

type AsyncParams struct{
  Type string `form:"type" binding:"required"`
  Items int `form:"items" binding:"required"`
  ItemsPerWorker int `form:"itemsPerWorker" binding:"required"`
}