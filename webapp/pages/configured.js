import { BASE_URL } from "utils/config";
import fetch from "isomorphic-unfetch";
import Container from "components/container";
import PlantConfigForm from "components/plantConfigForm";

const Configured = ({ configured }) => {
  return (
    <Container>
      <h2>Configured</h2>
      <ul>
        {configured.map(({ id, water, plant }) => (
          <div key={id}>
            <div>Plant: {plant}</div>
            <div>Water: {water}</div>
            <PlantConfigForm key={id} id={id} />
          </div>
        ))}
      </ul>
    </Container>
  );
};

Configured.getInitialProps = async ({ req }) => {
  const res = await fetch(`${BASE_URL}/api/configured`);
  const list = await res.json();
  return { configured: list };
};

export default Configured;
