package handlers

import (
	"net/http"
	"simplems/data"
)

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete by product ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"  Format(int64)
// @Success      204
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /products/{id} [delete]
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	prodIndex := r.Context().Value(ValidatedProductIndexKey{}).(int)
	data.ProductsInstance().DeleteProduct(prodIndex)
	rw.WriteHeader(http.StatusNoContent)
}
