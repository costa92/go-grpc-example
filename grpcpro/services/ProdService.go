package services

import (
	"context"
	"fmt"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	fmt.Println(in.ProdId)
	return &ProdResponse{ProdStock: 20}, nil
}
