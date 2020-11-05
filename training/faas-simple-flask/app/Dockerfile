FROM python:3-alpine

LABEL maintainer="John A. Fossum"

WORKDIR /app
COPY . .

RUN pip install -r requirements.txt

EXPOSE 8080

CMD ["python", "app.py"]