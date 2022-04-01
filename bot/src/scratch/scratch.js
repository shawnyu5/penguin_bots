const { Injectable, Logger } = require('@nestjs/common');
const { Cron } = require('@nestjs/schedule');

 @Cron('45 * * * * *')
function handleCron() {
   this.logger.debug('Called when the current second is 45');
}
