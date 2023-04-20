package main

import (
    "net/http"
    "os"
    "path/filepath"
    "html/template"
)

func main() {
    fs := http.FileServer(http.Dir("/home/chicmic/Desktop/Chat-App-Go/File/uploads"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            var files []string
            err := filepath.Walk("/home/chicmic/Desktop/Chat-App-Go/File/uploads", func(path string, info os.FileInfo, err error) error {
                if err != nil {
                    return err
                }
                if !info.IsDir() {
                    files = append(files, path)
                }
                return nil
            })
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            t, err := template.New("").Parse(`
            <html>
            <head>
                <title>File Downloads</title>
            </head>
            <body>
                <ul>
                    {{range .}}
                    <li><a href="/download?path={{.}}">{{.}}</a></li>
                    {{end}}
                </ul>
            </body>
            </html>
            `)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            if err := t.Execute(w, files); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }
        if r.URL.Path == "/download" {
            path := r.URL.Query().Get("path")
            http.ServeFile(w, r, path)
            return
        }
        http.NotFound(w, r)
    })
    http.ListenAndServe(":8080", nil)
}
