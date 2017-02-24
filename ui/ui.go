package ui

import (
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/higanworks/envmap"
	"github.com/qb0C80aE/clay/extensions"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

// BinaryFileSystem builds binaryFileSystem instance
func BinaryFileSystem(root string) static.ServeFileSystem {
	fs := &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
	return &binaryFileSystem{
		fs,
	}
}

func loadTemplates(filenames ...string) (*template.Template, error) {
	var t *template.Template
	for _, filename := range filenames {
		templateString, _ := Asset(filename)
		name := filepath.Base(filename)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err := tmpl.Parse(string(templateString))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

type routerInitializer struct {
}

func newRouterInitializer() *routerInitializer {
	return &routerInitializer{}
}

func versionInformation() map[interface{}]interface{} {
	programInformation := extensions.RegisteredProgramInformation()
	subModuleInformationList := programInformation.SubModuleInformationList()
	subModuleInformationMapList := make([]map[string]string, len(subModuleInformationList))
	for i, subModuleInformation := range subModuleInformationList {
		subModuleInformationMapList[i] = map[string]string{
			"Name":     subModuleInformation.Name(),
			"Revision": subModuleInformation.Revision(),
		}
	}
	result := map[interface{}]interface{}{
		"BuildTime":                programInformation.BuildTime(),
		"SubModuleInformationList": subModuleInformationMapList,
	}
	return result
}

func (routerInitializer *routerInitializer) InitializeEarly(r *gin.Engine) error {
	r.Use(static.Serve("/ui/files", BinaryFileSystem("ui/files")))

	templateBase := "ui/templates"
	templateFileNames := []string{
		fmt.Sprintf("%s/%s", templateBase, "design.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "diagram.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "dialog.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "footerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "functions.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "functions_node_port.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "headerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "index.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "navigation_bar.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "network.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "requirement.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "template.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "testscript.tmpl"),
	}
	htmlTemplate, err := loadTemplates(templateFileNames...)
	if err != nil {
		log.Fatal(err)
		return err
	}
	r.SetHTMLTemplate(htmlTemplate)

	envMap := envmap.All()
	ui := r.Group("/ui")
	{
		ui.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{"env": envMap, "version": versionInformation(), "category": "home"})
		})
		ui.GET("/network", func(c *gin.Context) {
			c.HTML(http.StatusOK, "network.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/design", func(c *gin.Context) {
			c.HTML(http.StatusOK, "design.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diagram.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/requirement", func(c *gin.Context) {
			c.HTML(http.StatusOK, "requirement.tmpl", gin.H{"env": envMap, "category": "design"})
		})
		ui.GET("/template", func(c *gin.Context) {
			c.HTML(http.StatusOK, "template.tmpl", gin.H{"env": envMap, "category": "process"})
		})
		ui.GET("/testscript", func(c *gin.Context) {
			c.HTML(http.StatusOK, "testscript.tmpl", gin.H{"env": envMap, "category": "process"})
		})
	}

	return nil
}

func (routerInitializer *routerInitializer) InitializeLate(r *gin.Engine) error {
	return nil
}

var uniqueRouterInitializer = newRouterInitializer()

func init() {
	extensions.RegisterRouterInitializer(uniqueRouterInitializer)
}
