import { graphql } from "react-relay/hooks";

// The $uuid variable is injected automatically from the route.
const UuidQuery = graphql`
  query UuidQuery($name: String!) {
    findNinja(name: $name) {
      name
      rank
    }
  }
`;

export default UuidQuery