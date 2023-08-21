package dto

type FindArticleDto struct {
	Id      string     `json:"id"`
	Title   TitleDto   `json:"titles"`
	Lead    Phrases    `json:"phrases"`
	Link    Urls       `json:"urls"`
	Section Categories `json:"categories"`
	Images  ImagesT    `json;"images"`
}

type TitleDto struct {
	Main string `json:"main"`
}

type Phrases struct {
	Main string `json:"main"`
}

type Urls struct {
	Main string `json:"main"`
}

type Categories struct {
	Main Category `json:"main"`
}

type Category struct {
	Slug string `json:"slug"`
}

type ImagesT struct {
	Types ImageType `json:"types"`
}

type ImageType struct {
	Square     Square     `json:"square"`
	Horizontal Horizontal `json:"horizontal"`
}

type Square struct {
	Medium string `json:"medium"`
}

type Horizontal struct {
	Big string `json:"big"`
}
