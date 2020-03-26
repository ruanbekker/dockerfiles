FROM golang
ADD postfix_exporter /postfix_exporter

ENTRYPOINT [ "/postfix_exporter" ]
CMD        [ "--web.listen-address=:9154", \
             "--web.telemetry-path=/metrics", \
             "--postfix.showq_path=/var/spool/postfix/public/showq", \
             "--postfix.logfile_path=/var/log/maillog" ]
