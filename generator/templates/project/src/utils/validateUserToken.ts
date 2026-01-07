import * as jwtDecode from "jwt-decode";

export function ValidateUserToken(token: string) {
  let valid = false;
  const gettoken = token.split(" ");
  if (gettoken[0] === "Bearer" && gettoken[1]) {
    const payload: any = jwtDecode(gettoken[1]);
    valid = payload.adm && payload.staff && payload.authorised;
  }

  return valid;
}
