package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	name  string
	price float64
	count int
}

type Store struct {
	products map[string]*Product
}

func NewStore() *Store {
	return &Store{
		products: make(map[string]*Product),
	}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	key := strings.ToLower(name)

	if _, exists := s.products[key]; exists {
		return fmt.Errorf("%s already exists", name)
	}
	if price <= 0 {
		return fmt.Errorf("price should be positive")
	}
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	s.products[key] = &Product{
		name:  strings.ToLower(name),
		price: price,
		count: count,
	}
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	key := strings.ToLower(name)
	p, exists := s.products[key]
	if !exists {
		return 0, fmt.Errorf("invalid product name")
	}
	return p.count, nil
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	key := strings.ToLower(name)
	p, exists := s.products[key]
	if !exists {
		return 0, fmt.Errorf("invalid product name")
	}
	return p.price, nil
}

func (s *Store) Order(name string, count int) error {
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	key := strings.ToLower(name)
	p, exists := s.products[key]
	if !exists {
		return fmt.Errorf("invalid product name")
	}

	if p.count == 0 {
		return fmt.Errorf("there is no %s in the store", name)
	}

	if p.count < count {
		return fmt.Errorf("not enough %s in the store. there are %d left", name, p.count)
	}

	p.count -= count
	return nil
}

func (s *Store) ProductsList() ([]string, error) {
	var names []string
	for _, p := range s.products {
		if p.count > 0 {
			names = append(names, p.name)
		}
	}
	if len(names) == 0 {
		return nil, errors.New("store is empty")
	}

	sort.Strings(names)
	return names, nil
}
