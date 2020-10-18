FROM balenalib/raspberry-pi

RUN apt-get update -qy && apt-get install -qy \
    python3-pip

WORKDIR /app

COPY app .

CMD ["python3", "single-pixel.py"]

