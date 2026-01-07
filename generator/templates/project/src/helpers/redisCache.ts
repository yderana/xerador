// tslint:disable-next-line:no-var-requires
const cacheManager = require("cache-manager");
// tslint:disable-next-line:no-var-requires
const redisStore = require("cache-manager-redis-store");
// tslint:disable-next-line:no-var-requires
require("dotenv").config();

const redisCache = cacheManager.caching({
  store: redisStore,
  host: process.env.REDIS_HOST,
  port: process.env.REDIS_PORT,
  auth_pass: process.env.REDIS_AUTH,
  prefix: process.env.REDIS_PREFIX
});

const redisClient = redisCache.store.getClient();

export function setRedis(key: any, value: any, withExp?: boolean, customExp?: string) {
  const result = null;
  /* istanbul ignore next */
  redisClient.on("error", error => {
    return error;
  });

  const ttl = process.env.REDIS_TTL;
  let properties = {};
  /* istanbul ignore next */
  if (ttl && withExp !== false) {
    properties = { ttl };
  }

  if (customExp) {
    properties = { ttl: customExp };
  }

  redisCache.set(key, value, properties, err => {
    /* istanbul ignore next */
    if (err) {
      throw err;
    }
  });
}

export async function getRedis(key) {
  return await redisCache.get(key);
}

/* istanbul ignore next */
export async function delRedis(key) {
  return await redisCache.del(key);
}

export async function flushRedis() {
  await redisCache.keys(`${process.env.REDIS_PREFIX}*`, (err, result) => {
    result.map(async data => {
      const key = data.split(`${process.env.REDIS_PREFIX}`)[1];
      await redisCache.del(key);
    });
  });
}
