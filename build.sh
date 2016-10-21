docker build -t swerter/latexapi_build -f Dockerfile.build .
docker run --rm -v "$PWD":/usr/src/myapp -it swerter/latexapi_build
docker build -t swerter/latexapi .
