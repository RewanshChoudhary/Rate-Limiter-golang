package algorithms
 ctx:=context.Background()

func main(){

	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	

}