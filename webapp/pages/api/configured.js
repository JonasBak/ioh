import { HUB_BASE_URL } from "utils/config";

export default async (req, res) => {
  const configured_req = await fetch(`${HUB_BASE_URL}/configured`);
  const configured = await configured_req.json();
  const tmp = await Promise.all(
    configured.map(id => fetch(`${HUB_BASE_URL}/config?id=${id}`))
  );
  const list = await Promise.all(tmp.map(r => r.json()));
  res.status(200);
  res.json(list.map((config, i) => ({ ...config, id: configured[i] })));
};
