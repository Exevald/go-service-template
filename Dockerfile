FROM debian:11-slim
COPY bin/guardian /app/bin/

ENTRYPOINT [ "/app/bin/servicetemplate" ]
CMD ["serve"]