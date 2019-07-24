const Unconfigured = ({ unconfigured }) => {
  return (
    <div>
      <h2>Unconfigured</h2>
      {unconfigured.map(id => (
        <div key={id}>{id}</div>
      ))}
    </div>
  );
};

Unconfigured.getInitialProps = async ({ req }) => {
  const res = await fetch("http://localhost:3000/api/unconfigured");
  const list = await res.json();
  return { unconfigured: list };
};

export default Unconfigured;
