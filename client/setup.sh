wget https://gist.githubusercontent.com/hayashiki/3ffc6d8658c68663d2ff7bb91d291f0d/raw/dcfc61ae6957497f203916b4c8fab2fcbfbac075/setup-01-must.sh | sh
wget https://gist.githubusercontent.com/hayashiki/3ffc6d8658c68663d2ff7bb91d291f0d/raw/dcfc61ae6957497f203916b4c8fab2fcbfbac075/setup-05-applo-client.sh | sh setup-05-applo-client.sh

yarn add graphql
# いらんかも
yarn add subscriptions-transport-ws
yarn add -D @apollo/client

yarn add tailwindcss
uarn add autoprefixer

# TODO lib/graphqlのWS更新
# アップデートしてもよいかも。failedしたときのオブザーバ対応とかが反映されてない
# cors origin

```
	acceptOrigins := []string{
		"http://localhost:3000",
		"https://hoge.com",
	}

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   acceptOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:            false,
	}).Handler)
```
