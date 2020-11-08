# Build image
docker build -f ./app/Dockerfile_ARM -t jafossum/faas-simple-flask ./app

# Push image
docker push jafossum/faas-simple-flask
