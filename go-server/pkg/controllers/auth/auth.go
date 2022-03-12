package auth

import (
	"fmt"
	"net/http"
	"tailor/pkg/helpers"

	"github.com/thedevsaddam/govalidator"
)

func Login(w http.ResponseWriter, req *http.Request) {

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"email":    []string{"required:email is requierd", "email:must be valid email address"},
		"password": []string{"required:password is requierd", "min: password must be min 6 characters"},
	}

	opts := govalidator.Options{
		Request:         req,      // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	fmt.Println("aaaaa")
	err := map[string]interface{}{"validationError": e}
	if err != nil {
		helpers.ResponseErrors(w, err, http.StatusBadRequest)
		return
	}
}
