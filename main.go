// This is a simple example of usage of Grafana client
// for importing dashboards from a bunch of JSON files (current dir used).
// You are can export dashboards with backup-dashboards utitity.
// NOTE: old dashboards with same names will be silently overrided!
//
// Usage:
//   import-dashboards http://sdk host:3000 api-key-string-here
//
// You need get API key with Admin rights from your Grafana!
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/grafana-tools/sdk"
)

/*

curl -H "Authorization: Bearer eyJrIjoiNWRXS201ZlZKcXE2b0lpQ2liWW9sa3RIMDVkYlMxWVYiLCJuIjoiZ3JhZmFuYS1hcGktZGVtbyIsImlkIjoxfQ==" http://localhost:3000/api/dashboards/home

*/
func main() {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
	)

	c := sdk.NewClient("http://localhost:3000", "eyJrIjoiNWRXS201ZlZKcXE2b0lpQ2liWW9sa3RIMDVkYlMxWVYiLCJuIjoiZ3JhZmFuYS1hcGktZGVtbyIsImlkIjoxfQ==", sdk.DefaultHTTPClient)
	dir := "dashboard-examples"
	filesInDir, err = ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			if rawBoard, err = ioutil.ReadFile(filepath.Join(dir, file.Name())); err != nil {
				log.Println(err)
				continue
			}
			var board sdk.Board
			if err = json.Unmarshal(rawBoard, &board); err != nil {
				log.Println(err)
				continue
			}
			c.DeleteDashboard(board.UpdateSlug())
			_, err := c.SetDashboard(board, false)
			if err != nil {
				log.Printf("error on importing dashboard %s", board.Title)
				continue
			}
		}
	}
}