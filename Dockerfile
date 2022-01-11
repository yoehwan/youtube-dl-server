FROM golang
ENV port 80
EXPOSE 80

# install youtube-dl
RUN apt-get update && apt-get upgrade
RUN apt-get install -y youtube-dl
#RUN apt install -y python3-pip
#RUN pip3 install --upgrade youtube-dl

# install ngrok
RUN apt-get install -y unzip wget
RUN wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip --no-check-certificate
RUN unzip ngrok-stable-linux-amd64.zip
RUN mv ./ngrok /usr/bin/ngrok

# install go app
WORKDIR /go/src/app
COPY . .
RUN go build -o app
RUN cp ./app /

CMD ["./app"]