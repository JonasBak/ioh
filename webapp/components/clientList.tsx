import { useAuth, AuthError } from "utils/auth";
import { useState, useEffect } from "react";
import { getClients } from "utils/req";
import { ClientType, ConfigType } from "utils/types";
import PlantConfigForm from "components/plantConfigForm";

const List = ({
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

const ClientList = () => {
  const { isAuthenticated, accessToken, authorizeUrl, logout } = useAuth();
  const [clients, setClients] = useState([]);
  const configured = clients.filter(c => !!c.config);
  const unconfigured = clients.filter(c => !c.config);

  useEffect(() => {
    const asyncWrapper = async () => {
      try {
        const queryResult = await getClients(accessToken);
        setClients(queryResult.data.clients);
      } catch (e) {
        if (e instanceof AuthError) {
          logout();
          location.href = authorizeUrl;
        } else {
          console.error(e);
        }
      }
    };
    if (isAuthenticated) asyncWrapper();
  }, [isAuthenticated]);

  const onPost = (id, config) =>
    setClients(
      clients.map(client => (client.id === id ? { ...client, config } : client))
    );
  return (
    <div>
      {unconfigured.length > 0 && (
        <List title="Unconfigured" clients={unconfigured} onPost={onPost} />
      )}
      {configured.length > 0 && (
        <List title="Configured" clients={configured} onPost={onPost} />
      )}
    </div>
  );
};

export default ClientList;
