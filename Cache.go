package main

import "errors"

const notExists = "key does not exist"

type data struct {
	value interface{}
}

type Cache struct {
	item map[string]data
}

func New() *Cache {
	return &Cache{
		item: make(map[string]data),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	c.item[key] = data{
		value: value,
	}
	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	if item, exists := c.item[key]; exists {
		return item.value, nil
	}
	return nil, errors.New(notExists)
}

func (c *Cache) Delete(key string) error {
	if _, exists := c.item[key]; exists {
		delete(c.item, key)
		return nil
	}
	return errors.New(notExists)
}
