module gincmf

go 1.14

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/gincmf/cmf v0.0.2
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	gopkg.in/oauth2.v3 v3.12.0
)

replace github.com/gincmf/cmf v0.0.2 => ../cmf
