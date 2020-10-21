# LP-2020

Projeto de Linguagens de Programação, UnB, semestre 01/2020.

## Pré-requisitos

Para executar o trabalho é preciso baixar as bases de dados referentes ao ENEM 2019 em http://inep.gov.br/web/guest/microdados e descompactar o arquivo no diretório atual.
Em seguida, todos os pacotes do diretório `./Packages` (opcional) podem ser movidos para o `$GOPATH` local com o seguinte comando:

```
env GIT_TERMINAL_PROMPT=1 go get github.com/munak98/LP-2020/Packages/extract
```

## Rodando o trabalho

Dentro da pasta Src, o trabalho pode ser executado com o comando:

``
go run main.go 
```

Para executar o trabalho com outra base de dados, por exemplo do ENEM 2017, basta baixar e descompactar a base no diretório e trocar o caminho `microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv` por outro equivalente. 

