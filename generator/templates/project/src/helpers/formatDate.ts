// tslint:disable-next-line:no-var-requires
const moment = require("moment");

/* istanbul ignore next */
export function ConvertDateTime(date: Date) {
  const newDate = date
    ? moment(date)
      .utcOffset(7)
      .format("YYYY-MM-DDTHH:mm:ss.sssZ")
    : null;
  return newDate;
}

export function DateTimeNow() {
  return moment()
    .utcOffset(7)
    .format("YYYY-MM-DDTHH:mm:ss.sssZ");
}
