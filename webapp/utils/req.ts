import { HUB_BASE_URL } from "utils/config";
import { QUERIES } from "utils/gql";
import fetch from "isomorphic-unfetch";
import { GetClientsType, GetConfigType, SetConfigType } from "utils/types";

const doQuery = async (accessToken, queryType, args = []) => {
  const req = await fetch(`${HUB_BASE_URL}/query`, {
    method: "POST",
    headers: {
      Authorization: accessToken && `Bearer ${accessToken}`
    },
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

export const getClients = (accessToken: string): Promise<GetClientsType> =>
  doQuery(accessToken, "getClients");

export const getConfig = (
  accessToken: string,
  clientId: string
): Promise<GetConfigType> => doQuery(accessToken, "getConfig", [clientId]);

export const setConfig = (
  accessToken: string,
  clientId: string,
  plant: string,
  water: number
): Promise<SetConfigType> =>
  doQuery(accessToken, "setConfig", [clientId, plant, water]);
