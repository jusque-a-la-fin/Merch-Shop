# Merch-Shop
## Как запустить:
```bash
git clone git@github.com:jusque-a-la-fin/Merch-Shop.git && cd Merch-Shop && sudo docker compose up
```
Тесты запускаются в сервисе test, который был добавлен в compose.yaml.
Тесты находятся в папке (https://github.com/jusque-a-la-fin/Merch-Shop/tree/main/internal/handlers/user/tests).
E2E-тест на сценарий покупки мерча: [](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/tests/e2e_buy_test.go).
E2E-тест на сценарий передачи монеток другим сотрудникам: [](https://github.com/jusque-a-la-fin/Merch-Shop/blob/main/internal/handlers/user/tests/e2e_send_test.go).
