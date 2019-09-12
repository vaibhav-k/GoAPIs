package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// DB is a struct for the database connection
type DB struct {
	DBCon *sql.DB
}

// User is struct for all users of the API
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ProductStruct is a struct for all products
type ProductStruct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CompanyName string  `json:"company_name"`
	Price       float32 `json:"price"`
	Category    int     `json:"category"`
	Quantity    int     `json:"quantity"`
}

// Cart is a struct for all products in the cart
type Cart struct {
	Username   string `json:"username"`
	Products   []int  `json:"products"`
	Quantities []int  `json:"quantities"`
}

// DateType is a typecasting of the Time struct
type DateType time.Time

func (t DateType) String() string {
	return time.Time(t).String()
}

// RegisterUser registers a new user
func (db DB) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering a user!")
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("SELECT user_id FROM `user_data` WHERE `user_name` = '%s'", user.Username)
	result, err := db.DBCon.Query(s)
	fmt.Println(user.Username)

	flag := 0

	if err != nil {
		panic(err.Error())
	} else if result != nil {
		for result.Next() {
			flag = 1
			id := 0
			err := result.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}

		if flag == 1 {
			fmt.Println("Someone already has taken the username. Please try with a different username.")
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"password": user.Password,
			})
			tokenString, error := token.SignedString([]byte("secret"))
			if error != nil {
				fmt.Println(error)
			}
			fmt.Println(tokenString)
			// if strings.TrimRight(r.Header.Get("Token"), "\n") == tokenString {
			// 	fmt.Println("Passed")
			// }
			//json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
			s = fmt.Sprintf("INSERT INTO `user_data` SET `user_name` = '%s', `password` = '%s'", user.Username, user.Password)
			result, err = db.DBCon.Query(s)

			if err == nil || result != nil {
				fmt.Println(user.Username, "is now registered!")
			} else {
				fmt.Println("Registration failed")
				fmt.Println(err)
			}
		}
	}
}

// FromNow Offset is a number that will multiply TimeUnit and the result will be added to the current time
type FromNow struct {
	Offset   int
	TimeUnit time.Duration
}

func (ts FromNow) String() string {
	t := time.Now().Add(ts.TimeUnit * time.Duration(ts.Offset))
	timeStamp := t.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%v", timeStamp)
}

// LoginUser logs a user in
func (db DB) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging user in!")
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("SELECT user_id FROM `user_data` WHERE `user_name` = '%s' AND `password` = '%s'", user.Username, user.Password)
	result, err := db.DBCon.Query(s)

	flag := 0

	if err != nil {
		panic(err.Error())
	} else if result != nil {
		for result.Next() {
			flag = 1
			id := 0
			err := result.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}
		if flag == 1 {
			// expr := time.Now().Local().Add(time.Hour * 48)
			// now := time.Now()

			//if now.Before(expr) {
			exampleQuery := fmt.Sprintf("UPDATE `user_data` SET `token_expiration` = (NOW() + INTERVAL 2 DAY) WHERE `user_name` = '%s'", user.Username)
			update, er := db.DBCon.Query(exampleQuery)

			if er == nil && update != nil {
				fmt.Println("Updated token expiration time and logged user in!")
			} else {
				panic(er.Error())
			}

			update.Close()
		}
	}
	result.Close()
}

// ViewCart lets the user view his/her cart
func (db DB) ViewCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Viewing items from cart!")
	decoder := json.NewDecoder(r.Body)
	var cart Cart
	err := decoder.Decode(&cart)
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)

	t := fmt.Sprintf("SELECT user_id FROM `user_data` WHERE user_name = '%s' AND token_expiration > NOW()", params["username"])
	rows, errpresent := db.DBCon.Query(t)

	var id string

	fmt.Println(id)
	if errpresent != nil {
		panic(errpresent.Error())
	} else if rows != nil {
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	s := fmt.Sprintf("SELECT product_id, quantity FROM `user_cart`")
	result, err := db.DBCon.Query(s)

	var ids []int
	var quants []int
	var names []string

	if err != nil {
		panic(err.Error())
	} else if result != nil {
		for result.Next() {
			var productID, quantity int
			err := result.Scan(&productID, &quantity)
			if err != nil {
				panic(err.Error())
			}
			ids = append(ids, productID)
			quants = append(quants, quantity)
		}

		for x := range ids {
			s = fmt.Sprintf("SELECT name FROM `products` WHERE `product_id` = %d", ids[x])
			name := ""
			result, err = db.DBCon.Query(s)
			for result.Next() {
				err := result.Scan(&name)
				if err != nil {
					panic(err.Error())
				}
				names = append(names, name)
			}
		}

		type Orders struct {
			Name     string
			Quantity int
		}

		var orders []Orders

		for i := range names {
			var order Orders
			order.Name = names[i]
			order.Quantity = quants[i]
			orders = append(orders, order)
		}
		json.NewEncoder(w).Encode(orders)
	}
}

