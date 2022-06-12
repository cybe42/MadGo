package main

func main() {
	app := Static()

	StartAPI(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
