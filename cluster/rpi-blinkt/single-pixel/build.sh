# Build image
docker build . -t single-pixel

# Tag image
docker tag pixel-controller jafossum/blinkt-single-pixel

# Push image
docker push jafossum/blinkt-single-pixel
