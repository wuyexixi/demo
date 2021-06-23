FROM image-registry.openshift-image-registry.svc:5000/openshift/devops-toolbox:v0.0

RUN yum install python38 -y && \
    yum install python3-pip -y

# We copy just the requirements.txt first to leverage Docker cache
COPY ./requirements.txt /app/requirements.txt

WORKDIR /app

RUN pip3 install -r requirements.txt

COPY . /app

ENTRYPOINT [ "python3" ]

CMD [ "app.py" ]