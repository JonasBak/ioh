export default (req, res) => {
  res.status(200);
  res.json([
    { Host: "ghi", Name: "Basil", Water: 3 },
    { Host: "jkl", Name: "Thyme", Water: 2 }
  ]);
};
