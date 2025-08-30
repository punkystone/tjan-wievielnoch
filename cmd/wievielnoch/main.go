package main

import "wievielnoch/internal"

func main() {
	env, err := internal.CheckEnv()
	if err != nil {
		panic(err)
	}
	bulletService := internal.NewBulletService(env.Session)
	internal.StartServer(bulletService)
}
