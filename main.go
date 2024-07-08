package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	consumerKey    = "4jz4O5jSggtjNGNvAgbapSUfyNh7WRMdMYAlvDHgztcCAdNM"
	consumerSecret = "3pFRpRMKEvxXSfbUEGOvkO355xE5lYRDcwbPblJcGmb2vcexXXYxDJulcVRlpRfw"
	oauthURL       = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	b2cURL         = "https://sandbox.safaricom.co.ke/mpesa/b2c/v3/paymentrequest"
)

// RequestBody structure to hold the request body
type RequestBody struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	InitiatorName            string `json:"InitiatorName"`
	SecurityCredential       string `json:"SecurityCredential"`
	CommandID                string `json:"CommandID"`
	Amount                   int    `json:"Amount"`
	PartyA                   int    `json:"PartyA"`
	PartyB                   int    `json:"PartyB"`
	Remarks                  string `json:"Remarks"`
	QueueTimeOutURL          string `json:"QueueTimeOutURL"`
	ResultURL                string `json:"ResultURL"`
	Occasion                 string `json:"occasion"`
}

func getOAuthToken() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", oauthURL, nil)
	if err != nil {
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("unable to get access token")
	}

	return token, nil
}

func sendPostRequest(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	token, err := getOAuthToken()
	if err != nil {
		return nil, fmt.Errorf("Error getting OAuth token: %s", err.Error())
	}

	requestBody := &RequestBody{
		OriginatorConversationID: "0cd41716-ff5c-4f7d-a5a2-7bc87b6f3378",
		InitiatorName:            "testapi",
		SecurityCredential:       "qEEbHtb8kcOfLu9VRpXz7o/i5KY/dmedf2CQi/muSpge0SpyKXCEBThJjyrOuYV1rUU1GXkq75ElpxmZvrdBlDdP3u+DaL9vi3mwhUtlb9qaK22jDGoRfgvxccXDB8TAya85fbhID5eDSGxFhEqdH8JfWe5Xltfm4tvtn1MjZej2wmxnY0LBqzk+jgawepRueOqFDMh0zeUcgkPJ3LepPgZr9sK7BVTYRBbRKqbYKNOFkpXYKpYeoJFMDmDm/zoHAeAhuB8FC2I3muuSY1e2LBOK5UctdIN2EylgjxFftIVmimf67jNkGphU3h3no/ZuTu1bGRoBoZcZpl/SLndlng==",
		CommandID:                "BusinessPayment",
		Amount:                   10,
		PartyA:                   600986,
		PartyB:                   254708374149,
		Remarks:                  "Test remarks",
		QueueTimeOutURL:          "https://mydomain.com/b2c/queue",
		ResultURL:                "https://mydomain.com/b2c/result",
		Occasion:                 "ok",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("Error encoding request body: %s", err.Error())
	}

	req, err := http.NewRequest("POST", b2cURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s", err.Error())
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error decoding response: %s", err.Error())
	}

	return result, nil
}

func sendHandler(w http.ResponseWriter, r *http.Request) {

	result, err := sendPostRequest(w, r)
	if err != nil {
		http.Error(w, "Error sending request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := getOAuthToken()
	if err != nil {
		http.Error(w, "Error getting OAuth token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OAuth Token:", token)
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var namePtr *string // Declaring a pointer to string
	fmt.Println("Results1 ", namePtr)
	name := "John Doe"
	fmt.Println("Results2 ", &name)
	namePtr = &name // Assigning the address of name to namePtr
	fmt.Println("Results3 ", *namePtr)

	// name1 := "John Doe"
	// namePtr := &name1
	// *namePtr = "Jane Smith" // Update the string value indirectly through namePtr
	// fmt.Println("Results ", namePtr)
	// fmt.Println("Results ", &name1)
	//p := Person{}
	// p.Age = 10
	// fmt.Println(p.Age)
	// //	fmt.Print("Rogers")
	// c1 := &Controller{}
	// c2 := Controller{}

	// // 	// Using the pointer receiver method
	// c1.setIdPointer(1)
	// c1.setNamePointer("Changed with Pointer")
	// //fmt.Println(c1.name) // Output: Changed with Pointer
	// fmt.Println("..ID...", c1.id)
	// // //fmt.Print("end...", c1)

	// // 	// Using the value receiver method
	// c2.setIdPointer(3)
	// c2.setNameValue("Changed with Value")
	// fmt.Println(c2.name) // Output: Original
	// fmt.Println("id value", c2.id)

	// http.HandleFunc("/token", tokenHandler)
	// http.HandleFunc("/send", sendHandler)
	// fmt.Println("Server starting on port 8080...")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
}

// package main

// import (
// 	"fmt"
// )

// type MyStruct struct {
// 	field  int
// 	field2 string
// 	field3 bool
// }

// // Value receiver method
// func (s MyStruct) ValueReceiverMethod() {
// 	s.field = 10
// 	//s.field2 = "rogers"
// }

// // Pointer receiver method
// func (s *MyStruct) PointerReceiverMethod() {
// 	//s.field = 20
// }

// func main() {
// 	// Create an instance of MyStruct
// 	s := MyStruct{}
// 	s.field = 5
// 	s.field2 = "nene"
// 	s.field3 = true

// 	// Call value receiver method
// 	s.ValueReceiverMethod()
// 	fmt.Println("After ValueReceiverMethod:", s.field2) // Output: 5

// 	// Call pointer receiver method
// 	s.PointerReceiverMethod()
// 	fmt.Println("After PointerReceiverMethod:", s) // Output: 20
// }
