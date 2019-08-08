import Head from "components/head";
import Navbar from "components/navbar";
import { AuthWrapper } from "components/auth";

const Container = ({ children }) => (
  <div className="container">
    <AuthWrapper>
      <Head />
      <Navbar />
      {children}
    </AuthWrapper>
  </div>
);

export default Container;
