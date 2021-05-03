package handle
import
	"fmt"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Эта строка должна быть видна в браузере")
}

func login(l http.ResponseWriter, k *http.Request) {
	l.Header().Set("Content-Type", "application/json")

	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := User{}

	err := json.NewDecoder(k.Body).Decode(&params)

	if err == io.EOF {
		log.Println("The error:", err)
		fmt.Fprint(l, "Please send a request body")
		return
	} else if err != nil {
		log.Println("The error", err)
		http.Error(l, err.Error(), 500)
		return
	}

	fmt.Fprint(l, "запрос на логин с параметрами:", params)

	email := params.Email
	password := params.Password
	fmt.Println("the user:", email)
	fmt.Println("the password:", password)

	resp, err2 := cli.GenerateToken(context.Background(),
		&ps.Request{
			Email:    email,
			Password: password,
		})

	if err2 != nil {
		log.Fatalf("could not get answer: %v", err2)
	}
	log.Println("Token and expires_at are:", resp.Token, resp.ExpiresAt)
	fmt.Fprint(l, "Token and expires_at are:", resp.Token, resp.ExpiresAt)
}