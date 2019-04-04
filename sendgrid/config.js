module.exports = logger =>
  require('common-env/withLogger')(logger)
    .on('env:fallback', function(fullKeyName, $default, $secure) {
      // see https://github.com/FGRibreau/common-env/issues/27
      if ($secure && !process.env[fullKeyName]) {
        throw new Error(`${fullKeyName} environment variable must be defined`);
      }
    })
    .getOrElseAll({
      smtp: {
        port: 25,
      },
      sendgrid: {
        api: {
          $default: 'SENDGRID_API_KEY',
          $secure: true,
        },
      },
    });
