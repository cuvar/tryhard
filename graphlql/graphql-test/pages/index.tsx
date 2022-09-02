import type { NextPage } from "next";
import { useQuery, gql } from "@apollo/client";

const GET_LOCATIONS = gql`
  query GetLocations {
    locations {
      id
      name
      description
      photo
    }
  }
`;

const GET_TODOS = gql`
  query GetTodos {
    todos {
      text
      done
      user {
        name
      }
    }
  }
`;

const Home: NextPage = () => {
  const { loading, error, data } = useQuery(GET_TODOS);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error {error.message}</p>;

  return (
    <div>
      <div>Hello world</div>
      <div>{data}</div>
    </div>
  );
};

export default Home;
