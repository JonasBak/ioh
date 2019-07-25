import fetch from "isomorphic-unfetch";
import Container from "components/container";
import Link from "next/link";

const Index = ({ unconfigured, configured }) => {
  return (
    <Container>
      <h2>IOH</h2>
      {unconfigured > 0 && (
        <div>
          <Link href="/unconfigured">
            <a>Click here</a>
          </Link>{" "}
          to configure your {unconfigured} unconfigured plants!
        </div>
      )}
      {configured > 0 && (
        <div>
          <Link href="/configured">
            <a>Click here</a>
          </Link>{" "}
          to look at your configured plants
        </div>
      )}
    </Container>
  );
};

Index.getInitialProps = async ({ req }) => {
  const res = await Promise.all([
    fetch("http://localhost:3000/api/unconfigured"),
    fetch("http://localhost:3000/api/configured")
  ]);
  const list = await Promise.all(res.map(r => r.json()));
  return { unconfigured: list[0].length, configured: list[1].length };
};

export default Index;
