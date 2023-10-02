package {{.PkgName}}

import (
"net/http"
"github.com/zeromicro/go-zero/rest/httpx"
{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
{{if .HasRequest}}var req types.{{.RequestType}}
if err := httpx.Parse(r, &req); err != nil {
httpx.ErrorCtx(r.Context(), w, err)
return
}

{{end}}

{{if .HasResp}}resp, {{end}}err := svcCtx.Deps.{.LazyGroupName}().{{.Call}}(r.Context(){{if .HasRequest}}, &req{{end}})
if err != nil {
httpx.ErrorCtx(r.Context(), w, err)
} else {
{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
}
}
}
