const simpleParser = require('mailparser').simpleParser;
const debugParser = require('debug')('parser');
const debugTransform = require('debug')('transform');

function parse(buffer) {
  debugParser('got buffer %s', buffer);
  return simpleParser(buffer);
}

const makeAddress = (address, name) => ({
  name: name || address,
  email: address,
});

function formatAddresses(addressObject) {
  return addressObject.value.map(({ address, name }) =>
    makeAddress(address, name)
  );
}

const toSendGridEmail = session => email => {
  debugParser('got parsed email %s', JSON.stringify(email));

  const mail = {
    subject: email.subject,
    text: email.text ? email.text : undefined,
    html: email.html ? email.html : undefined,
    content: [],
    from: {},
    to: [],
  };

  ['cc', 'to', 'bcc'].forEach(part => {
    if (email[part]) {
      mail[part] = formatAddresses(email[part]);
    }
  });

  if (email['reply-to']) {
    mail.replyTo = formatAddresses(email['reply-to'])[0];
  }

  if (email.from) {
    mail.from = formatAddresses(email.from)[0];
  }

  if (session.envelope.mailFrom && session.envelope.mailFrom.address) {
    mail.from = makeAddress(session.envelope.mailFrom.address);
  }

  if (session.envelope.rcptTo) {
    mail.to = session.envelope.rcptTo.map(({ address }) =>
      makeAddress(address)
    );
  }

  return mail;
};

const bufferToSendGridEmail = (buffer, session) =>
  parse(buffer).then(toSendGridEmail(session));

// parse -> toSendGridEmail
module.exports = {
  bufferToSendGridEmail,
};
