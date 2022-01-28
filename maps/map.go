package maps

type Dictionary map[string]string

func (d Dictionary) Lookup(s string) string {
    return d[s]
}

func (d Dictionary) Add (key string, value string) {
    d[key] = value
}