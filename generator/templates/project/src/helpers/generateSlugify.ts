import { snakeCase } from "lodash";
export function generateSlugify(text: string) {
  const slug = snakeCase(text).replace(/\_/g, "-");
  return slug;
}
