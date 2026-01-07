import * as AJV from "ajv";

export const validate = (JSC: object, data: object) => {
  const ajv = new AJV({ allErrors: true });
  Object.keys(data).forEach((key: any) => {
    if (typeof data[key] == "string") {
      data[key] = data[key].trim();
    }
  });

  const valid = ajv.validate(JSC, data);
  const errorText =
    ajv.errorsText() && ajv.errorsText().toLocaleLowerCase() !== "no errors"
      ? ajv.errors.map((error: any) => `${error.dataPath.substring(1)} ${error.message}`).join(', ')
      : "";

  return {
    errorText,
    valid: !!valid
  };
};
