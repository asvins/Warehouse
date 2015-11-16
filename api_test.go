package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/asvins/router"
)

var (
	_headers map[string]string
	products map[string]Product
)

func _addProduct(p Product) {

	products[p.Name] = p
}

func populateProducts() {
	_addProduct(Product{Name: "coke", Description: "From Coke", Type: 1, CurrQuantity: 60, MinQuantity: 50})
	_addProduct(Product{Name: "h2oh", Description: "From AmBev", Type: 2, CurrQuantity: 100, MinQuantity: 50})
	_addProduct(Product{Name: "pepsi", Description: "From Pepsico", Type: 3, CurrQuantity: 10, MinQuantity: 20})
	_addProduct(Product{Name: "original", Description: "From AmBev", Type: 4, CurrQuantity: 70, MinQuantity: 50})
	_addProduct(Product{Name: "kuat", Description: "From Coke", Type: 5, CurrQuantity: 80, MinQuantity: 90})
	_addProduct(Product{Name: "guarana", Description: "From From AmBev", Type: 6, CurrQuantity: 5, MinQuantity: 10})
	_addProduct(Product{Name: "mate", Description: "From Sei la", Type: 7, CurrQuantity: 60, MinQuantity: 50})
	_addProduct(Product{Name: "soda", Description: "From Coke company", Type: 8, CurrQuantity: 70, MinQuantity: 60})
	_addProduct(Product{Name: "juice", Description: "From Mother Nature", Type: 9, CurrQuantity: 90, MinQuantity: 110})
}

func getBytes(p Product) []byte {
	bjson, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return bjson
}

func init() {
	_headers = make(map[string]string)
	products = make(map[string]Product)
	populateProducts()
}

func makeRequest(httpMethod string, url string, requestObj []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(requestObj))
	addHeaders(req, headers)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

// get - http://127.0.0.1:8080/api/inventory/product/:id
func productExists(id int) bool {
	response, err := makeRequest(router.GET, "http://127.0.0.1:8080/api/inventory/product/"+strconv.Itoa(id), make([]byte, 1), _headers)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] productExists: ", string(body), " StatusCode: ", response.StatusCode)

	return response.StatusCode == http.StatusOK

}

func getOpenOrder() *Order {
	orderResponse, err := makeRequest(router.GET, "http://127.0.0.1:8080/api/inventory/order/open", make([]byte, 1), _headers)
	if err != nil {
		panic(err)
	}

	defer orderResponse.Body.Close()
	orderBody, _ := ioutil.ReadAll(orderResponse.Body)

	order := Order{}
	if err := json.Unmarshal(orderBody, &order); err != nil {
		panic(err)
	}

	return &order
}

// GET - http://127.0.0.1:8080/api/inventory/order/open
func openOrderExists() bool {
	response, err := makeRequest(router.GET, "http://127.0.0.1:8080/api/inventory/order/open", make([]byte, 1), _headers)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] getOpenOrder: ", string(body))

	return response.StatusCode == http.StatusOK
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////// DON'T CHANGE THE TESTS ORDER //////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

// POST - http://127.0.0.1:8080/api/inventory/product
func TestAddProduct(t *testing.T) {
	fmt.Println("[INFO] -- TestAddProduct begin -- ")
	fmt.Println("[INFO] TestAddProduct Should just add a new product to database")

	postProductResponse, err := makeRequest(router.POST, "http://127.0.0.1:8080/api/inventory/product", getBytes(products["coke"]), _headers)
	if err != nil {
		t.Error(err)
	}

	if postProductResponse.StatusCode != http.StatusOK {
		t.Error("[ERROR] Status code should be: ", http.StatusOK, " Got: ", postProductResponse.StatusCode)
	}

	defer postProductResponse.Body.Close()
	body, _ := ioutil.ReadAll(postProductResponse.Body)
	fmt.Println("[INFO] Response: ", string(body))

	coke := &Product{}
	err = json.Unmarshal(body, coke)
	if err != nil {
		t.Error(err)
		return
	}

	products["coke"] = *coke
	if products["coke"].ID == 0 {
		t.Error("[ERROR] Coke id not updated")
	}

	if openOrderExists() {
		t.Error("[ERROR] Order should not have been created!")
	}

	fmt.Println("[INFO] -- TestAddProduct end --\n")
}

