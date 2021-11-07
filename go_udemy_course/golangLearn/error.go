package main
import	(
	"fmt"

//"io/ioutil"
"log"
"errors"

)
	var ErrNorgateMath = errors.New("norgate math: square root of")

func main(){
	fmt.Println("%Tn", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f< 0{
		return 0, ErrNorgateMath
	}
	return 42, nil
}

//error fmt.Println()
//_, err := os.Open("no-file.txt")
	// if err != nil {
	// 	fmt.Println("err  happened", err)
	// }


// func OpenFile(){
// 	f, err := os.Open("names.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return 
// 	}

// 	defer f.Close()

// 	bs, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bs))
// }

// func WriteFile(){

// 	f, err := os.Create("names.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	r := strings.NewReader("Wassup")

// 	io.Copy(f, r)
// }