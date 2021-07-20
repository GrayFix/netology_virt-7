FROM amazoncorretto

RUN curl -L https://pkg.jenkins.io/redhat-stable/jenkins.repo -o /etc/yum.repos.d/jenkins.repo

RUN rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key

RUN yum update -y

RUN yum install -y jenkins

ENV JENKINS_JAVA_OPTIONS="-Djava.awt.headless=true"
ENV JENKINS_HOME="/var/lib/jenkins"
ENV JENKINS_USER="jenkins"
ENV JENKINS_PORT="5000"
ENV JENKINS_ARGS=""

EXPOSE 5000/tcp

WORKDIR /home/app
COPY run_jenkins.sh /home/app/run_jenkins.sh

CMD ["/home/app/run_jenkins.sh"]
