import Head from "components/head";

const Container = ({ children }) => (
  <div className="container">
    <Head />
    {children}
  </div>
);

export default Container;
