
# Projeto LP-2020-UnB

Projeto consiste em analisar dados sobre o ENEM de 2019 de todos estados do Brasil, através dos microdados 
fornecidos pelo inep, os dados coletados são:

### Análise Estadual

- Número de participantes 
- Médias de cada área de conhecimento 
- Número de participantes de Escola Pública e Privada

### Análise por raça

- Número de participantes de cada raça 
- Médias por área de conhecimento de cada raça
- Número de participantes de Escola Pública e Privada de cada raça

Objetivo é demonstrar os diferenciais e pontos fortes da linguagem Golang, 
ao realizar a leitura dos microdados e armazenamento dos dados obtidos utilizando 
processos de execuções paralelas, buscando obter um menor tempo de execução e melhor performance.

## Pré-requisitos

Baixar pacotes externos utilizados atravé dos comandos:

```bash
go get golang.org/x/text/encoding/charmap

go get github.com/montanaflynn/stats
```


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

