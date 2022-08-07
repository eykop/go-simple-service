package handlers

import (
	"fmt"
	"net/http"
	"simplems/data"

	"go.uber.org/zap"
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

	productId, err := getProductId(r, p.l)
	if err != nil {
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}

	p.l.Info("Delete Procut", zap.Int("id", productId))
	delErr := data.DeleteProduct(productId)
	if delErr != nil {
		p.l.Error("Failed to delete product", zap.Error(delErr))
		http.Error(rw, fmt.Sprintf("Bad request, could not delete product %v", delErr), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}
