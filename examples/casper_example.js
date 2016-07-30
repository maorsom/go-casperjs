var casper = require('casper').create();
casper.start('{{.Url}}');

casper.then(function() {
    this.echo(this.getTitle());
});
