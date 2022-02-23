package gen

//go:generate go run github.com/99designs/gqlgen generate
//go:generate sqlboiler -c ../sqlboiler.toml -o ../models mysql
