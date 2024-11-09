FROM ubuntu:latest
LABEL maintainer="jeneksero@gmail.com"
RUN apt-get update -y
RUN apt-get install -y python3-pip python3-dev build-essential python3-full python3-venv
RUN python3 -m venv /venv
COPY . .
ENV PATH="/venv/bin:$PATH"
RUN pip install --no-cache-dir -r requirements.txt
ENTRYPOINT ["python3"]
CMD ["main.py"]