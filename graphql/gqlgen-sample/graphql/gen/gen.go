package gen

//go:generate sqlboiler -c ../sqlboiler.toml -o ../models mysql
//go:generate go run github.com/99designs/gqlgen
