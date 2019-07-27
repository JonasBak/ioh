export default async (req, res) => {
  const tmp_res = await fetch("http://hub:5151/configured");
  const list = await tmp_res.json();
  res.status(200);
  res.json(list);
};
