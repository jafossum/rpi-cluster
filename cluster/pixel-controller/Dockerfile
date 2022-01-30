FROM arm32v7/python:3-buster

LABEL maintainer="John A. Fossum"
LABEL org.opencontainers.image.source="https://github.com/jafossum/rpi-cluster"

RUN pip3 install --no-cache-dir blinkt

WORKDIR /app

COPY app .

CMD ["python", "pixel-controller.py"]