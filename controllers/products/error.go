package products

type ResponseError struct {
	Message string `json:"message" binding:"required"`
	Status  int    `json:"status" binding:"required"`
}
