package main

import (
	"bytes"
	"io/ioutil"
	"testing"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	
	//easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

const filePath string = "./data/users.txt"

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Browsers []string `json:"browsers"`
}

func Decode(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func Encode(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	Encode(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	Encode(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	Decode(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	Decode(l, v)
}


// func FastSlowSearch(out io.Writer) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	seenBrowsers := map[string]bool{}

// 	totalUsers := -1

// 	// scanner := bufio.NewScanner(file)
//     fmt.Fprintln(out, "found users:")
    
//     // users := [1000]User{}
//     // currentUser := 0
	
// 	dec := json.NewDecoder(file)
// 	otherState := true
// 	emailState := false
// 	nameState := false
// 	browsersState := false

// 	name := ""
// 	email := ""
// 	isAndroid := false
// 	isMSIE := false

// 	for {
// 		t, err := dec.Token()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
		
// 		switch t.(type){
		
// 		case json.Delim:
// 			value := t.(json.Delim)
// 			if (value == '}'){
// 				if (isAndroid && isMSIE){
// 					fmt.Fprintf(out ,"[%d] %s <%s>\n", totalUsers, name, strings.Replace(email, "@", " [at] ", 1))
// 				}
// 				isAndroid = false
// 				isMSIE = false
// 			}else if (value == '{'){
// 				totalUsers += 1
// 				otherState = true
// 				emailState = false
// 				nameState = false
// 				browsersState = false
// 			}else if (value == ']'){
// 				otherState = true
// 				otherState = true
// 				emailState = false
// 				nameState = false
// 				browsersState = false
// 			}
// 			continue
		
// 		case string:			
// 			value := t.(string)
// 			if emailState{
// 				email = value
// 				emailState = false
// 				otherState = true
// 			}else if nameState{
// 				name = value
// 				nameState = false
// 				otherState = true
// 			}else if browsersState{
// 				if strings.Contains(value, "Android") {
// 					isAndroid = true
// 					if !(seenBrowsers[value]){
// 						seenBrowsers[value] = true
// 					}
// 				}
// 				if strings.Contains(value, "MSIE"){
// 					isMSIE = true
// 					if !(seenBrowsers[value]){
// 						seenBrowsers[value] = true
// 					}
// 				}

// 			}else if otherState{
// 				switch value{
// 				case "name":
// 					nameState = true
// 					otherState = false
// 					continue
// 				case "browsers":
// 					browsersState = true
// 					otherState = false
// 					continue
// 				case "email":
// 					emailState = true
// 					otherState = false
// 					continue
// 				}
// 			}
// 		}
// 	}

// 	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
// }

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	seenBrowsers := map[string]bool{}
	

	totalUsers := -1

	scanner := bufio.NewScanner(file)
    fmt.Fprintln(out, "found users:")
    
    users := [1000]User{}
    currentUser := 0

    for scanner.Scan() {

		user := users[currentUser]
		currentUser += 1
		user.UnmarshalJSON(scanner.Bytes())
		
		isAndroid := false
		isMSIE := false

		totalUsers += 1

		for _, browser := range user.Browsers {
			
			if strings.Contains(browser, "Android") {
				isAndroid = true
				if !(seenBrowsers[browser]){
					seenBrowsers[browser] = true
				}
			}

			if strings.Contains(browser, "MSIE"){
				isMSIE = true
				if !(seenBrowsers[browser]){
					seenBrowsers[browser] = true
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		fmt.Fprintf(out ,"[%d] %s <%s>\n", totalUsers, user.Name, strings.Replace(user.Email, "@", " [at] ", 1))

	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}

// ======================================================

func SlowSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("@")
	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""

	lines := strings.Split(string(fileContents), "\n")

	users := make([]map[string]interface{}, 0)
	for _, line := range lines {
		user := make(map[string]interface{})
		// fmt.Printf("%v %v\n", err, line)
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers, ok := user["browsers"].([]interface{})
		if !ok {
			// log.Println("cant cast browsers")
			continue
		}

		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			if ok, err := regexp.MatchString("Android", browser); ok && err == nil {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			if ok, err := regexp.MatchString("MSIE", browser); ok && err == nil {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user["email"].(string), " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	SlowSearch(ioutil.Discard)
	FastSearch(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	SlowSearch(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch(ioutil.Discard)
	}
}
