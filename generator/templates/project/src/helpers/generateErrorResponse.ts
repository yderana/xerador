// import { config, captureException } from "raven";
// config(process.env.SENTRY_URL).install();

/* istanbul ignore next */
export function errorResponse(error: any) {
  // captureException(error);
  let message = error.message;
  let statusCode = 400;
  if (error.response && error.response.data) {
    message =
      error.response.data.message ||
      (error.response.statusText || error.message);
  }

  if (error.response && error.response.status) {
    statusCode = error.response.status || 400;
  }

  return { statusCode, message };
}
