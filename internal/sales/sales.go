package sales

import "fmt"

type Checkout interface {
	Checkout([]int) (string, error)
}

func Result(checkout Checkout, list []int) (string, error) {
	return checkout.Checkout(list)
}

type Bank struct{}


func (data Bank) Checkout(items []int) (string, error) {
	var total uint16 = 0
	for _, v := range items {
		total += uint16(v)
	}
	return fmt.Sprintf("Total harga %d dan pembayaran dengan Bank", total), nil
}

type Online struct{}

func (data Online) Checkout(items []int) (string, error) {
	var total uint16 = 0
	for _, v := range items {
		total += uint16(v)
	}
	return fmt.Sprintf("Total harga %d dan pembayaran dengan Online", total), nil
}

type Fictional struct{
	list []uint
}

func (data Fictional) GetList() string {
	var total uint16 
	for _, v := range data.list {
		total += uint16(v)
	}
	return fmt.Sprintf("Total %d", total)
}


func (data Fictional) Checkout(items []int) (string, error) {
	var addition uint16

	for _,v := range items {
		if v <= 0 {
			return "", fmt.Errorf("Error bosku")
		}
	}

	for _,v := range items {
		addition += uint16(v)
	}
	data.list = append(data.list, uint(addition))
	return "", nil
}
