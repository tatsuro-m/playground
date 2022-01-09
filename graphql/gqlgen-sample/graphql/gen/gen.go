package gen

//go:generate go run github.com/99designs/gqlgen
//go:generate sqlboiler -c ../sqlboiler.toml -o ../models mysql
