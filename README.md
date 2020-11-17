
# Projeto LP-2020-UnB

Projeto consiste em analisar dados sobre o ENEM dos anos de 2017, 2018 e 2019, através dos microdados fornecidos pelo [inep](http://inep.gov.br/web/guest/microdados), os dados coletados propiciam as seguintes análises:

### Análise Anual

- Número de participantes 
- Médias das áreas de conhecimento 
- Número de participantes de cada Raça 
- Médias das áreas de conhecimento de cada Raça
- Número de participantes de Escola Pública e Privada
- Médias das áreas de conhecimento de cada Tipo de Escola

### Análise Estadual

- Número de participantes 
- Médias das áreas de conhecimento 
- Número de participantes de cada Raça 
- Médias das áreas de conhecimento de cada Raça
- Número de participantes de Escola Pública e Privada
- Número de participantes de Escola Pública e Privada de cada Raça

Objetivo é demonstrar os diferenciais e pontos fortes da linguagem Golang, 
ao realizar a extração de uma grande quantidade de dados dos microdados e armazenamento dos dados obtidos utilizando processos de execuções paralelas e não paralelas, buscando comparar os desempenhos e verificar o aumento de performance que o paralelismo pode fornecer.

## Pré-requisitos

Para executar o trabalho é preciso baixar as bases de dados referentes ao ENEM 2019 em http://inep.gov.br/web/guest/microdados e descompactar o arquivo no diretório atual.
Em seguida, todos os pacotes do diretório `./Packages` (opcional) podem ser movidos para o `$GOPATH` local com o seguinte comando:

```
env GIT_TERMINAL_PROMPT=1 go get github.com/munak98/LP-2020/Packages/extract
```

## Rodando o trabalho

Dentro da pasta Src, o trabalho pode ser executado com o comando:

```
go run main.go 
```

Para executar o trabalho com outra base de dados, por exemplo do ENEM 2017, basta baixar e descompactar a base no diretório e trocar o caminho `microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv` por outro equivalente. 

