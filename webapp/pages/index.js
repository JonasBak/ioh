import { BASE_URL } from "utils/config";
import Container from "components/container";
import { getClients } from "utils/req";
import Link from "next/link";
import PlantConfigForm from "components/plantConfigForm";

const Index = ({ clients }) => {
  const configured = clients.filter(c => !!c.config);
  const unconfigured = clients.filter(c => !c.config);
  return (
    <Container>
      <h2>IOH</h2>
      {unconfigured.length > 0 && (
        <div>
          <h2>Unconfigured</h2>
          {unconfigured.map(client => (
            <PlantConfigForm key={client.id} {...client} />
          ))}
        </div>
      )}
      {configured.length > 0 && (
        <div>
          <h2>Configured</h2>
          {configured.map(client => (
            <PlantConfigForm key={client.id} {...client} />
          ))}
        </div>
      )}
    </Container>
  );
};

Index.getInitialProps = async ({ req }) => {
  const queryResult = await getClients();
  return queryResult.data;
};

export default Index;
