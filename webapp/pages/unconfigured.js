import { BASE_URL } from "utils/config";
import fetch from "isomorphic-unfetch";
import Container from "components/container";
import PlantConfigForm from "components/plantConfigForm";

const Unconfigured = ({ unconfigured }) => (
  <Container>
    <h2>Unconfigured</h2>
    {unconfigured.map(({ id }) => (
      <PlantConfigForm key={id} id={id} />
    ))}
  </Container>
);

Unconfigured.getInitialProps = async ({ req }) => {
  const res = await fetch(`${BASE_URL}/api/unconfigured`);
  const list = await res.json();
  return { unconfigured: list };
};

export default Unconfigured;