// PUT http://127.0.0.1:8080/api/inventory/product/:id
func TestUpdateProductAndCreateOrder(t *testing.T) {
	fmt.Println("[INFO] -- TestUpdateProductAndCreateOrder start --")

	coke := products["coke"]
	coke.CurrQuantity = 30
	products["coke"] = coke

	response, err := makeRequest(router.PUT, "http://127.0.0.1:8080/api/inventory/product/"+strconv.Itoa(products["coke"].ID), getBytes(products["coke"]), _headers)
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Error("[ERROR] Status code should be: ", http.StatusOK, " Got: ", response.StatusCode)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] Response: ", string(body))

	if !openOrderExists() {
		t.Error("[ERROR] Order should have been created!")
	}

	fmt.Println("[INFO] -- TestUpdateProductAndCreateOrder end --\n")
}

// DELETE http://127.0.0.1:8080/api/inventory/product/:id
func TestDeleteProduct(t *testing.T) {
	fmt.Println("[INFO] -- TestDeleteProduct start --")

	response, err := makeRequest(router.DELETE, "http://127.0.0.1:8080/api/inventory/product/"+strconv.Itoa(products["coke"].ID), getBytes(products["coke"]), _headers)
	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] Response: ", string(body))

	if productExists(products["coke"].ID) {
		t.Error("[ERROR] Product with id", products["coke"].ID, " Should have been deleted")
	}

	fmt.Println("[INFO] -- TestDeleteProduct end --\n")
}

func TestAddMultipleProducts(t *testing.T) {
	fmt.Println("[INFO] -- TestDeleteProduct start --")

	for _, value := range products {
		postProductResponse, err := makeRequest(router.POST, "http://127.0.0.1:8080/api/inventory/product", getBytes(value), _headers)
		if err != nil {
			t.Error(err)
		}

		if postProductResponse.StatusCode != http.StatusOK {
			t.Error("[ERROR] Status code should be: ", http.StatusOK, " Got: ", postProductResponse.StatusCode)
		}

		defer postProductResponse.Body.Close()
		body, _ := ioutil.ReadAll(postProductResponse.Body)

		currProd := &Product{}
		json.Unmarshal(body, currProd)
		products[value.Name] = *currProd
	}

	fmt.Println("[INFO] -- TestDeleteProduct end --\n")
}

// PUT http://127.0.0.1:8080/api/inventory/product/:id
func TestUpdateProductAndRemoveFromOrder(t *testing.T) {
	fmt.Println("[INFO] -- TestUpdateProductAndRemoveFromOrder start --")

	juice := products["juice"]
	juice.CurrQuantity = 200
	products["juice"] = juice

	response, err := makeRequest(router.PUT, "http://127.0.0.1:8080/api/inventory/product/"+strconv.Itoa(products["juice"].ID), getBytes(products["juice"]), _headers)
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Error("[ERROR] Status code should be: ", http.StatusOK, " Got: ", response.StatusCode)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] Response: ", string(body))

	order := getOpenOrder()
	for _, prod := range order.Pproducts {
		if prod.ProductId == products["juice"].ID {
			t.Error("[ERROR] product should have been deleted!")
		}
	}

	fmt.Println("[INFO] -- TestUpdateProductAndRemoveFromOrde end --\n")
}

// GET http://127.0.0.1:8080/api/inventory/product/:id/consume/:quantity
func TestConsumeProduct(t *testing.T) {
	fmt.Println("[INFO] -- TestConsumeProduct start --")

	id := strconv.Itoa(products["h2oh"].ID)
	quantity := "60"
	response, err := makeRequest(router.GET, "http://127.0.0.1:8080/api/inventory/product/"+id+"/consume/"+quantity, make([]byte, 1), _headers)
	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("[INFO] getOpenOrder: ", string(body))

	if response.StatusCode != http.StatusOK {
		t.Error("[ERROR] statudCode should be: ", http.StatusOK, " Got: ", response.StatusCode)
	}

	order := getOpenOrder()

	productAdded := false
	for _, prod := range order.Pproducts {
		if prod.ProductId == products["h2oh"].ID {
			productAdded = true
		}
	}

	if !productAdded {
		t.Error("[EROR] Product NOT added after consume request")
	}

	fmt.Println("[INFO] -- TestConsumeProduct end --\n")
}

// PUT http://127.0.0.1:8080/api/inventory/order/:id/approve
func TestApproveOrder(t *testing.T) {
	order := getOpenOrder()
	id := strconv.Itoa(order.ID)
	response, err := makeRequest(router.PUT, "http://127.0.0.1:8080/api/inventory/order/"+id+"/approve", make([]byte, 1), _headers)

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Error("[ERROR] Status code should be: ", http.StatusOK, " Got: ", response.StatusCode)
	}

	if openOrderExists() {
		t.Error("[ERROR] Open order shouldn't exist")
	}
}
