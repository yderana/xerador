import Hashids from "hashids";
const hash = new Hashids("iORRj83UnOw53nfg", 7);
export function encodeHash(input: number) {
  return hash.encode(input);
}

export function decodeHash(value: string) {
  return hash.decode(value)[0];
}