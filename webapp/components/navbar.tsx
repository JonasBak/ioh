import { useAuth } from "utils/auth";

const Navbar = () => {
  const { isAuthenticated, authorizeUrl, logout, user } = useAuth();
  return (
    <div style={{ display: "flex", justifyContent: "space-between" }}>
      <h3>IOH</h3>
      <div>
        {isAuthenticated ? (
          <button onClick={() => logout()}>Sign out</button>
        ) : (
          <a href={authorizeUrl}>
            <button>Sign in</button>
          </a>
        )}
      </div>
    </div>
  );
};

export default Navbar;
