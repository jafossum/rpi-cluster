FROM arm32v7/python:3-buster

LABEL maintainer="John A. Fossum"
LABEL org.opencontainers.image.source="https://github.com/jafossum/rpi-cluster"

WORKDIR /app
COPY app .

CMD ["python", "single-pixel.py"]
