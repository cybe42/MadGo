FROM alpine:latest
COPY bin/MadGo /bin/
EXPOSE 8080
RUN chmod u+x /bin/MadGo
CMD [ "MadGo" ]