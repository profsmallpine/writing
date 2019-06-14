package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// handler is http struct for passing services to the router.
type handler struct {
	DB     *gorm.DB
	Logger *log.Logger
}

// goHome is used for handling requests to "/".
func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/index.tmpl", nil)
}

// zenOfWritingGoodCode is used for handling requests to "/writing/zen-of-writing-good-code".
func (h *handler) zenOfWritingGoodCode(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/writing/zen_of_writing_good_code.tmpl", nil)
}

// lessonsFromFailedDocuSignIntegration is used for handling requests to "/writing/lessons-from-api-design".
func (h *handler) lessonsFromFailedDocuSignIntegration(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/writing/lessons_from_failed_docusign_integration.tmpl", nil)
}

// majorStripeFunctionalityDowngrade is used for handling requests to "/writing/major-stripe-functionality-downgrade".
func (h *handler) majorStripeFunctionalityDowngrade(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/writing/major_stripe_functionality_downgrade.tmpl", nil)
}

// devicesStrategy is used for handling requests to "/writing/devices-strategy".
func (h *handler) devicesStrategy(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/writing/devices_strategy.tmpl", nil)
}

// respond is used to parse a base template.
func respond(logger *log.Logger, w http.ResponseWriter, r *http.Request, layout string, data interface{}) {
	// Parse static files.
	tmpl := template.Must(template.New("base.tmpl").ParseFiles(
		"./tmpl/base.tmpl",
		layout,
	))
	err := tmpl.Funcs(template.FuncMap{}).Execute(w, data)

	// Log template compilation failure.
	if err != nil {
		logger.Println("Template execution error: ", err.Error(), layout)
		return
	}
}
