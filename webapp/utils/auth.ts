import { createContext, useState, useEffect, useContext } from "react";
import { AUTH_CONFIG } from "./config";
import { UserType } from "./types";

const { domain, clientId, redirectURI, audience } = AUTH_CONFIG;

export class AuthError extends Error {}

export const authorizeUrl = `https://${domain}/authorize?
response_type=id_token token&
client_id=${clientId}&
redirect_uri=${redirectURI}&
scope=openid profile email&
audience=${audience}&
state=xyzABC123&
nonce=eq...hPmz`;

type StateType = {
  isAuthenticated: boolean;
  accessToken?: string;
  user?: UserType;
  authorizeUrl: string;
  logout: () => void;
};

const initialState: StateType = {
  isAuthenticated: false,
  accessToken: undefined,
  user: undefined,
  authorizeUrl,
  logout: () => {}
};

export const AuthContext = createContext(initialState);

export const useAuth = (): StateType => useContext(AuthContext);

export const checkToken = (tokenType: string) => {
  const url = window.location.href;
  const auth_index = url.indexOf(tokenType + "=");
  if (auth_index === -1) {
    return;
  }
  const start = auth_index + (tokenType + "=").length;
  const token = url.substring(start).split("&")[0];
  return token;
};

const TOKEN_KEY = "access_token";
const USER_KEY = "user";

export const loadLocalStorage = (
  setAccessToken,
  setUser,
  setIsAuthenticated
) => {
  const accessToken = window.localStorage.getItem(TOKEN_KEY);
  const user = window.localStorage.getItem(USER_KEY);
  if (accessToken && user) {
    setAccessToken(accessToken);
    setUser(JSON.parse(user));
    setIsAuthenticated(true);
    return true;
  }
  return false;
};

export const setLocalStorage = (accessToken, user) => {
  window.localStorage.setItem(TOKEN_KEY, accessToken);
  window.localStorage.setItem(USER_KEY, JSON.stringify(user));
};

export const clearLocalStorage = () => {
  window.localStorage.removeItem(TOKEN_KEY);
  window.localStorage.removeItem(USER_KEY);
};
