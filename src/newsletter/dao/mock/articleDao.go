package mock

import (
	"pulzo/src/newsletter/view/dto"
	"strconv"
	"strings"
)

type ArticleDao struct {
	arrayArticle []dto.ArticleDto
}

var urls []string = []string{
	"/entretenimiento/masterchef-fuerte-pelea-entre-martha-isabel-bolanos-zulma-rey-PP2892262",
	"/economia/claro-movistar-recibieron-orden-gobierno-nacional-beneficiar-usuarios-PP2892094A",
	"/nacion/atentado-estacion-policia-bucaramanga-hoy-asi-estan-6-heridos-PP2892035A",
	"/deportes/luis-fernando-muriel-estara-despedida-sebastian-viera-metropolitano-PP2891869",
	"/entretenimiento/locutora-deysa-rayo-expuso-eps-por-grave-problema-pidio-ayuda-con-urgencia-PP2891871",
	"/deportes/independiente-santa-fe-confirmo-fichaje-volante-enrique-serje-junior-PP2892244A",
	"/vivir-bien/planta-ayudara-tener-tu-casa-fria-necesitar-aire-acondiconado-PP2892030",
	"/economia/cual-es-inversion-que-mas-hacen-colombianos-con-bancos-que-se-trata-PP2891956A",
	"/economia/pension-colombia-dane-dice-que-90-migrantes-no-cotiza-colpensiones-PP2891786",
	"/deportes/tour-francia-2023-etapa-5-vingegaard-luce-jai-hindley-nuevo-lider-PP2890720",
}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		arrayArticle: initializeArticleData(),
	}
}

func (a *ArticleDao) FindArticleById(id string) dto.ArticleDto {
	for _, item := range a.arrayArticle {
		if item.Id == id {
			return item
		}
	}
	return dto.ArticleDto{}
}

func (a *ArticleDao) ArticlesList() []dto.ArticleDto {
	return a.arrayArticle
}

func initializeArticleData() (arrayArticle []dto.ArticleDto) {
	for index, url := range urls {
		strIndex := strconv.Itoa(index)

		newArticle := dto.ArticleDto{
			Id:      getId(url),
			Section: getSection(url),
			Title:   "Title " + strIndex,
			Lead:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			Image:   "https://d2yoo3qu6vrk5d.cloudfront.net/pulzo-lite/images-resized/PP2890720-s-o.webp",
			Link:    url,
		}
		arrayArticle = append(arrayArticle, newArticle)
	}
	return arrayArticle
}

func getId(url string) string {
	result := strings.Split(url, "-PP")
	if len(result) > 1 {
		return "PP" + result[1]
	}
	return ""
}

func getSection(url string) string {
	result := strings.Split(url, "/")
	if len(result) > 1 {
		return result[1]
	}
	return ""
}
