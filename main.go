package main

import (
	"cmp"
	"fmt"
	_ "go-language/utils"
	"net/http"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}
type Content interface {
	GetContent() string
}

func Share(c Content) {
	fmt.Println(c.GetContent())
}

type Twitter struct {
	Username string
	Followers int
}

type PostBlog struct {
	Username string
	Followers int
}

func (t *Twitter) GetContent() string {
	return fmt.Sprintf("Twitter: @%s has %d followers", t.Username, t.Followers)
}


func (p *PostBlog) GetContent() string {
	return fmt.Sprintf("PostBlog: @%s has %d followers", p.Username, p.Followers)
}



func isGreater[T cmp.Ordered](a, b T) bool {
	return a > b
}

func addbytwo[T int|float64](num T) T {
	return num + 2
}


// score := []int{90, 80, 70, 60, 50}

// score = append(score, 40)

// func (t *Employee) String() info() {
// 	t.name = "John Doe"
// 	fmt.Println("Employee: %s, Age: %d, Salary: %.2f", t.Name, t.Age, t.Salary)
	
// }

func (t *Employee) updatename() {
	fmt.Println("Employee: %s, Age: %d, Salary: %.2f", t.Name, t.Age, t.Salary)
}

func main(){
	num := 100
	var num2 int = 200
	var char string = "Hello, Go!"
	fmt.Println("Hello, World!")
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(char)

	 var recordemap map[string]int
	 recordemap = make(map[string]int)
	 recordemap["Alice"] = 25
	 recordemap["Bob"] = 30
	 recordemap["Charlie"] = 35


	emp := Employee{Name: "John Doe", Age: 30, Salary: 50000.0}
	fmt.Println(emp)
	// u.Add(5, 10)

	score := []int{90, 80, 70, 60, 50}

	
    score = append(score, 40)

	score = score[0:2]


	score[0] = 100
    for _, s := range score {
		fmt.Println(s)
	}

   pb := Twitter{Username: "john_doe", Followers: 1000}
   blog := PostBlog{Username: "jane_doe", Followers: 500}
   Share(&pb)
   Share(&blog)


   if val, ok := recordemap["Alice"]; ok {
		fmt.Println("Alice's age is:", val)
	} else {
		fmt.Println("Alice not found in the map.")
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "https://example.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}else {
		fmt.Println("Request created successfully:", req)
		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Response status:", resp.Status)
	}

}