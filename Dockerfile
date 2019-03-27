FROM ubuntu:xenial
RUN apt-get update && apt-get -y upgrade
RUN apt-get install -y golang git apache2 openssh-server php php7.0-mysql libapache2-mod-php7.0 curl lynx-cur

# Old SSHd server stuff TB removed
# EXPOSE 22
# ENTRYPOINT ["/entrypoint.sh"]
# COPY rootfs/ /
# RUN sed -i s/#PermitRootLogin.*/PermitRootLogin\ yes/ /etc/ssh/sshd_config && echo "root:root" | chpasswd


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

# Add GoPotUrself Golang code
WORKDIR $GOPATH/src/github.com/
RUN git clone https://github.com/Mustard1/GoPotUrself.git
WORKDIR /

# Enable SSHd server
# RUN mkdir /var/run/sshd
# RUN echo 'root:root' | chpasswd
# RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
# RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

# ENV NOTVISIBLE "in users profile"
# RUN echo "export VISIBLE=now" >> /etc/profile

# Messing around with adding additional users 
#RUN useradd -ms /bin/bash admin
#USER admin
#WORKDIR /home/admin

