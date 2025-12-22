package main

import (
	"html/template"
	"log"
	"net/http"

	"intersection-go/intersection"
)

var page = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="id">
<head>
	<meta charset="UTF-8">
	<title>Intersection Go App</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background: linear-gradient(135deg, #667eea, #764ba2);
			color: #333;
			min-height: 100vh;
			display: flex;
			align-items: center;
			justify-content: center;
		}
		.card {
			background: #fff;
			padding: 30px;
			border-radius: 12px;
			width: 420px;
			box-shadow: 0 10px 30px rgba(0,0,0,0.2);
		}
		h1 {
			text-align: center;
			color: #4a4a4a;
		}
		.status {
			text-align: center;
			color: #2ecc71;
			font-weight: bold;
			margin-bottom: 20px;
		}
		.box {
			background: #f4f6f8;
			padding: 12px;
			border-radius: 8px;
			margin-bottom: 10px;
			font-family: monospace;
		}
		.result {
			background: #e8f8f5;
			color: #117a65;
			padding: 15px;
			border-radius: 8px;
			font-weight: bold;
			text-align: center;
		}
		.footer {
			margin-top: 20px;
			text-align: center;
			font-size: 12px;
			color: #888;
		}
	</style>
</head>
<body>
	<div class="card">
		<h1>Intersection Go</h1>
		<div class="status">✅ Deploy Success</div>

		<div class="box">Set A: {{.SetA}}</div>
		<div class="box">Set B: {{.SetB}}</div>

		<div class="result">
			Hasil Intersection:<br>
			{{.Result}}
		</div>

		<div class="footer">
			Test Jenkins Pipeline<br>
			Deployed on Render.com
		</div>
	</div>
</body>
</html>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setA := "1,3,4,7,13"
		setB := "1,2,4,13,15"

		result := intersection.FindIntersection(
			[]string{setA, setB},
		)

		data := map[string]string{
			"SetA":   setA,
			"SetB":   setB,
			"Result": result,
		}

		page.Execute(w, data)
	})

	log.Println("Server running on :3232")
	log.Fatal(http.ListenAndServe(":3232", nil))
}
