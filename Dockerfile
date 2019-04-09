FROM ubuntu:xenial
RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y golang git apache2 openssh-server php php7.0-mysql libapache2-mod-php7.0 curl lynx-cur

WORKDIR /
RUN mkdir go
RUN mkdir go/src
RUN mkdir go/bin
RUN mkdir go/pkg
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH="/go/bin:${PATH}"

# Add GoPotUrself Golang code
WORKDIR $GOPATH/src/github.com/
RUN git clone -b packaging https://github.com/Mustard1/GoPotUrself.git

WORKDIR $GOPATH/src/github.com/GoPotUrself/shell/
RUN go build
RUN go install
WORKDIR $GOPATH/src/github.com/GoPotUrself/
RUN go build
RUN go install

WORKDIR /

# Expose port 8080 for shell
EXPOSE 8080

# Enable apache mods.
RUN a2enmod php7.0
RUN a2enmod rewrite

# Update the PHP.ini file, enable <? ?> tags and quieten logging.
RUN sed -i "s/short_open_tag = Off/short_open_tag = On/" /etc/php/7.0/apache2/php.ini
RUN sed -i "s/error_reporting = .*$/error_reporting = E_ERROR | E_WARNING | E_PARSE/" /etc/php/7.0/apache2/php.ini

# Manually set up the apache environment variables
ENV APACHE_RUN_USER www-data
ENV APACHE_RUN_GROUP www-data
ENV APACHE_LOG_DIR /var/log/apache2
ENV APACHE_LOCK_DIR /var/lock/apache2
ENV APACHE_PID_FILE /var/run/apache2.pid

# Expose apache.
EXPOSE 80

# Copy this repo into place.
ADD www /var/www/site
RUN chmod 777 /var/www/site/images 

# Update the default apache site with the config we created.
ADD ./apache-config.conf /etc/apache2/sites-enabled/000-default.conf

# By default start up apache in the foreground, override with /bin/bash for interative.
CMD /usr/sbin/apache2ctl -D FOREGROUND

