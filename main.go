package main

func main() {
	a := App{}
	a.Initialize("user", "pass", "dbname")
	a.Run(":8080")
}
