import fetch from "isomorphic-unfetch";
import Container from "components/container";

const Configured = ({ configured }) => {
  return (
    <Container>
      <h2>Configured</h2>
      <ul>
        {configured.map(({ Host, Name, Water }) => (
          <li key={Host}>{Name}</li>
        ))}
      </ul>
    </Container>
  );
};

Configured.getInitialProps = async ({ req }) => {
  const res = await fetch("http://localhost:3000/api/configured");
  const list = await res.json();
  return { configured: list };
};

export default Configured;
