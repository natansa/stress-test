## Testes do projeto
- Execute para gerar a imagem docker: docker build -t stress-test --no-cache .
- Execute para teste do stress-test no docker: docker run --rm stress-test --url="http://google.com" --requests=20 --concurrency=5