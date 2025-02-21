const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const { startStandaloneServer } = require("@apollo/server/standalone");

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: "user", url: "http://localhost:4001/graphql" },
            { name: "order", url: "http://localhost:4002/graphql" },
            { name: "delivery", url: "http://localhost:4003/graphql" },
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
