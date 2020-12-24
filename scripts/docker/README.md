### Деплой релеев ChainBridge

Сеть **Equilibrium** может взаимодействовать с другими сетями, на
данный момент только с **Ethereum**. Для этого используется подсистема
Cross-chain Bridge, разработанная компанией ChainSafe и состоящая из
трёх совместно работающих компонентов: специального Substrate-палета,
набора смарт-контрактов в **Ethereum** и релей-демона
```ChainBridge```. Данный документ описывает установку и запуск
релеев. Особо надо отметить, что одновременно будет работать несколько
инстансов релея, отличающихся конфигурационным файлом.

### Файлы

Необходимы три файла из репозитория
https://github.com/equilibrium-eosdt/ChainBridge, из каталога
scripts/docker:

- chainbridge.sh
- config.json
- Dockerfile.chainbridg

### Установка: построение Docker-образа

В предположении, что текущий каталог scripts/docker:
docker build --force-rm --file Dockerfile.chainbridge --tag chainbridge0 .

### Запуск контейнера

docker run --name chainbridge0 --network host chainbridge0
