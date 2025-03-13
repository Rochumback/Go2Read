package structs

import "github.com/lib/pq"

type Manga struct {
    Id                 int
    Title              string
    Alternative_titles pq.StringArray
    Searchable         string
    Synopsis           string
    Created_at         string
    Updated_at         string
    Created_by         int
}
