import fetch from "isomorphic-unfetch";
import Container from "components/container";

const Unconfigured = ({ unconfigured }) => {
  return (
    <Container>
      <h2>Unconfigured</h2>
      <ul>
        {unconfigured.map(id => (
          <li key={id}>{id}</li>
        ))}
      </ul>
    </Container>
  );
};

Unconfigured.getInitialProps = async ({ req }) => {
  const res = await fetch("http://localhost:3000/api/unconfigured");
  const list = await res.json();
  return { unconfigured: list };
};

export default Unconfigured;
