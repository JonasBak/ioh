import Container from "components/container";
import ClientList from "components/clientList";
import { ClientType } from "utils/types";

const Index = ({ clients: currentClients }: { clients: ClientType[] }) => (
  <Container>
    <ClientList />
  </Container>
);

export default Index;
