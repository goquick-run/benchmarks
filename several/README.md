# teststress
Testes stress no quick e comparando com alguns frameworks 


### Vegeta
```sh
$ sudo apt install vegeta  # Ubuntu/Debian

$ cat targets.txt | vegeta attack -duration=30s -rate=1000 -insecure -http2 | tee results.bin | vegeta report

$ echo "GET https://localhost:443/v1/user/345" | vegeta attack -duration=30s -rate=1000 -insecure -http2 | tee results.bin | vegeta report

Explicação dos Parâmetros:
	•	-duration=30s → O teste rodará por 30 segundos.
	•	-rate=1000 → Dispara 1000 requisições por segundo.
	•	-insecure → Ignora problemas de certificado TLS (se autoassinado).
	•	-http2 → Força o uso do HTTP/2.
	•	tee results.bin → Salva os resultados para análise posterior.
	•	vegeta report → Exibe um resumo das estatísticas.

$ vegeta plot < results.bin > plot.html && open plot.html

 ```


### k6 script GET
 ```js
import http from 'k6/http';
import { check } from 'k6';

export let options = {
    stages: [
        { duration: '30s', target: 500 }, // Ramp-up para 500 VUs
        { duration: '1m', target: 500 },  // Mantém 500 VUs
        { duration: '30s', target: 0 },   // Ramp-down
    ],
};

export default function () {
    let res = http.get('http://localhost:8080/v1/user/123'); // Alterado para HTTP sem TLS

    check(res, {
        'status is 200': (r) => r.status === 200,
        'no errors': (r) => !r.error,
    });
}
```

```sh
$ k6 run k6.script.get.http.j
```

### hey
```sh
$ hey -n 100000 -c 100 -h2 https://localhost:443/v1/user/123
```