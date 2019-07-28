import { HUB_BASE_URL } from "utils/config";

export default async (req, res) => {
  if (req.method !== "POST") {
    return;
  }
  const body = JSON.parse(req.body);
  const post = await fetch(`${HUB_BASE_URL}/config?id=${body["id"]}`, {
    method: "POST",
    body: req.body
  });
  res.status(post.status);
  res.end(await post.text());
};
