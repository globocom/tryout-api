package api

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	greetings := `this is the tryout api!
to create a new challenge use POST /challenge with {
	'name': string,
	'endpoints': [
		{
			'path': '/example',
			'input': 'json string',
			'output': 'string output',
			'http_status': http status,
			'http_method': 'get',
			'throughput': requests/sec,
		}
	]
}.
`
	w.Write([]byte(greetings))
}
