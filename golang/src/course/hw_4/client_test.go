package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io"
	"time"
	"encoding/xml"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type RawUsers struct {
	XMLName xml.Name `xml:"root"`
	Users   []RawUser   `xml:"row"`
}

type RawUser struct {
	Id 		int  `xml:"id"`
	Name    string   `xml:"first_name"`
	Age     int   `xml:"age"`
	About 	string `xml:"about"`
	Gender  string `xml:gender`
}

const testToken = "token"

type searchServer func(http.ResponseWriter, *http.Request)

func StatusBadRequestCantUnpackSearchServer(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `"strange{{}}json"`)
}

func StatusBadRequestUnknownErrorSearchServer(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `{"Error": "666"}`)
}

func StatusBadRequestErrorBadOrderFieldSearchServer(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `{"Error": "ErrorBadOrderField"}`)
}

func SearchServer(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("AccessToken")
	
	if (token != testToken){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	xmlFile, _ := os.Open("dataset.xml")
	defer xmlFile.Close()

	rawUsers := RawUsers{}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &rawUsers)

	limit, _ := strconv.Atoi(r.FormValue("limit"))
	query := r.FormValue("query")

	filtered := []RawUser{}
	rawUsersMarshalled := []byte{}

	if (query != ""){
		for _, user := range rawUsers.Users{
			if user.Name == query{
				filtered = append(filtered, user)
			}
		}
	}else{
		filtered = rawUsers.Users
	}

	if limit > len(filtered){
		rawUsersMarshalled, _ = json.Marshal(filtered)
	}else{
		rawUsersMarshalled, _ = json.Marshal(filtered[:limit])
	}
	io.WriteString(w, string(rawUsersMarshalled))
}



func CantMarshallSearchServer(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("AccessToken")
	
	if (token != testToken){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	xmlFile, _ := os.Open("dataset.xml")
	defer xmlFile.Close()

	rawUsers := RawUsers{}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &rawUsers)

	rawUsersMarshalled, _ := json.Marshal(rawUsers)
	io.WriteString(w, string(rawUsersMarshalled))
}

func TestFindUsersReqLimit(t *testing.T) {
	client := SearchClient{}
	req := SearchRequest{Limit: -10}
	_, err := client.FindUsers(req)
	if (err != nil) {
		if (err.Error() != "limit must be > 0"){
			t.Errorf("limit check failed")
		}
	}
}

func TestFindUsersOffsetLimit(t *testing.T){
	client := SearchClient{}
	req := SearchRequest{Offset: -10}
	_, err := client.FindUsers(req)
	if (err != nil) {
		if (err.Error() != "offset must be > 0"){
			t.Errorf("limit check failed")
		}
	}
}

func TestFindUsersBadAccessToken(t *testing.T){
	client := getClientForFunc(SearchServer)
	client.AccessToken = "bad_token"

	_, err := client.FindUsers(SearchRequest{})
	if (err != nil) {
		if (err.Error() != "Bad AccessToken"){
			t.Errorf("Access token check failed")
		}
	}	
}

func TestFindUsersInternalServerError(t *testing.T){
	server := func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusInternalServerError)}
	client := getClientForFunc(server) 

	_, err := client.FindUsers(SearchRequest{})
	if (err != nil) {
		if (err.Error() != "SearchServer fatal error"){
			t.Errorf("internal error is not raised")
		}
	}
}

func TestFindUsersStatusBadRequest(t *testing.T){
	client := getClientForFunc(StatusBadRequestCantUnpackSearchServer)
	_, err := client.FindUsers(SearchRequest{})
	if (err != nil) {
		if (err.Error() != "cant unpack error json: json: cannot unmarshal string into Go value of type main.SearchErrorResponse"){
			t.Errorf("bad request is not raised")
		}
	}
}

func TestFindUsersStatusBadRequestUnknownError(t *testing.T){
	client := getClientForFunc(StatusBadRequestUnknownErrorSearchServer)
	_, err := client.FindUsers(SearchRequest{})
	if (err != nil) {
		if (err.Error() != "unknown bad request error: 666"){
			t.Errorf("unknown bad request not raised")
		}
	}	
}

func TestFindUsersTimeOut(t *testing.T){
	server := func(w http.ResponseWriter, r *http.Request){
		time.Sleep(2 * time.Second)}
	client := getClientForFunc(server)

	_, err := client.FindUsers(SearchRequest{})
	if (err != nil) {
		if (err.Error() != "timeout for limit=1&offset=0&order_by=0&order_field=&query="){
			t.Errorf("no timeout check")
		}
	} 
}

func TestFindUsersBadOrderField(t *testing.T){
	client := getClientForFunc(StatusBadRequestErrorBadOrderFieldSearchServer)
	_, err := client.FindUsers(SearchRequest{OrderField: "invalid"})
	if (err != nil) {
		if (err.Error() != "OrderFeld invalid invalid"){
			t.Errorf("no bad order field check")
		}
	} 	
}

func TestFindUsersUnknownSendError(t *testing.T){
	client := getClientForFunc(SearchServer)
	client.URL += "zzzz"
	_, err := client.FindUsers(SearchRequest{OrderField: "invalid"})
	if (err != nil) {
	} 	
}

func getTestServer(serverFunc searchServer) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(serverFunc))
}

func getClientForFunc(serverFunc searchServer) SearchClient{
	return SearchClient{AccessToken: testToken, URL: getTestServer(serverFunc).URL}
}



func TestFindUsers(t *testing.T){
	client := getClientForFunc(SearchServer)
	client.FindUsers(SearchRequest{OrderField: "invalid", Limit: 50})
}

func TestFindUsersCantMarshall(t *testing.T){
	client := getClientForFunc(CantMarshallSearchServer)
	_, err := client.FindUsers(SearchRequest{Limit: 50})
	if (err != nil){
		if (err.Error() != "cant unpack result json: json: cannot unmarshal object into Go value of type []main.User"){
			t.Errorf("no unmarshal test")
		}
	}
}

func TestFindUsersMatchLimit(t *testing.T){
	client := getClientForFunc(SearchServer)
	client.FindUsers(SearchRequest{Limit: 0})
}

func TestFindUsersMatch(t *testing.T){
	client := getClientForFunc(SearchServer)
	client.FindUsers(SearchRequest{Limit: 10, Query: "Boyd"})
}