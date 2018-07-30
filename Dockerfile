FROM iron/go

WORKDIR /harambe_data_service

ADD ./bin/harambe_data_service /harambe_data_service

ARG PORT_ARG=3000
ENV PORT=${PORT_ARG}

EXPOSE ${PORT_ARG}

CMD ./harambe_data_service $PORT