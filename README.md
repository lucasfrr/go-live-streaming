## Go Live Streaming :brazil:

Projetinho de servidor de live streaming utilizando Go, Nginx, Postgres e Docker.

#### Execução

Obs: certifique-se de ter `Docker` e `Docker compose` instalado em sua máquina.

- Baixe o repositório
- Depois, execute `docker compose up --build`

Quando a instância subir, teremos 3 serviços online:

- Um servidor RMTP usando NGINX, responsável pela transmissão
- Uma aplicação de autenticação escrita em Go, que recebe as credenciais da live e realiza a autenticação
- Um banco de dados Postgres, que tem duas lives salvas com `name` e `stream_key`

Para usar o servidor, faça uma transmissão usando OBS ou outro programa similar que suporte o protecolo RMTP.
Depois, vá até as configurações de transmissão. Na opção de serviços de streaming, escolha "personalizado". O link do servidor a ser passsado será `rtmp://localhost:1935/live` e a `stream_key` será umas das streamkeys salvas no banco

## Go Live Streaming :uk:

Simple server live streaming project using Go, Nginx, Postgres and Docker.