// AddToCart lets users add products to cart
func (db DB) AddToCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding items to cart!")
	decoder := json.NewDecoder(r.Body)
	var cart Cart
	err := decoder.Decode(&cart)
	if err != nil {
		panic(err)
	}

	s := fmt.Sprintf("SELECT user_id FROM `user_data` WHERE `user_name` = '%s'", cart.Username)
	result, err := db.DBCon.Query(s)

	id := ""

	if err != nil {
		panic(err.Error())
	} else if result != nil {
		for result.Next() {
			err := result.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	arr := [1]int{0}
	for x := range arr {
		s = fmt.Sprintf("INSERT INTO `user_cart`(`user_id`, `product_id`, `quantity`) VALUES (%s, %d, %d)", id, cart.Products[x], cart.Quantities[x])
		result, err = db.DBCon.Query(s)
		if err != nil {
			fmt.Println("Got error")
			panic(err.Error())
		} else if result != nil {
			fmt.Println("Products inserted into cart!")
		}
	}
}

// GetProducts shows a JSON of all available products
func (db DB) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Viewing all the products!")
	var products []ProductStruct
	result, err := db.DBCon.Query("SELECT * FROM `products`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var product ProductStruct
		err := result.Scan(&product.ID, &product.Name, &product.CompanyName, &product.Price, &product.Category, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

// CreateProducts adds a new product to the database
func (db DB) CreateProducts(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product ProductStruct
	err := decoder.Decode(&product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating a new product!")

	// fmt.Println(product)
	// fmt.Println(product.ID, product.Name, product.CompanyName, product.Price, product.Category)
	s := fmt.Sprintf("INSERT INTO `products` SET `product_id` = %d, `name` = '%s', `company_name` = '%s', `price` = %f, `category` = %d, `quantity` = %d", product.ID, product.Name, product.CompanyName, product.Price, product.Category, product.Quantity)
	result, err := db.DBCon.Query(s)

	if err == nil || result != nil {
		fmt.Println(product.Name, "is now inserted!")
	} else {
		fmt.Println("Insertion failed")
		fmt.Println(err)
	}
}

// GetProduct fetches a single product with the given ID
func (db DB) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	s := fmt.Sprintf("SELECT * FROM `products` WHERE product_id = %s", params["id"])
	result, err := db.DBCon.Query(s)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var product ProductStruct
	for result.Next() {
		err := result.Scan(&product.ID, &product.Name, &product.CompanyName, &product.Price, &product.Category, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct updates details of a prodct
func (db DB) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var product ProductStruct
	err := decoder.Decode(&product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Name, product.CompanyName, product.Price, product.Category, product.Quantity)
	s := fmt.Sprintf("UPDATE `products` SET `name` = '%s', `company_name` = '%s', `price` = %f, `category` = %d, `quantity` = %d WHERE product_id = '%s'", product.Name, product.CompanyName, product.Price, product.Category, product.Quantity, params["id"])
	result, err := db.DBCon.Query(s)

	if err == nil || result != nil {
		fmt.Println(product.Name, "is now updated!")
	} else {
		fmt.Println("Updating failed")
		fmt.Println(err)
	}
}

// DeleteProduct deletes a product from the database
func (db DB) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	s := fmt.Sprintf("DELETE FROM `products` WHERE product_id = %s", params["id"])
	result, err := db.DBCon.Query(s)

	if err == nil || result != nil {
		fmt.Println(params["id"], "is now deleted!")
	} else {
		fmt.Println("Deleting failed")
		fmt.Println(err)
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
