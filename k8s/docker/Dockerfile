FROM ubuntu

RUN apt-get update && \
    apt-get install -y python3 python3-pip python3-venv

WORKDIR /opt/source-code

COPY . /opt/source-code

RUN python3 -m venv /opt/venv
RUN /opt/venv/bin/pip install flask flask-mysql

ENV PATH="/opt/venv/bin:$PATH"
ENV FLASK_APP=app.py

ENTRYPOINT ["flask", "run", "--host=0.0.0.0"]