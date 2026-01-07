// tslint:disable-next-line:no-var-requires
const aws = require("aws-sdk");
// tslint:disable-next-line:no-var-requires
const sharp = require("sharp");
// tslint:disable-next-line:no-var-requires
require("dotenv").config();
const config = {
  accessKeyId: process.env.AWS_KEY,
  secretAccessKey: process.env.AWS_SECRET_KEY,
  region: process.env.AWS_REGION,
  endpoint: process.env.AWS_ENDPOINT
};
aws.config.update(config);

/* istanbul ignore next */
export async function uploadS3(fileName: string, file: any, directory?: string) {
  let original: any = null;
  if (file.mimetype.split('/')[0] == 'image') {
    original = await sharp(file.buffer);
  } else {
    original = file.buffer;
  }

  const params = {
    Key: fileName,
    Body: original,
    Bucket: `${process.env.AWS_BUCKET}${directory ? `/${directory}` : ''}`,
    ACL: 'public-read'
  };
  const s3bucket = new aws.S3();

  const upload = await s3bucket.upload(params).promise();
  const response = {
    path: `https://${upload.Bucket}/${upload.Key}`,
    location: upload.Location
  };
  return response;
}
/* istanbul ignore next */
export async function deleteS3File(fileName: string, directory?: string) {
  const params = {
    Key: fileName,
    Bucket: `${process.env.AWS_BUCKET}${directory ? `/${directory}` : ''}`
  };
  const s3bucket = new aws.S3();

  const remove = await s3bucket.deleteObject(params).promise();
  return remove;
}
