package function

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		defer r.Body.Close()
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("body is empty"))
		return
	}

	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	var input Input
	err = json.Unmarshal(dataBytes, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	if input.UserID == nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("user id is empty"))
		return
	}

	// Add a sleep time for simulating database connection
	time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)

	userInfo := UserInfo{
		ID: *input.UserID,
	}
	switch *input.UserID {
	case 10:
		userInfo = UserInfo{
			FirstName:   "Amin",
			LastName:    "Ghasvari",
			PhoneNumber: "09336205449",
			CurrentAddressLocation: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
			Addresses: []string{"Azadegan ST. 13"},
		}
		break
	case 20:
		userInfo = UserInfo{
			FirstName:   "Arman",
			LastName:    "Heydari",
			PhoneNumber: "09362817764",
			CurrentAddressLocation: Point{
				Lat: 22.11,
				Lon: 22.22,
			},
			Addresses: []string{"Saadat abad ST. 52"},
		}
		break
	case 30:
		userInfo = UserInfo{
			FirstName:   "Masoud",
			LastName:    "Golestane",
			PhoneNumber: "09197846219",
			CurrentAddressLocation: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
			Addresses: []string{"Fadak ST. 92", "Tajrish ST. 32"},
		}
		break
	case 40:
		userInfo = UserInfo{
			FirstName:   "Ali",
			LastName:    "Sedaghi",
			PhoneNumber: "09376199092",
			CurrentAddressLocation: Point{
				Lat: 44.11,
				Lon: 44.22,
			},
			Addresses: []string{"Fallah ST. 11", "Gheytarieh ST. 82"},
		}
		break
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("can not find the user"))
		return
	}

	outputByte, err := json.Marshal(userInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(outputByte)
}
