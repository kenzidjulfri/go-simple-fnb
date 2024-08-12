package services

import (
	"errors"
	"log"

	"github.com/kenzidjulfri/config"
	"github.com/kenzidjulfri/constant"
	"github.com/kenzidjulfri/model"
	"github.com/kenzidjulfri/repository"
	"github.com/kenzidjulfri/request"
	"github.com/kenzidjulfri/response"
	"gorm.io/gorm"
)

type BaseService interface {
	CreateOrder(req request.OrderRequest) (*response.OrderResponse, error)
	GetBill(tableId uint) (*response.BillResponse, error)
	GetMenu() ([]response.MenuResponse, error)
}

type BaseServiceImpl struct {
	printerRepo repository.PrinterRepository
	productRepo repository.ProductRepository
	variantRepo repository.VariantRepository
	promoRepo   repository.PromoRepository
	orderRepo   repository.OrderRepository
	db          *gorm.DB
}

var BaseSvc BaseService

func GetBaseService() BaseService {
	return &BaseServiceImpl{
		orderRepo:   repository.GetOrderRepository(),
		printerRepo: repository.GetPrinterRepository(),
		productRepo: repository.GetProductRepository(),
		variantRepo: repository.GetVariantRepository(),
		promoRepo:   repository.GetPromoRepository(),
		db:          config.DB,
	}
}

func (s *BaseServiceImpl) CreateOrder(req request.OrderRequest) (*response.OrderResponse, error) {
	order := &model.Order{
		TableID: req.TableID,
	}

	orderDetails := []model.OrderDetail{}
	printers := make(map[string]string, 0)
	for _, detail := range req.Details {
		orderDetail := &model.OrderDetail{
			ItemID: detail.ItemID,
			Qty:    detail.Qty,
		}

		switch detail.ItemType {
		case constant.ItemBase.String():
			product, err := s.productRepo.GetByID(s.db, detail.ItemID)
			if err != nil {
				return nil, err
			}

			orderDetail.ItemName = product.Name
			orderDetail.BasePrice = product.Price
			printers = s.updatePrinters(printers, product.Category)
		case constant.ItemVariant.String():
			variant, err := s.variantRepo.GetByID(s.db, detail.ItemID)
			if err != nil {
				return nil, err
			}

			orderDetail.ItemName = variant.Name
			orderDetail.BasePrice = variant.Product.Price + variant.AdditionalPrice
			printers = s.updatePrinters(printers, variant.Product.Category)
		case constant.ItemPromo.String():
			promo, err := s.promoRepo.GetByID(s.db, detail.ItemID)
			if err != nil {
				return nil, err
			}

			for _, promoDetail := range promo.Details {
				printers = s.updatePrinters(printers, promoDetail.Product.Category)
			}

			orderDetail.ItemName = promo.Name
			orderDetail.BasePrice = promo.Price
		default:
			return nil, errors.New("invalid item type")
		}

		orderDetail.SumPrice = orderDetail.BasePrice * float64(orderDetail.Qty)
		orderDetails = append(orderDetails, *orderDetail)
	}

	order.Details = orderDetails
	err := s.orderRepo.Create(s.db, order)
	if err != nil {
		return nil, err
	}

	respPrinters := []response.Printer{}
	for k, v := range printers {
		respPrinters = append(respPrinters, response.Printer{
			Code: k,
			Name: v,
		})
	}

	resp := &response.OrderResponse{
		OrderID:  order.ID,
		Printers: respPrinters,
	}

	return resp, nil
}

func checkCategory(categoryName string) *string {
	if categoryName != constant.CategoryFood.String() && categoryName != constant.CategoryBeverage.String() {
		return nil
	}

	return &categoryName
}

