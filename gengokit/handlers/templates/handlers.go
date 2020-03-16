package templates

const HandlerMethods = `
{{ with $te := .}}
		{{range $i := .Methods}}
		// {{.Name}} implements Service.
		func (s {{ToLower $te.ServiceName}}Service) {{.Name}}(ctx context.Context, in *pb.{{GoName .RequestType.Name}}) (*pb.{{GoName .ResponseType.Name}}, error){
			var resp pb.{{GoName .ResponseType.Name}}
			resp = pb.{{GoName .ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
		{{end}}
{{- end}}
`

const Handlers = `
package handlers

import (
	"context"

	// 3d Party
	"github.com/go-kit/kit/log"
	stdzipkin "github.com/openzipkin/zipkin-go"

	// This Service
	pb "{{.PBImportPath -}}"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService(logger log.Logger, zipkinTracer *stdzipkin.Tracer) pb.{{GoName .Service.Name}}Server {
	return {{ToLower .Service.Name}}Service{
		logger:       logger,
		zipkinTracer: zipkinTracer,
	}
}

type {{ToLower .Service.Name}}Service struct{
	logger       log.Logger
	zipkinTracer *stdzipkin.Tracer
}

{{with $te := . }}
	{{range $i := $te.Service.Methods}}
		// {{$i.Name}} implements Service.
		func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*pb.{{GoName $i.ResponseType.Name}}, error){
			var resp pb.{{GoName $i.ResponseType.Name}}
			resp = pb.{{GoName $i.ResponseType.Name}}{
				{{range $j := $i.ResponseType.Message.Fields -}}
					// {{GoName $j.Name}}:
				{{end -}}
			}
			return &resp, nil
		}
	{{end}}
{{- end}}
`
