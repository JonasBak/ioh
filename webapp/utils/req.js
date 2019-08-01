import { HUB_BASE_URL } from "utils/config";
import { QUERIES } from "utils/gql";
import fetch from "isomorphic-unfetch";

const doQuery = async (type, args = []) => {
  const req = await fetch(`${HUB_BASE_URL}/query`, {
    method: "POST",
    body: JSON.stringify({
      operationName: QUERIES[type].operationName,
      query: QUERIES[type].query,
      variables: QUERIES[type].variables
        .map((name, i) => ({ [name]: args[i] }))
        .reduce((obj, a) => ({ ...obj, ...a }), {})
    })
  });
  return await req.json();
};

export const getClients = () => doQuery("getClients");
export const getConfig = clientId => doQuery("getConfig", [clientId]);
export const setConfig = (clientId, plant, water) =>
  doQuery("setConfig", [clientId, plant, water]);
