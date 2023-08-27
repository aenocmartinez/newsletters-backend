package domain_test

import (
	"pulzo/src/newsletter/dao/mock"
	"pulzo/src/newsletter/domain"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestListadoNewsletter(t *testing.T) {

	connectionDatabase := mock.NewNewsletterDao()
	list := connectionDatabase.NewsletterList()
	totalRegister := len(list)
	expected := 3

	assert.Equal(t, totalRegister, expected)
}

func TestCrearNewsletter(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()

	newsletter := domain.NewNewsletter("Nación")
	newsletter.SetSection("nacion")

	var expected error = nil

	err := connectionDatabase.CreateNewsletter(*newsletter)
	assert.Equal(t, err, expected)
}

func TestCrearNewsletterYaExistente(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()

	newsletter := domain.NewNewsletter("Entretenimiento")
	var expected error = nil

	err := connectionDatabase.CreateNewsletter(*newsletter)
	assert.NotEqual(t, err, expected)
}

func TestBuscarNewsletterPorId(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	newsletter := connectionDatabase.FindNewsletterById(2)
	expected := true

	assert.Equal(t, newsletter.Exists(), expected)
}

func TestBuscarNewsletterQueNoExistaPorId(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	newsletter := connectionDatabase.FindNewsletterById(5)
	expected := false

	assert.Equal(t, newsletter.Exists(), expected)
}

func TestBuscarNewsletterPorNombre(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	newsletter := connectionDatabase.FindNewsletterByName("DEPORTES")
	expected := true

	assert.Equal(t, newsletter.Exists(), expected)
}

func TestBuscarNewsletterQueNoExistaPorNombre(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	newsletter := connectionDatabase.FindNewsletterByName("Ultima hora")
	expected := false

	assert.Equal(t, newsletter.Exists(), expected)
}

func TestEliminarUnNewsletter(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	newsletter := domain.Newsletter{}
	newsletter.SetId(3)
	var expected error = nil

	err := connectionDatabase.DeleteNewsletter(newsletter)

	assert.Equal(t, err, expected)
}

func TestEliminarUnNewsletterNoExistente(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	var expected error = nil

	newsletter := domain.Newsletter{}
	newsletter.SetId(5)

	err := connectionDatabase.DeleteNewsletter(newsletter)

	assert.NotEqual(t, err, expected)
}

func TestActualizarNewsletter(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	var expected error = nil

	newsletter := connectionDatabase.FindNewsletterById(2)
	newsletter.SetName("NUEVO NOMBRE")
	err := connectionDatabase.UpdateNewsletter(newsletter)
	assert.Equal(t, err, expected)
}

func TestActualizarNewsletterConUnNombreYaExistente(t *testing.T) {
	connectionDatabase := mock.NewNewsletterDao()
	var expected error = nil

	newsletter := connectionDatabase.FindNewsletterById(2)
	newsletter.SetName(strings.ToUpper("entretenimiento"))
	err := connectionDatabase.UpdateNewsletter(newsletter)
	assert.NotEqual(t, err, expected)
}
