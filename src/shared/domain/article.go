package domain

import (
	"errors"
	"pulzo/src/shared/infraestructure/util"
)

type Article struct {
	Id                      string     `json:"id"`
	Titles                  Titles     `json:"titles"`
	Phrases                 Phrases    `json:"phrases"`
	Categories              Section    `json:"categories"`
	Urls                    Url        `json:"urls"`
	Author                  Author     `json:"author"`
	Published               string     `json:"published"`
	Type                    string     `json:"type"`
	Censorship              string     `json:"censorship"`
	UnpublishImage          string     `json:"unpublishImage"`
	Ext                     string     `json:"ext"`
	Metadescripcion         string     `json:"metadescripcion"`
	Created                 int32      `json:"created"`
	Updated                 int32      `json:"updated"`
	Images                  Images     `json:"images"`
	Videos                  Video      `json:"videos"`
	ExternalUrl             string     `json:"externalUrl"`
	Carousel                string     `json:"carousel"`
	TrinoTW                 string     `json:"trinoTW"`
	LiveBlog                string     `json:"liveBlog"`
	Shorthand               string     `json:"shorthand"`
	Newsroom                string     `json:"newsroom_url"`
	Estilo                  int        `json:"estilo"`
	Alianza                 int        `json:"alianza"`
	Audio                   string     `json:"audio"`
	Tags                    []string   `json:"tags"`
	Sources                 []Sources  `json:"sources"`
	Related                 []string   `json:"related"`
	Body2                   string     `json:"body2"`
	Comercial               Commercial `json:"comercial"`
	Directtv                Directv    `json:"directtv"`
	Custom_fields           CustomFields
	Tambientepuedeinteresar []TPIStruct   `json:"tambientepuedeinteresar"`
	Amp                     Amp           `json:"amp"`
	Gallery                 []GalleryItem `json:"gallery"`
	Gallery2                Gallery2      `json:"gallery2"`
	Embeded                 []string      `json:"embed"`
	Autores                 string        `json:"autores"`
	Duration                string        `json:"duration"`
	Receta                  Receta        `json:"recetas"`
	Coordinates             Coordinates   `json:"coordinates"`
	Hermes                  interface{}   `json:"hermes"`
	IsContext               interface{}   `json:"isContext"`
	Score                   int           `json:"score"`
	Icon                    string        `json:"icon"`
	InitialScore            int           `json:"initialScore"`
	RelatedSeo              []string      `json:"relatedSeo"`
	ArticlesSeo             []RelatedSeo  `json:"articlesSeo"`
	articleRepository       ArticleRepository
}

type Url struct {
	Main string `json:"main"`
}

type Titles struct {
	Main     string `json:"main"`
	Facebook string `json:"facebook"`
	Seo      string `json:"seo"`
	Pulzo    string `json:"pulzo"`
}

type Phrases struct {
	Main string `json:"main"`
}

type Section struct {
	Main DescSection `json:"main"`
}

type DescSection struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Images struct {
	Meta  Meta             `json:"meta"`
	Types CropAndResizeImg `json:"types"`
}

type Meta struct {
	Title       string `json:"title"`
	Credit      string `json:"credit"`
	Description string `json:"description"`
	Alt         string `json:"alt"`
	File        string `json:"file"`
	Location    string `json:"location"`
}

type Types struct {
	Thumb    string `json:"thumb"`
	Mobile   string `json:"mobile"`
	Medio    string `json:"medio"`
	Large    string `json:"large"`
	Vertical string `json:"vertical"`
	Opening  string `json:"opening"`
	Gif      string `json:"gif"`
}

type NewTypes struct {
	Horizontal NewSizes `json:"horizontal"`
	Vertical   NewSizes `json:"vertical"`
	Square     NewSizes `json:"square"`
}

type NewSizes struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Big      string `json:"big"`
	Original string `json:"original"`
}

type Creator struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

type Owner struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Profile      string   `json:"profile"`
	Url          string   `json:"url"`
	Twitter      TwitterS `json:"twitter"`
	UrlFacebook  string   `json:"urlFacebook"`
	Email        string   `json:"email"`
	LinkPersonal string   `json:"linkPersonal"`
	TeamPulzo    string   `json:"teamPulzo"`
	Short_Text   string   `json:"short_text"`
	Large_Text   string   `json:"large_text"`
}

type TwitterS struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Author struct {
	Creator Creator `json:"creator"`
	Owner   Owner   `json:"owner"`
}

type Sources struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Source string `json:"source"`
}

