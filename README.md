## Go Live Streaming :brazil:

Projetinho de servidor de live streaming utilizando Go, Nginx, Postgres e Docker.

#### Execução

Obs: certifique-se de ter `Docker` e `Docker compose` instalado em sua máquina.

- Baixe o repositório
- Depois, execute `docker compose up --build`

Quando a instância subir, teremos 4 serviços online:

- Um servidor RTMP usando NGINX, responsável pela transmissão
- Uma serviço de autenticação escrita em Go, que recebe as credenciais da live e realiza a autenticação
- Um serviço de playback, que serve os arquivos gerados pelo servidor RTMP através de um endpoint
- Um banco de dados Postgres, que tem duas lives salvas com `name` e `stream_key`

![Diagrama do projeto](https://drive.google.com/file/d/16dn3i6jV8bIBe-6PaZ81pO6zOyPs46T7/view?usp=sharing)

Para usar o servidor, faça uma transmissão usando OBS ou outro aplicativo similar que suporte o protecolo RMTP.
Depois, vá até as configurações de transmissão. Na opção de serviços de streaming, escolha "personalizado". O link do servidor a ser passsado será `rtmp://localhost:1935/live` e a `stream_key` a ser utilizada será uma das instâncias salvas no banco. O Formato da `stream_key` deve ser: `<nome_da_live>_<chave>`. Como, por exemplo: `livedego_087f601d-c4d6-48e5-b69e-3aa31ca62ad6`.

Feito isso, no seu aplicativo de transmissão, inicie a live. Ocorrendo como o esperado, o servidor RTMP vai começar a gerar os arquivos `.ts` e gera um arquivo `index.m38u`, que é uma playlist que indexa os fragmentos gerados e os reproduz na sequência correta.

Se quiser visualizar os arquivos gerados no volume, só precisa acessar o container do NGINX via terminal (`docker exec -it <id do container> bash`) e navegar até o diretório onde os arquivos são gerados, que fica em `hls/live/<nome da live>`. Ou você pode utilizar o Docker Desktop, caso tenha instalado.

Com os serviços rodando e a live ligada, você pode reproduzir a live na URL: `http://localhost:8001/live/<nomedalive>/index.m3u8`. No seu navegador, instale uma extensão que reproduza arquivos `.m3u8`, cole essa URL no navegador e você poderá ver a live sendo tocada.

## Go Live Streaming :uk:

Simple server live streaming project using Go, Nginx, Postgres and Docker.