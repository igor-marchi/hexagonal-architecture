package cli

import (
	"fmt"

	"github.com/igor-marchi/go-hexagonal/application"
)

func Run(
	service application.IProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product ID %s with name %s has been created with the price %f and status %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
		break
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been enabled", res.GetName())
		break
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product %s has been disabled", res.GetName())
		break
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
		break
	}

	return result, nil
}
