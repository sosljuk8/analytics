package {{.pkgName}}

import (
	{{.imports}}
)

func (cmd *Group) {{.function}}(_ context.Context {{ if .request }}, {{.request}}{{ end -}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
