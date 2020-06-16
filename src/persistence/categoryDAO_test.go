package persistence

import "testing"

func TestGetCategory(t *testing.T) {
	c,err := GetCategory("BIRDS")
	if err != nil {
		t.Error(err)
	}
	t.Log(c,"categoryId:",c.CategoryId)
}

func TestGetCategoryList(t *testing.T) {
	r,err := GetCategoryList()
	if err != nil {
		t.Error(err)
	}
	t.Log(r[1])
}
