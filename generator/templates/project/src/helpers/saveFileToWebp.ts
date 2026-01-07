import { uploadS3 } from "@helpers/s3Upload";

// tslint:disable-next-line:no-var-requires
const sharp = require("sharp");
import { snakeCase } from "lodash";
/* istanbul ignore next */
export async function saveFileToWebp(params: {
  file: any;
  customName?: string;
  width?: number;
  height?: number;
  directory?: string;
}) {
  const { file, width, height, directory } = params;
  let response = null;

  let size = {};
  if (width) {
    size = { ...size, width };
  }
  if (height) {
    size = { ...size, height };
  }

  size = { ...size, fit: sharp.fit.cover, position: sharp.strategy.entropy };

  // resize file before upload
  const original = await sharp(file.buffer).resize(size);

  // resize and convert file to webp
  const webp = await sharp(file.buffer)
    .resize(size)
    .webp();

  // rename file upload
  const timeNow = new Date();
  const originalName = file.originalname.split(".");
  const name = `${snakeCase(originalName[0])}_${timeNow.getTime()}`;

  // upload file
  const uploadOriginalFormat = await uploadS3(
    `${name}.${originalName[originalName.length - 1]}`,
    original,
    directory
  );

  // upload file webp
  const uploadWebp = await uploadS3(`${name}.webp`, webp, directory);

  response = {
    originalFormat: uploadOriginalFormat.path,
    webp: uploadWebp.path
  };

  return response;
}
