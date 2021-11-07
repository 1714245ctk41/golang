FROM alpine:3.8

WORKDIR /home/
COPY crawl.bin .
RUN chmod +x crawl.bin

EXPOSE 80
CMD ["./crawl.bin"]