func (s *BaseServiceImpl) updatePrinters(printers map[string]string, usage string) map[string]string {
	if checkCategory(usage) != nil {
		printer, err := s.printerRepo.GetByUsage(s.db, usage)
		if err != nil {
			return printers
		}

		if _, found := printers[printer.Code]; !found {
			printers[printer.Code] = printer.Name
		}
	}

	return printers
}

func (s *BaseServiceImpl) GetBill(id uint) (*response.BillResponse, error) {
	order, err := s.orderRepo.GetByID(s.db, id)
	if err != nil {
		return nil, err
	}

	printer, err := s.printerRepo.GetByUsage(s.db, "BILL")
	if err != nil {
		return nil, err
	}

	billDetails := []response.BillDetailResponse{}
	bill := response.BillResponse{
		Table: order.Table.Name,
		Printers: []response.Printer{
			{
				Code: printer.Code,
				Name: printer.Name,
			},
		},
	}

	var totalPrice float64 = 0
	for _, detail := range order.Details {
		log.Printf("detail: %v", detail)
		billDetails = append(billDetails, response.BillDetailResponse{
			ItemName:  detail.ItemName,
			Qty:       detail.Qty,
			BasePrice: detail.BasePrice,
			SumPrice:  detail.SumPrice,
		})

		totalPrice += detail.SumPrice
	}

	bill.Details = billDetails
	bill.TotalPrice = totalPrice

	return &bill, nil
}

func (s *BaseServiceImpl) GetMenu() ([]response.MenuResponse, error) {
	menus := []response.MenuResponse{}
	foods, err := s.productRepo.GetAllByCategory(s.db, constant.CategoryFood.String())
	if err != nil {
		return nil, err
	}

	foodItems := []response.SubMenuResponse{}
	for _, food := range foods {
		if food.Variants != nil && len(food.Variants) > 0 {
			for _, variant := range food.Variants {
				foodItems = append(foodItems, response.SubMenuResponse{
					ItemID:    variant.ID,
					ItemName:  variant.Name,
					ItemType:  constant.ItemVariant.String(),
					BasePrice: food.Price + variant.AdditionalPrice,
				})
			}
		} else {
			foodItems = append(foodItems, response.SubMenuResponse{
				ItemID:    food.ID,
				ItemName:  food.Name,
				ItemType:  constant.ItemBase.String(),
				BasePrice: food.Price,
			})
		}
	}

	menus = append(menus, response.MenuResponse{
		Category: constant.CategoryFood.String(),
		Items:    foodItems,
	})

	beverages, err := s.productRepo.GetAllByCategory(s.db, constant.CategoryBeverage.String())
	if err != nil {
		return nil, err
	}

	beverageItems := []response.SubMenuResponse{}
	for _, beverage := range beverages {
		if beverage.Variants != nil && len(beverage.Variants) > 0 {
			for _, variant := range beverage.Variants {
				beverageItems = append(beverageItems, response.SubMenuResponse{
					ItemID:    variant.ID,
					ItemName:  variant.Name,
					ItemType:  constant.ItemVariant.String(),
					BasePrice: beverage.Price + variant.AdditionalPrice,
				})
			}
		} else {
			beverageItems = append(beverageItems, response.SubMenuResponse{
				ItemID:    beverage.ID,
				ItemName:  beverage.Name,
				ItemType:  constant.ItemBase.String(),
				BasePrice: beverage.Price,
			})
		}
	}

	menus = append(menus, response.MenuResponse{
		Category: constant.CategoryBeverage.String(),
		Items:    beverageItems,
	})

	promos, err := s.promoRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}

	promoItems := []response.SubMenuResponse{}
	for _, promo := range promos {
		promoItems = append(promoItems, response.SubMenuResponse{
			ItemID:    promo.ID,
			ItemName:  promo.Name,
			ItemType:  constant.ItemPromo.String(),
			BasePrice: promo.Price,
		})
	}

	menus = append(menus, response.MenuResponse{
		Category: "PROMO",
		Items:    promoItems,
	})

	return menus, nil
}
