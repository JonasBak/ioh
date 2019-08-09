import jwt_decode from "jwt-decode";
import { createContext, useState, useEffect, useContext } from "react";
import {
  authorizeUrl,
  AuthContext,
  loadLocalStorage,
  checkToken,
  setLocalStorage,
  clearLocalStorage
} from "utils/auth";
import { AUTH_CONFIG } from "utils/config";

export const AuthWrapper = ({ children }) => {
  const [accessToken, setAccessToken] = useState();
  const [user, setUser] = useState();
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  useEffect(() => {
    if (isAuthenticated || !window) return;

    const encodedAccessToken = checkToken("access_token");
    const encodedIdToken = checkToken("id_token");
    if (encodedAccessToken && encodedIdToken) {
      const user = jwt_decode(encodedIdToken);

      setAccessToken(encodedAccessToken);
      setUser(user);
      setIsAuthenticated(true);

      setLocalStorage(encodedAccessToken, user);
      window.history.pushState(null, "", "/");
    } else {
      loadLocalStorage(setAccessToken, setUser, setIsAuthenticated);
    }
  });
  const logout = () => {
    clearLocalStorage();
    setAccessToken(undefined);
    setUser(undefined);
    setIsAuthenticated(false);
  };
  return (
    <AuthContext.Provider
      value={{
        accessToken,
        isAuthenticated,
        authorizeUrl,
        user,
        logout
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
