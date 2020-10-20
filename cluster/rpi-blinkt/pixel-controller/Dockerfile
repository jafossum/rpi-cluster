FROM balenalib/raspberry-pi-python:3.7-run

RUN apt-get update -qy && apt-get install -qy \
    python3-blinkt

WORKDIR /app

COPY app .

CMD ["python", "pixel-controller.py"]
