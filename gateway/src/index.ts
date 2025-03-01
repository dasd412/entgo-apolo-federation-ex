import {ApolloGateway, IntrospectAndCompose} from "@apollo/gateway";
import {ApolloServer} from "@apollo/server";
import {startStandaloneServer} from "@apollo/server/standalone";

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: "user", url: "http://localhost:8081/graphql" },
            { name: "order", url: "http://localhost:8082/graphql" },
            { name: "delivery", url: "http://localhost:8083/graphql" },
        ],
    }),
});

async function startServer() {
    const server = new ApolloServer({ gateway });

    const { url } = await startStandaloneServer(server, {
        listen: { port: 4000 },
    });

    console.log(`🚀 Apollo Gateway running at ${url}`);
}

startServer();