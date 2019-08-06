export type QueryType = {
  operationName: string;
  query: string;
  variables: string[];
};

export type ConfigType = {
  plant: string;
  water: number;
};

export type ClientType = {
  id: string;
  active: boolean;
  config?: ConfigType;
};

export type GetClientsType = {
  data: {
    clients: ClientType[];
  };
};

export type GetConfigType = {
  data: {
    config: ConfigType;
  };
};

export type SetConfigType = {
  data: {
    setConfig: ConfigType;
  };
};
