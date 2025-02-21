import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    // количество сотрудников
    vus: 20000, 
    duration: '1m',
    // RPS — 1k
    rps: 1000, 
    thresholds: {
        // 99.99% всех запросов должны обрабатываться за время менее 50 мс
        'http_req_duration': ['p(99.99)<50'],
        // SLI успешности ответа — 99.99%
        'http_req_failed': ['rate<0.001'], 
    },
};

export default function () {
    const userIndex = __VU; 
    var url = 'http://avito-shop-service:8080/api/auth'; 
    const payload = JSON.stringify({
        username: `word${userIndex+5}`, 
        password: `string${userIndex+5}`, 
    });

    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    var result = http.post(url, payload, params);
    
    check(result, {
        'is status 200 for /api/auth': (r) => r.status === 200,
    });

    const jsonResponse = JSON.parse(result.body);
    const authToken = jsonResponse.token;

    url = 'http://avito-shop-service:8080/api/info';
    params = {
        headers: {
            'Authorization': `${authToken}`,
        },
    };

    result = http.get(url, params);

    check(result, {
        'is status 200 for /api/info': (r) => r.status === 200,
    });

    sleep(1);
}