import { HUB_BASE_URL } from "utils/config";

export default async (req, res) => {
  const tmp_res = await fetch(`${HUB_BASE_URL}/unconfigured`);
  const list = await tmp_res.json();
  res.status(200);
  res.json(list);
};
