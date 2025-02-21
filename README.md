# Merch-Shop
## Как запустить:
```bash
git clone git@github.com:jusque-a-la-fin/Merch-Shop.git && cd Merch-Shop && sudo docker compose up
```
Тесты находятся в папке [tests](https://github.com/jusque-a-la-fin/Merch-Shop/tree/main/internal/handlers/user).  
[E2E-тест на сценарий покупки мерча](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/e2e_buy_test.go).  
[E2E-тест на сценарий передачи монеток другим сотрудникам](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/e2e_send_test.go).  
[E2E-тест на сценарий аутентификации](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/e2e_auth_test.go).  
[E2E-тест на сценарий получения информации](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/e2e_get_test.go).  

В [compose.yaml](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/golangci.yml) были добавлены три сервиса: test, linters и k6.  
В сервисе test запускаются тесты.  
В сервисе linters запускаются линтеры.  
В сервисе k6 запускается нагрузочное тестирование (Grafana k6) через 15 секунд после завершения тестирования.  