FROM ubuntu:latest

RUN apt-get update --yes && apt install --yes curl gnupg

RUN curl -L https://pkg.jenkins.io/debian-stable/jenkins.io.key | apt-key add -

RUN echo "deb https://pkg.jenkins.io/debian-stable binary/" >> /etc/apt/sources.list

RUN apt-get update --yes && apt-get install --yes openjdk-8-jre jenkins

ENV JENKINS_JAVA_OPTIONS="-Djava.awt.headless=true"
ENV JENKINS_HOME="/var/lib/jenkins"
ENV JENKINS_USER="jenkins"
ENV JENKINS_PORT="5000"
ENV JENKINS_ARGS=""

EXPOSE 5000/tcp

WORKDIR /home/app
COPY run_jenkins.sh /home/app/run_jenkins.sh

CMD ["/home/app/run_jenkins.sh"]
