const logger = require('winston');
const config = require('./config')(logger);

const sgMail = require('@sendgrid/mail');
sgMail.setApiKey(config.sendgrid.api);

const SMTPServer = require('smtp-server').SMTPServer;
const { bufferToSendGridEmail } = require('./utils');
const debug = require('debug')('server');

// start-up the SMTP server (no auth)
new SMTPServer({
  secure: false,
  disabledCommands: ['AUTH'],
  onData: function(stream, session, callback) {
    let buffer = '';
    let _mail;

    // parse a message as it is received
    stream.setEncoding('utf8');
    stream.on('data', function(part) {
      buffer += part;
    });

    // message fully received
    stream.on('end', function() {
      bufferToSendGridEmail(buffer, session)
        .then(mail => {
          _mail = mail;
          return sgMail.send(mail, false);
        })
        .then(() => {
          logger.info('forwarded', JSON.stringify(_mail));
          callback();
        })
        .catch(err => {
          debug('Error', err.message);
          debug('Error stack', err.stack);

          const { message, code, response } = err;

          logger.error(
            'failed forwarding',
            JSON.stringify(
              {
                error: { message, code, response },
                mail: _mail,
              },
              null,
              2
            )
          );

          callback();
        });
    });
  },
}).listen(config.smtp.port);

logger.info(
  'SMTP Forward listening on port %s for SMTP traffic.',
  config.smtp.port
);