type MainVideo struct {
	Image   string `json:"image"`
	Url     string `json:"url"`
	Title   string `json:"title"`
	Embeded string `json:"embeded"`
}

type Video struct {
	Main MainVideo `json:"main"`
}

type Promoted struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Color       string `json:"color"`
	Facebook    string `json:"facebook"`
	Image       string `json:"image"`
	Imageletter string `json:"imageletter"`
	Imageblack  string `json:"imageblack"`
}

type Commercial struct {
	Promovido Promoted `json:"promovido"`
}

type Directv struct {
	Mobile  string `json:"mobile"`
	Desktop string `json:"desktop"`
	Link    string `json:"link"`
}

type CustomFields struct {
	Directv   string
	Promovido string
}

type TPIStruct struct {
	Image_Desktop string      `json:"image_desktop"`
	Image_Mobile  string      `json:"image_mobile"`
	Image_Msn     string      `json:"image_msn"`
	Type          string      `json:"type"`
	Link          string      `json:"link"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	Author_Img    string      `json:"author_img"`
	Alt           string      `json:"alt"`
	Section       DescSection `json:"section"`
}
type Amp struct {
	Body        string      `json:"body"`
	DesktopBody string      `json:"desktopBody"`
	Libraries   interface{} `json:"libraries"`
}

type GalleryItem struct {
	File    string `json:"file"`
	ALtText string `json:"alt_text"`
	Title   string `json:"title"`
	Credits string `json:"credits"`
}

type Receta struct {
	Banners     BannersReceta `json:"banners"`
	Preparacion string        `json:"preparacion"`
	Patrocinado string        `json:"patrocinado"`
	Logo        string        `json:"logo"`
}

type BannersReceta struct {
	Top    BannerImg `json:"top"`
	Bottom BannerImg `json:"bottom"`
}

type BannerImg struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
}

type BodyArticleJson struct {
	FechaMinuto  string `json:"fechaMinuto"`
	TituloMinuto string `json:"tituloMinuto"`
}

type Coordinates struct {
	Horizontal Ratio  `json:"horizontal"`
	Vertical   Ratio  `json:"vertical"`
	Square     Ratio  `json:"square"`
	PostId     string `json:"post_id"`
	Url        string `json:"url"`
}

type Ratio struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type ContentDetails struct {
	Duration string `json:"duration"`
}

type Item struct {
	ContentDetails ContentDetails `json:"contentDetails"`
}

type Youtube struct {
	Items []Item `json:"items"`
}

type Gallery2 struct {
	Id        int64          `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	CreatedAt string         `json:"createdAt,omitempty"`
	UpdatedAt string         `json:"updatedAt,omitempty"`
	Images    []GalleryItem2 `json:"images,omitempty"`
}

type GalleryItem2 struct {
	// Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Alt         string `json:"alt,omitempty"`
	Description string `json:"description,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	PathDesktop string `json:"pathDesktop,omitempty"`
	PathMobile  string `json:"pathMobile,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type JsonFileGalleryLatest struct {
	Id        string           `json:"id,omitempty"`
	Title     string           `json:"title,omitempty"`
	Section   string           `json:"section,omitempty"`
	Slug      string           `json:"slug,omitempty"`
	Alt       string           `json:"alt,omitempty"`
	GalleryId int64            `json:"galleryId,omitempty"`
	Url       string           `json:"url,omitempty"`
	Images    CropAndResizeImg `json:"images"`
}

type RelatedSeo struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type CropAndResizeImg struct {
	CropUrls   CropUrls `json:"cropUrls"`
	Horizontal Sizes    `json:"horizontal"`
	Vertical   Sizes    `json:"vertical"`
	Square     Sizes    `json:"square"`
}

type CropUrls struct {
	Horizontal string `json:"horizontal"`
	Vertical   string `json:"vertical"`
	Square     string `json:"square"`
}

type ResizeUrls struct {
	Horizontal Sizes `json:"horizontal"`
	Vertical   Sizes `json:"vertical"`
	Square     Sizes `json:"square"`
}

type Sizes struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Big      string `json:"big"`
	Original string `json:"original"`
}

func (a *Article) SetRepository(articleRepository ArticleRepository) {
	a.articleRepository = articleRepository
}

func (a *Article) Exists() bool {
	return a.Id != ""
}

func FindArticleByURL(url string, articleRepository ArticleRepository) (article Article, err error) {
	idArticle := util.ExtractIdArticleFromURL(url)
	if len(idArticle) == 0 {
		return article, errors.New("no se puede encontrar el id del artículo en la url")
	}

	article = articleRepository.FindArticleById(idArticle)

	return article, err
}
