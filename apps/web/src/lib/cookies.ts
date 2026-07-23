import Cookies from 'js-cookie';

export const TOKEN_KEYS = {
  ACCESS_TOKEN: 'access_token',
  REFRESH_TOKEN: 'refresh_token',
  SESSION_TOKEN: 'session_token',
};

export const setAuthCookies = (accessToken?: string, refreshToken?: string, sessionToken?: string) => {
  if (accessToken) {
    Cookies.set(TOKEN_KEYS.ACCESS_TOKEN, accessToken, { expires: 30, sameSite: 'lax' });
  }
  if (refreshToken) {
    Cookies.set(TOKEN_KEYS.REFRESH_TOKEN, refreshToken, { expires: 30, sameSite: 'lax' });
  }
  if (sessionToken) {
    Cookies.set(TOKEN_KEYS.SESSION_TOKEN, sessionToken, { expires: 30, sameSite: 'lax' });
  }
};

export const getAccessTokenCookie = (): string | undefined => {
  return Cookies.get(TOKEN_KEYS.ACCESS_TOKEN);
};

export const removeAuthCookies = () => {
  Cookies.remove(TOKEN_KEYS.ACCESS_TOKEN);
  Cookies.remove(TOKEN_KEYS.REFRESH_TOKEN);
  Cookies.remove(TOKEN_KEYS.SESSION_TOKEN);
};
