### Деплой релеев ChainBridge

Сеть **Equilibrium** может взаимодействовать с другими сетями, на
данный момент только с **Ethereum**. Для этого используется подсистема
Cross-chain Bridge, разработанная компанией ChainSafe и состоящая из
трёх совместно работающих компонентов: специального Substrate-палета,
набора смарт-контрактов в **Ethereum** и релей-демона
```ChainBridge```. Данный документ описывает установку и запуск
релеев. Особо надо отметить, что одновременно может работать несколько
инстансов релея, отличающихся конфигурационным файлом. Количество
одновременно запущенных релеев определяется параметром relayThreshold
и количеством зарегистрированных спец-аккаунтов — релейеров.

### Файлы

Необходимы три файла из репозитория
https://github.com/equilibrium-eosdt/ChainBridge, из каталога
scripts/docker:

- ```chainbridge.sh```
- ```config.json```
- ```Dockerfile.chainbridge```

### Установка: построение Docker-образа

В предположении, что текущий каталог scripts/docker:
```
docker build --force-rm --file Dockerfile.chainbridge --tag chainbridge0 .
```

### Запуск контейнера

```
docker run --name chainbridge0 --network host chainbridge0
```

### Конфигурация

Общая:

- Эндпоинты нод сетей Substrate и Ethereum (указываются в ```config.json```)
- Адреса смарт-контрактов сети Ethereum, обеспечивающих обмен (указываются в ```config.json```):
```
"bridge": "0x...",
"erc20Handler": "0x...",
"genericHandler": "0x...",
```
- Путь к каталогу keystore (указывается в скрипте ```chainbridge.sh```); подробнее в секции "Keystore"
- Путь к каталогу blockstore (указывается в скрипте ```chainbridge.sh```); не нужен, если запускать с опцией ```--latest```

Кроме того, каждый инстанс должен иметь свою, отличную от других,
конфигурацию: адреса двух аккаунтов (по одному на каждую сеть),
играющих особую роль релейеров (указываются в ```config.json```):
```
"from": "0x...",
"from": "0x...",
```

### Keystore

```ChainBridge``` требует наличия ключей для подписывания и отправки транзакций, а также для идентификации.

Как работать с ключами: ```chainbridge accounts --help```.

Пароль для keystore может быть передан через переменную окружения ```KEYSTORE_PASSWORD```.

Импортировать внешние ключи Ethereum:
```chainbridge accounts import --ethereum /path/to/key```.

Импортировать в keystore секретный ключ:
```chainbridge account import --privateKey key```.
