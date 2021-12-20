package main

//if the declaration is in capital letters then it is accessible to all the packages
// var I int = 27

func main() {
	// fmt.Println(I)
	// var n bool
	// m := 2 == 1
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T", m, m)
	// a := 10 		//1010
	// b := 3			//0011
	// fmt.Println(a & b)	//0010
	// fmt.Println(a | b)	//1011
	// fmt.Println(a ^ b)	//1001 (reverse of & if both are same the not counted)
	// fmt.Println(a &^ b)	//0100 (only those are count in which both are zero)
	// 	a := 10
	// 	fmt.Println(a << 2)
	// 	fmt.Println(a >> 2)
	// var n complex128 = complex(5, 6)
	// fmt.Printf("%v, %T", n, n)

	// s := "this is a string"
	// b := []byte(s)
	// fmt.Printf("%v, %T\n", b, b)
	// const (
	// 	_  = iota
	// 	KB = 1 << (10 * iota)
	// 	MB
	// 	GB
	// 	TB
	// 	PB
	// 	ZB
	// 	YB
	// )
	// fileSize := 4000000000.
	// fmt.Println("KB: ", KB)
	// fmt.Println("MB: ", MB)
	// fmt.Println("GB: ", GB)
	// fmt.Println("TB: ", TB, "\n")

	// fmt.Println("%0.4fTB", fileSize/TB)

	// var a = [...]int{1, 2, 3, 4, 5}
	// b := append(a[:1], a[2:]...)
	// fmt.Println(b)
	// fmt.Println(a)
	// fmt.Printf("Cap: %v, Len: %v", cap(b), len(b))

	// c := b[1:]
	// d := b[:len(b)-1]
	// fmt.Println(c)
	// fmt.Println(d)
	// statePopulations := map[string]int{
	// 	"California":   1232132131,
	// 	"Texas":        23242432,
	// 	"Florida":      231312313,
	// 	"New York":     31323432,
	// 	"Pennsylvania": 34354645,
	// }
	// fmt.Println(statePopulations)

	//STRUCTS

	// type Doctor struct {
	// 	number     int
	// 	actorName  string
	// 	companions []string
	// }
	// aDoctor := struct{ name string }{name: "Karan"}
	// anotherDoctor := &aDoctor
	// anotherDoctor.name = "Tom"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)

	//EMBEDDING IN STRUCT
	// type Animal struct {
	// 	origin string
	// 	name   string
	// }
	// type Bird struct {
	// 	Animal
	// 	speed  float32
	// 	canFly bool
	// }

	// b := Bird{speed: 34.4,
	// 	canFly: true,
	// 	Animal: Animal{name: "Emu",
	// 		origin: "Australia"}}
	// fmt.Println(b)

	//IF-ELSE

	// statePopulations := map[string]int{
	// 	"California":   1232132131,
	// 	"Texas":        23242432,
	// 	"Florida":      231312313,
	// 	"New York":     31323432,
	// 	"Pennsylvania": 34354645,
	// }
	// if pop, ok := statePopulations["Florida"]; ok {
	// 	fmt.Println(pop)
	// }
	// number := 50
	// guess := 1
	// if guess < 1 || guess > 100 {
	// 	fmt.Println("Guess must be between 1 and 100!")
	// }

	// if guess < number {
	// 	fmt.Println("Too low")
	// }
	// if guess > number {
	// 	fmt.Println("Too high")
	// }
	// if guess == number {
	// 	fmt.Println("You got it!!")
	// }
	// myNum := 0.123
	// if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
	// 	fmt.Println(math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2) - 1))
	// 	fmt.Println("These are same")
	// } else {
	// 	fmt.Println("These are not same")
	// }

	// SWITCH STATEMENTS

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(10))
	// switch rand.Intn(10) {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("one, three, five, seven, nine")
	// case 2, 4, 6, 8:
	// 	fmt.Println("two, four, six, eight")
	// default:
	// 	fmt.Println("10")
	// }
	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less than or equal to ten")
	// case i <= 20:
	// 	fmt.Println("less than or equal to twenty")
	// default:
	// 	fmt.Println("greater than twenty")
	// }
	//3:12:43
}
