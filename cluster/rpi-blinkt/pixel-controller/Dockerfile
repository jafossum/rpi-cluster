FROM balenalib/raspberry-pi

RUN apt-get update -qy && apt-get install -qy \
    python3-pip \
    python3-blinkt

WORKDIR /app

COPY app .

CMD ["python3", "pixel-controller.py"]