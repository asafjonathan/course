package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tailor/pkg/driver"
	"tailor/pkg/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/streadway/amqp"
)

type Product struct {
	Id    int
	Title string
	Image string
	Likes int
}

type Response struct {
	Products []Product `json:"products"`
}

var validate *validator.Validate

type ProductRequest struct {
	Title string `validate:"required"`
	Image string `validate:"required"`
}

func create(connection *driver.DB, w http.ResponseWriter, productRequest ProductRequest) Product {
	lastInsertId := 0
	err := connection.SQL.QueryRow("INSERT INTO products (title, image) VALUES($1,$2) RETURNING id", productRequest.Title, productRequest.Image).Scan(&lastInsertId)
	if err != nil {
		helpers.ServerError(w, err)
	}
	product := Product{
		Id:    lastInsertId,
		Image: productRequest.Image,
		Title: productRequest.Title,
		Likes: 0,
	}
	return product
}
func getProductById(connection *driver.DB, w http.ResponseWriter, productId int) (Product, int) {
	rows, err := connection.SQL.Query("SELECT id, title, image, likes FROM products WHERE id = $1 LIMIT 1", productId)
	if err != nil {
		helpers.ServerError(w, err)
	}
	defer rows.Close()
	var product Product
	var title, image string
	var id, likes int
	counter := 0
	for rows.Next() {
		err := rows.Scan(&id, &title, &image, &likes)
		log.Println(err)
		if err != nil {
			helpers.ServerError(w, err)
		}
		counter++
		product = Product{Id: id, Title: title, Image: image, Likes: likes}
	}
	if err = rows.Err(); err != nil {
		helpers.ServerError(w, err)
	}
	return product, counter
}
func getAllProducts(connection *driver.DB, w http.ResponseWriter) []Product {
	rows, err := connection.SQL.Query("SELECT id, title, image, likes FROM products")
	if err != nil {
		helpers.ServerError(w, err)
	}
	defer rows.Close()
	var products []Product
	var title, image string
	var id, likes int
	for rows.Next() {
		err := rows.Scan(&id, &title, &image, &likes)
		if err != nil {
			helpers.ServerError(w, err)
		}
		products = append(products, Product{Id: id, Title: title, Image: image, Likes: likes})
	}
	if err = rows.Err(); err != nil {
		helpers.ServerError(w, err)
	}
	return products
}
func GetProductsEndPoint(w http.ResponseWriter, req *http.Request) {
	conn, _ := driver.ConnectSQL()
	products := getAllProducts(conn, w)
	defer conn.SQL.Close()
	helpers.ResponseSuccess(w, products, "products")
}

func likePublisher(product Product) bool {
	conn, ch := driver.RabbitChannel()
	bytes, err := json.Marshal(Product{
		Id:    product.Id,
		Image: product.Image,
		Title: product.Title,
		Likes: product.Likes,
	})
	if err != nil {
		panic(err)
	}
	err = ch.Publish("", "ProductQueue", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	})
	if err != nil {
		log.Println(err)
		return false
	}

	conn.Close()
	ch.Close()
	return true
}
func LikeEndPoint(w http.ResponseWriter, req *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(req, "productId"))
	if err != nil {
		helpers.ResponseError(w, "Bad request", 400)
		return
	}
	conn, _ := driver.ConnectSQL()
	product, affectedRows := getProductById(conn, w, productId)
	defer conn.SQL.Close()
	if affectedRows == 0 {
		helpers.ResponseError(w, "Product not found", 404)
		return
	}
	isPublish := likePublisher(product)
	if !isPublish {
		helpers.ServerError(w, err)
	}
	helpers.ResponseSuccess(w, "Success to publish to queue", "message")
}
func CreateEndPoint(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var productRequest ProductRequest
	err := decoder.Decode(&productRequest)
	if err != nil {
		helpers.ResponseError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = validateStruct(productRequest)
	if err != nil {
		helpers.ResponseError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	conn, _ := driver.ConnectSQL()
	product := create(conn, w, productRequest)
	defer conn.SQL.Close()

	helpers.ResponseSuccess(w, product, "product")
}

func validateStruct(product ProductRequest) error {
	validate = validator.New()
	err := validate.Struct(product)
	if err != nil {
		return err
	}
	return nil
}
