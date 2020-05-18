module gincmf

go 1.14

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gincmf/cmf v0.0.2
	github.com/jinzhu/gorm v1.9.12
	gopkg.in/oauth2.v3 v3.12.0
)

replace github.com/gincmf/cmf v0.0.2 => ../cmf
