package main

type Alias struct {
	ID    int32
	Alias string
	Name  string
}

func ToAliasMap[F any](values []F, transform func(F) Alias) map[int32]Alias {
	var result = make(map[int32]Alias, len(values))

	for _, value := range values {
		var alias = transform(value)

		result[alias.ID] = alias
	}

	return result
}
