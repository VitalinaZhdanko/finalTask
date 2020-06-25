FROM webdevops/liquibase:postgres

USER root

WORKDIR /liquibase

COPY ./entripoints/migrator.sh .
RUN chmod +x ./migrator.sh

COPY ./app/models/sql/ .

ENTRYPOINT ["./migrator.sh"]

CMD ["update"]