import { HUB_BASE_URL } from "utils/config";
import { QUERIES } from "utils/gql";
import fetch from "isomorphic-unfetch";
import { GetClientsType, GetConfigType, SetConfigType } from "utils/types";

const doQuery = async (queryType, args = []) => {
  const req = await fetch(`${HUB_BASE_URL}/query`, {
    method: "POST",
    body: JSON.stringify({
      operationName: QUERIES[queryType].operationName,
      query: QUERIES[queryType].query,
      variables: QUERIES[queryType].variables
        .map((name, i) => ({ [name]: args[i] }))
        .reduce((obj, a) => ({ ...obj, ...a }), {})
    })
  });
  return await req.json();
};

export const getClients = (): Promise<GetClientsType> => doQuery("getClients");

export const getConfig = (clientId: string): Promise<GetConfigType> =>
  doQuery("getConfig", [clientId]);

export const setConfig = (
  clientId: string,
  plant: string,
  water: number
): Promise<SetConfigType> => doQuery("setConfig", [clientId, plant, water]);
