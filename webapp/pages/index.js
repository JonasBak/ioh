import Link from "next/link";
import Head from "next/head";
import fetch from "isomorphic-unfetch";

const Index = ({ unconfigured, configured }) => {
  return (
    <div>
      <Head>
        <title>IOH</title>
        <link
          rel="stylesheet"
          href="https://cdn.rawgit.com/Chalarangelo/mini.css/v3.0.1/dist/mini-default.min.css"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
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
    </div>
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
