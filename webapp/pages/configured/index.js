const Configured = ({ configured }) => {
  return (
    <div>
      <h2>Configured</h2>
      {configured.map(id => (
        <div key={id}>{id}</div>
      ))}
    </div>
  );
};

Configured.getInitialProps = async ({ req }) => {
  const res = await fetch("http://localhost:3000/api/configured");
  const list = await res.json();
  return { configured: list };
};

export default Configured;
