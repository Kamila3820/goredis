import http from 'k6/http'

export let options = {
    vus: 1,
    duration: '5s',

}

export default function() {
    http.get('http://host.docker.internal:8000/hello')
}