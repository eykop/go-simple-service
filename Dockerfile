FROM golang:alpine AS base

WORKDIR /usr/src/app

FROM base AS alpine_builder


RUN apk update && apk add --no-cache gcc libffi-dev musl-dev postgresql-dev


FROM alpine_builder AS go_builder

COPY . .
RUN go get .
RUN go build -o webApp .
RUN ls -la

FROM base AS preflight
# preflight stage, prepare everything for the final stage on a clean image:
# * User
# * Dirs
# * Install required dependencies binaries and libs

# create the app user
RUN addgroup -S app && adduser -S app -G app

# create the appropriate directories
ENV HOME=/home/app
ENV APP_HOME=/home/app/web
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

# install dependencies
RUN apk update && apk add --no-cache libpq libffi


# get our prebuilt go app
COPY --from=go_builder /usr/src/app/webApp $APP_HOME

RUN chown -R app:app $APP_HOME


FROM preflight AS final
# final stage

# get our project entry point file!
COPY docker-entrypoint.sh .
RUN chmod +x "${APP_HOME}/docker-entrypoint.sh"

EXPOSE 3000

# chown all the files(including the prepared virtual env) to the app user
RUN chown -R app:app $APP_HOME

# change to the app user
USER app

ENTRYPOINT ["/home/app/web/docker-entrypoint.sh"]
