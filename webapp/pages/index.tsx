import { useState } from "react";
import Container from "components/container";
import { getClients } from "utils/req";
import { ClientType, ConfigType } from "utils/types";
import Link from "next/link";
import PlantConfigForm from "components/plantConfigForm";

const ClientList = ({
  clients,
  title,
  onPost
}: {
  clients: ClientType[];
  title: string;
  onPost: (id: string, config: ConfigType) => void;
}) => (
  <div>
    <h2>{title}</h2>
    {clients.map(client => (
      <PlantConfigForm key={client.id} client={client} onPost={onPost} />
    ))}
  </div>
);

const Index = ({ clients: currentClients }: { clients: ClientType[] }) => {
  const [clients, setClients] = useState(currentClients);
  const configured = clients.filter(c => !!c.config);
  const unconfigured = clients.filter(c => !c.config);
  const onPost = (id, config) =>
    setClients(
      clients.map(client => (client.id === id ? { ...client, config } : client))
    );
  return (
    <Container>
      <h2>IOH</h2>
      {unconfigured.length > 0 && (
        <ClientList
          title="Unconfigured"
          clients={unconfigured}
          onPost={onPost}
        />
      )}
      {configured.length > 0 && (
        <ClientList title="Configured" clients={configured} onPost={onPost} />
      )}
    </Container>
  );
};

Index.getInitialProps = async ({ req }) => {
  const queryResult = await getClients();
  return queryResult.data;
};

export default Index;
