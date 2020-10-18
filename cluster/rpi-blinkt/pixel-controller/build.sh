# Build image
docker build . -t pixel-controller

# Tag image
docker tag pixel-controller jafossum/blinkt-pixel-controller

# Push image
docker push jafossum/blinkt-pixel-controller
