# Build image
docker build . -t single-pixel

# Tag image
docker tag single-pixel jafossum/blinkt-single-pixel

# Push image
docker push jafossum/blinkt-single-pixel
