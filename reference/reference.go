/* package main */
/* every declare var need to be used, case if not, the program don't compile*/

var var1 string 

var (
	var2 string
	var3 int
	Var4 int /* Vars with first letter in uppercase can be access for everyone*/
	var5 float32
)

var a, b, c int 32 /*if you want to declare multiple vars as the same type*/

var a1, b1, c1 = "a1", "b1", "c1" /* on this way you can put a attribute in multiple vars*/

var x, y = 1, 2

x, y = y, x /* this is how switch the values of 3 or more variables*/

message := "is like this you infer the var type"

func hello (nome string){
	fmt.println("hello", nome, "space") /* we need import fmt to print data*/
}

func sum(num1 int, num2 int) int { /*! the int after () is the type of return*/
	return num1 + num2  		   /*(num1 int, num2 int) int > (num1,num2 int) int !reduce way */
}

func convertAndSum(num3 int, str string) int {
	i, err := strconv.Atoi(str) /*to omit the error is  i, _ := strconv.Atoi(str)*/
	return num3 + i
}

/* 
	this functions are just available at the package to make
	then accessible from other packages, the first letter need
	we in upper case
	func ConvertAndSum
*/

